import { DeviceIdentity, loadOrCreateDeviceIdentity, signDevicePayload } from "./device-identity";

class HandshakeError extends Error {
  code: string;
  details?: {
    code?: string;
    reason?: string;
    requestId?: string;
  };

  constructor(code: string, message: string, details?: { code?: string; reason?: string; requestId?: string }) {
    super(message);
    this.name = 'HandshakeError';
    this.code = code;
    this.details = details;
  }
}

type ChallengeEvent = {
  type: "event";
  event: "connect.challenge";
  payload: {
    nonce: string;
    ts: number;
  };
};

type ConnectRequest = {
  type: "req";
  id: string;
  method: "connect";
  params: {
    minProtocol: number;
    maxProtocol: number;
    client: {
      id: string;
      version: string;
      platform: string;
      mode: string;
    };
    role: string;
    scopes: string[];
    caps: string[];
    commands: string[];
    permissions: Record<string, boolean>;
    auth: {
      token?: string;
    };
    locale: string;
    userAgent: string;
    device: {
      id: string;
      publicKey: string;
      signature: string;
      signedAt: number;
      nonce: string;
    };
  };
};

type ConnectResponse = {
  type: "res";
  id: string;
  ok: boolean;
  payload?: {
    type: "hello-ok";
    protocol: number;
    server: {
      version: string;
      connId: string;
    };
    features: {
      methods: string[];
      events: string[];
    };
    snapshot: Record<string, unknown>;
    policy: {
      maxPayload: number;
      maxBufferedBytes: number;
      tickIntervalMs: number;
    };
    auth?: {
      deviceToken: string;
      role: string;
      scopes: string[];
    };
  };
  error?: {
    code: string;
    message: string;
    details?: {
      code?: string;
      reason?: string;
      requestId?: string;
    };
  };
};

type SignaturePayloadParams = {
  deviceId: string;
  clientId: string;
  clientMode: string;
  role: string;
  scopes: string[];
  token: string;
  nonce: string;
  signedAtMs: number;
};

function buildDeviceAuthPayload(params: SignaturePayloadParams): string {
  const scopes = params.scopes.join(",");
  const token = params.token ?? "";
  return [
    "v2",
    params.deviceId,
    params.clientId,
    params.clientMode,
    params.role,
    scopes,
    String(params.signedAtMs),
    token,
    params.nonce,
  ].join("|");
}

type ConnectParams = {
  clientId: string;
  clientVersion: string;
  mode: string;
  role: string;
  scopes: string[];
  platform?: string;
  deviceFamily?: string;
  token?: string;
  gatewayUrl: string;
};

async function buildConnectRequest(
  params: ConnectParams,
  identity: DeviceIdentity,
  challengeNonce: string
): Promise<ConnectRequest> {
  const signedAtMs = Date.now();

  const signaturePayload = buildDeviceAuthPayload({
    deviceId: identity.deviceId,
    clientId: params.clientId,
    clientMode: params.mode,
    role: params.role,
    scopes: params.scopes,
    token: params.token || "",
    nonce: challengeNonce,
    signedAtMs,
  });

  console.log('Signature payload:', signaturePayload);
  console.log('Device ID:', identity.deviceId);
  console.log('Public key:', identity.publicKey);
  console.log('Auth token:', params.token || '(none)');

  const signature = await signDevicePayload(identity.privateKey, signaturePayload);
  console.log('Signature:', signature);

  return {
    type: "req",
    id: crypto.randomUUID(),
    method: "connect",
    params: {
      minProtocol: 3,
      maxProtocol: 3,
      client: {
        id: params.clientId,
        version: params.clientVersion,
        platform: params.platform || "web",
        mode: params.mode,
      },
      role: params.role,
      scopes: params.scopes,
      caps: [],
      commands: [],
      permissions: {},
      auth: {
        token: params.token,
      },
      locale: navigator.language || "en-US",
      userAgent: `${params.clientId}/${params.clientVersion}`,
      device: {
        id: identity.deviceId,
        publicKey: identity.publicKey,
        signature,
        signedAt: signedAtMs,
        nonce: challengeNonce,
      },
    },
  };
}

function waitForChallenge(ws: WebSocket, timeoutMs: number = 10_000): Promise<ChallengeEvent> {
  return new Promise((resolve, reject) => {
    const timeout = setTimeout(() => {
      ws.removeEventListener("message", handler);
      reject(new Error("Timeout waiting for connect.challenge"));
    }, timeoutMs);

    const handler = (event: MessageEvent) => {
      try {
        const msg = JSON.parse(event.data);
        if (msg.type === "event" && msg.event === "connect.challenge") {
          clearTimeout(timeout);
          ws.removeEventListener("message", handler);
          resolve(msg as ChallengeEvent);
        }
      } catch (err) {
        clearTimeout(timeout);
        ws.removeEventListener("message", handler);
        reject(new Error(`Failed to parse challenge: ${err}`));
      }
    };

    ws.addEventListener("message", handler);
  });
}

function waitForConnectResponse(
  ws: WebSocket,
  requestId: string,
  timeoutMs: number = 30_000
): Promise<ConnectResponse> {
  return new Promise((resolve, reject) => {
    const timeout = setTimeout(() => {
      ws.removeEventListener("message", handler);
      reject(new Error("Timeout waiting for connect response"));
    }, timeoutMs);

    const handler = (event: MessageEvent) => {
      try {
        const msg = JSON.parse(event.data);
        if (msg.type === "res" && msg.id === requestId) {
          clearTimeout(timeout);
          ws.removeEventListener("message", handler);
          resolve(msg as ConnectResponse);
        }
      } catch (err) {
        clearTimeout(timeout);
        ws.removeEventListener("message", handler);
        reject(new Error(`Failed to parse connect response: ${err}`));
      }
    };

    ws.addEventListener("message", handler);
  });
}

export type HandshakeResult = {
  deviceToken?: string;
  protocol: number;
  server: {
    version: string;
    connId: string;
  };
  role: string;
  scopes: string[];
};

export async function performHandshake(params: ConnectParams): Promise<HandshakeResult> {
  return new Promise((resolve, reject) => {
    const ws = new WebSocket(params.gatewayUrl);

    ws.addEventListener("error", (err) => {
      reject(new Error(`WebSocket error: ${err}`));
    });

    ws.addEventListener("open", async () => {
      try {
        // Step 1: Wait for challenge
        const challenge = await waitForChallenge(ws);

        // Step 2: Load or create device identity
        const identity = await loadOrCreateDeviceIdentity();

        // Step 3: Build and send connect request
        const connectRequest = await buildConnectRequest(params, identity, challenge.payload.nonce);
        console.log('Sending connect request:', JSON.stringify(connectRequest, null, 2));
        ws.send(JSON.stringify(connectRequest));

        // Step 4: Wait for response
        const response = await waitForConnectResponse(ws, connectRequest.id);

        if (!response.ok) {
          const errorCode = response.error?.code || "UNKNOWN_ERROR";
          const errorMessage = response.error?.message || "Connect failed";
          reject(new HandshakeError(errorCode, errorMessage, response.error?.details));
          return;
        }

        if (!response.payload) {
          reject(new Error("Connect succeeded but no payload received"));
          return;
        }

        resolve({
          deviceToken: response.payload.auth?.deviceToken,
          protocol: response.payload.protocol,
          server: response.payload.server,
          role: response.payload.auth?.role || params.role,
          scopes: response.payload.auth?.scopes || params.scopes,
        });
      } catch (err) {
        reject(err);
      }
    });
  });
}
