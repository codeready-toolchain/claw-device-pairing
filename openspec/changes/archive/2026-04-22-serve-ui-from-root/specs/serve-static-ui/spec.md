## ADDED Requirements

### Requirement: Serve UI index from root path
The system SHALL serve the UI's index.html file when a GET request is made to the root path `/`.

#### Scenario: Request to root path
- **WHEN** a GET request is sent to `/`
- **THEN** the server responds with status 200 and serves the index.html file

#### Scenario: Root path returns HTML content type
- **WHEN** a GET request is sent to `/`
- **THEN** the response Content-Type header is `text/html` or `text/html; charset=utf-8`

### Requirement: Serve static assets
The system SHALL serve static assets (JavaScript, CSS, images) from the UI build directory with correct MIME types.

#### Scenario: Serve JavaScript file
- **WHEN** a GET request is sent to a path ending in `.js` (e.g., `/assets/main.js`)
- **THEN** the server responds with status 200 and Content-Type `application/javascript` or `text/javascript`

#### Scenario: Serve CSS file
- **WHEN** a GET request is sent to a path ending in `.css` (e.g., `/assets/style.css`)
- **THEN** the server responds with status 200 and Content-Type `text/css`

#### Scenario: Serve image file
- **WHEN** a GET request is sent to an image path (e.g., `/assets/logo.png`)
- **THEN** the server responds with status 200 and appropriate image Content-Type (e.g., `image/png`)

### Requirement: SPA routing fallback
The system SHALL serve index.html for routes that do not match static files or API endpoints, to support client-side routing.

#### Scenario: Non-existent route fallback
- **WHEN** a GET request is sent to a path that doesn't match a static file or API route (e.g., `/pairing/device/123`)
- **THEN** the server responds with status 200 and serves the index.html file

#### Scenario: API routes not affected by fallback
- **WHEN** a request is sent to an API endpoint (e.g., POST `/pair-device`)
- **THEN** the API handler processes the request normally and does not serve index.html

### Requirement: Handle missing UI build directory
The system SHALL log a clear error and exit if the UI build directory does not exist at server startup.

#### Scenario: Missing UI directory
- **WHEN** the server starts and the UI build directory is not present
- **THEN** the server logs an error message indicating the missing directory path and exits with non-zero status
