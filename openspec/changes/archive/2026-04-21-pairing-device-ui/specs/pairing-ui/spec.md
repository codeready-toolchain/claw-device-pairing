## ADDED Requirements

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
The card body SHALL display a "Pairing device..." message with a loading indicator.

#### Scenario: Pairing message is shown
- **WHEN** the application renders
- **THEN** the card body displays the text "Pairing device..."

#### Scenario: Loading spinner is visible
- **WHEN** the pairing message is displayed
- **THEN** a Patternfly Spinner component with size "md" is shown

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
The project SHALL include Vite build configuration for development and production.

#### Scenario: Development server can start
- **WHEN** `npm run dev` is executed
- **THEN** Vite starts a development server with hot module replacement

#### Scenario: Production build succeeds
- **WHEN** `npm run build` is executed
- **THEN** Vite creates an optimized production bundle in `dist/`

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
