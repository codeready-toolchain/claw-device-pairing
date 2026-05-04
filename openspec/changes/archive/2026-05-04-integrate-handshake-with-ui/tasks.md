## 1. Import Dependencies

- [x] 1.1 Import useState and useEffect from React in App.jsx
- [x] 1.2 Import performHandshake from './handshake' in App.jsx

## 2. Add State Management

- [x] 2.1 Add handshakeStatus state ('idle' | 'loading' | 'success' | 'error')
- [x] 2.2 Add handshakeError state (string | null)
- [x] 2.3 Add deviceToken state (string | null)

## 3. Implement Handshake Trigger

- [x] 3.1 Add useEffect hook with empty dependency array
- [x] 3.2 Set handshakeStatus to 'loading' before calling performHandshake
- [x] 3.3 Call performHandshake with gateway URL from import.meta.env.VITE_GATEWAY_URL or default 'ws://localhost:8080/ws'
- [x] 3.4 Pass clientId as 'device-pairing-ui', role as 'operator', scopes as empty array, and clientVersion
- [x] 3.5 On success, set handshakeStatus to 'success' and store deviceToken
- [x] 3.6 On error, set handshakeStatus to 'error' and store error message

## 4. Update ProgressStepper for Step Completion

- [x] 4.1 Add variant="success" to first ProgressStep when handshakeStatus === 'success'
- [x] 4.2 Keep first ProgressStep in default/pending state when handshakeStatus !== 'success'

## 5. Implement Conditional Rendering

- [x] 5.1 Show loading message/spinner in CardBody when handshakeStatus === 'loading'
- [x] 5.2 Show ProgressStepper in CardBody when handshakeStatus === 'success'
- [x] 5.3 Show error message in CardBody when handshakeStatus === 'error'
- [x] 5.4 Include error details from handshakeError in error display

## 6. Testing and Verification

- [x] 6.1 Test page load triggers handshake automatically
- [x] 6.2 Verify loading state displays during handshake
- [x] 6.3 Verify success state shows ProgressStepper with first step complete
- [x] 6.4 Verify error state shows error message (test with invalid gateway URL)
- [x] 6.5 Test with actual gateway if available or mock WebSocket for end-to-end flow
