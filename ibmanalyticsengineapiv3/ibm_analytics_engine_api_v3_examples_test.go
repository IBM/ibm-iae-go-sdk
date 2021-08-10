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
		It(`GetInstanceByID request example`, func() {
			fmt.Println("\nGetInstanceByID() result:")
			// begin-get_instance_by_id

			getInstanceByIdOptions := ibmAnalyticsEngineApiService.NewGetInstanceByIdOptions(
				"testString",
			)

			instanceDetails, response, err := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDetails, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDetails).ToNot(BeNil())

		})
		It(`CreateApplication request example`, func() {
			fmt.Println("\nCreateApplication() result:")
			// begin-create_application

			createApplicationOptions := ibmAnalyticsEngineApiService.NewCreateApplicationOptions(
				"testString",
			)

			applicationResponse, response, err := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_application

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(applicationResponse).ToNot(BeNil())

		})
		It(`GetApplications request example`, func() {
			fmt.Println("\nGetApplications() result:")

			// begin-get_applications

			getApplicationsOptions := ibmAnalyticsEngineApiService.NewGetApplicationsOptions(
				"testString",
			)

			applicationCollection, response, err := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationCollection, "", "  ")
			fmt.Println(string(b))

			// end-get_applications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationCollection).ToNot(BeNil())

		})
		It(`GetApplicationByID request example`, func() {

			fmt.Println("\nGetApplicationByID() result:")


			getApplicationByIdOptions := ibmAnalyticsEngineApiService.NewGetApplicationByIdOptions(
				"testString",
				"testString",
			)

			applicationGetResponse, response, err := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(applicationGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_application_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(applicationGetResponse).ToNot(BeNil())

		})
		It(`GetApplicationState request example`, func() {
			fmt.Println("\nGetApplicationState() result:")
			// begin-get_application_state

			getApplicationStateOptions := ibmAnalyticsEngineApiService.NewGetApplicationStateOptions(
				"testString",
				"testString",
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
		It(`DeleteApplicationByID request example`, func() {
			// begin-delete_application_by_id

			deleteApplicationByIdOptions := ibmAnalyticsEngineApiService.NewDeleteApplicationByIdOptions(
				"testString",
				"testString",
			)

			response, err := ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_application_by_id
			fmt.Printf("\nDeleteApplicationByID() response status code: %d\n", response.StatusCode)


			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
