// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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
	"github.com/IBM/ibm-iae-go-sdk/v2/ibmanalyticsengineapiv3"
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
		instanceGuid string
		instanceGuidWithoutInstanceHome string
        applicationId string
		hmacAccessKey string
		hmacSecretKey string
		alternateHmacAccessKey string
		alternateHmacSecretKey string
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
			instanceGuid = config["INSTANCE_GUID"]
			instanceGuidWithoutInstanceHome = config["INSTANCE_GUID_WO_INSTANCE_HOME"]
			hmacAccessKey = config["HMAC_ACCESS_KEY"]
			hmacSecretKey = config["HMAC_SECRET_KEY"]
			alternateHmacAccessKey = config["ALTERNATE_HMAC_ACCESS_KEY"]
			alternateHmacSecretKey = config["ALTERNATE_HMAC_SECRET_KEY"]


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
				InstanceID: core.StringPtr(instanceGuid),
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
				InstanceID: core.StringPtr(instanceGuid),
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
				InstanceID: core.StringPtr(instanceGuidWithoutInstanceHome),
				NewInstanceID: core.StringPtr("testString"),
				NewProvider: core.StringPtr("ibm-cos"),
				NewType: core.StringPtr("objectstore"),
				NewRegion: core.StringPtr("us-south"),
				NewEndpoint: core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud"),
				NewHmacAccessKey: core.StringPtr(hmacAccessKey),
				NewHmacSecretKey: core.StringPtr(hmacSecretKey),
			}

			instanceHomeResponse, response, err := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceHomeResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceHomeCredentials - Update instance home credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceHomeCredentials(updateInstanceHomeCredentialsOptions *UpdateInstanceHomeCredentialsOptions)`, func() {
			updateInstanceHomeCredentialsOptions := &ibmanalyticsengineapiv3.UpdateInstanceHomeCredentialsOptions{
				InstanceID: core.StringPtr(instanceGuid),
				HmacAccessKey: core.StringPtr(alternateHmacAccessKey),
				HmacSecretKey: core.StringPtr(alternateHmacSecretKey),
			}

			instanceHomeResponse, response, err := ibmAnalyticsEngineApiService.UpdateInstanceHomeCredentials(updateInstanceHomeCredentialsOptions)
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
				InstanceID: core.StringPtr(instanceGuid),
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
			defaultConfigs := map[string]string {
				"spark.driver.memory": "8G",
				"spark.driver.cores": "2",
			}
			replaceInstanceDefaultConfigsOptions := &ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions{
				InstanceID: core.StringPtr(instanceGuid),
				Body: defaultConfigs,
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
				InstanceID: core.StringPtr(instanceGuid),
				Body: make(map[string]interface{}),
			}

			result, response, err := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceDefaultRuntime - Get instance default runtime`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceDefaultRuntime(getInstanceDefaultRuntimeOptions *GetInstanceDefaultRuntimeOptions)`, func() {
			getInstanceDefaultRuntimeOptions := &ibmanalyticsengineapiv3.GetInstanceDefaultRuntimeOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			runtime, response, err := ibmAnalyticsEngineApiService.GetInstanceDefaultRuntime(getInstanceDefaultRuntimeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runtime).ToNot(BeNil())
		})
	})

	Describe(`ReplaceInstanceDefaultRuntime - Replace instance default runtime`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceInstanceDefaultRuntime(replaceInstanceDefaultRuntimeOptions *ReplaceInstanceDefaultRuntimeOptions)`, func() {
			replaceInstanceDefaultRuntimeOptions := &ibmanalyticsengineapiv3.ReplaceInstanceDefaultRuntimeOptions{
				InstanceID: core.StringPtr(instanceGuid),
				SparkVersion: core.StringPtr("3.3"),
			}

			runtime, response, err := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultRuntime(replaceInstanceDefaultRuntimeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runtime).ToNot(BeNil())
		})
	})

	Describe(`CreateApplication - Deploy a Spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApplication(createApplicationOptions *CreateApplicationOptions)`, func() {
			runtimeModel := &ibmanalyticsengineapiv3.Runtime{
				SparkVersion: core.StringPtr("3.1"),
			}

			createApplicationOptions := &ibmanalyticsengineapiv3.CreateApplicationOptions{
				InstanceID: core.StringPtr(instanceGuid),
				Application: core.StringPtr("/opt/ibm/spark/examples/src/main/python/wordcount.py"),
				Arguments: []string{"/opt/ibm/spark/examples/src/main/resources/people.txt"},
				Runtime: runtimeModel,
			}

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())

			applicationId = *applicationResponse.ID
			fmt.Printf("applicationResponse application_id : %v \n",applicationId)	
		})
	})

	Describe(`ListApplications - List all Spark applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApplications(listApplicationsOptions *ListApplicationsOptions) with pagination`, func(){
			listApplicationsOptions := &ibmanalyticsengineapiv3.ListApplicationsOptions{
				InstanceID: core.StringPtr(instanceGuid),
				State: []string{"accepted","running","finished","stopped","failed"},
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
			}

			listApplicationsOptions.Start = nil
			listApplicationsOptions.Limit = core.Int64Ptr(1)

			var allResults []ibmanalyticsengineapiv3.Application
			for {
				applicationCollection, response, err := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(applicationCollection).ToNot(BeNil())
				allResults = append(allResults, applicationCollection.Applications...)

				listApplicationsOptions.Start, err = applicationCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listApplicationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListApplications(listApplicationsOptions *ListApplicationsOptions) using ApplicationsPager`, func(){
			listApplicationsOptions := &ibmanalyticsengineapiv3.ListApplicationsOptions{
				InstanceID: core.StringPtr(instanceGuid),
				State: []string{"accepted","running","finished","stopped","failed"},
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := ibmAnalyticsEngineApiService.NewApplicationsPager(listApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []ibmanalyticsengineapiv3.Application
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = ibmAnalyticsEngineApiService.NewApplicationsPager(listApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListApplications() returned a total of %d item(s) using ApplicationsPager.\n", len(allResults))
		})
	})

	Describe(`GetApplication - Retrieve the details of a given Spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApplication(getApplicationOptions *GetApplicationOptions)`, func() {
			getApplicationOptions := &ibmanalyticsengineapiv3.GetApplicationOptions{
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
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
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
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
				InstanceID: core.StringPtr(instanceGuid),
			}

			currentResourceConsumptionResponse, response, err := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(currentResourceConsumptionResponse).ToNot(BeNil())
		})
	})

	Describe(`GetResourceConsumptionLimits - Get resource consumption limits`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceConsumptionLimits(getResourceConsumptionLimitsOptions *GetResourceConsumptionLimitsOptions)`, func() {
			getResourceConsumptionLimitsOptions := &ibmanalyticsengineapiv3.GetResourceConsumptionLimitsOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			resourceConsumptionLimitsResponse, response, err := ibmAnalyticsEngineApiService.GetResourceConsumptionLimits(getResourceConsumptionLimitsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceConsumptionLimitsResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceLogForwardingConfig - Replace log forwarding configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions)`, func() {
			replaceLogForwardingConfigOptions := &ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions{
				InstanceID: core.StringPtr(instanceGuid),
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
				InstanceID: core.StringPtr(instanceGuid),
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
				InstanceGuid: core.StringPtr(instanceGuid),
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
				InstanceGuid: core.StringPtr(instanceGuid),
			}

			loggingConfigurationResponse, response, err := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loggingConfigurationResponse).ToNot(BeNil())
		})
	})

	Describe(`StartSparkHistoryServer - Start Spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`StartSparkHistoryServer(startSparkHistoryServerOptions *StartSparkHistoryServerOptions)`, func() {
			startSparkHistoryServerOptions := &ibmanalyticsengineapiv3.StartSparkHistoryServerOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			sparkHistoryServerResponse, response, err := ibmAnalyticsEngineApiService.StartSparkHistoryServer(startSparkHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(sparkHistoryServerResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSparkHistoryServer - Get Spark history server details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkHistoryServer(getSparkHistoryServerOptions *GetSparkHistoryServerOptions)`, func() {
			getSparkHistoryServerOptions := &ibmanalyticsengineapiv3.GetSparkHistoryServerOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			sparkHistoryServerResponse, response, err := ibmAnalyticsEngineApiService.GetSparkHistoryServer(getSparkHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkHistoryServerResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteApplication - Stop application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions)`, func() {
			deleteApplicationOptions := &ibmanalyticsengineapiv3.DeleteApplicationOptions{
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
			}

			response, err := ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`StopSparkHistoryServer - Stop Spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`StopSparkHistoryServer(stopSparkHistoryServerOptions *StopSparkHistoryServerOptions)`, func() {
			stopSparkHistoryServerOptions := &ibmanalyticsengineapiv3.StopSparkHistoryServerOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			response, err := ibmAnalyticsEngineApiService.StopSparkHistoryServer(stopSparkHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
