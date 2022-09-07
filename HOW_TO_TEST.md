# How to run tests

## Pre-requisites

## Integration Tests

1. Create a `ibmanalyticsengine-service.env` file in the `ibm-iae-go-sdk` directory using `ibmanalyticsengine-service.env.hide` as an example.
2. Update `ibmanalyticsengine-service.env` file with your own **APIKEY**.
3. export your instance guid as `IBM_ANALYTICS_ENGINE_INSTANCE_GUID` environment variable.
4. export your instance guid  for updating instance home as `IBM_ANALYTICS_ENGINE_INSTANCE_HOME_GUID` environment variable.
5. export your HMAC Access Key as `IBM_ANALYTICS_ENGINE_HMAC_ID` environment variable.
6. export your HMAC Secret Key as `IBM_ANALYTICS_ENGINE_HMAC_KEY` environment variable.
1. In root directory `ibm-iae-go-sdk/`. Run the following:
    ```
    go test ./... -tags=integration -v
    ```

## Unit Tests

1. In root directory `ibm-iae-go-sdk/`. Run the following:
    ```
    go test ./...
    ```
Docs update