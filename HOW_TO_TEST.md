# How to run tests

## Pre-requisites

## Integration Tests

1. Create a `ibmanalyticsengine-service.env` file in the `ibm-iae-go-sdk` directory using `ibmanalyticsengine-service.env.hide` as an example.
1. Update `ibmanalyticsengine-service.env` file with your own **APIKEY**.
1. export your instance guid as `IBM_ANALYTICS_ENGINE_INSTANCE_GUID` environment variable.
1. In root directory `ibm-iae-go-sdk/`. Run the following:
    ```
    go test ./... -tags=integration
    ```

## Unit Tests

1. In root directory `ibm-iae-go-sdk/`. Run the following:
    ```
    go test ./...
    ```
