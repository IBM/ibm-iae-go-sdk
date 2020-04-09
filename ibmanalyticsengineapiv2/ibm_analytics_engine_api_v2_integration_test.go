package ibmanalyticsengineapiv2_test

import (
    "github.com/joho/godotenv"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "os"

	. "github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv2"
)

/**
 * This class contains an integration test for the IBM Analytics Engine service.
 *
 * Notes:
 *
 * 1. This IBM Analytics Engine integration test shows how to automatically skip tests if the required config file
 *    is not available.
 *
 * 2. Before running this test:
 *    a. "cp example-service.env.hide ibmanalyticsengine-service.env"
 *    b. start up the IbmAnalyticsEngineApi service by following the instructions here:
 *    https://github.com/IBM/ibm-iae-go-sdk/blob/master/HOW_TO_TEST.md
 */
const externalConfigFile = "../ibmanalyticsengine-service.env"

var (
    service      *ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2
    err          error
    configLoaded bool = false
)

func shouldSkipTest() {
    if !configLoaded {
        Skip("External configuration is not available, skipping...")
    }
}

var _ = Describe(`IbmAnalyticsEngineApiV2`, func() {
    It("Successfully load the configuration", func() {
        err = godotenv.Load(externalConfigFile)
        if err == nil {
            //
            // Retrieve any test-specific properties from the environment.
            // ...

            // Set the flag to allow tests to execute.
            configLoaded = true
        }

        if !configLoaded {
            Skip("External configuration could not be loaded, skipping...")
        }
    })

    It("Successfully create the service client instance", func() {
        shouldSkipTest()

        service, err = ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(
            &ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{})
        Expect(err).To(BeNil())
        Expect(service).ToNot(BeNil())
        Expect(service.Service.Options.URL).To(Not(Equal("")))
    })

    Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {
        It(`Successfully list all resources`, func() {
            shouldSkipTest()

            getAnalyticsEngineByIdOptions := &ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions {
            	InstanceGuid: os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID"),
            }
            result, detailedResponse, err := service.GetAnalyticsEngineByID(&ibmanalyticsengineapiv2.getAnalyticsEngineByIdOptions{})
            Expect(err).To(BeNil())
            Expect(detailedResponse).ToNot(BeNil())
            Expect(detailedResponse.StatusCode).To(Equal(200))
            Expect(result).ToNot(BeNil())

            // resources := result.Resources
        })
    })

})