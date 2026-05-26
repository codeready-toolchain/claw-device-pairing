### Requirement: Health endpoint returns version information
The `/health` endpoint SHALL return a `200 OK` response with a JSON body containing the application's `commit_hash` and `build_time` fields sourced from `internal/version/version.go`, along with a `status` field set to `"ok"`.

#### Scenario: Successful health check
- **WHEN** a GET request is made to `/health`
- **THEN** the response status code SHALL be `200 OK` and the JSON body SHALL contain `status` set to `"ok"`, `commit_hash` set to the value of `version.CommitHash`, and `build_time` set to the value of `version.BuildTime`

#### Scenario: Health check with default build values
- **WHEN** a GET request is made to `/health` and the application was built without ldflags
- **THEN** the response SHALL contain `commit_hash` set to `"unknown"` and `build_time` set to `"unknown"`

### Requirement: Health handler is a standalone function
The health handler SHALL be implemented as a standalone function in `internal/handlers/health.go` rather than an inline closure, following the project's handler organization pattern.

#### Scenario: Handler registered on Echo router
- **WHEN** the server starts
- **THEN** the `/health` GET route SHALL be served by the handler function defined in `internal/handlers/health.go`
