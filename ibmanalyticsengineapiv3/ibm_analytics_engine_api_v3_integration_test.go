// +build integration

package ibmanalyticsengineapiv3_test

import (
    "github.com/joho/godotenv"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "os"
   	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv3"
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
    service      *ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3
    err          error
    configLoaded bool = false
    instanceGuid string
)

func shouldSkipTest() {
    if !configLoaded {
        Skip("External configuration is not available, skipping...")
    }
}


var _ = Describe(`IbmAnalyticsEngineApiV3`, func() {
    It("Successfully load the configuration", func() {
        err = godotenv.Load(externalConfigFile)
        if err == nil {
            //
            // Retrieve any test-specific properties from the environment.
            instanceGuid = os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID")

            // Set the flag to allow tests to execute.
            configLoaded = true
        }

        if !configLoaded {
            Skip("External configuration could not be loaded, skipping...")
        }
    })

    It("Successfully create the service client instance", func() {
        shouldSkipTest()

        service, err = ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(
            &ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{})
        Expect(err).To(BeNil())
        Expect(service).ToNot(BeNil())
        Expect(service.Service.Options.URL).To(Not(Equal("")))
    })

   //  Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {
   //      It(`Invoke GetAnalyticsEngineByID successfully`, func() {
   //          shouldSkipTest()

			// // Construct an instance of the GetAnalyticsEngineByIdOptions model
			// getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
			// getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

   //          result, response, err := service.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
   //          Expect(err).To(BeNil())
   //          Expect(response).ToNot(BeNil())
   //          Expect(response.StatusCode).To(Equal(200))
   //          Expect(result).ToNot(BeNil())

   //          // resources := result.Resources
   //      })
   //  })

	Describe(`GetInstanceByID(getInstanceByIdOptions *GetInstanceByIdOptions)`, func() {
		It(`Invoke GetInstanceByID successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the GetInstanceByIdOptions model
			getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
			getInstanceByIdOptionsModel.InstanceID = core.StringPtr(instanceGuid)
			getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

			// Expect response parsing to fail since we are receiving a text/plain response
			result, response, err := service.GetInstanceByID(getInstanceByIdOptionsModel)
            Expect(err).To(BeNil())
            Expect(response).ToNot(BeNil())
            Expect(response.StatusCode).To(Equal(200))
            Expect(result).ToNot(BeNil())

		})
	})

})