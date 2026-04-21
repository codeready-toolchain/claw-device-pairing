## 1. Project Structure Setup

- [x] 1.1 Create `ui/` directory at project root
- [x] 1.2 Create `ui/src/` directory for source files
- [x] 1.3 Create `ui/public/` directory for static assets
- [x] 1.4 Create `ui/src/components/` directory for React components

## 2. Configuration Files

- [x] 2.1 Create `ui/package.json` with React, Vite, and Patternfly dependencies matching claw-signup versions
- [x] 2.2 Create `ui/vite.config.js` with @vitejs/plugin-react configuration
- [x] 2.3 Create `ui/eslint.config.js` with React linting rules
- [x] 2.4 Create `ui/.gitignore` to exclude node_modules and dist directories

## 3. HTML Entry Point

- [x] 3.1 Create `ui/index.html` with viewport meta tag and charset
- [x] 3.2 Add `<div id="root"></div>` to index.html
- [x] 3.3 Add `<script type="module" src="/src/main.jsx"></script>` to index.html
- [x] 3.4 Set page title to "Claw Device Pairing"

## 4. React Application Bootstrap

- [x] 4.1 Create `ui/src/main.jsx` entry point file
- [x] 4.2 Import `@patternfly/react-core/dist/styles/base.css` in main.jsx
- [x] 4.3 Import custom `index.css` in main.jsx
- [x] 4.4 Set up React root rendering with StrictMode in main.jsx

## 5. App Component

- [x] 5.1 Create `ui/src/App.jsx` as root component
- [x] 5.2 Import Card, CardTitle, CardBody, and Spinner from Patternfly
- [x] 5.3 Add Card component with CardTitle "Device Pairing"
- [x] 5.4 Add CardBody with Spinner (size "md") and "Pairing device..." text
- [x] 5.5 Style the card body to align spinner and text horizontally

## 6. Custom Styles

- [x] 6.1 Create `ui/src/index.css` for custom application styles
- [x] 6.2 Add centering styles for the card layout
- [x] 6.3 Add styles for spinner and text alignment

## 7. Static Assets

- [x] 7.1 Create or copy favicon to `ui/public/favicon.svg`
- [x] 7.2 Update `ui/index.html` to reference favicon

## 8. Development Environment Setup

- [x] 8.1 Run `npm install` in `ui/` directory to install dependencies
- [x] 8.2 Add `dev`, `build`, `lint`, and `preview` scripts to package.json
- [x] 8.3 Test that `npm run dev` starts development server successfully
- [x] 8.4 Verify hot module replacement works during development

## 9. Verification

- [x] 9.1 Test that opening http://localhost:5173 shows the card with pairing message
- [x] 9.2 Verify Patternfly styling is applied correctly to Card components
- [x] 9.3 Verify Spinner is visible and animated
- [x] 9.4 Run `npm run build` and verify production build succeeds
- [x] 9.5 Run `npm run preview` and verify built application works
