## MODIFIED Requirements

### Requirement: Build Configuration
The project SHALL include Vite build configuration for development and production with support for flexible deployment paths.

#### Scenario: Development server can start
- **WHEN** `npm run dev` is executed
- **THEN** Vite starts a development server with hot module replacement

#### Scenario: Production build succeeds
- **WHEN** `npm run build` is executed
- **THEN** Vite creates an optimized production bundle in `dist/`

#### Scenario: Production build uses relative asset paths
- **WHEN** `npm run build` is executed
- **THEN** all asset references in the build output use relative paths

#### Scenario: Build supports deployment at any base path
- **WHEN** the production build is deployed at a subpath
- **THEN** the application loads and runs correctly without rebuilding
