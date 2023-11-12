# Contribution Guide

## Test

The project contains a suite of both unit tests and integration tests for ensuring the API endpoints function as expected.

### Unit Tests

To run the unit test suit, run:

```bash
docker compose --profile unit-test run unit-test
```

This will run the unit test suite of the project in a Go container.

### Integration Tests

The integration tests for the project are run through Docker Compose and involve all the application components
run as separate containers. To run the integration test suite, run the command:

```bash
docker compose --profile integration-test run integration-test
```

This will start the integration test container after bringing up its dependencies, which will run integration tests against the available application endpoints.
