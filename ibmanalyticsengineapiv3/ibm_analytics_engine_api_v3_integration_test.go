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
		// !!! Start of custom content to be copied !!!  
		// Declartion of configuration variables
		instanceGuid string
		applicationId string
		instanceGuidInstanceHome string
		hmacAccessKey string
		hmacSecretKey string
		// !!! End of custom content to be copied !!! 
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

			// !!! Start of custom content to be copied !!!  
			// Assignment of configuration variables from environment
			instanceGuid = os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID")
			instanceGuidInstanceHome = os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID_INSTANCE_HOME")
			hmacAccessKey = os.Getenv("IBM_ANALYTICS_ENGINE_HMAC_ACCESS_KEY")
			hmacSecretKey = os.Getenv("IBM_ANALYTICS_ENGINE_HMAC_SECRET_KEY")
			// !!! End of custom content to be copied !!! 

			fmt.Printf("Service URL: %s\n", serviceURL)

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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
		})
	})

	Describe(`SetInstanceHome - Set instance home`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetInstanceHome(setInstanceHomeOptions *SetInstanceHomeOptions)`, func() {

			setInstanceHomeOptions := &ibmanalyticsengineapiv3.SetInstanceHomeOptions{
				InstanceID: core.StringPtr(instanceGuidInstanceHome),
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 409
			// 500
			//
		})
	})

	Describe(`CreateApplication - Deploy a Spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApplication(createApplicationOptions *CreateApplicationOptions)`, func() {

			createApplicationOptions := &ibmanalyticsengineapiv3.CreateApplicationOptions{
				InstanceID: core.StringPtr(instanceGuid),
				Application: core.StringPtr("/opt/ibm/spark/examples/src/main/python/wordcount.py"),
				// Jars: core.StringPtr("cos://cloud-object-storage/jars/tests.jar"),
				// Packages: core.StringPtr("testString"),
				// Repositories: core.StringPtr("testString"),
				// Files: core.StringPtr("testString"),
				// Archives: core.StringPtr("testString"),
				// Name: core.StringPtr("spark-app"),
				// Class: core.StringPtr("com.company.path.ClassName"),
				Arguments: []string{"/opt/ibm/spark/examples/src/main/resources/people.txt"},
				// Conf: make(map[string]interface{}),
				// Env: make(map[string]interface{}),
			}

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())
			// !!! Start of custom content to be copied !!! 
			applicationId = *applicationResponse.ID
			fmt.Printf("applicationResponse application_id : %v \n",applicationId)
			// !!! End of custom content to be copied !!! 
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
		})
	})

	Describe(`ListApplications - Retrieve all Spark applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListApplications(listApplicationsOptions *ListApplicationsOptions)`, func() {

			listApplicationsOptions := &ibmanalyticsengineapiv3.ListApplicationsOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			applicationCollection, response, err := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationCollection).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 404
			// 500
			//
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

			sparkHistoryServerStartResponse, response, err := ibmAnalyticsEngineApiService.StartSparkHistoryServer(startSparkHistoryServerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkHistoryServerStartResponse).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
		})
	})

	Describe(`GetSparkHistoryServer - Retrieve Spark history server details`, func() {
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
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

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 404
			// 500
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//
