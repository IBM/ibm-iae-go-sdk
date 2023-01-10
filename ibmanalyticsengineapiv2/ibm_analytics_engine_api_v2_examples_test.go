// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/v2/ibmanalyticsengineapiv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM Analytics Engine API service.
//
// The following configuration properties are assumed to be defined:
// IBM_ANALYTICS_ENGINE_API_URL=<service base url>
// IBM_ANALYTICS_ENGINE_API_AUTH_TYPE=iam
// IBM_ANALYTICS_ENGINE_API_APIKEY=<IAM apikey>
// IBM_ANALYTICS_ENGINE_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../ibm_analytics_engine_api_v2.env"

var (
	ibmAnalyticsEngineApiService *ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IbmAnalyticsEngineApiV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmanalyticsengineapiv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			ibmAnalyticsEngineApiServiceOptions := &ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{}

			ibmAnalyticsEngineApiService, err = ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(ibmAnalyticsEngineApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
		})
	})

	Describe(`IbmAnalyticsEngineApiV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllAnalyticsEngines request example`, func() {
			// begin-getAllAnalyticsEngines

			getAllAnalyticsEnginesOptions := ibmAnalyticsEngineApiService.NewGetAllAnalyticsEnginesOptions()

			response, err := ibmAnalyticsEngineApiService.GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions)
			if err != nil {
				panic(err)
			}

			// end-getAllAnalyticsEngines
			fmt.Printf("\nGetAllAnalyticsEngines() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetAnalyticsEngineByID request example`, func() {
			fmt.Println("\nGetAnalyticsEngineByID() result:")
			// begin-getAnalyticsEngineById

			getAnalyticsEngineByIdOptions := ibmAnalyticsEngineApiService.NewGetAnalyticsEngineByIdOptions(
				"testString",
			)

			analyticsEngine, response, err := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngine, "", "  ")
			fmt.Println(string(b))

			// end-getAnalyticsEngineById

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngine).ToNot(BeNil())

		})
		It(`GetAnalyticsEngineStateByID request example`, func() {
			fmt.Println("\nGetAnalyticsEngineStateByID() result:")
			// begin-getAnalyticsEngineStateById

			getAnalyticsEngineStateByIdOptions := ibmAnalyticsEngineApiService.NewGetAnalyticsEngineStateByIdOptions(
				"testString",
			)

			analyticsEngineState, response, err := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineState, "", "  ")
			fmt.Println(string(b))

			// end-getAnalyticsEngineStateById

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineState).ToNot(BeNil())

		})
		It(`CreateCustomizationRequest request example`, func() {
			fmt.Println("\nCreateCustomizationRequest() result:")
			// begin-createCustomizationRequest

			analyticsEngineCustomActionModel := &ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{
				Name: core.StringPtr("testString"),
			}

			createCustomizationRequestOptions := ibmAnalyticsEngineApiService.NewCreateCustomizationRequestOptions(
				"testString",
				"all",
				[]ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel},
			)

			analyticsEngineCreateCustomizationResponse, response, err := ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineCreateCustomizationResponse, "", "  ")
			fmt.Println(string(b))

			// end-createCustomizationRequest

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCreateCustomizationResponse).ToNot(BeNil())

		})
		It(`GetAllCustomizationRequests request example`, func() {
			fmt.Println("\nGetAllCustomizationRequests() result:")
			// begin-getAllCustomizationRequests

			getAllCustomizationRequestsOptions := ibmAnalyticsEngineApiService.NewGetAllCustomizationRequestsOptions(
				"testString",
			)

			analyticsEngineCustomizationRequestCollectionItem, response, err := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineCustomizationRequestCollectionItem, "", "  ")
			fmt.Println(string(b))

			// end-getAllCustomizationRequests

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCustomizationRequestCollectionItem).ToNot(BeNil())

		})
		It(`GetCustomizationRequestByID request example`, func() {
			fmt.Println("\nGetCustomizationRequestByID() result:")
			// begin-getCustomizationRequestById

			getCustomizationRequestByIdOptions := ibmAnalyticsEngineApiService.NewGetCustomizationRequestByIdOptions(
				"testString",
				"testString",
			)

			analyticsEngineCustomizationRunDetails, response, err := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineCustomizationRunDetails, "", "  ")
			fmt.Println(string(b))

			// end-getCustomizationRequestById

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineCustomizationRunDetails).ToNot(BeNil())

		})
		It(`ResizeCluster request example`, func() {
			fmt.Println("\nResizeCluster() result:")
			// begin-resizeCluster

			resizeClusterRequestModel := &ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest{
			}

			resizeClusterOptions := ibmAnalyticsEngineApiService.NewResizeClusterOptions(
				"testString",
				resizeClusterRequestModel,
			)

			analyticsEngineResizeClusterResponse, response, err := ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineResizeClusterResponse, "", "  ")
			fmt.Println(string(b))

			// end-resizeCluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineResizeClusterResponse).ToNot(BeNil())

		})
		It(`ResetClusterPassword request example`, func() {
			fmt.Println("\nResetClusterPassword() result:")
			// begin-resetClusterPassword

			resetClusterPasswordOptions := ibmAnalyticsEngineApiService.NewResetClusterPasswordOptions(
				"testString",
			)

			analyticsEngineResetClusterPasswordResponse, response, err := ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineResetClusterPasswordResponse, "", "  ")
			fmt.Println(string(b))

			// end-resetClusterPassword

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineResetClusterPasswordResponse).ToNot(BeNil())

		})
		It(`ConfigureLogging request example`, func() {
			// begin-configureLogging

			analyticsEngineLoggingNodeSpecModel := &ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{
				NodeType: core.StringPtr("management"),
				Components: []string{"ambari-server"},
			}

			analyticsEngineLoggingServerModel := &ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer{
				Type: core.StringPtr("logdna"),
				Credential: core.StringPtr("testString"),
				ApiHost: core.StringPtr("testString"),
				LogHost: core.StringPtr("testString"),
			}

			configureLoggingOptions := ibmAnalyticsEngineApiService.NewConfigureLoggingOptions(
				"testString",
				[]ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel},
				analyticsEngineLoggingServerModel,
			)

			response, err := ibmAnalyticsEngineApiService.ConfigureLogging(configureLoggingOptions)
			if err != nil {
				panic(err)
			}

			// end-configureLogging
			fmt.Printf("\nConfigureLogging() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`GetLoggingConfig request example`, func() {
			fmt.Println("\nGetLoggingConfig() result:")
			// begin-getLoggingConfig

			getLoggingConfigOptions := ibmAnalyticsEngineApiService.NewGetLoggingConfigOptions(
				"testString",
			)

			analyticsEngineLoggingConfigDetails, response, err := ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineLoggingConfigDetails, "", "  ")
			fmt.Println(string(b))

			// end-getLoggingConfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineLoggingConfigDetails).ToNot(BeNil())

		})
		It(`UpdatePrivateEndpointWhitelist request example`, func() {
			fmt.Println("\nUpdatePrivateEndpointWhitelist() result:")
			// begin-updatePrivateEndpointWhitelist

			updatePrivateEndpointWhitelistOptions := ibmAnalyticsEngineApiService.NewUpdatePrivateEndpointWhitelistOptions(
				"testString",
				[]string{"testString"},
				"add",
			)

			analyticsEngineWhitelistResponse, response, err := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyticsEngineWhitelistResponse, "", "  ")
			fmt.Println(string(b))

			// end-updatePrivateEndpointWhitelist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyticsEngineWhitelistResponse).ToNot(BeNil())

		})
		It(`DeleteLoggingConfig request example`, func() {
			// begin-deleteLoggingConfig

			deleteLoggingConfigOptions := ibmAnalyticsEngineApiService.NewDeleteLoggingConfigOptions(
				"testString",
			)

			response, err := ibmAnalyticsEngineApiService.DeleteLoggingConfig(deleteLoggingConfigOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteLoggingConfig
			fmt.Printf("\nDeleteLoggingConfig() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})
})
