## Context

The project builds a Go server with an embedded React UI, packaged as a container image via a multi-stage Containerfile. The Makefile already provides `test`, `build-image`, and `push-image` targets using Podman. There are no existing GitHub Actions workflows. The default branch is `master`. The image is published to `quay.io/codeready-toolchain/claw-device-pairing`.

## Goals / Non-Goals

**Goals:**
- Automate test and build validation on every PR to `master`
- Automate image build and push to quay.io on every merge to `master`
- Tag pushed images with both the short commit SHA and `latest`

**Non-Goals:**
- Deploying to a cluster (out of scope for now)
- Building/testing on multiple architectures (single `linux/amd64` target)
- Caching Go modules or npm dependencies in CI (can be added later)
- Branch protection rules or required status checks configuration

## Decisions

### Use Red Hat GitHub Actions for container operations
- **`redhat-actions/buildah-build`** for building container images from the Containerfile. This replaces calling `make build-image` directly, providing better GitHub Actions integration (outputs, logging, caching).
- **`redhat-actions/push-to-registry`** for pushing images to quay.io. Handles multi-tag pushes natively.
- **`redhat-actions/podman-login`** for authenticating to quay.io. Wraps `podman login` with proper secret handling.

**Alternative considered:** Using raw `make build-image` / `make push-image` / `podman login` commands. Rejected because the Red Hat actions provide better integration with GitHub Actions (structured outputs, proper error handling, secret masking).

### Use `ubuntu-latest` runners
GitHub-hosted `ubuntu-latest` runners include Podman and Buildah by default. No special setup step is needed.

### Authenticate to quay.io using `redhat-actions/podman-login` with `QUAY_ROBOT` and `QUAY_TOKEN` secrets
The CD workflow will use the `podman-login` action with `QUAY_ROBOT` (robot account username) and `QUAY_TOKEN` (robot account token) repository secrets.

### CI jobs run in parallel (test and build-image are independent)
The `test` and `build-image` jobs have no data dependency on each other, so they run concurrently to minimize PR feedback time.

### CD workflow triggers on `push` to `master` (not `pull_request` closed/merged)
Using `push` to `master` is the idiomatic GitHub Actions pattern for post-merge workflows. It's simpler and more reliable than filtering `pull_request` events for `merged == true`.

### Image tags: short commit SHA + `latest`
The `buildah-build` action supports multiple tags natively. The CD workflow will build the image with both the short commit SHA and `latest` tags, then push both using `push-to-registry`.

## Risks / Trade-offs

- **[Buildah version drift]** → GitHub runner Buildah version may differ from local Podman builds. Mitigated by using standard OCI build commands and a standard Containerfile.
- **[Secret exposure]** → `QUAY_TOKEN` could leak in logs. Mitigated by GitHub's built-in secret masking and the `podman-login` action's secure handling.
- **[No caching]** → CI will do full `go mod download` and `npm ci` on every run, which is slower. Acceptable for now; caching can be added later as an optimization.
