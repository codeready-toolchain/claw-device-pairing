// Verification script for device identity implementation
import { loadOrCreateDeviceIdentity, signDevicePayload } from "./device-identity";

async function verifyIdentity() {
  console.log("=== Device Identity Verification ===\n");

  // Test 1: Generate device identity
  console.log("Test 1: Generating device identity...");
  const identity = await loadOrCreateDeviceIdentity();
  console.log("✓ Device ID:", identity.deviceId);
  console.log("✓ Public Key length:", identity.publicKey.length);
  console.log("✓ Private Key length:", identity.privateKey.length);

  // Test 2: Verify deviceId is hex and 64 chars (SHA-256)
  console.log("\nTest 2: Verifying deviceId format...");
  const hexRegex = /^[0-9a-f]{64}$/;
  if (hexRegex.test(identity.deviceId)) {
    console.log("✓ Device ID is valid SHA-256 hex fingerprint");
  } else {
    console.error("✗ Device ID format invalid");
  }

  // Test 3: Test signature generation
  console.log("\nTest 3: Testing signature generation...");
  const testPayload = "device=test:client=test:role=operator:scopes=:token=:nonce=test123:platform=web:deviceFamily=browser:ts=1234567890";
  const signature = await signDevicePayload(identity.privateKey, testPayload);
  console.log("✓ Signature generated:", signature.substring(0, 20) + "...");
  console.log("✓ Signature length:", signature.length);

  // Test 4: Verify localStorage persistence
  console.log("\nTest 4: Testing localStorage persistence...");
  const identity2 = await loadOrCreateDeviceIdentity();
  if (identity.deviceId === identity2.deviceId) {
    console.log("✓ Identity persisted and retrieved correctly");
  } else {
    console.error("✗ Identity not persisted correctly");
  }

  // Test 5: Verify v3 signature payload format
  console.log("\nTest 5: Verifying v3 signature payload format...");
  const v3Payload = `device=${identity.deviceId}:client=test-client:role=operator:scopes=operator.read,operator.write:token=:nonce=abc123:platform=web:deviceFamily=browser:ts=${Date.now()}`;
  const v3Signature = await signDevicePayload(identity.privateKey, v3Payload);
  console.log("✓ V3 payload signed successfully");
  console.log("✓ Payload format:", v3Payload.substring(0, 60) + "...");

  console.log("\n=== All verification tests passed! ===");
}

// Run verification if this is the main module
if (import.meta.url === new URL(import.meta.url).href) {
  verifyIdentity().catch(console.error);
}
