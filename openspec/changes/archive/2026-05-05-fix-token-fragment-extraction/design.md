## Context

The UI extracts an authentication token from the URL fragment to pass to the WebSocket handshake. Currently, the code removes the `#` or `#/` prefix from the fragment but does not handle the `token=` prefix, causing the entire string `token=abc123` to be passed instead of just `abc123`.

The current implementation in `App.jsx`:
```javascript
const fragment = window.location.hash
const token = fragment ? fragment.replace(/^#\/?/, '') : undefined
```

This works for fragments like `#abc123` but fails for `#token=abc123`.

## Goals / Non-Goals

**Goals:**
- Extract only the token value from URL fragments in the format `#token=<value>`
- Support both `#<value>` and `#token=<value>` formats for backward compatibility
- Pass only the token value to `performHandshake()` in the `token` parameter

**Non-Goals:**
- Supporting complex query string parsing in the fragment (only the `token=` prefix matters)
- Changing the handshake protocol or server-side token validation
- Supporting multiple parameters in the URL fragment

## Decisions

### Decision 1: Parse token with simple string replacement

**Rationale**: The URL fragment format is controlled and simple - either `#<value>` or `#token=<value>`. A regex or simple string replacement is sufficient and avoids adding URL parsing dependencies.

**Alternatives considered**:
- URLSearchParams: Would work but is overkill for a single parameter and requires prepending `?` to the fragment
- Complex regex: Unnecessary complexity for this simple case

**Decision**: Use a regex that removes both `#/?` prefix and optional `token=` prefix:
```javascript
const token = fragment.replace(/^#\/?(?:token=)?/, '')
```

### Decision 2: Maintain backward compatibility

**Rationale**: Some deployments or test scenarios might use `#<value>` directly. The regex pattern should support both formats.

**Implementation**: The `(?:token=)?` part makes the `token=` prefix optional, so both formats work.

## Risks / Trade-offs

**[Risk]** Fragment might contain URL-encoded characters (e.g., `#token=abc%2B123`)
→ **Mitigation**: The extracted value is passed directly to the handshake without decoding, which is correct - the server will handle any necessary decoding.

**[Risk]** Fragment might have unexpected format (e.g., `#token=value&other=data`)
→ **Mitigation**: The current simple parsing will extract `value&other=data` as the token value. This is acceptable because the server will reject invalid tokens. Adding full query string parsing is out of scope for this fix.

**[Trade-off]** Using regex instead of URL parsing library
→ **Acceptable**: Keeps the code simple and dependency-free. The fragment format is controlled and well-defined.
