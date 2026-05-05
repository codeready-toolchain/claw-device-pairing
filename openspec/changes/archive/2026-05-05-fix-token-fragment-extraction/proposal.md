## Why

The UI currently extracts the authentication token from the URL fragment incorrectly, passing the entire fragment value (e.g., `"token=abc123"`) instead of just the token value (e.g., `"abc123"`) to the WebSocket handshake. This causes authentication failures because the server expects only the token value in `params.auth.token`.

## What Changes

- Update token extraction logic in `App.jsx` to parse the URL fragment correctly
- Remove the `token=` prefix when extracting the token value from the URL fragment
- Ensure only the token value is passed to `performHandshake()` in the `token` parameter

## Capabilities

### New Capabilities
<!-- None - this is a bug fix -->

### Modified Capabilities
- `auto-handshake`: Token extraction from URL fragment now correctly strips the `token=` prefix

## Impact

- `ui/src/App.jsx` - token extraction logic
- WebSocket authentication flow - will now work correctly when token is provided in URL fragment
- No breaking changes - this fixes existing broken behavior
