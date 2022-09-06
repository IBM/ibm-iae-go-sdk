// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmanalyticsengineapiv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmAnalyticsEngineApiV3 Integration Tests`, func() {
	const externalConfigFile = "../ibm_analytics_engine_api_v3.env"

	var (
		err          error
		ibmAnalyticsEngineApiService *ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(ibmanalyticsengineapiv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			ibmAnalyticsEngineApiServiceOptions := &ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{}

			ibmAnalyticsEngineApiService, err = ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(ibmAnalyticsEngineApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
			Expect(ibmAnalyticsEngineApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			ibmAnalyticsEngineApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetInstance - Find Analytics Engine by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {
			getInstanceOptions := &ibmanalyticsengineapiv3.GetInstanceOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			instance, response, err := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceState - Find Analytics Engine state by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceState(getInstanceStateOptions *GetInstanceStateOptions)`, func() {
			getInstanceStateOptions := &ibmanalyticsengineapiv3.GetInstanceStateOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			instanceGetStateResponse, response, err := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGetStateResponse).ToNot(BeNil())
		})
	})

	Describe(`SetInstanceHome - Set instance home`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetInstanceHome(setInstanceHomeOptions *SetInstanceHomeOptions)`, func() {
			setInstanceHomeOptions := &ibmanalyticsengineapiv3.SetInstanceHomeOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				NewInstanceID: core.StringPtr("testString"),
				NewProvider: core.StringPtr("ibm-cos"),
				NewType: core.StringPtr("objectstore"),
				NewRegion: core.StringPtr("us-south"),
				NewEndpoint: core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud"),
				NewHmacAccessKey: core.StringPtr("821**********0ae"),
				NewHmacSecretKey: core.StringPtr("03e****************4fc3"),
			}

			instanceHomeResponse, response, err := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceHomeResponse).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceDefaultConfigs - Get instance default Spark configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptions *GetInstanceDefaultConfigsOptions)`, func() {
			getInstanceDefaultConfigsOptions := &ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			result, response, err := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ReplaceInstanceDefaultConfigs - Replace instance default Spark configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptions *ReplaceInstanceDefaultConfigsOptions)`, func() {
			replaceInstanceDefaultConfigsOptions := &ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				Body: make(map[string]string),
			}

			result, response, err := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceDefaultConfigs - Update instance default Spark configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions *UpdateInstanceDefaultConfigsOptions)`, func() {
			updateInstanceDefaultConfigsOptions := &ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				Body: make(map[string]interface{}),
			}

			result, response, err := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CreateApplication - Deploy a Spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApplication(createApplicationOptions *CreateApplicationOptions)`, func() {
			createApplicationOptions := &ibmanalyticsengineapiv3.CreateApplicationOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				Application: core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py"),
				Jars: core.StringPtr("cos://cloud-object-storage/jars/tests.jar"),
				Packages: core.StringPtr("testString"),
				Repositories: core.StringPtr("testString"),
				Files: core.StringPtr("testString"),
				Archives: core.StringPtr("testString"),
				Name: core.StringPtr("spark-app"),
				Class: core.StringPtr("com.company.path.ClassName"),
				Arguments: []string{"[arg1, arg2, arg3]"},
				Conf: make(map[string]interface{}),
				Env: make(map[string]interface{}),
			}

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())
		})
	})

	Describe(`ListApplications - Retrieve all Spark applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApplications(listApplicationsOptions *ListApplicationsOptions)`, func() {
			listApplicationsOptions := &ibmanalyticsengineapiv3.ListApplicationsOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			applicationCollection, response, err := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationCollection).ToNot(BeNil())
		})
	})

	Describe(`GetApplication - Retrieve the details of a given Spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApplication(getApplicationOptions *GetApplicationOptions)`, func() {
			getApplicationOptions := &ibmanalyticsengineapiv3.GetApplicationOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				ApplicationID: core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b"),
			}

			applicationGetResponse, response, err := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetResponse).ToNot(BeNil())
		})
	})

	Describe(`GetApplicationState - Get the status of the application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApplicationState(getApplicationStateOptions *GetApplicationStateOptions)`, func() {
			getApplicationStateOptions := &ibmanalyticsengineapiv3.GetApplicationStateOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				ApplicationID: core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b"),
			}

			applicationGetStateResponse, response, err := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetStateResponse).ToNot(BeNil())
		})
	})

	Describe(`GetCurrentResourceConsumption - Get current resource consumption`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions *GetCurrentResourceConsumptionOptions)`, func() {
			getCurrentResourceConsumptionOptions := &ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			currentResourceConsumptionResponse, response, err := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(currentResourceConsumptionResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceLogForwardingConfig - Replace log forwarding configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions)`, func() {
			replaceLogForwardingConfigOptions := &ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				Enabled: core.BoolPtr(true),
				Sources: []string{"spark-driver", "spark-executor"},
				Tags: []string{"<tag_1>", "<tag_2>", "<tag_n"},
			}

			logForwardingConfigResponse, response, err := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logForwardingConfigResponse).ToNot(BeNil())
		})
	})

	Describe(`GetLogForwardingConfig - Get log forwarding configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLogForwardingConfig(getLogForwardingConfigOptions *GetLogForwardingConfigOptions)`, func() {
			getLogForwardingConfigOptions := &ibmanalyticsengineapiv3.GetLogForwardingConfigOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			logForwardingConfigResponse, response, err := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logForwardingConfigResponse).ToNot(BeNil())
		})
	})

	Describe(`ConfigurePlatformLogging - Enable or disable log forwarding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ConfigurePlatformLogging(configurePlatformLoggingOptions *ConfigurePlatformLoggingOptions)`, func() {
			configurePlatformLoggingOptions := &ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions{
				InstanceGuid: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				Enable: core.BoolPtr(true),
			}

			loggingConfigurationResponse, response, err := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loggingConfigurationResponse).ToNot(BeNil())
		})
	})

	Describe(`GetLoggingConfiguration - Retrieve the logging configuration for a given instance id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoggingConfiguration(getLoggingConfigurationOptions *GetLoggingConfigurationOptions)`, func() {
			getLoggingConfigurationOptions := &ibmanalyticsengineapiv3.GetLoggingConfigurationOptions{
				InstanceGuid: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
			}

			loggingConfigurationResponse, response, err := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loggingConfigurationResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteApplication - Stop application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions)`, func() {
			deleteApplicationOptions := &ibmanalyticsengineapiv3.DeleteApplicationOptions{
				InstanceID: core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09"),
				ApplicationID: core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b"),
			}

			response, err := ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
