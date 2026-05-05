## 1. Update Token Extraction Logic

- [x] 1.1 Read current token extraction code in ui/src/App.jsx
- [x] 1.2 Update the regex to remove both `#/?` prefix and optional `token=` prefix
- [x] 1.3 Change the regex from `/^#\/?/` to `/^#\/?(?:token=)?/`
- [x] 1.4 Verify the code handles all fragment formats: `#value`, `#/value`, `#token=value`, `#/token=value`

## 2. Test Token Extraction

- [x] 2.1 Test with URL fragment `#abc123` - should extract `abc123`
- [x] 2.2 Test with URL fragment `#token=abc123` - should extract `abc123`
- [x] 2.3 Test with URL fragment `#/token=abc123` - should extract `abc123`
- [x] 2.4 Test with no URL fragment - should pass undefined to performHandshake
- [x] 2.5 Verify console log shows "Using auth token from URL fragment" when token is present

## 3. Verify WebSocket Handshake

- [x] 3.1 Start the application with a token in the URL fragment
- [x] 3.2 Verify the WebSocket handshake includes the correct token value in params.auth.token
- [x] 3.3 Verify authentication succeeds with the extracted token
- [x] 3.4 Check that no `token=` prefix is sent to the server
