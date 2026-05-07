## 1. Environment Configuration

- [x] 1.1 Add NAMESPACE environment variable to deployment manifest using downward API
- [x] 1.2 Add CLAW_INSTANCE environment variable to deployment manifest for instance label value
- [x] 1.3 Verify environment variable defaults and fallback behavior in code

## 2. Kubernetes Client Setup

- [x] 2.1 Add Kubernetes client-go dependencies to go.mod
- [x] 2.2 Initialize in-cluster Kubernetes client at server startup
- [x] 2.3 Add graceful handling for non-cluster environments (local development)
- [x] 2.4 Store client reference for reuse across requests

## 3. CR Type Definitions

- [x] 3.1 Verify ClawDevicePairingRequest CRD exists or create API types
- [x] 3.2 Define Go struct for ClawDevicePairingRequest with spec fields (RequestID, Selector)
- [x] 3.3 Add appropriate Kubernetes metadata and type registration

## 4. Request ID Sanitization

- [x] 4.1 Implement DNS-1123 validation for requestId
- [x] 4.2 Implement sanitization logic (lowercase, alphanumeric with hyphens, max 63 chars)
- [x] 4.3 Add unit tests for sanitization edge cases

## 5. Instance Label Retrieval

- [x] 5.1 Read CLAW_INSTANCE environment variable at startup
- [x] 5.2 Construct label selector from instance label key and value
- [x] 5.3 Add logging for missing instance label scenario
- [x] 5.4 Define fallback behavior when instance label is unavailable

## 6. CR Creation Logic

- [x] 6.1 Add CR creation function that accepts requestId and returns error
- [x] 6.2 Set CR name using sanitized requestId
- [x] 6.3 Populate spec.RequestID field
- [x] 6.4 Populate spec.Selector field with instance label selector
- [x] 6.5 Handle AlreadyExists error gracefully
- [x] 6.6 Add comprehensive error logging for CR creation failures
- [x] 6.7 Return error to caller when CR creation fails

## 7. Integration with HandlePairDevice

- [x] 7.1 Call CR creation function from HandlePairDevice method
- [x] 7.2 Pass requestId from incoming PairingRequest to CR creation
- [x] 7.3 Add logging for successful CR creation
- [x] 7.4 Return error to UI when CR creation fails
- [x] 7.5 Return succinct error message to UI: "Something wrong happened, could not pair the device"

## 8. UI Error Handling

- [x] 8.1 Update UI pairing handler to catch CR creation errors from backend
- [x] 8.2 Display user-friendly error message when pairing fails
- [x] 8.3 Add appropriate error UI state (error banner, toast, or inline message)
- [x] 8.4 Ensure error message matches backend error: "Something wrong happened, could not pair the device"

## 9. Testing

- [x] 9.1 Add unit tests for CR creation function
- [x] 9.2 Add unit tests for requestId sanitization
- [x] 9.3 Add unit tests for selector construction
- [x] 9.4 Add integration test for HandlePairDevice with CR creation
- [x] 9.5 Test error scenarios (missing env vars, CR creation failure, AlreadyExists)
- [x] 9.6 Test UI error handling when backend returns CR creation error

## 10. Documentation

- [x] 10.1 Update deployment documentation with new environment variables
- [x] 10.2 Document CR lifecycle and operator integration expectations
- [x] 10.3 Add troubleshooting guide for common CR creation issues
