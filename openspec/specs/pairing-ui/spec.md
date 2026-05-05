## Purpose

This specification defines the user interface for device pairing with the OpenClaw system. It establishes the structure, components, and visual presentation of the web-based pairing interface using React and Patternfly.
## Requirements
### Requirement: UI Directory Structure
The system SHALL create a `ui/` directory at the project root with a standard Vite + React project structure.

#### Scenario: Directory structure matches claw-signup
- **WHEN** the UI is initialized
- **THEN** the `ui/` directory contains `src/`, `public/`, `index.html`, `package.json`, `vite.config.js`, and `eslint.config.js`

#### Scenario: Source files are organized
- **WHEN** developers navigate the project
- **THEN** `src/` contains `main.jsx`, `App.jsx`, `index.css`, and a `components/` directory

### Requirement: Patternfly Card Layout
The application SHALL display a single Patternfly Card component as the main UI element.

#### Scenario: Card is rendered on page load
- **WHEN** the application loads
- **THEN** a Patternfly Card component is visible in the viewport

#### Scenario: Card has a title
- **WHEN** the card is displayed
- **THEN** the card shows a CardTitle with text "Device Pairing"

### Requirement: Pairing Status Display
The card body SHALL display a progress stepper showing the stages of device pairing.

#### Scenario: Stepper replaces simple message
- **WHEN** the application renders
- **THEN** the card body displays a Stepper component instead of just "Pairing device..." text

#### Scenario: Loading spinner is removed
- **WHEN** the card body is displayed
- **THEN** no Spinner component is shown (replaced by the Stepper)

### Requirement: Patternfly Styling
The application SHALL import and apply Patternfly base CSS styles.

#### Scenario: Patternfly CSS is imported
- **WHEN** the application initializes
- **THEN** `@patternfly/react-core/dist/styles/base.css` is imported in `main.jsx`

#### Scenario: Components use Patternfly styling
- **WHEN** Patternfly components are rendered
- **THEN** they display with default Patternfly visual styles

### Requirement: Development Dependencies
The project SHALL use the same core dependencies as claw-signup for consistency.

#### Scenario: React version matches
- **WHEN** package.json is examined
- **THEN** React version is `^19.2.4`

#### Scenario: Patternfly version matches
- **WHEN** package.json is examined
- **THEN** `@patternfly/react-core` version is `^6.4.1`

#### Scenario: Vite version matches
- **WHEN** package.json is examined
- **THEN** Vite version is `^8.0.4`

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

### Requirement: HTML Entry Point
The application SHALL use an HTML file as the entry point with proper meta tags.

#### Scenario: HTML file includes root element
- **WHEN** `index.html` is loaded
- **THEN** it contains a `<div id="root"></div>` element

#### Scenario: HTML file includes title
- **WHEN** `index.html` is loaded
- **THEN** the page title is "Claw Device Pairing"

#### Scenario: HTML file includes module script
- **WHEN** `index.html` is loaded
- **THEN** it includes `<script type="module" src="/src/main.jsx"></script>`

### Requirement: Dynamic Content Based on Handshake State
The card body SHALL display different content based on the current handshake state.

#### Scenario: Loading state content
- **WHEN** the handshake is in progress (loading state)
- **THEN** a loading message or spinner is displayed in the card body

#### Scenario: Success state content
- **WHEN** the handshake succeeds
- **THEN** the ProgressStepper is displayed with the first step marked as complete

#### Scenario: Error state content
- **WHEN** the handshake fails
- **THEN** an error message is displayed instead of the ProgressStepper

### Requirement: State-Driven Rendering
The application SHALL use React state to control UI rendering based on handshake lifecycle.

#### Scenario: UseEffect for handshake
- **WHEN** the App component mounts
- **THEN** a useEffect hook triggers the handshake

#### Scenario: State updates trigger re-renders
- **WHEN** handshake state changes
- **THEN** the UI re-renders to reflect the new state

