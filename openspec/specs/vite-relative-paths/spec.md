# vite-relative-paths Specification

## Purpose
TBD - created by archiving change use-relative-asset-paths. Update Purpose after archive.
## Requirements
### Requirement: Vite Base Configuration
The Vite build configuration SHALL use relative paths for all asset references.

#### Scenario: Base path is relative
- **WHEN** vite.config.js is read
- **THEN** the base option is set to './'

#### Scenario: Production build uses relative paths
- **WHEN** npm run build is executed
- **THEN** all asset references in the output use relative paths (e.g., './assets/...' instead of '/assets/...')

### Requirement: Asset Reference Resolution
All asset types SHALL be referenced using relative paths in the production build.

#### Scenario: JavaScript bundles use relative paths
- **WHEN** the production build generates JavaScript files
- **THEN** script tags and imports use relative paths

#### Scenario: CSS files use relative paths
- **WHEN** the production build generates CSS files
- **THEN** stylesheet links use relative paths

#### Scenario: Static assets use relative paths
- **WHEN** the production build references images, fonts, or other static assets
- **THEN** all references use relative paths from the HTML entry point

### Requirement: Development Server Compatibility
The development server SHALL continue to work correctly with relative path configuration.

#### Scenario: Dev server starts successfully
- **WHEN** npm run dev is executed
- **THEN** the Vite development server starts without errors

#### Scenario: HMR works with relative paths
- **WHEN** a file is modified during development
- **THEN** Hot Module Replacement updates the browser correctly

### Requirement: Deployment Path Independence
The production build SHALL work correctly when deployed at any base path.

#### Scenario: Works at root path
- **WHEN** the build output is served from /
- **THEN** all assets load correctly

#### Scenario: Works at subpath
- **WHEN** the build output is served from a subpath (e.g., /pair-device/, /ui/)
- **THEN** all assets load correctly using relative path resolution

