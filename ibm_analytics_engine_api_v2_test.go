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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IbmAnalyticsEngineApiV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmAnalyticsEngineApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				URL: "https://ibmanalyticsengineapiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmAnalyticsEngineApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_ANALYTICS_ENGINE_API_URL": "https://ibmanalyticsengineapiv2/api",
				"IBM_ANALYTICS_ENGINE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				})
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmAnalyticsEngineApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmAnalyticsEngineApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmAnalyticsEngineApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmAnalyticsEngineApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmAnalyticsEngineApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmAnalyticsEngineApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmAnalyticsEngineApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmAnalyticsEngineApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				})
				err := ibmAnalyticsEngineApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmAnalyticsEngineApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmAnalyticsEngineApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmAnalyticsEngineApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmAnalyticsEngineApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_ANALYTICS_ENGINE_API_URL": "https://ibmanalyticsengineapiv2/api",
				"IBM_ANALYTICS_ENGINE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmAnalyticsEngineApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_ANALYTICS_ENGINE_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2UsingExternalConfig(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmAnalyticsEngineApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmanalyticsengineapiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions)`, func() {
		getAllAnalyticsEnginesPath := "/v2/analytics_engines"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllAnalyticsEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAllAnalyticsEngines successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmAnalyticsEngineApiService.GetAllAnalyticsEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetAllAnalyticsEnginesOptions model
				getAllAnalyticsEnginesOptionsModel := new(ibmanalyticsengineapiv2.GetAllAnalyticsEnginesOptions)
				getAllAnalyticsEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmAnalyticsEngineApiService.GetAllAnalyticsEngines(getAllAnalyticsEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetAllAnalyticsEngines with error: Operation request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAllAnalyticsEnginesOptions model
				getAllAnalyticsEnginesOptionsModel := new(ibmanalyticsengineapiv2.GetAllAnalyticsEnginesOptions)
				getAllAnalyticsEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmAnalyticsEngineApiService.GetAllAnalyticsEngines(getAllAnalyticsEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions) - Operation response error`, func() {
		getAnalyticsEngineByIDPath := "/v2/analytics_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAnalyticsEngineByID with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {
		getAnalyticsEngineByIDPath := "/v2/analytics_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "service_plan": "ServicePlan", "hardware_size": "HardwareSize", "software_package": "SoftwarePackage", "domain": "Domain", "creation_time": "2019-01-01T12:00:00.000Z", "commision_time": "2019-01-01T12:00:00.000Z", "decommision_time": "2019-01-01T12:00:00.000Z", "deletion_time": "2019-01-01T12:00:00.000Z", "state_change_time": "2019-01-01T12:00:00.000Z", "state": "State", "nodes": [{"id": 2, "fqdn": "Fqdn", "type": "Type", "state": "State", "public_ip": "PublicIp", "private_ip": "PrivateIp", "state_change_time": "2019-01-01T12:00:00.000Z", "commission_time": "2019-01-01T12:00:00.000Z"}], "user_credentials": {"user": "User"}, "service_endpoints": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}, "service_endpoints_ip": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}, "private_endpoint_whitelist": ["PrivateEndpointWhitelist"]}`)
				}))
			})
			It(`Invoke GetAnalyticsEngineByID successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineByIDWithContext(ctx, getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineByIDWithContext(ctx, getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "service_plan": "ServicePlan", "hardware_size": "HardwareSize", "software_package": "SoftwarePackage", "domain": "Domain", "creation_time": "2019-01-01T12:00:00.000Z", "commision_time": "2019-01-01T12:00:00.000Z", "decommision_time": "2019-01-01T12:00:00.000Z", "deletion_time": "2019-01-01T12:00:00.000Z", "state_change_time": "2019-01-01T12:00:00.000Z", "state": "State", "nodes": [{"id": 2, "fqdn": "Fqdn", "type": "Type", "state": "State", "public_ip": "PublicIp", "private_ip": "PrivateIp", "state_change_time": "2019-01-01T12:00:00.000Z", "commission_time": "2019-01-01T12:00:00.000Z"}], "user_credentials": {"user": "User"}, "service_endpoints": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}, "service_endpoints_ip": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}, "private_endpoint_whitelist": ["PrivateEndpointWhitelist"]}`)
				}))
			})
			It(`Invoke GetAnalyticsEngineByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAnalyticsEngineByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAnalyticsEngineByIdOptions model with no property values
				getAnalyticsEngineByIdOptionsModelNew := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions) - Operation response error`, func() {
		getAnalyticsEngineStateByIDPath := "/v2/analytics_engines/testString/state"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineStateByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAnalyticsEngineStateByID with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineStateByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions)`, func() {
		getAnalyticsEngineStateByIDPath := "/v2/analytics_engines/testString/state"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineStateByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"state": "State"}`)
				}))
			})
			It(`Invoke GetAnalyticsEngineStateByID successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineStateByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByIDWithContext(ctx, getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByIDWithContext(ctx, getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAnalyticsEngineStateByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"state": "State"}`)
				}))
			})
			It(`Invoke GetAnalyticsEngineStateByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineStateByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAnalyticsEngineStateByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAnalyticsEngineStateByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAnalyticsEngineStateByIdOptions model with no property values
				getAnalyticsEngineStateByIdOptionsModelNew := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions) - Operation response error`, func() {
		createCustomizationRequestPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomizationRequestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomizationRequest with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the AnalyticsEngineCustomAction model
				analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
				analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
				analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
				analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
				analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}

				// Construct an instance of the CreateCustomizationRequestOptions model
				createCustomizationRequestOptionsModel := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
				createCustomizationRequestOptionsModel.InstanceGuid = core.StringPtr("testString")
				createCustomizationRequestOptionsModel.Target = core.StringPtr("all")
				createCustomizationRequestOptionsModel.CustomActions = []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}
				createCustomizationRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions)`, func() {
		createCustomizationRequestPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomizationRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"request_id": 9}`)
				}))
			})
			It(`Invoke CreateCustomizationRequest successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the AnalyticsEngineCustomAction model
				analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
				analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
				analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
				analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
				analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}

				// Construct an instance of the CreateCustomizationRequestOptions model
				createCustomizationRequestOptionsModel := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
				createCustomizationRequestOptionsModel.InstanceGuid = core.StringPtr("testString")
				createCustomizationRequestOptionsModel.Target = core.StringPtr("all")
				createCustomizationRequestOptionsModel.CustomActions = []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}
				createCustomizationRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.CreateCustomizationRequestWithContext(ctx, createCustomizationRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.CreateCustomizationRequestWithContext(ctx, createCustomizationRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomizationRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"request_id": 9}`)
				}))
			})
			It(`Invoke CreateCustomizationRequest successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateCustomizationRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the AnalyticsEngineCustomAction model
				analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
				analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
				analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
				analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
				analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}

				// Construct an instance of the CreateCustomizationRequestOptions model
				createCustomizationRequestOptionsModel := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
				createCustomizationRequestOptionsModel.InstanceGuid = core.StringPtr("testString")
				createCustomizationRequestOptionsModel.Target = core.StringPtr("all")
				createCustomizationRequestOptionsModel.CustomActions = []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}
				createCustomizationRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomizationRequest with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the AnalyticsEngineCustomAction model
				analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
				analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
				analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
				analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
				analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}

				// Construct an instance of the CreateCustomizationRequestOptions model
				createCustomizationRequestOptionsModel := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
				createCustomizationRequestOptionsModel.InstanceGuid = core.StringPtr("testString")
				createCustomizationRequestOptionsModel.Target = core.StringPtr("all")
				createCustomizationRequestOptionsModel.CustomActions = []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}
				createCustomizationRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomizationRequestOptions model with no property values
				createCustomizationRequestOptionsModelNew := new(ibmanalyticsengineapiv2.CreateCustomizationRequestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateCustomizationRequest(createCustomizationRequestOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions) - Operation response error`, func() {
		getAllCustomizationRequestsPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllCustomizationRequestsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAllCustomizationRequests with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAllCustomizationRequestsOptions model
				getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAllCustomizationRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions)`, func() {
		getAllCustomizationRequestsPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllCustomizationRequestsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"id": "ID"}]`)
				}))
			})
			It(`Invoke GetAllCustomizationRequests successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetAllCustomizationRequestsOptions model
				getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAllCustomizationRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetAllCustomizationRequestsWithContext(ctx, getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetAllCustomizationRequestsWithContext(ctx, getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllCustomizationRequestsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"id": "ID"}]`)
				}))
			})
			It(`Invoke GetAllCustomizationRequests successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllCustomizationRequestsOptions model
				getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAllCustomizationRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAllCustomizationRequests with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetAllCustomizationRequestsOptions model
				getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr("testString")
				getAllCustomizationRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAllCustomizationRequestsOptions model with no property values
				getAllCustomizationRequestsOptionsModelNew := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions) - Operation response error`, func() {
		getCustomizationRequestByIDPath := "/v2/analytics_engines/testString/customization_requests/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomizationRequestByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomizationRequestByID with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomizationRequestByIdOptions model
				getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions)`, func() {
		getCustomizationRequestByIDPath := "/v2/analytics_engines/testString/customization_requests/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomizationRequestByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "run_status": "RunStatus", "run_details": {"overall_status": "OverallStatus", "details": [{"node_name": "NodeName", "node_type": "NodeType", "start_time": "StartTime", "end_time": "EndTime", "time_taken": "TimeTaken", "status": "Status", "log_file": "LogFile"}]}}`)
				}))
			})
			It(`Invoke GetCustomizationRequestByID successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomizationRequestByIdOptions model
				getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetCustomizationRequestByIDWithContext(ctx, getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetCustomizationRequestByIDWithContext(ctx, getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomizationRequestByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "run_status": "RunStatus", "run_details": {"overall_status": "OverallStatus", "details": [{"node_name": "NodeName", "node_type": "NodeType", "start_time": "StartTime", "end_time": "EndTime", "time_taken": "TimeTaken", "status": "Status", "log_file": "LogFile"}]}}`)
				}))
			})
			It(`Invoke GetCustomizationRequestByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomizationRequestByIdOptions model
				getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomizationRequestByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomizationRequestByIdOptions model
				getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCustomizationRequestByIdOptions model with no property values
				getCustomizationRequestByIdOptionsModelNew := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResizeCluster(resizeClusterOptions *ResizeClusterOptions) - Operation response error`, func() {
		resizeClusterPath := "/v2/analytics_engines/testString/resize"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resizeClusterPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResizeCluster with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest model
				resizeClusterRequestModel := new(ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
				resizeClusterRequestModel.ComputeNodesCount = core.Int64Ptr(int64(38))

				// Construct an instance of the ResizeClusterOptions model
				resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				resizeClusterOptionsModel.InstanceGuid = core.StringPtr("testString")
				resizeClusterOptionsModel.Body = resizeClusterRequestModel
				resizeClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResizeCluster(resizeClusterOptions *ResizeClusterOptions)`, func() {
		resizeClusterPath := "/v2/analytics_engines/testString/resize"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resizeClusterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"request_id": "RequestID"}`)
				}))
			})
			It(`Invoke ResizeCluster successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest model
				resizeClusterRequestModel := new(ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
				resizeClusterRequestModel.ComputeNodesCount = core.Int64Ptr(int64(38))

				// Construct an instance of the ResizeClusterOptions model
				resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				resizeClusterOptionsModel.InstanceGuid = core.StringPtr("testString")
				resizeClusterOptionsModel.Body = resizeClusterRequestModel
				resizeClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ResizeClusterWithContext(ctx, resizeClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ResizeClusterWithContext(ctx, resizeClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resizeClusterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"request_id": "RequestID"}`)
				}))
			})
			It(`Invoke ResizeCluster successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ResizeCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest model
				resizeClusterRequestModel := new(ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
				resizeClusterRequestModel.ComputeNodesCount = core.Int64Ptr(int64(38))

				// Construct an instance of the ResizeClusterOptions model
				resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				resizeClusterOptionsModel.InstanceGuid = core.StringPtr("testString")
				resizeClusterOptionsModel.Body = resizeClusterRequestModel
				resizeClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResizeCluster with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest model
				resizeClusterRequestModel := new(ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
				resizeClusterRequestModel.ComputeNodesCount = core.Int64Ptr(int64(38))

				// Construct an instance of the ResizeClusterOptions model
				resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				resizeClusterOptionsModel.InstanceGuid = core.StringPtr("testString")
				resizeClusterOptionsModel.Body = resizeClusterRequestModel
				resizeClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResizeClusterOptions model with no property values
				resizeClusterOptionsModelNew := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResizeCluster(resizeClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions) - Operation response error`, func() {
		resetClusterPasswordPath := "/v2/analytics_engines/testString/reset_password"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetClusterPasswordPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResetClusterPassword with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ResetClusterPasswordOptions model
				resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr("testString")
				resetClusterPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions)`, func() {
		resetClusterPasswordPath := "/v2/analytics_engines/testString/reset_password"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetClusterPasswordPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "user_credentials": {"user": "User", "password": "Password"}}`)
				}))
			})
			It(`Invoke ResetClusterPassword successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ResetClusterPasswordOptions model
				resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr("testString")
				resetClusterPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ResetClusterPasswordWithContext(ctx, resetClusterPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ResetClusterPasswordWithContext(ctx, resetClusterPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetClusterPasswordPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "user_credentials": {"user": "User", "password": "Password"}}`)
				}))
			})
			It(`Invoke ResetClusterPassword successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ResetClusterPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResetClusterPasswordOptions model
				resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr("testString")
				resetClusterPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResetClusterPassword with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ResetClusterPasswordOptions model
				resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr("testString")
				resetClusterPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResetClusterPasswordOptions model with no property values
				resetClusterPasswordOptionsModelNew := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ResetClusterPassword(resetClusterPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions)`, func() {
		configureLoggingPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(configureLoggingPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(202)
				}))
			})
			It(`Invoke ConfigureLogging successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmAnalyticsEngineApiService.ConfigureLogging(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AnalyticsEngineLoggingNodeSpec model
				analyticsEngineLoggingNodeSpecModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec)
				analyticsEngineLoggingNodeSpecModel.NodeType = core.StringPtr("management")
				analyticsEngineLoggingNodeSpecModel.Components = []string{"ambari-server"}

				// Construct an instance of the AnalyticsEngineLoggingServer model
				analyticsEngineLoggingServerModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer)
				analyticsEngineLoggingServerModel.Type = core.StringPtr("logdna")
				analyticsEngineLoggingServerModel.Credential = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.ApiHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.LogHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.Owner = core.StringPtr("testString")

				// Construct an instance of the ConfigureLoggingOptions model
				configureLoggingOptionsModel := new(ibmanalyticsengineapiv2.ConfigureLoggingOptions)
				configureLoggingOptionsModel.InstanceGuid = core.StringPtr("testString")
				configureLoggingOptionsModel.LogSpecs = []ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel}
				configureLoggingOptionsModel.LogServer = analyticsEngineLoggingServerModel
				configureLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmAnalyticsEngineApiService.ConfigureLogging(configureLoggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ConfigureLogging with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the AnalyticsEngineLoggingNodeSpec model
				analyticsEngineLoggingNodeSpecModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec)
				analyticsEngineLoggingNodeSpecModel.NodeType = core.StringPtr("management")
				analyticsEngineLoggingNodeSpecModel.Components = []string{"ambari-server"}

				// Construct an instance of the AnalyticsEngineLoggingServer model
				analyticsEngineLoggingServerModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer)
				analyticsEngineLoggingServerModel.Type = core.StringPtr("logdna")
				analyticsEngineLoggingServerModel.Credential = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.ApiHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.LogHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.Owner = core.StringPtr("testString")

				// Construct an instance of the ConfigureLoggingOptions model
				configureLoggingOptionsModel := new(ibmanalyticsengineapiv2.ConfigureLoggingOptions)
				configureLoggingOptionsModel.InstanceGuid = core.StringPtr("testString")
				configureLoggingOptionsModel.LogSpecs = []ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel}
				configureLoggingOptionsModel.LogServer = analyticsEngineLoggingServerModel
				configureLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmAnalyticsEngineApiService.ConfigureLogging(configureLoggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ConfigureLoggingOptions model with no property values
				configureLoggingOptionsModelNew := new(ibmanalyticsengineapiv2.ConfigureLoggingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmAnalyticsEngineApiService.ConfigureLogging(configureLoggingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions) - Operation response error`, func() {
		getLoggingConfigPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoggingConfig with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLoggingConfigOptions model
				getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				getLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions)`, func() {
		getLoggingConfigPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"log_specs": [{"node_type": "management", "components": ["ambari-server"]}], "log_server": {"type": "logdna", "credential": "Credential", "api_host": "ApiHost", "log_host": "LogHost", "owner": "Owner"}, "log_config_status": [{"node_type": "management", "node_id": "NodeID", "action": "Action", "status": "Status"}]}`)
				}))
			})
			It(`Invoke GetLoggingConfig successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLoggingConfigOptions model
				getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				getLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfigWithContext(ctx, getLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfigWithContext(ctx, getLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"log_specs": [{"node_type": "management", "components": ["ambari-server"]}], "log_server": {"type": "logdna", "credential": "Credential", "api_host": "ApiHost", "log_host": "LogHost", "owner": "Owner"}, "log_config_status": [{"node_type": "management", "node_id": "NodeID", "action": "Action", "status": "Status"}]}`)
				}))
			})
			It(`Invoke GetLoggingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoggingConfigOptions model
				getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				getLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLoggingConfig with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLoggingConfigOptions model
				getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				getLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoggingConfigOptions model with no property values
				getLoggingConfigOptionsModelNew := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfig(getLoggingConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions)`, func() {
		deleteLoggingConfigPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoggingConfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteLoggingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmAnalyticsEngineApiService.DeleteLoggingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoggingConfigOptions model
				deleteLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.DeleteLoggingConfigOptions)
				deleteLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				deleteLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteLoggingConfig(deleteLoggingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLoggingConfig with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLoggingConfigOptions model
				deleteLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.DeleteLoggingConfigOptions)
				deleteLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")
				deleteLoggingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmAnalyticsEngineApiService.DeleteLoggingConfig(deleteLoggingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLoggingConfigOptions model with no property values
				deleteLoggingConfigOptionsModelNew := new(ibmanalyticsengineapiv2.DeleteLoggingConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteLoggingConfig(deleteLoggingConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions) - Operation response error`, func() {
		updatePrivateEndpointWhitelistPath := "/v2/analytics_engines/testString/private_endpoint_whitelist"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePrivateEndpointWhitelistPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePrivateEndpointWhitelist with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
				updatePrivateEndpointWhitelistOptionsModel := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
				updatePrivateEndpointWhitelistOptionsModel.InstanceGuid = core.StringPtr("testString")
				updatePrivateEndpointWhitelistOptionsModel.IpRanges = []string{"testString"}
				updatePrivateEndpointWhitelistOptionsModel.Action = core.StringPtr("add")
				updatePrivateEndpointWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions)`, func() {
		updatePrivateEndpointWhitelistPath := "/v2/analytics_engines/testString/private_endpoint_whitelist"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePrivateEndpointWhitelistPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"private_endpoint_whitelist": ["PrivateEndpointWhitelist"]}`)
				}))
			})
			It(`Invoke UpdatePrivateEndpointWhitelist successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
				updatePrivateEndpointWhitelistOptionsModel := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
				updatePrivateEndpointWhitelistOptionsModel.InstanceGuid = core.StringPtr("testString")
				updatePrivateEndpointWhitelistOptionsModel.IpRanges = []string{"testString"}
				updatePrivateEndpointWhitelistOptionsModel.Action = core.StringPtr("add")
				updatePrivateEndpointWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelistWithContext(ctx, updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelistWithContext(ctx, updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePrivateEndpointWhitelistPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"private_endpoint_whitelist": ["PrivateEndpointWhitelist"]}`)
				}))
			})
			It(`Invoke UpdatePrivateEndpointWhitelist successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
				updatePrivateEndpointWhitelistOptionsModel := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
				updatePrivateEndpointWhitelistOptionsModel.InstanceGuid = core.StringPtr("testString")
				updatePrivateEndpointWhitelistOptionsModel.IpRanges = []string{"testString"}
				updatePrivateEndpointWhitelistOptionsModel.Action = core.StringPtr("add")
				updatePrivateEndpointWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdatePrivateEndpointWhitelist with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
				updatePrivateEndpointWhitelistOptionsModel := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
				updatePrivateEndpointWhitelistOptionsModel.InstanceGuid = core.StringPtr("testString")
				updatePrivateEndpointWhitelistOptionsModel.IpRanges = []string{"testString"}
				updatePrivateEndpointWhitelistOptionsModel.Action = core.StringPtr("add")
				updatePrivateEndpointWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePrivateEndpointWhitelistOptions model with no property values
				updatePrivateEndpointWhitelistOptionsModelNew := new(ibmanalyticsengineapiv2.UpdatePrivateEndpointWhitelistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmAnalyticsEngineApiService, _ := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				URL:           "http://ibmanalyticsengineapiv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAnalyticsEngineCustomAction successfully`, func() {
				name := "testString"
				_model, err := ibmAnalyticsEngineApiService.NewAnalyticsEngineCustomAction(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAnalyticsEngineLoggingNodeSpec successfully`, func() {
				nodeType := "management"
				components := []string{"ambari-server"}
				_model, err := ibmAnalyticsEngineApiService.NewAnalyticsEngineLoggingNodeSpec(nodeType, components)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAnalyticsEngineLoggingServer successfully`, func() {
				typeVar := "logdna"
				credential := "testString"
				apiHost := "testString"
				logHost := "testString"
				_model, err := ibmAnalyticsEngineApiService.NewAnalyticsEngineLoggingServer(typeVar, credential, apiHost, logHost)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigureLoggingOptions successfully`, func() {
				// Construct an instance of the AnalyticsEngineLoggingNodeSpec model
				analyticsEngineLoggingNodeSpecModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec)
				Expect(analyticsEngineLoggingNodeSpecModel).ToNot(BeNil())
				analyticsEngineLoggingNodeSpecModel.NodeType = core.StringPtr("management")
				analyticsEngineLoggingNodeSpecModel.Components = []string{"ambari-server"}
				Expect(analyticsEngineLoggingNodeSpecModel.NodeType).To(Equal(core.StringPtr("management")))
				Expect(analyticsEngineLoggingNodeSpecModel.Components).To(Equal([]string{"ambari-server"}))

				// Construct an instance of the AnalyticsEngineLoggingServer model
				analyticsEngineLoggingServerModel := new(ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer)
				Expect(analyticsEngineLoggingServerModel).ToNot(BeNil())
				analyticsEngineLoggingServerModel.Type = core.StringPtr("logdna")
				analyticsEngineLoggingServerModel.Credential = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.ApiHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.LogHost = core.StringPtr("testString")
				analyticsEngineLoggingServerModel.Owner = core.StringPtr("testString")
				Expect(analyticsEngineLoggingServerModel.Type).To(Equal(core.StringPtr("logdna")))
				Expect(analyticsEngineLoggingServerModel.Credential).To(Equal(core.StringPtr("testString")))
				Expect(analyticsEngineLoggingServerModel.ApiHost).To(Equal(core.StringPtr("testString")))
				Expect(analyticsEngineLoggingServerModel.LogHost).To(Equal(core.StringPtr("testString")))
				Expect(analyticsEngineLoggingServerModel.Owner).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigureLoggingOptions model
				instanceGuid := "testString"
				configureLoggingOptionsLogSpecs := []ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{}
				var configureLoggingOptionsLogServer *ibmanalyticsengineapiv2.AnalyticsEngineLoggingServer = nil
				configureLoggingOptionsModel := ibmAnalyticsEngineApiService.NewConfigureLoggingOptions(instanceGuid, configureLoggingOptionsLogSpecs, configureLoggingOptionsLogServer)
				configureLoggingOptionsModel.SetInstanceGuid("testString")
				configureLoggingOptionsModel.SetLogSpecs([]ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel})
				configureLoggingOptionsModel.SetLogServer(analyticsEngineLoggingServerModel)
				configureLoggingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(configureLoggingOptionsModel).ToNot(BeNil())
				Expect(configureLoggingOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(configureLoggingOptionsModel.LogSpecs).To(Equal([]ibmanalyticsengineapiv2.AnalyticsEngineLoggingNodeSpec{*analyticsEngineLoggingNodeSpecModel}))
				Expect(configureLoggingOptionsModel.LogServer).To(Equal(analyticsEngineLoggingServerModel))
				Expect(configureLoggingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomizationRequestOptions successfully`, func() {
				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				Expect(analyticsEngineCustomActionScriptModel).ToNot(BeNil())
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = map[string]interface{}{"anyKey": "anyValue"}
				Expect(analyticsEngineCustomActionScriptModel.SourceType).To(Equal(core.StringPtr("http")))
				Expect(analyticsEngineCustomActionScriptModel.ScriptPath).To(Equal(core.StringPtr("testString")))
				Expect(analyticsEngineCustomActionScriptModel.SourceProps).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the AnalyticsEngineCustomAction model
				analyticsEngineCustomActionModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomAction)
				Expect(analyticsEngineCustomActionModel).ToNot(BeNil())
				analyticsEngineCustomActionModel.Name = core.StringPtr("testString")
				analyticsEngineCustomActionModel.Type = core.StringPtr("bootstrap")
				analyticsEngineCustomActionModel.Script = analyticsEngineCustomActionScriptModel
				analyticsEngineCustomActionModel.ScriptParams = []string{"testString"}
				Expect(analyticsEngineCustomActionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(analyticsEngineCustomActionModel.Type).To(Equal(core.StringPtr("bootstrap")))
				Expect(analyticsEngineCustomActionModel.Script).To(Equal(analyticsEngineCustomActionScriptModel))
				Expect(analyticsEngineCustomActionModel.ScriptParams).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateCustomizationRequestOptions model
				instanceGuid := "testString"
				createCustomizationRequestOptionsTarget := "all"
				createCustomizationRequestOptionsCustomActions := []ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{}
				createCustomizationRequestOptionsModel := ibmAnalyticsEngineApiService.NewCreateCustomizationRequestOptions(instanceGuid, createCustomizationRequestOptionsTarget, createCustomizationRequestOptionsCustomActions)
				createCustomizationRequestOptionsModel.SetInstanceGuid("testString")
				createCustomizationRequestOptionsModel.SetTarget("all")
				createCustomizationRequestOptionsModel.SetCustomActions([]ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel})
				createCustomizationRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomizationRequestOptionsModel).ToNot(BeNil())
				Expect(createCustomizationRequestOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(createCustomizationRequestOptionsModel.Target).To(Equal(core.StringPtr("all")))
				Expect(createCustomizationRequestOptionsModel.CustomActions).To(Equal([]ibmanalyticsengineapiv2.AnalyticsEngineCustomAction{*analyticsEngineCustomActionModel}))
				Expect(createCustomizationRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoggingConfigOptions successfully`, func() {
				// Construct an instance of the DeleteLoggingConfigOptions model
				instanceGuid := "testString"
				deleteLoggingConfigOptionsModel := ibmAnalyticsEngineApiService.NewDeleteLoggingConfigOptions(instanceGuid)
				deleteLoggingConfigOptionsModel.SetInstanceGuid("testString")
				deleteLoggingConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoggingConfigOptionsModel).ToNot(BeNil())
				Expect(deleteLoggingConfigOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoggingConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllAnalyticsEnginesOptions successfully`, func() {
				// Construct an instance of the GetAllAnalyticsEnginesOptions model
				getAllAnalyticsEnginesOptionsModel := ibmAnalyticsEngineApiService.NewGetAllAnalyticsEnginesOptions()
				getAllAnalyticsEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllAnalyticsEnginesOptionsModel).ToNot(BeNil())
				Expect(getAllAnalyticsEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllCustomizationRequestsOptions successfully`, func() {
				// Construct an instance of the GetAllCustomizationRequestsOptions model
				instanceGuid := "testString"
				getAllCustomizationRequestsOptionsModel := ibmAnalyticsEngineApiService.NewGetAllCustomizationRequestsOptions(instanceGuid)
				getAllCustomizationRequestsOptionsModel.SetInstanceGuid("testString")
				getAllCustomizationRequestsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllCustomizationRequestsOptionsModel).ToNot(BeNil())
				Expect(getAllCustomizationRequestsOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(getAllCustomizationRequestsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAnalyticsEngineByIdOptions successfully`, func() {
				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				instanceGuid := "testString"
				getAnalyticsEngineByIdOptionsModel := ibmAnalyticsEngineApiService.NewGetAnalyticsEngineByIdOptions(instanceGuid)
				getAnalyticsEngineByIdOptionsModel.SetInstanceGuid("testString")
				getAnalyticsEngineByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAnalyticsEngineByIdOptionsModel).ToNot(BeNil())
				Expect(getAnalyticsEngineByIdOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(getAnalyticsEngineByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAnalyticsEngineStateByIdOptions successfully`, func() {
				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				instanceGuid := "testString"
				getAnalyticsEngineStateByIdOptionsModel := ibmAnalyticsEngineApiService.NewGetAnalyticsEngineStateByIdOptions(instanceGuid)
				getAnalyticsEngineStateByIdOptionsModel.SetInstanceGuid("testString")
				getAnalyticsEngineStateByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAnalyticsEngineStateByIdOptionsModel).ToNot(BeNil())
				Expect(getAnalyticsEngineStateByIdOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(getAnalyticsEngineStateByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomizationRequestByIdOptions successfully`, func() {
				// Construct an instance of the GetCustomizationRequestByIdOptions model
				instanceGuid := "testString"
				requestID := "testString"
				getCustomizationRequestByIdOptionsModel := ibmAnalyticsEngineApiService.NewGetCustomizationRequestByIdOptions(instanceGuid, requestID)
				getCustomizationRequestByIdOptionsModel.SetInstanceGuid("testString")
				getCustomizationRequestByIdOptionsModel.SetRequestID("testString")
				getCustomizationRequestByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomizationRequestByIdOptionsModel).ToNot(BeNil())
				Expect(getCustomizationRequestByIdOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(getCustomizationRequestByIdOptionsModel.RequestID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomizationRequestByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoggingConfigOptions successfully`, func() {
				// Construct an instance of the GetLoggingConfigOptions model
				instanceGuid := "testString"
				getLoggingConfigOptionsModel := ibmAnalyticsEngineApiService.NewGetLoggingConfigOptions(instanceGuid)
				getLoggingConfigOptionsModel.SetInstanceGuid("testString")
				getLoggingConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoggingConfigOptionsModel).ToNot(BeNil())
				Expect(getLoggingConfigOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(getLoggingConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResetClusterPasswordOptions successfully`, func() {
				// Construct an instance of the ResetClusterPasswordOptions model
				instanceGuid := "testString"
				resetClusterPasswordOptionsModel := ibmAnalyticsEngineApiService.NewResetClusterPasswordOptions(instanceGuid)
				resetClusterPasswordOptionsModel.SetInstanceGuid("testString")
				resetClusterPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resetClusterPasswordOptionsModel).ToNot(BeNil())
				Expect(resetClusterPasswordOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(resetClusterPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResizeClusterOptions successfully`, func() {
				// Construct an instance of the ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest model
				resizeClusterRequestModel := new(ibmanalyticsengineapiv2.ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
				Expect(resizeClusterRequestModel).ToNot(BeNil())
				resizeClusterRequestModel.ComputeNodesCount = core.Int64Ptr(int64(38))
				Expect(resizeClusterRequestModel.ComputeNodesCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the ResizeClusterOptions model
				instanceGuid := "testString"
				var body ibmanalyticsengineapiv2.ResizeClusterRequestIntf = nil
				resizeClusterOptionsModel := ibmAnalyticsEngineApiService.NewResizeClusterOptions(instanceGuid, body)
				resizeClusterOptionsModel.SetInstanceGuid("testString")
				resizeClusterOptionsModel.SetBody(resizeClusterRequestModel)
				resizeClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resizeClusterOptionsModel).ToNot(BeNil())
				Expect(resizeClusterOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(resizeClusterOptionsModel.Body).To(Equal(resizeClusterRequestModel))
				Expect(resizeClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePrivateEndpointWhitelistOptions successfully`, func() {
				// Construct an instance of the UpdatePrivateEndpointWhitelistOptions model
				instanceGuid := "testString"
				updatePrivateEndpointWhitelistOptionsIpRanges := []string{"testString"}
				updatePrivateEndpointWhitelistOptionsAction := "add"
				updatePrivateEndpointWhitelistOptionsModel := ibmAnalyticsEngineApiService.NewUpdatePrivateEndpointWhitelistOptions(instanceGuid, updatePrivateEndpointWhitelistOptionsIpRanges, updatePrivateEndpointWhitelistOptionsAction)
				updatePrivateEndpointWhitelistOptionsModel.SetInstanceGuid("testString")
				updatePrivateEndpointWhitelistOptionsModel.SetIpRanges([]string{"testString"})
				updatePrivateEndpointWhitelistOptionsModel.SetAction("add")
				updatePrivateEndpointWhitelistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePrivateEndpointWhitelistOptionsModel).ToNot(BeNil())
				Expect(updatePrivateEndpointWhitelistOptionsModel.InstanceGuid).To(Equal(core.StringPtr("testString")))
				Expect(updatePrivateEndpointWhitelistOptionsModel.IpRanges).To(Equal([]string{"testString"}))
				Expect(updatePrivateEndpointWhitelistOptionsModel.Action).To(Equal(core.StringPtr("add")))
				Expect(updatePrivateEndpointWhitelistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
