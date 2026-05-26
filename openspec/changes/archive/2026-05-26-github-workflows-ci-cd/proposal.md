## Why

The project has no automated CI/CD pipeline. Tests and image builds are run manually, which is error-prone and slows down the development loop. Adding GitHub Actions workflows ensures every PR is validated and every merge to `master` produces a pushed container image.

## What Changes

- Add a CI workflow (`.github/workflows/ci.yaml`) that runs on PRs targeting `master`:
  - Job 1: `make test` (Go tests with race detection and coverage)
  - Job 2: `make build-image` (container image build via Podman, validates the Containerfile)
- Add a CD workflow (`.github/workflows/cd.yaml`) that runs when a PR is merged to `master`:
  - Builds the container image tagged with both the short commit SHA and `latest`
  - Pushes the image to `quay.io/codeready-toolchain/claw-device-pairing` using a `QUAY_TOKEN` secret

## Capabilities

### New Capabilities
- `ci-workflow`: GitHub Actions workflow for continuous integration (test + build validation on PRs)
- `cd-workflow`: GitHub Actions workflow for continuous deployment (build + push image on merge)

### Modified Capabilities

None.

## Impact

- New files: `.github/workflows/ci.yaml`, `.github/workflows/cd.yaml`
- Requires GitHub repository secret: `QUAY_TOKEN` (Quay.io robot account token)
- Uses Red Hat GitHub Actions: `redhat-actions/buildah-build`, `redhat-actions/push-to-registry`, `redhat-actions/podman-login`
- No changes to existing application code or Makefile
