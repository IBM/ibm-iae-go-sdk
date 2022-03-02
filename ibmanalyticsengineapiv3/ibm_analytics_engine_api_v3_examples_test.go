// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
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

package ibmanalyticsengineapiv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv3"
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
const externalConfigFile = "../ibm_analytics_engine_api_v3.env"

var (
	ibmAnalyticsEngineApiService *ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`IbmAnalyticsEngineApiV3 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmanalyticsengineapiv3.DefaultServiceName)
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

			ibmAnalyticsEngineApiServiceOptions := &ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{}

			ibmAnalyticsEngineApiService, err = ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(ibmAnalyticsEngineApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
		})
	})

	Describe(`IbmAnalyticsEngineApiV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstance request example`, func() {
			fmt.Println("\nGetInstance() result:")
			// begin-get_instance

			getInstanceOptions := ibmAnalyticsEngineApiService.NewGetInstanceOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			instance, response, err := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instance, "", "  ")
			fmt.Println(string(b))

			// end-get_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`GetInstanceState request example`, func() {
			fmt.Println("\nGetInstanceState() result:")
			// begin-get_instance_state

			getInstanceStateOptions := ibmAnalyticsEngineApiService.NewGetInstanceStateOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			instanceGetStateResponse, response, err := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGetStateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGetStateResponse).ToNot(BeNil())

		})
		It(`CreateInstanceHome request example`, func() {
			fmt.Println("\nCreateInstanceHome() result:")
			// begin-create_instance_home

			createInstanceHomeOptions := ibmAnalyticsEngineApiService.NewCreateInstanceHomeOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			instanceHomeResponse, response, err := ibmAnalyticsEngineApiService.CreateInstanceHome(createInstanceHomeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceHomeResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_home

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceHomeResponse).ToNot(BeNil())

		})
		It(`CreateApplication request example`, func() {
			fmt.Println("\nCreateApplication() result:")
			// begin-create_application

			createApplicationOptions := ibmAnalyticsEngineApiService.NewCreateApplicationOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_application

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())

		})
		It(`ListApplications request example`, func() {
			fmt.Println("\nListApplications() result:")
			// begin-list_applications

			listApplicationsOptions := ibmAnalyticsEngineApiService.NewListApplicationsOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			applicationCollection, response, err := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_applications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationCollection).ToNot(BeNil())

		})
		It(`GetApplication request example`, func() {
			fmt.Println("\nGetApplication() result:")
			// begin-get_application

			getApplicationOptions := ibmAnalyticsEngineApiService.NewGetApplicationOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
				"ff48cc19-0e7e-4627-aac6-0b4ad080397b",
			)

			applicationGetResponse, response, err := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_application

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetResponse).ToNot(BeNil())

		})
		It(`GetApplicationState request example`, func() {
			fmt.Println("\nGetApplicationState() result:")
			// begin-get_application_state

			getApplicationStateOptions := ibmAnalyticsEngineApiService.NewGetApplicationStateOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
				"ff48cc19-0e7e-4627-aac6-0b4ad080397b",
			)

			applicationGetStateResponse, response, err := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationGetStateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_application_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetStateResponse).ToNot(BeNil())

		})
		It(`ConfigurePlatformLogging request example`, func() {
			fmt.Println("\nConfigurePlatformLogging() result:")
			// begin-configure_platform_logging

			configurePlatformLoggingOptions := ibmAnalyticsEngineApiService.NewConfigurePlatformLoggingOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			loggingConfigurationResponse, response, err := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loggingConfigurationResponse, "", "  ")
			fmt.Println(string(b))

			// end-configure_platform_logging

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loggingConfigurationResponse).ToNot(BeNil())

		})
		It(`GetLoggingConfiguration request example`, func() {
			fmt.Println("\nGetLoggingConfiguration() result:")
			// begin-get_logging_configuration

			getLoggingConfigurationOptions := ibmAnalyticsEngineApiService.NewGetLoggingConfigurationOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			loggingConfigurationResponse, response, err := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loggingConfigurationResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_logging_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loggingConfigurationResponse).ToNot(BeNil())

		})
		It(`StartSparkHistoryServer request example`, func() {
			fmt.Println("\nStartSparkHistoryServer() result:")
			// begin-start_spark_history_server

			startSparkHistoryServerOptions := ibmAnalyticsEngineApiService.NewStartSparkHistoryServerOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			sparkHistoryServerStartResponse, response, err := ibmAnalyticsEngineApiService.StartSparkHistoryServer(startSparkHistoryServerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sparkHistoryServerStartResponse, "", "  ")
			fmt.Println(string(b))

			// end-start_spark_history_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkHistoryServerStartResponse).ToNot(BeNil())

		})
		It(`GetSparkHistoryServer request example`, func() {
			fmt.Println("\nGetSparkHistoryServer() result:")
			// begin-get_spark_history_server

			getSparkHistoryServerOptions := ibmAnalyticsEngineApiService.NewGetSparkHistoryServerOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			sparkHistoryServerResponse, response, err := ibmAnalyticsEngineApiService.GetSparkHistoryServer(getSparkHistoryServerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sparkHistoryServerResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_spark_history_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkHistoryServerResponse).ToNot(BeNil())

		})
		It(`StopSparkHistoryServer request example`, func() {
			// begin-stop_spark_history_server

			stopSparkHistoryServerOptions := ibmAnalyticsEngineApiService.NewStopSparkHistoryServerOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			response, err := ibmAnalyticsEngineApiService.StopSparkHistoryServer(stopSparkHistoryServerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from StopSparkHistoryServer(): %d\n", response.StatusCode)
			}

			// end-stop_spark_history_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoggingConfiguration request example`, func() {
			// begin-delete_logging_configuration

			deleteLoggingConfigurationOptions := ibmAnalyticsEngineApiService.NewDeleteLoggingConfigurationOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
			)

			response, err := ibmAnalyticsEngineApiService.DeleteLoggingConfiguration(deleteLoggingConfigurationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteLoggingConfiguration(): %d\n", response.StatusCode)
			}

			// end-delete_logging_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteApplication request example`, func() {
			// begin-delete_application

			deleteApplicationOptions := ibmAnalyticsEngineApiService.NewDeleteApplicationOptions(
				"e64c907a-e82f-46fd-addc-ccfafbd28b09",
				"ff48cc19-0e7e-4627-aac6-0b4ad080397b",
			)

			response, err := ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteApplication(): %d\n", response.StatusCode)
			}

			// end-delete_application

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
