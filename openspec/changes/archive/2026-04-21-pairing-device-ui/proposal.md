## Why

The claw-device-pairing project needs a web-based user interface to provide visual feedback during the device pairing process. Users need to see the pairing status in a clean, professional interface that matches the existing claw-signup project's design.

## What Changes

- Create a new `ui` directory with a Patternfly-based React application
- Implement a single-card layout displaying pairing status
- Use the same project structure and dependencies as the claw-signup project for consistency
- Set up Vite build tooling, ESLint, and standard React development tools

## Capabilities

### New Capabilities
- `pairing-ui`: Web interface for displaying device pairing status with a single card showing "pairing device..." text

### Modified Capabilities
<!-- No existing capabilities are being modified -->

## Impact

- New `ui/` directory at project root containing the complete React application
- New dependencies: React, Vite, Patternfly React components
- Development workflow additions: `npm install`, `npm run dev`, `npm run build` in the `ui` directory
- No impact on existing Go backend code or APIs
