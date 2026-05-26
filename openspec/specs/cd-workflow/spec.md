## ADDED Requirements

### Requirement: CD workflow triggers on push to master
The CD workflow SHALL be triggered when commits are pushed to the `master` branch (i.e., when a PR is merged).

#### Scenario: PR merged to master
- **WHEN** a pull request is merged into the `master` branch
- **THEN** the CD workflow SHALL start

#### Scenario: Push to non-master branch
- **WHEN** commits are pushed to a branch other than `master`
- **THEN** the CD workflow SHALL NOT run

### Requirement: CD authenticates to quay.io
The CD workflow SHALL authenticate to quay.io using the `redhat-actions/podman-login` action with `QUAY_ROBOT` and `QUAY_TOKEN` repository secrets.

#### Scenario: Successful authentication
- **WHEN** the `QUAY_ROBOT` and `QUAY_TOKEN` secrets are valid
- **THEN** the workflow SHALL authenticate to `quay.io` before building or pushing

#### Scenario: Authentication failure
- **WHEN** the `QUAY_TOKEN` secret is invalid or missing
- **THEN** the login step SHALL fail and the workflow SHALL report failure

### Requirement: CD builds container image
The CD workflow SHALL build the container image using the `redhat-actions/buildah-build` action with the project's Containerfile.

#### Scenario: Image builds successfully
- **WHEN** the build step runs
- **THEN** the container image SHALL be built with the current commit's SHA and build time passed as build args

### Requirement: CD tags image with commit SHA and latest
The CD workflow SHALL tag the built image with both the short git commit SHA and `latest` using the `redhat-actions/buildah-build` action's `tags` input.

#### Scenario: Image is tagged correctly
- **WHEN** the image is built from commit `abc1234`
- **THEN** the image SHALL be available as `quay.io/codeready-toolchain/claw-device-pairing:abc1234` and `quay.io/codeready-toolchain/claw-device-pairing:latest`

### Requirement: CD pushes image to quay.io
The CD workflow SHALL push the built image to the quay.io registry using the `redhat-actions/push-to-registry` action.

#### Scenario: Successful push
- **WHEN** the image is built and tagged
- **THEN** the workflow SHALL push both image tags to `quay.io/codeready-toolchain/claw-device-pairing`

### Requirement: CD build and push run sequentially
The image MUST be built before it can be pushed. The push step SHALL depend on the successful completion of the build step.

#### Scenario: Build fails
- **WHEN** the build step fails
- **THEN** the push step SHALL NOT execute and the workflow SHALL report failure
