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
var _ = Describe(`IbmAnalyticsEngineApiV3 Examples Tests`, func() {

	const externalConfigFile = "../ibm_analytics_engine_api_v3.env"

	var (
		ibmAnalyticsEngineApiService *ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmanalyticsengineapiv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
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
				"dc0e9889-eab2-4t9e-9441-566209499546",
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
				"dc0e9889-eab2-4t9e-9441-566209499546",
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
		It(`SetInstanceHome request example`, func() {
			fmt.Println("\nSetInstanceHome() result:")
			// begin-set_instance_home

			setInstanceHomeOptions := ibmAnalyticsEngineApiService.NewSetInstanceHomeOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			  )
			setInstanceHomeOptions.SetNewHmacAccessKey("b9****************************4b")
			setInstanceHomeOptions.SetNewHmacSecretKey("fa********************************************8a")

			instanceHomeResponse, response, err := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceHomeResponse, "", "  ")
			fmt.Println(string(b))

			// end-set_instance_home

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceHomeResponse).ToNot(BeNil())
		})
		It(`GetInstanceDefaultConfigs request example`, func() {
			fmt.Println("\nGetInstanceDefaultConfigs() result:")
			// begin-get_instance_default_configs

			getInstanceDefaultConfigsOptions := ibmAnalyticsEngineApiService.NewGetInstanceDefaultConfigsOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			instanceDefaultConfigs, response, err := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDefaultConfigs, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_default_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDefaultConfigs).ToNot(BeNil())
		})
		It(`ReplaceInstanceDefaultConfigs request example`, func() {
			fmt.Println("\nReplaceInstanceDefaultConfigs() result:")
			// begin-replace_instance_default_configs

			defaultConfigs := map[string]string {
				"spark.driver.memory": "8G",
				"spark.driver.cores": "2",
			}
			
			replaceInstanceDefaultConfigsOptions := ibmAnalyticsEngineApiService.NewReplaceInstanceDefaultConfigsOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
				defaultConfigs,
			)

			instanceDefaultConfigs, response, err := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDefaultConfigs, "", "  ")
			fmt.Println(string(b))

			// end-replace_instance_default_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDefaultConfigs).ToNot(BeNil())
		})
		It(`UpdateInstanceDefaultConfigs request example`, func() {
			fmt.Println("\nUpdateInstanceDefaultConfigs() result:")
			// begin-update_instance_default_configs

			defaultConfigs := map[string]interface{} {
				"ae.spark.history-server.cores": "1",
				"ae.spark.history-server.memory": "4G",
			}

			updateInstanceDefaultConfigsOptions := ibmAnalyticsEngineApiService.NewUpdateInstanceDefaultConfigsOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
				defaultConfigs,
			)

			instanceDefaultConfigs, response, err := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDefaultConfigs, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_default_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDefaultConfigs).ToNot(BeNil())
		})
		It(`GetInstanceDefaultRuntime request example`, func() {
			fmt.Println("\nGetInstanceDefaultRuntime() result:")
			// begin-get_instance_default_runtime

			getInstanceDefaultRuntimeOptions := ibmAnalyticsEngineApiService.NewGetInstanceDefaultRuntimeOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			runtime, response, err := ibmAnalyticsEngineApiService.GetInstanceDefaultRuntime(getInstanceDefaultRuntimeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(runtime, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_default_runtime

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runtime).ToNot(BeNil())
		})
		It(`ReplaceInstanceDefaultRuntime request example`, func() {
			fmt.Println("\nReplaceInstanceDefaultRuntime() result:")
			// begin-replace_instance_default_runtime

			replaceInstanceDefaultRuntimeOptions := ibmAnalyticsEngineApiService.NewReplaceInstanceDefaultRuntimeOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)
			replaceInstanceDefaultRuntimeOptions.SetSparkVersion("3.3")

			runtime, response, err := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultRuntime(replaceInstanceDefaultRuntimeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(runtime, "", "  ")
			fmt.Println(string(b))

			// end-replace_instance_default_runtime

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runtime).ToNot(BeNil())
		})
		It(`CreateApplication request example`, func() {
			fmt.Println("\nCreateApplication() result:")
			// begin-create_application

			createApplicationOptions := ibmAnalyticsEngineApiService.NewCreateApplicationOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)
			createApplicationOptions.SetApplication("/opt/ibm/spark/examples/src/main/python/wordcount.py")
			createApplicationOptions.SetArguments([]string{"/opt/ibm/spark/examples/src/main/resources/people.txt"})
			createApplicationOptions.SetRuntime(&ibmanalyticsengineapiv3.Runtime{
				SparkVersion: core.StringPtr("3.3"),
			})

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
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)
			listApplicationsOptions.SetState([]string{"accepted", "submitted", "waiting", "running", "finished", "failed"})

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
				"dc0e9889-eab2-4t9e-9441-566209499546",
				"db933645-0b68-4dcb-80d8-7b71a6c8e542",
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
				"dc0e9889-eab2-4t9e-9441-566209499546",
				"db933645-0b68-4dcb-80d8-7b71a6c8e542",
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
		It(`GetCurrentResourceConsumption request example`, func() {
			fmt.Println("\nGetCurrentResourceConsumption() result:")
			// begin-get_current_resource_consumption

			getCurrentResourceConsumptionOptions := ibmAnalyticsEngineApiService.NewGetCurrentResourceConsumptionOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			currentResourceConsumptionResponse, response, err := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(currentResourceConsumptionResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_current_resource_consumption

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(currentResourceConsumptionResponse).ToNot(BeNil())
		})
		It(`GetResourceConsumptionLimits request example`, func() {
			fmt.Println("\nGetResourceConsumptionLimits() result:")
			// begin-get_resource_consumption_limits

			getResourceConsumptionLimitsOptions := ibmAnalyticsEngineApiService.NewGetResourceConsumptionLimitsOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			resourceConsumptionLimitsResponse, response, err := ibmAnalyticsEngineApiService.GetResourceConsumptionLimits(getResourceConsumptionLimitsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceConsumptionLimitsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_consumption_limits

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceConsumptionLimitsResponse).ToNot(BeNil())
		})
		It(`ReplaceLogForwardingConfig request example`, func() {
			fmt.Println("\nReplaceLogForwardingConfig() result:")
			// begin-replace_log_forwarding_config

			replaceLogForwardingConfigOptions := ibmAnalyticsEngineApiService.NewReplaceLogForwardingConfigOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			logForwardingConfigResponse, response, err := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logForwardingConfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_log_forwarding_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logForwardingConfigResponse).ToNot(BeNil())
		})
		It(`GetLogForwardingConfig request example`, func() {
			fmt.Println("\nGetLogForwardingConfig() result:")
			// begin-get_log_forwarding_config

			getLogForwardingConfigOptions := ibmAnalyticsEngineApiService.NewGetLogForwardingConfigOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
			)

			logForwardingConfigResponse, response, err := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logForwardingConfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_log_forwarding_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logForwardingConfigResponse).ToNot(BeNil())
		})
		It(`ConfigurePlatformLogging request example`, func() {
			fmt.Println("\nConfigurePlatformLogging() result:")
			// begin-configure_platform_logging

			configurePlatformLoggingOptions := ibmAnalyticsEngineApiService.NewConfigurePlatformLoggingOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
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
				"dc0e9889-eab2-4t9e-9441-566209499546",
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
		It(`DeleteApplication request example`, func() {
			// begin-delete_application

			deleteApplicationOptions := ibmAnalyticsEngineApiService.NewDeleteApplicationOptions(
				"dc0e9889-eab2-4t9e-9441-566209499546",
				"db933645-0b68-4dcb-80d8-7b71a6c8e542",
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
