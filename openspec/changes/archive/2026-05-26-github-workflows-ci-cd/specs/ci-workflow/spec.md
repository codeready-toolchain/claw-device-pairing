## ADDED Requirements

### Requirement: CI workflow triggers on pull requests to master
The CI workflow SHALL be triggered when a pull request is opened, synchronized, or reopened targeting the `master` branch.

#### Scenario: PR opened against master
- **WHEN** a pull request is opened targeting the `master` branch
- **THEN** the CI workflow SHALL start both the test and build-image jobs

#### Scenario: PR targets a non-master branch
- **WHEN** a pull request is opened targeting a branch other than `master`
- **THEN** the CI workflow SHALL NOT run

### Requirement: CI runs Go tests
The CI workflow SHALL include a job that runs `make test` to execute all Go tests with race detection and coverage.

#### Scenario: All tests pass
- **WHEN** the test job runs and all Go tests pass
- **THEN** the job SHALL report success

#### Scenario: A test fails
- **WHEN** the test job runs and one or more Go tests fail
- **THEN** the job SHALL report failure and the PR status check SHALL be marked as failed

### Requirement: CI validates container image build
The CI workflow SHALL include a job that uses the `redhat-actions/buildah-build` action to validate that the Containerfile builds successfully.

#### Scenario: Image builds successfully
- **WHEN** the build-image job runs and the container image builds without errors
- **THEN** the job SHALL report success

#### Scenario: Image build fails
- **WHEN** the build-image job runs and the container image build fails
- **THEN** the job SHALL report failure and the PR status check SHALL be marked as failed

### Requirement: CI test and build-image jobs run in parallel
The test and build-image jobs SHALL have no dependency on each other and SHALL run concurrently.

#### Scenario: Both jobs start simultaneously
- **WHEN** the CI workflow is triggered
- **THEN** the test and build-image jobs SHALL start without waiting for each other
