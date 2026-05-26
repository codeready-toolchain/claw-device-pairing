## 1. CI Workflow

- [x] 1.1 Create `.github/workflows/ci.yaml` with `pull_request` trigger on `master` branch
- [x] 1.2 Add `test` job: checkout, setup Go 1.25, run `make test`
- [x] 1.3 Add `lint-and-build-ui` job: checkout, setup Node 20, run `npm ci`, `npm run lint`, `npm run build` in `ui/`
- [x] 1.4 Add `build-image` job: checkout, use `redhat-actions/buildah-build` to build the Containerfile (build validation only, no push)

## 2. CD Workflow

- [x] 2.1 Create `.github/workflows/cd.yaml` with `push` trigger on `master` branch
- [x] 2.2 Add login step: use `redhat-actions/podman-login` to authenticate to quay.io with `QUAY_ROBOT` and `QUAY_TOKEN` secrets
- [x] 2.3 Add build step: use `redhat-actions/buildah-build` to build the image with tags `<short-sha>` and `latest`, passing `COMMIT_HASH` and `BUILD_TIME` as build args
- [x] 2.4 Add push step: use `redhat-actions/push-to-registry` to push both tags to `quay.io/codeready-toolchain/claw-device-pairing`
