## 1. Update Vite Configuration

- [x] 1.1 Read current ui/vite.config.js file
- [x] 1.2 Add `base: './'` option to the Vite config
- [x] 1.3 Ensure configuration is properly formatted and valid

## 2. Verify Development Server

- [x] 2.1 Start development server with `npm run dev` in ui/ directory
- [x] 2.2 Verify server starts without errors
- [x] 2.3 Test Hot Module Replacement by modifying a source file
- [x] 2.4 Confirm application loads correctly in browser

## 3. Build and Verify Production Bundle

- [x] 3.1 Run `npm run build` in ui/ directory to create production bundle
- [x] 3.2 Inspect build output in ui/dist/ directory
- [x] 3.3 Verify index.html uses relative paths for script tags (should be './assets/...' not '/assets/...')
- [x] 3.4 Verify CSS and JavaScript bundles use relative paths for asset references

## 4. Test Deployment Scenarios

- [x] 4.1 Test production build served from root path (/)
- [x] 4.2 Test production build served from a subpath (e.g., /pair-device/)
- [x] 4.3 Verify all assets (JS, CSS, images) load correctly in both scenarios
- [x] 4.4 Check browser console for any 404 errors or failed asset loads
