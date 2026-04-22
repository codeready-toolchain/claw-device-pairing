## 1. Setup and Configuration

- [x] 1.1 Verify UI build directory structure and identify correct path (ui/dist or ui/build)
- [x] 1.2 Add startup validation to check UI build directory exists before starting server

## 2. Static File Serving Implementation

- [x] 2.1 Add Echo Static middleware to serve files from UI build directory at root path
- [x] 2.2 Configure route registration order to register API routes after static file serving
- [x] 2.3 Add SPA fallback route to serve index.html for non-matching paths

## 3. Testing and Validation

- [x] 3.1 Test GET request to `/` returns index.html with correct Content-Type
- [x] 3.2 Test static assets (.js, .css, images) are served with correct MIME types
- [x] 3.3 Test SPA routing fallback - non-existent paths return index.html
- [x] 3.4 Test API endpoints still work correctly and are not affected by static serving
- [x] 3.5 Test server startup error when UI build directory is missing
- [x] 3.6 Update or remove CORS middleware configuration for production vs development
