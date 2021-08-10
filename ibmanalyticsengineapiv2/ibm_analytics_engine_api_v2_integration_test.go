// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibmanalyticsengineapiv2_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmanalyticsengineapiv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmAnalyticsEngineApiV2 Integration Tests`, func() {

	const externalConfigFile = "../ibmanalyticsengine-service.env"

	var (
		err          error
		ibmAnalyticsEngineApiService *ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2
		serviceURL   string
		config       map[string]string
		instanceGuid string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmanalyticsengineapiv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			instanceGuid = os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID")
			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			ibmAnalyticsEngineApiServiceOptions := &ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{}

			ibmAnalyticsEngineApiService, err = ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(ibmAnalyticsEngineApiServiceOptions)

			Expect(err).To(BeNil())
			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
			Expect(ibmAnalyticsEngineApiService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetAllAnalyticsEngines - List all Analytics Engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions)`, func() {

			getAllAnalyticsEnginesOptions := &ibmanalyticsengineapiv2.GetAllAnalyticsEnginesOptions{
			}

			response, err := ibmAnalyticsEngineApiService.GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetAnalyticsEngineByID - Get details of Analytics Engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {

			getAnalyticsEngineByIdOptions := &ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			analyticsEngine, response, err := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngine).ToNot(BeNil())

		})
	})

	Describe(`GetAnalyticsEngineStateByID - Get state of Analytics Engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions)`, func() {

			getAnalyticsEngineStateByIdOptions := &ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			analyticsEngineState, response, err := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineState).ToNot(BeNil())

		})
	})

	Describe(`CreateCustomizationRequest - Create an adhoc customization request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions)`, func() {

			analyticsEngineCustomActionScriptModel := &ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript{
				SourceType: core.StringPtr("http"),
				ScriptPath: core.StringPtr("testString"),
				SourceProps: map[string]interface{}{"anyKey": "anyValue"},
			}

			analyticsEngineCustomActionModel := &ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{
				Name: core.StringPtr("testString"),
				Type: core.StringPtr("bootstrap"),
				Script: analyticsEngineCustomActionScriptModel,
				ScriptParams: []string{"testString"},
			}

			createCustomizationRequestOptions := &ibmanalyticsengineapiv2.CreateCustomizationRequestOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
				Target: core.StringPtr("all"),
				CustomActions: []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel},
			}

			analyticsEngineCreateCustomizationResponse, response, err := ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCreateCustomizationResponse).ToNot(BeNil())

		})
	})

	Describe(`GetAllCustomizationRequests - Get all customization requests run on an Analytics Engine cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions)`, func() {

			getAllCustomizationRequestsOptions := &ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			analyticsEngineCustomizationRequestCollectionItem, response, err := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCustomizationRequestCollectionItem).ToNot(BeNil())

		})
	})

	Describe(`GetCustomizationRequestByID - Retrieve details of specified customization request ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions)`, func() {

			getCustomizationRequestByIdOptions := &ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
				RequestID: core.StringPtr("testString"),
			}

			analyticsEngineCustomizationRunDetails, response, err := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCustomizationRunDetails).ToNot(BeNil())

		})
	})

	Describe(`ResizeCluster - Resize the cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResizeCluster(resizeClusterOptions *ResizeClusterOptions)`, func() {

			resizeClusterRequestModel := &ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest{
				ComputeNodesCount: core.Int64Ptr(int64(38)),
			}

			resizeClusterOptions := &ibmanalyticsengineapiv2.ResizeClusterOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
				Body: resizeClusterRequestModel,
			}

			analyticsEngineResizeClusterResponse, response, err := ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineResizeClusterResponse).ToNot(BeNil())

		})
	})

	Describe(`ResetClusterPassword - Reset cluster password`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions)`, func() {

			resetClusterPasswordOptions := &ibmanalyticsengineapiv2.ResetClusterPasswordOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			analyticsEngineResetClusterPasswordResponse, response, err := ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineResetClusterPasswordResponse).ToNot(BeNil())

		})
	})

	Describe(`ConfigureLogging - Configure log aggregation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions)`, func() {

			analyticsEngineLoggingNodeSpecModel := &ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{
				NodeType: core.StringPtr("management"),
				Components: []string{"ambari-server"},
			}

			analyticsEngineLoggingServerModel := &ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer{
				Type: core.StringPtr("logdna"),
				Credential: core.StringPtr("testString"),
				ApiHost: core.StringPtr("testString"),
				LogHost: core.StringPtr("testString"),
				Owner: core.StringPtr("testString"),
			}

			configureLoggingOptions := &ibmanalyticsengineapiv2.ConfigureLoggingOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
				LogSpecs: []ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel},
				LogServer: analyticsEngineLoggingServerModel,
			}

			response, err := ibmAnalyticsEngineApiService.ConfigureLogging(configureLoggingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`GetLoggingConfig - Retrieve the status of log configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions)`, func() {

			getLoggingConfigOptions := &ibmanalyticsengineapiv2.GetLoggingConfigOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			analyticsEngineLoggingConfigDetails, response, err := ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineLoggingConfigDetails).ToNot(BeNil())

		})
	})

	Describe(`UpdatePrivateEndpointWhitelist - Update private endpoint whitelist`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions)`, func() {

			updatePrivateEndpointWhitelistOptions := &ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
				IpRanges: []string{"testString"},
				Action: core.StringPtr("add"),
			}

			analyticsEngineWhitelistResponse, response, err := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineWhitelistResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteLoggingConfig - Delete the log configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions)`, func() {

			deleteLoggingConfigOptions := &ibmanalyticsengineapiv2.DeleteLoggingConfigOptions{
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			response, err := ibmAnalyticsEngineApiService.DeleteLoggingConfig(deleteLoggingConfigOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
