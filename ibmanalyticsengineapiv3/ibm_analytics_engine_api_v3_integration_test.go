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
			//fmt.Printf("GetInstance response: %s\n", response)
			//fmt.Printf("GetInstance instance: %v\n", instance)
			//fmt.Printf("GetInstance err: %s\n", err)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

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
				//Class: core.StringPtr("com.company.path.ClassName"),
				Arguments: []string{"/opt/ibm/spark/examples/src/main/resources/people.txt"},
				//Conf: make(map[string]interface{}),
				//Env: make(map[string]interface{}),
			}

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)

			//fmt.Printf("CreateApplication applicationResponse: %v\n", applicationResponse)
			//fmt.Printf("CreateApplication response: %v\n", response)
			//fmt.Printf("CreateApplication err: %s\n", err)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(applicationResponse).ToNot(BeNil())
			applicationId = *applicationResponse.ID
			fmt.Printf("applicationResponse application_id : %v \n",applicationId)	
			
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
			//fmt.Printf("ListApplications applicationCollection: %v\n", applicationCollection)
			//fmt.Printf("ListApplications response: %s\n", response)
			//fmt.Printf("ListApplications err: %s\n", err)
			
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
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
			}

			applicationGetResponse, response, err := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptions)

			//fmt.Printf("GetApplication applicationGetResponse: %v\n", applicationGetResponse)
			//fmt.Printf("GetApplication response: %s\n", response)
			//fmt.Printf("GetApplication err: %s\n", err)
			
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


			//fmt.Printf("GetApplicationState applicationGetStateResponse: %v\n", applicationGetStateResponse)
			//fmt.Printf("GetApplicationState response: %s\n", response)
			//fmt.Printf("GetApplicationState err: %s\n", err)
			
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetStateResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteApplicationByID - Stop application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteApplicationByID(deleteApplicationByIdOptions *DeleteApplicationByIdOptions)`, func() {

			deleteApplicationByIdOptions := &ibmanalyticsengineapiv3.DeleteApplicationByIdOptions{
				InstanceID: core.StringPtr(instanceGuid),
				ApplicationID: core.StringPtr(applicationId),
			}


			
			response, err := ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptions)

			//fmt.Printf("DeleteApplicationByID response: %v\n", response)
			//fmt.Printf("DeleteApplicationByID err: %s\n", err)
			
			
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
	
})

//
// Utility functions are declared in the unit test file
//
