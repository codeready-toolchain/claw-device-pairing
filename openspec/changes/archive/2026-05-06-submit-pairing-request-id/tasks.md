## 1. Rename Backend Endpoint

- [x] 1.1 Rename route from /pair-device to /pairing-requests in cmd/main.go
- [x] 1.2 Rename PairingHandler function to PairingRequestsHandler
- [x] 1.3 Update handler function references in route registration
- [x] 1.4 Update test file names and function names
- [x] 1.5 Update endpoint references in test assertions

## 2. State Management

- [x] 2.1 Add state variable for pairing submission status (pending, success, error)
- [x] 2.2 Add state variable to store pairing request ID
- [x] 2.3 Add state variable to store pairing error message

## 3. Error Handling and Request ID Extraction

- [x] 3.1 Update handshake error handler to check for NOT_PAIRED error code
- [x] 3.2 Extract requestId from error.details.requestId when NOT_PAIRED occurs
- [x] 3.3 Handle case when requestId is missing from error.details
- [x] 3.4 Store extracted requestId in state

## 4. Pairing Submission

- [x] 4.1 Create function to submit pairing request via POST to /pairing-requests
- [x] 4.2 Set Content-Type header to application/json in POST request
- [x] 4.3 Include requestId in JSON request body
- [x] 4.4 Call pairing submission function when requestId is available

## 5. Response Handling

- [x] 5.1 Handle successful pairing response (200 OK)
- [x] 5.2 Handle pairing submission failure responses (non-200 status)
- [x] 5.3 Handle network errors during pairing submission
- [x] 5.4 Update pairing status state based on response

## 6. UI Updates

- [x] 6.1 Update second progress step variant based on pairing status
- [x] 6.2 Show loading indicator during pairing submission
- [x] 6.3 Show success state when pairing completes
- [x] 6.4 Show error message when pairing fails
