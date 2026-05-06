## 1. Backend Status Endpoint

- [x] 1.1 Add status response model to internal/models/pairing.go
- [x] 1.2 Add GET handler method to PairingRequestsHandler in internal/handlers/pairing.go
- [x] 1.3 Implement status check logic (return 202 for pending, 200 for approved)
- [x] 1.4 Handle request ID extraction from URL path parameter
- [x] 1.5 Handle not found case for invalid request IDs
- [x] 1.6 Register GET route in cmd/main.go for /pairing-requests/:id
- [x] 1.7 Add tests for GET endpoint in internal/handlers/pairing_requests_test.go

## 2. UI State Management

- [x] 2.1 Add state variable for approval status (pending/approved/timeout)
- [x] 2.2 Add state variable for polling active flag
- [x] 2.3 Add state variable for navigation button enabled state

## 3. Status Polling Implementation

- [x] 3.1 Create polling function that makes GET request to pairing-requests/:id
- [x] 3.2 Implement 1-second interval polling with setInterval
- [x] 3.3 Handle 202 response (continue polling)
- [x] 3.4 Handle 200 response (stop polling, mark as approved)
- [x] 3.5 Handle error responses (stop polling, show error)
- [x] 3.6 Implement 30-second timeout with automatic stop
- [x] 3.7 Add useEffect hook to start polling when pairing request ID is available
- [x] 3.8 Add cleanup function to clear interval on unmount

## 4. Navigation Button

- [x] 4.1 Import Button component from @patternfly/react-core
- [x] 4.2 Add Button component to JSX below ProgressStepper
- [x] 4.3 Set button label to "Go to OpenClaw"
- [x] 4.4 Bind button disabled state to approval status
- [x] 4.5 Implement onClick handler for navigation

## 5. Navigation Logic

- [x] 5.1 Create function to extract token from URL fragment
- [x] 5.2 Create function to construct root URL with token
- [x] 5.3 Remove /integration/device-pairing path from URL
- [x] 5.4 Preserve #token=... fragment in new URL
- [x] 5.5 Use window.location.href for navigation
- [x] 5.6 Handle case when token is missing (log warning, navigate without fragment)

## 6. Integration and Error Handling

- [x] 6.1 Trigger polling after pairing submission completes
- [x] 6.2 Update pairing status display during polling
- [x] 6.3 Display timeout message when 30s elapses
- [x] 6.4 Handle network errors during polling
- [x] 6.5 Test full flow: submit → poll → approve → navigate
