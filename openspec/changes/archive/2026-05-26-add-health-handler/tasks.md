## 1. Health Handler Implementation

- [x] 1.1 Create `internal/handlers/health.go` with a `HandleHealth` function that returns a JSON response containing `status`, `commit_hash`, and `build_time` from `internal/version`
- [x] 1.2 Add a response model struct in `internal/models/` for the health response (with `json` tags using `snake_case`)

## 2. Route Wiring

- [x] 2.1 Replace the inline `/health` handler in `cmd/main.go` with a call to `handlers.HandleHealth`

## 3. Tests

- [x] 3.1 Add unit tests in `internal/handlers/health_test.go` verifying the response status code and JSON body fields
