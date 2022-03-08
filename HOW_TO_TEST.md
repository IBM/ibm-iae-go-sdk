# How to run tests

## Pre-requisites
1. Ensure that your JAVA_HOME is set correctly
    `export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home` 
2. Set GOPATH correctly
    `export GOPATH=$(go env GOPATH)`
3. Download [SDK generation code](https://github.ibm.com/CloudEngineering/openapi-sdkgen/releases) 
4. Clone [`GO SDK`] https://github.com/IBM/ibm-iae-go-sdk/
5. Clone [`Cloud API Doc`] https://github.ibm.com/cloud-api-docs/analytics-engine
6. Generate the sdk code with: 
`openapi-sdkgen/openapi-sdkgen.sh  generate -g ibm-go -i ibmanalyticsengine-v3.json  -o /Users/git/ibm-iae-go-sdk  --genITs --genExamples --api-package github.com/IBM/ibm-iae-go-sdk`


## Integration Tests

1. The SDK has been generated from the latest ibmanalyticsengine-v3.json which contains modifications and newly added api methods.
The Integration tests have to be run on the generated SDK not on the sdk from git. 
2. To merge with remote version pick the code marked between and update `ibm_analytics_engine_api_v3_integration_test.go`
    `// !!! Start of custom content to be copied !!!`
    `// !!! End of custom content to be copied !!!`
3. Create a `ibm_analytics_engine_api_v3.env` file in the `ibm-iae-go-sdk` directory using `ibmanalyticsengine-service.env.hide` as an example.
4. Update `ibm_analytics_engine_api_v3.env` file with your own **APIKEY** and other details.
5. export instance guid as `IBM_ANALYTICS_ENGINE_INSTANCE_GUID` environment variable.
6. export instance guid to cretae instance home as `IBM_ANALYTICS_ENGINE_INSTANCE_GUID_INSTANCE_HOME` environment variable.
7. export HMAC access key as `IBM_ANALYTICS_ENGINE_HMAC_ACCESS_KEY` environment variable.
8. export HMAC secret key as `IBM_ANALYTICS_ENGINE_HMAC_SECRET_KEY` environment variable.
9. In root directory `ibm-iae-go-sdk/ibmanalyticsengineapiv3`. Run the following:
    ```
    go test ./... -tags=integration -v
    ```

## Unit Tests

1. In root directory `ibm-iae-go-sdk/`. Run the following:
    ```
    go test ./...
    ```
