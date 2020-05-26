// +build integration

package ibmanalyticsengineapiv2_test

import (
    "github.com/joho/godotenv"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "os"
   	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv2"
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
    instanceGuid string
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

        service, err = ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(
            &ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{})
        Expect(err).To(BeNil())
        Expect(service).ToNot(BeNil())
        Expect(service.Service.Options.URL).To(Not(Equal("")))
    })

    Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {
        It(`Invoke GetAnalyticsEngineByID successfully`, func() {
            shouldSkipTest()

			// Construct an instance of the GetAnalyticsEngineByIdOptions model
			getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
			getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

            result, response, err := service.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
            Expect(err).To(BeNil())
            Expect(response).ToNot(BeNil())
            Expect(response.StatusCode).To(Equal(200))
            Expect(result).ToNot(BeNil())

            // resources := result.Resources
        })
    })

    Describe(`GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions)`, func() {
        It(`Invoke GetAnalyticsEngineStateByID successfully`, func() {
            shouldSkipTest()

			getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
			getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			result, response, err := service.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
            Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

            // resources := result.Resources
        })
    })

    Describe(`CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions)`, func() {
    	It(`Invoke CreateCustomizationRequest successfully`, func() {
    		shouldSkipTest()

			// Construct an instance of the AnalyticsEngineCustomActionScript model
			analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
			analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
			analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
			analyticsEngineCustomActionScriptModel.SourceProps = CreateMockMap()

			// Construct an instance of the AnalyticsEngineCustomAction model
			analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
			analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
			analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
			analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
			analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}

			// Construct an instance of the CreateCustomizationRequestOptions model
			createCustomizationRequestOptionsModel := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
			createCustomizationRequestOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)
			createCustomizationRequestOptionsModel.Target = core.StringPtr("all")
			createCustomizationRequestOptionsModel.CustomActions = []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}

			// Invoke operation with valid options model (positive test)
			result, response, err := service.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions)`, func() {
		It(`Invoke GetAllCustomizationRequests successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the GetAllCustomizationRequestsOptions model
			getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
			getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			// Invoke operation with valid options model (positive test)
			result, response, err := service.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})
	Describe(`GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions)`, func() {
		It(`Invoke GetCustomizationRequestByID successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the GetAllCustomizationRequestsOptions model
			getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
			getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			// Invoke operation with valid options model (positive test)
			result_id, response_id, err_id := service.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
			Expect(err_id).To(BeNil())
			Expect(response_id).ToNot(BeNil())

			// Construct an instance of the GetCustomizationRequestByIdOptions model
			getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
			getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)
			getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr((*(*result_id)[0].ID))

			// Invoke operation with valid options model (positive test)
			result, response, err := service.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ResizeCluster(resizeClusterOptions *ResizeClusterOptions)`, func() {
		It(`Invoke ResizeCluster successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the ResizeClusterOptions model
			resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
			resizeClusterOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)
			resizeClusterOptionsModel.ComputeNodesCount = core.Int64Ptr(int64(2))

			// Invoke operation with valid options model (positive test)
			result, response, err := service.ResizeCluster(resizeClusterOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions)`, func() {
		It(`Invoke ResetClusterPassword successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the ResetClusterPasswordOptions model
			resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
			resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			// Invoke operation with valid options model (positive test)
			result, response, err := service.ResetClusterPassword(resetClusterPasswordOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions)`, func() {
		It(`Invoke ConfigureLogging successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the AnalyticsEngineLoggingNodeSpec model
			analyticsEngineLoggingNodeSpecModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec)
			analyticsEngineLoggingNodeSpecModel.NodeType = core.StringPtr("management")
			analyticsEngineLoggingNodeSpecModel.Components = []string{"ambari-server"}

			// Construct an instance of the AnalyticsEngineLoggingServer model
			analyticsEngineLoggingServerModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer)
			analyticsEngineLoggingServerModel.Type = core.StringPtr("logdna")
			analyticsEngineLoggingServerModel.Credential = core.StringPtr("testString")
			analyticsEngineLoggingServerModel.ApiHost = core.StringPtr("testString")
			analyticsEngineLoggingServerModel.LogHost = core.StringPtr("testString")
			analyticsEngineLoggingServerModel.Owner = core.StringPtr("testString")

			// Construct an instance of the ConfigureLoggingOptions model
			configureLoggingOptionsModel := new(ibmanalyticsengineapiv2.ConfigureLoggingOptions)
			configureLoggingOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)
			configureLoggingOptionsModel.LogSpecs = []ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel}
			configureLoggingOptionsModel.LogServer = analyticsEngineLoggingServerModel

			// Invoke operation with valid options model (positive test)
			response, err := service.ConfigureLogging(configureLoggingOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions)`, func() {
		It(`Invoke GetLoggingConfig successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the GetLoggingConfigOptions model
			getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
			getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			// Invoke operation with valid options model (positive test)
			result, response, err := service.GetLoggingConfig(getLoggingConfigOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions)`, func() {
		It(`Invoke DeleteLoggingConfig successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the DeleteLoggingConfigOptions model
			deleteLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.DeleteLoggingConfigOptions)
			deleteLoggingConfigOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)

			// Invoke operation with valid options model (positive test)
			response, err := service.DeleteLoggingConfig(deleteLoggingConfigOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions)`, func() {
		It(`Invoke UpdatePrivateEndpointWhitelist successfully`, func() {
			shouldSkipTest()

			// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
			updatePrivateEndpointWhitelistOptionsModel := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
			updatePrivateEndpointWhitelistOptionsModel.InstanceGuid = core.StringPtr(instanceGuid)
			updatePrivateEndpointWhitelistOptionsModel.IpRanges = []string{"testString"}
			updatePrivateEndpointWhitelistOptionsModel.Action = core.StringPtr("add")

			// Invoke operation with valid options model (positive test)
			result, response, err := service.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})
})