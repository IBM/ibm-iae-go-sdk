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

package ibmanalyticsengineapiv3_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"reflect"
)

/**
 * This file contains an integration test for the ibmanalyticsengineapiv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmAnalyticsEngineApiV3 Integration Tests`, func() {

	const externalConfigFile = "../ibmanalyticsengine-service.env"

	var (
		err          error
		ibmAnalyticsEngineApiService *ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3
		serviceURL   string
		config       map[string]string
		// Retrieve any test-specific properties from the environment.
        instanceGuid string
        applicationId string
		
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
			instanceGuid = os.Getenv("IBM_ANALYTICS_ENGINE_INSTANCE_GUID")
			fmt.Printf("Service URL: %s\n", serviceURL)
			fmt.Printf("instanceGuid: %s\n", instanceGuid)
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
		})
	})

	Describe(`GetInstanceByID - Find instance by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceByID(getInstanceByIdOptions *GetInstanceByIdOptions)`, func() {
			fmt.Fprint(os.Stdout, fmt.Sprintf("instance id  value is - %v", instanceGuid))
			
			getInstanceByIdOptions := &ibmanalyticsengineapiv3.GetInstanceByIdOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			instanceDetails, response, err := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDetails).ToNot(BeNil())

		})
	})

	Describe(`CreateApplication - Deploys a Spark application in the Analytics Engine instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApplication(createApplicationOptions *CreateApplicationOptions)`, func() {

			createApplicationOptions := &ibmanalyticsengineapiv3.CreateApplicationOptions{
				InstanceID: core.StringPtr(instanceGuid),
				Application: core.StringPtr("/opt/ibm/spark/examples/src/main/python/wordcount.py"),
				//Class: core.StringPtr("testString"),
				ApplicationArguments: []string{"/opt/ibm/spark/examples/src/main/resources/people.txt"},
				//Conf: make(map[string]interface{}),
				//Env: make(map[string]interface{}),
				
				
			}
			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())			
			Expect(err).To(BeNil())
			applicationId = *applicationResponse.ApplicationID
			fmt.Printf("applicationResponse application_id : %v \n",applicationId)			
		})
	})

	Describe(`GetApplications - Gets all applications submitted in an instance with a specified inst_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApplications(getApplicationsOptions *GetApplicationsOptions)`, func() {

			getApplicationsOptions := &ibmanalyticsengineapiv3.GetApplicationsOptions{
				InstanceID: core.StringPtr(instanceGuid),
			}

			applicationCollection, response, err := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationCollection).ToNot(BeNil())

		})
	})

	Describe(`GetApplicationByID - Gets the details of the application identified by the app_id identifier`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApplicationByID(getApplicationByIdOptions *GetApplicationByIdOptions)`, func() {

			getApplicationByIdOptions := &ibmanalyticsengineapiv3.GetApplicationByIdOptions{
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
			}

			applicationGetResponse, response, err := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetResponse).ToNot(BeNil())

		})
	})

	Describe(`GetApplicationState - Gets the status of the application identified by the app_id identifier`, func() {
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

	Describe(`DeleteApplicationByID - Stops the specified application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApplicationByID(deleteApplicationByIdOptions *DeleteApplicationByIdOptions)`, func() {

			deleteApplicationByIdOptions := &ibmanalyticsengineapiv3.DeleteApplicationByIdOptions{
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
			}

			response, err := ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
