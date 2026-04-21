## Context

The claw-device-pairing project currently lacks a user interface. The claw-signup project has established a UI pattern using React, Vite, and Patternfly components with a card-based layout. To maintain consistency across the Claw ecosystem, this pairing UI should mirror that established structure.

Current state:
- No existing UI directory in claw-device-pairing
- claw-signup reference UI available at `../claw-signup/ui` with proven Vite + React + Patternfly stack
- Project structure: main.jsx → App.jsx → component hierarchy
- Patternfly CSS imported in main.jsx for global styling

## Goals / Non-Goals

**Goals:**
- Create a minimal pairing status UI matching claw-signup's layout structure
- Display a single card with "pairing device..." message
- Use identical technology stack (Vite, React 19, Patternfly) for consistency
- Enable future expansion to full pairing flow UI

**Non-Goals:**
- Full pairing flow implementation (this change focuses on basic UI structure only)
- WebSocket integration or real-time status updates
- Multi-step pairing wizard or complex state management
- Backend API integration

## Decisions

### 1. Technology Stack
**Decision**: Use Vite + React 19 + Patternfly React Components v6
**Rationale**: Matches claw-signup exactly, proven stack, modern tooling
**Alternatives considered**:
- Next.js: Overkill for simple SPA, adds SSR complexity
- Plain HTML/CSS: No component reuse, harder to maintain consistency

### 2. Project Structure
**Decision**: Mirror claw-signup's `ui/` directory structure
```
ui/
├── src/
│   ├── main.jsx          (entry point, imports Patternfly CSS)
│   ├── App.jsx           (root component)
│   ├── components/       (UI components)
│   └── index.css         (custom styles)
├── public/               (static assets)
├── index.html            (HTML template)
├── package.json
├── vite.config.js
└── eslint.config.js
```
**Rationale**: Developers familiar with claw-signup can navigate immediately, tooling configs identical
**Alternatives considered**:
- Monorepo with shared components: Premature for two projects
- Embedded in Go binary: Complicates development workflow

### 3. Component Design
**Decision**: Single Card component in `App.jsx` with minimal logic
**Rationale**: Simplest possible implementation, easy to extend later
**Structure**:
```jsx
<Card>
  <CardTitle>Device Pairing</CardTitle>
  <CardBody>
    <Spinner size="md" />
    <span>Pairing device...</span>
  </CardBody>
</Card>
```

### 4. Dependencies
**Decision**: Copy exact versions from claw-signup package.json
**Rationale**: Avoid version drift, proven compatibility
**Key dependencies**:
- `react@^19.2.4`, `react-dom@^19.2.4`
- `@patternfly/react-core@^6.4.1`
- `@patternfly/react-icons@^6.4.0`
- `vite@^8.0.4`

## Risks / Trade-offs

**[Risk]** Dependency version drift between claw-signup and claw-device-pairing
→ **Mitigation**: Document source project in package.json comments, periodic sync checks

**[Risk]** Duplicated configuration files (eslint, vite) in both projects
→ **Mitigation**: Accept duplication for now, extract to shared config package if 3+ projects

**[Trade-off]** Simple card vs. full pairing flow
→ **Decision**: Start simple, iterate. Card structure supports future expansion (add form, progress stepper, etc.)

**[Trade-off]** Separate UI directory vs. embedded Go templates
→ **Decision**: Separate directory for better developer experience (hot reload, modern tooling), static build served by Go in production
