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
	"github.com/IBM/ibm-iae-go-sdk/ibmanalyticsengineapiv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IbmAnalyticsEngineApiV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmAnalyticsEngineApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
				URL: "https://ibmanalyticsengineapiv3/api",
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
				"IBM_ANALYTICS_ENGINE_API_URL": "https://ibmanalyticsengineapiv3/api",
				"IBM_ANALYTICS_ENGINE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
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
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
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
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
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
				"IBM_ANALYTICS_ENGINE_API_URL": "https://ibmanalyticsengineapiv3/api",
				"IBM_ANALYTICS_ENGINE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
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
			ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3UsingExternalConfig(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
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
			url, err = ibmanalyticsengineapiv3.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://api.us-south.ae.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = ibmanalyticsengineapiv3.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://api.eu-de.ae.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = ibmanalyticsengineapiv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions) - Operation response error`, func() {
		getInstancePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstance with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {
		getInstancePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "state": "created", "state_change_time": "2021-01-30T08:30:00.000Z", "default_runtime": {"spark_version": "SparkVersion"}, "instance_home": {"id": "ID", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "bucket": "Bucket", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}, "default_config": {"key": "Key"}}`)
				}))
			})
			It(`Invoke GetInstance successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetInstanceWithContext(ctx, getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetInstanceWithContext(ctx, getInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "href": "Href", "state": "created", "state_change_time": "2021-01-30T08:30:00.000Z", "default_runtime": {"spark_version": "SparkVersion"}, "instance_home": {"id": "ID", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "bucket": "Bucket", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}, "default_config": {"key": "Key"}}`)
				}))
			})
			It(`Invoke GetInstance successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstance with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceOptions model with no property values
				getInstanceOptionsModelNew := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInstance successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceOptions model
				getInstanceOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceOptions)
				getInstanceOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstance(getInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceState(getInstanceStateOptions *GetInstanceStateOptions) - Operation response error`, func() {
		getInstanceStatePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/state"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceState with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceStateOptions model
				getInstanceStateOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				getInstanceStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceState(getInstanceStateOptions *GetInstanceStateOptions)`, func() {
		getInstanceStatePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/state"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "created"}`)
				}))
			})
			It(`Invoke GetInstanceState successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceStateOptions model
				getInstanceStateOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				getInstanceStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetInstanceStateWithContext(ctx, getInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetInstanceStateWithContext(ctx, getInstanceStateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "created"}`)
				}))
			})
			It(`Invoke GetInstanceState successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceStateOptions model
				getInstanceStateOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				getInstanceStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceState with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceStateOptions model
				getInstanceStateOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				getInstanceStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceStateOptions model with no property values
				getInstanceStateOptionsModelNew := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInstanceState successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceStateOptions model
				getInstanceStateOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceStateOptions)
				getInstanceStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceState(getInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetInstanceHome(setInstanceHomeOptions *SetInstanceHomeOptions) - Operation response error`, func() {
		setInstanceHomePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/instance_home"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setInstanceHomePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetInstanceHome with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the SetInstanceHomeOptions model
				setInstanceHomeOptionsModel := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				setInstanceHomeOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.NewInstanceID = core.StringPtr("testString")
				setInstanceHomeOptionsModel.NewProvider = core.StringPtr("ibm-cos")
				setInstanceHomeOptionsModel.NewType = core.StringPtr("objectstore")
				setInstanceHomeOptionsModel.NewRegion = core.StringPtr("us-south")
				setInstanceHomeOptionsModel.NewEndpoint = core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.NewHmacAccessKey = core.StringPtr("821**********0ae")
				setInstanceHomeOptionsModel.NewHmacSecretKey = core.StringPtr("03e****************4fc3")
				setInstanceHomeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetInstanceHome(setInstanceHomeOptions *SetInstanceHomeOptions)`, func() {
		setInstanceHomePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/instance_home"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setInstanceHomePath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance_id": "InstanceID", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}`)
				}))
			})
			It(`Invoke SetInstanceHome successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the SetInstanceHomeOptions model
				setInstanceHomeOptionsModel := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				setInstanceHomeOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.NewInstanceID = core.StringPtr("testString")
				setInstanceHomeOptionsModel.NewProvider = core.StringPtr("ibm-cos")
				setInstanceHomeOptionsModel.NewType = core.StringPtr("objectstore")
				setInstanceHomeOptionsModel.NewRegion = core.StringPtr("us-south")
				setInstanceHomeOptionsModel.NewEndpoint = core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.NewHmacAccessKey = core.StringPtr("821**********0ae")
				setInstanceHomeOptionsModel.NewHmacSecretKey = core.StringPtr("03e****************4fc3")
				setInstanceHomeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.SetInstanceHomeWithContext(ctx, setInstanceHomeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.SetInstanceHomeWithContext(ctx, setInstanceHomeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setInstanceHomePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance_id": "InstanceID", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}`)
				}))
			})
			It(`Invoke SetInstanceHome successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.SetInstanceHome(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetInstanceHomeOptions model
				setInstanceHomeOptionsModel := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				setInstanceHomeOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.NewInstanceID = core.StringPtr("testString")
				setInstanceHomeOptionsModel.NewProvider = core.StringPtr("ibm-cos")
				setInstanceHomeOptionsModel.NewType = core.StringPtr("objectstore")
				setInstanceHomeOptionsModel.NewRegion = core.StringPtr("us-south")
				setInstanceHomeOptionsModel.NewEndpoint = core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.NewHmacAccessKey = core.StringPtr("821**********0ae")
				setInstanceHomeOptionsModel.NewHmacSecretKey = core.StringPtr("03e****************4fc3")
				setInstanceHomeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetInstanceHome with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the SetInstanceHomeOptions model
				setInstanceHomeOptionsModel := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				setInstanceHomeOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.NewInstanceID = core.StringPtr("testString")
				setInstanceHomeOptionsModel.NewProvider = core.StringPtr("ibm-cos")
				setInstanceHomeOptionsModel.NewType = core.StringPtr("objectstore")
				setInstanceHomeOptionsModel.NewRegion = core.StringPtr("us-south")
				setInstanceHomeOptionsModel.NewEndpoint = core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.NewHmacAccessKey = core.StringPtr("821**********0ae")
				setInstanceHomeOptionsModel.NewHmacSecretKey = core.StringPtr("03e****************4fc3")
				setInstanceHomeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetInstanceHomeOptions model with no property values
				setInstanceHomeOptionsModelNew := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetInstanceHome successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the SetInstanceHomeOptions model
				setInstanceHomeOptionsModel := new(ibmanalyticsengineapiv3.SetInstanceHomeOptions)
				setInstanceHomeOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.NewInstanceID = core.StringPtr("testString")
				setInstanceHomeOptionsModel.NewProvider = core.StringPtr("ibm-cos")
				setInstanceHomeOptionsModel.NewType = core.StringPtr("objectstore")
				setInstanceHomeOptionsModel.NewRegion = core.StringPtr("us-south")
				setInstanceHomeOptionsModel.NewEndpoint = core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.NewHmacAccessKey = core.StringPtr("821**********0ae")
				setInstanceHomeOptionsModel.NewHmacSecretKey = core.StringPtr("03e****************4fc3")
				setInstanceHomeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.SetInstanceHome(setInstanceHomeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptions *GetInstanceDefaultConfigsOptions)`, func() {
		getInstanceDefaultConfigsPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/default_configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceDefaultConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke GetInstanceDefaultConfigs successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceDefaultConfigsOptions model
				getInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions)
				getInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigsWithContext(ctx, getInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetInstanceDefaultConfigsWithContext(ctx, getInstanceDefaultConfigsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceDefaultConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke GetInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceDefaultConfigsOptions model
				getInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions)
				getInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceDefaultConfigs with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceDefaultConfigsOptions model
				getInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions)
				getInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceDefaultConfigsOptions model with no property values
				getInstanceDefaultConfigsOptionsModelNew := new(ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceDefaultConfigsOptions model
				getInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceDefaultConfigsOptions)
				getInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptions *ReplaceInstanceDefaultConfigsOptions)`, func() {
		replaceInstanceDefaultConfigsPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/default_configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceInstanceDefaultConfigsPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke ReplaceInstanceDefaultConfigs successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceInstanceDefaultConfigsOptions model
				replaceInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions)
				replaceInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceInstanceDefaultConfigsOptionsModel.Body = make(map[string]string)
				replaceInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigsWithContext(ctx, replaceInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigsWithContext(ctx, replaceInstanceDefaultConfigsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceInstanceDefaultConfigsPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke ReplaceInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceInstanceDefaultConfigsOptions model
				replaceInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions)
				replaceInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceInstanceDefaultConfigsOptionsModel.Body = make(map[string]string)
				replaceInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceInstanceDefaultConfigs with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceInstanceDefaultConfigsOptions model
				replaceInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions)
				replaceInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceInstanceDefaultConfigsOptionsModel.Body = make(map[string]string)
				replaceInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceInstanceDefaultConfigsOptions model with no property values
				replaceInstanceDefaultConfigsOptionsModelNew := new(ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceInstanceDefaultConfigsOptions model
				replaceInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.ReplaceInstanceDefaultConfigsOptions)
				replaceInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceInstanceDefaultConfigsOptionsModel.Body = make(map[string]string)
				replaceInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions *UpdateInstanceDefaultConfigsOptions)`, func() {
		updateInstanceDefaultConfigsPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/default_configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceDefaultConfigsPath))
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
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke UpdateInstanceDefaultConfigs successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdateInstanceDefaultConfigsOptions model
				updateInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions)
				updateInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				updateInstanceDefaultConfigsOptionsModel.Body = make(map[string]interface{})
				updateInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigsWithContext(ctx, updateInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigsWithContext(ctx, updateInstanceDefaultConfigsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceDefaultConfigsPath))
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
					fmt.Fprintf(res, "%s", `{"mapKey": "Inner"}`)
				}))
			})
			It(`Invoke UpdateInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateInstanceDefaultConfigsOptions model
				updateInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions)
				updateInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				updateInstanceDefaultConfigsOptionsModel.Body = make(map[string]interface{})
				updateInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceDefaultConfigs with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceDefaultConfigsOptions model
				updateInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions)
				updateInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				updateInstanceDefaultConfigsOptionsModel.Body = make(map[string]interface{})
				updateInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceDefaultConfigsOptions model with no property values
				updateInstanceDefaultConfigsOptionsModelNew := new(ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateInstanceDefaultConfigs successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceDefaultConfigsOptions model
				updateInstanceDefaultConfigsOptionsModel := new(ibmanalyticsengineapiv3.UpdateInstanceDefaultConfigsOptions)
				updateInstanceDefaultConfigsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				updateInstanceDefaultConfigsOptionsModel.Body = make(map[string]interface{})
				updateInstanceDefaultConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApplication(createApplicationOptions *CreateApplicationOptions) - Operation response error`, func() {
		createApplicationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApplicationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApplication with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the CreateApplicationOptions model
				createApplicationOptionsModel := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				createApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.Application = core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.Jars = core.StringPtr("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.Packages = core.StringPtr("testString")
				createApplicationOptionsModel.Repositories = core.StringPtr("testString")
				createApplicationOptionsModel.Files = core.StringPtr("testString")
				createApplicationOptionsModel.Archives = core.StringPtr("testString")
				createApplicationOptionsModel.Name = core.StringPtr("spark-app")
				createApplicationOptionsModel.Class = core.StringPtr("com.company.path.ClassName")
				createApplicationOptionsModel.Arguments = []string{"[arg1, arg2, arg3]"}
				createApplicationOptionsModel.Conf = make(map[string]interface{})
				createApplicationOptionsModel.Env = make(map[string]interface{})
				createApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApplication(createApplicationOptions *CreateApplicationOptions)`, func() {
		createApplicationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApplicationPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "accepted"}`)
				}))
			})
			It(`Invoke CreateApplication successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateApplicationOptions model
				createApplicationOptionsModel := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				createApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.Application = core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.Jars = core.StringPtr("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.Packages = core.StringPtr("testString")
				createApplicationOptionsModel.Repositories = core.StringPtr("testString")
				createApplicationOptionsModel.Files = core.StringPtr("testString")
				createApplicationOptionsModel.Archives = core.StringPtr("testString")
				createApplicationOptionsModel.Name = core.StringPtr("spark-app")
				createApplicationOptionsModel.Class = core.StringPtr("com.company.path.ClassName")
				createApplicationOptionsModel.Arguments = []string{"[arg1, arg2, arg3]"}
				createApplicationOptionsModel.Conf = make(map[string]interface{})
				createApplicationOptionsModel.Env = make(map[string]interface{})
				createApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.CreateApplicationWithContext(ctx, createApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.CreateApplicationWithContext(ctx, createApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createApplicationPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "accepted"}`)
				}))
			})
			It(`Invoke CreateApplication successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApplicationOptions model
				createApplicationOptionsModel := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				createApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.Application = core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.Jars = core.StringPtr("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.Packages = core.StringPtr("testString")
				createApplicationOptionsModel.Repositories = core.StringPtr("testString")
				createApplicationOptionsModel.Files = core.StringPtr("testString")
				createApplicationOptionsModel.Archives = core.StringPtr("testString")
				createApplicationOptionsModel.Name = core.StringPtr("spark-app")
				createApplicationOptionsModel.Class = core.StringPtr("com.company.path.ClassName")
				createApplicationOptionsModel.Arguments = []string{"[arg1, arg2, arg3]"}
				createApplicationOptionsModel.Conf = make(map[string]interface{})
				createApplicationOptionsModel.Env = make(map[string]interface{})
				createApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateApplication with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the CreateApplicationOptions model
				createApplicationOptionsModel := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				createApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.Application = core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.Jars = core.StringPtr("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.Packages = core.StringPtr("testString")
				createApplicationOptionsModel.Repositories = core.StringPtr("testString")
				createApplicationOptionsModel.Files = core.StringPtr("testString")
				createApplicationOptionsModel.Archives = core.StringPtr("testString")
				createApplicationOptionsModel.Name = core.StringPtr("spark-app")
				createApplicationOptionsModel.Class = core.StringPtr("com.company.path.ClassName")
				createApplicationOptionsModel.Arguments = []string{"[arg1, arg2, arg3]"}
				createApplicationOptionsModel.Conf = make(map[string]interface{})
				createApplicationOptionsModel.Env = make(map[string]interface{})
				createApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateApplicationOptions model with no property values
				createApplicationOptionsModelNew := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateApplication successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the CreateApplicationOptions model
				createApplicationOptionsModel := new(ibmanalyticsengineapiv3.CreateApplicationOptions)
				createApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.Application = core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.Jars = core.StringPtr("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.Packages = core.StringPtr("testString")
				createApplicationOptionsModel.Repositories = core.StringPtr("testString")
				createApplicationOptionsModel.Files = core.StringPtr("testString")
				createApplicationOptionsModel.Archives = core.StringPtr("testString")
				createApplicationOptionsModel.Name = core.StringPtr("spark-app")
				createApplicationOptionsModel.Class = core.StringPtr("com.company.path.ClassName")
				createApplicationOptionsModel.Arguments = []string{"[arg1, arg2, arg3]"}
				createApplicationOptionsModel.Conf = make(map[string]interface{})
				createApplicationOptionsModel.Env = make(map[string]interface{})
				createApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.CreateApplication(createApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListApplications(listApplicationsOptions *ListApplicationsOptions) - Operation response error`, func() {
		listApplicationsPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListApplications with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				listApplicationsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListApplications(listApplicationsOptions *ListApplicationsOptions)`, func() {
		listApplicationsPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"id": "ID", "href": "Href", "spark_application_id": "SparkApplicationID", "spark_application_name": "SparkApplicationName", "state": "State", "start_time": "StartTime", "end_time": "EndTime", "finish_time": "FinishTime"}]}`)
				}))
			})
			It(`Invoke ListApplications successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				listApplicationsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ListApplicationsWithContext(ctx, listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ListApplicationsWithContext(ctx, listApplicationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"id": "ID", "href": "Href", "spark_application_id": "SparkApplicationID", "spark_application_name": "SparkApplicationName", "state": "State", "start_time": "StartTime", "end_time": "EndTime", "finish_time": "FinishTime"}]}`)
				}))
			})
			It(`Invoke ListApplications successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ListApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				listApplicationsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListApplications with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				listApplicationsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListApplicationsOptions model with no property values
				listApplicationsOptionsModelNew := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListApplications successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ListApplicationsOptions model
				listApplicationsOptionsModel := new(ibmanalyticsengineapiv3.ListApplicationsOptions)
				listApplicationsOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.ListApplications(listApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplication(getApplicationOptions *GetApplicationOptions) - Operation response error`, func() {
		getApplicationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications/ff48cc19-0e7e-4627-aac6-0b4ad080397b"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplication with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				getApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplication(getApplicationOptions *GetApplicationOptions)`, func() {
		getApplicationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications/ff48cc19-0e7e-4627-aac6-0b4ad080397b"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application_details": {"application": "cos://bucket_name.my_cos/my_spark_app.py", "jars": "cos://cloud-object-storage/jars/tests.jar", "packages": "Packages", "repositories": "Repositories", "files": "Files", "archives": "Archives", "name": "spark-app", "class": "com.company.path.ClassName", "arguments": ["[arg1, arg2, arg3]"], "conf": {"mapKey": "anyValue"}, "env": {"mapKey": "anyValue"}}, "id": "2b83d31c-397b-48ad-ad76-b83347c982db", "spark_application_id": "SparkApplicationID", "spark_application_name": "SparkApplicationName", "state": "accepted", "start_time": "2021-01-30T08:30:00.000Z", "end_time": "2021-01-30T08:30:00.000Z", "finish_time": "2021-01-30T08:30:00.000Z"}`)
				}))
			})
			It(`Invoke GetApplication successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				getApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetApplicationWithContext(ctx, getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetApplicationWithContext(ctx, getApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application_details": {"application": "cos://bucket_name.my_cos/my_spark_app.py", "jars": "cos://cloud-object-storage/jars/tests.jar", "packages": "Packages", "repositories": "Repositories", "files": "Files", "archives": "Archives", "name": "spark-app", "class": "com.company.path.ClassName", "arguments": ["[arg1, arg2, arg3]"], "conf": {"mapKey": "anyValue"}, "env": {"mapKey": "anyValue"}}, "id": "2b83d31c-397b-48ad-ad76-b83347c982db", "spark_application_id": "SparkApplicationID", "spark_application_name": "SparkApplicationName", "state": "accepted", "start_time": "2021-01-30T08:30:00.000Z", "end_time": "2021-01-30T08:30:00.000Z", "finish_time": "2021-01-30T08:30:00.000Z"}`)
				}))
			})
			It(`Invoke GetApplication successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				getApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplication with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				getApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationOptions model with no property values
				getApplicationOptionsModelNew := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetApplication successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationOptions model
				getApplicationOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationOptions)
				getApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplication(getApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions)`, func() {
		deleteApplicationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications/ff48cc19-0e7e-4627-aac6-0b4ad080397b"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteApplicationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteApplication successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmAnalyticsEngineApiService.DeleteApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApplicationOptions model
				deleteApplicationOptionsModel := new(ibmanalyticsengineapiv3.DeleteApplicationOptions)
				deleteApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				deleteApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				deleteApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApplication with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the DeleteApplicationOptions model
				deleteApplicationOptionsModel := new(ibmanalyticsengineapiv3.DeleteApplicationOptions)
				deleteApplicationOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				deleteApplicationOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				deleteApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteApplicationOptions model with no property values
				deleteApplicationOptionsModelNew := new(ibmanalyticsengineapiv3.DeleteApplicationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteApplication(deleteApplicationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationState(getApplicationStateOptions *GetApplicationStateOptions) - Operation response error`, func() {
		getApplicationStatePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications/ff48cc19-0e7e-4627-aac6-0b4ad080397b/state"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplicationState with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationStateOptions model
				getApplicationStateOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationState(getApplicationStateOptions *GetApplicationStateOptions)`, func() {
		getApplicationStatePath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/spark_applications/ff48cc19-0e7e-4627-aac6-0b4ad080397b/state"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "State", "start_time": "StartTime", "end_time": "EndTime", "finish_time": "FinishTime"}`)
				}))
			})
			It(`Invoke GetApplicationState successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationStateOptions model
				getApplicationStateOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetApplicationStateWithContext(ctx, getApplicationStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetApplicationStateWithContext(ctx, getApplicationStateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "state": "State", "start_time": "StartTime", "end_time": "EndTime", "finish_time": "FinishTime"}`)
				}))
			})
			It(`Invoke GetApplicationState successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationStateOptions model
				getApplicationStateOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplicationState with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationStateOptions model
				getApplicationStateOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationStateOptions model with no property values
				getApplicationStateOptionsModelNew := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetApplicationState successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationStateOptions model
				getApplicationStateOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationStateOptions)
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationState(getApplicationStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions *GetCurrentResourceConsumptionOptions) - Operation response error`, func() {
		getCurrentResourceConsumptionPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/current_resource_consumption"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCurrentResourceConsumptionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCurrentResourceConsumption with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				getCurrentResourceConsumptionOptionsModel := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				getCurrentResourceConsumptionOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions *GetCurrentResourceConsumptionOptions)`, func() {
		getCurrentResourceConsumptionPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/current_resource_consumption"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCurrentResourceConsumptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cores": "Cores", "memory": "Memory"}`)
				}))
			})
			It(`Invoke GetCurrentResourceConsumption successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				getCurrentResourceConsumptionOptionsModel := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				getCurrentResourceConsumptionOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumptionWithContext(ctx, getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetCurrentResourceConsumptionWithContext(ctx, getCurrentResourceConsumptionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCurrentResourceConsumptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cores": "Cores", "memory": "Memory"}`)
				}))
			})
			It(`Invoke GetCurrentResourceConsumption successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				getCurrentResourceConsumptionOptionsModel := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				getCurrentResourceConsumptionOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCurrentResourceConsumption with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				getCurrentResourceConsumptionOptionsModel := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				getCurrentResourceConsumptionOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCurrentResourceConsumptionOptions model with no property values
				getCurrentResourceConsumptionOptionsModelNew := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCurrentResourceConsumption successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				getCurrentResourceConsumptionOptionsModel := new(ibmanalyticsengineapiv3.GetCurrentResourceConsumptionOptions)
				getCurrentResourceConsumptionOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetCurrentResourceConsumption(getCurrentResourceConsumptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions) - Operation response error`, func() {
		replaceLogForwardingConfigPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/log_forwarding_config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceLogForwardingConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceLogForwardingConfig with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				replaceLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				replaceLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.Enabled = core.BoolPtr(true)
				replaceLogForwardingConfigOptionsModel.Sources = []string{"spark-driver", "spark-executor"}
				replaceLogForwardingConfigOptionsModel.Tags = []string{"<tag_1>", "<tag_2>", "<tag_n"}
				replaceLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions)`, func() {
		replaceLogForwardingConfigPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/log_forwarding_config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceLogForwardingConfigPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": ["Sources"], "tags": ["Tags"], "log_server": {"type": "ibm-log-analysis"}, "enabled": true}`)
				}))
			})
			It(`Invoke ReplaceLogForwardingConfig successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				replaceLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				replaceLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.Enabled = core.BoolPtr(true)
				replaceLogForwardingConfigOptionsModel.Sources = []string{"spark-driver", "spark-executor"}
				replaceLogForwardingConfigOptionsModel.Tags = []string{"<tag_1>", "<tag_2>", "<tag_n"}
				replaceLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfigWithContext(ctx, replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ReplaceLogForwardingConfigWithContext(ctx, replaceLogForwardingConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceLogForwardingConfigPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": ["Sources"], "tags": ["Tags"], "log_server": {"type": "ibm-log-analysis"}, "enabled": true}`)
				}))
			})
			It(`Invoke ReplaceLogForwardingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				replaceLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				replaceLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.Enabled = core.BoolPtr(true)
				replaceLogForwardingConfigOptionsModel.Sources = []string{"spark-driver", "spark-executor"}
				replaceLogForwardingConfigOptionsModel.Tags = []string{"<tag_1>", "<tag_2>", "<tag_n"}
				replaceLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceLogForwardingConfig with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				replaceLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				replaceLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.Enabled = core.BoolPtr(true)
				replaceLogForwardingConfigOptionsModel.Sources = []string{"spark-driver", "spark-executor"}
				replaceLogForwardingConfigOptionsModel.Tags = []string{"<tag_1>", "<tag_2>", "<tag_n"}
				replaceLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceLogForwardingConfigOptions model with no property values
				replaceLogForwardingConfigOptionsModelNew := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceLogForwardingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				replaceLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.ReplaceLogForwardingConfigOptions)
				replaceLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.Enabled = core.BoolPtr(true)
				replaceLogForwardingConfigOptionsModel.Sources = []string{"spark-driver", "spark-executor"}
				replaceLogForwardingConfigOptionsModel.Tags = []string{"<tag_1>", "<tag_2>", "<tag_n"}
				replaceLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.ReplaceLogForwardingConfig(replaceLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogForwardingConfig(getLogForwardingConfigOptions *GetLogForwardingConfigOptions) - Operation response error`, func() {
		getLogForwardingConfigPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/log_forwarding_config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogForwardingConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogForwardingConfig with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLogForwardingConfigOptions model
				getLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				getLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogForwardingConfig(getLogForwardingConfigOptions *GetLogForwardingConfigOptions)`, func() {
		getLogForwardingConfigPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/log_forwarding_config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogForwardingConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": ["Sources"], "tags": ["Tags"], "log_server": {"type": "ibm-log-analysis"}, "enabled": true}`)
				}))
			})
			It(`Invoke GetLogForwardingConfig successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogForwardingConfigOptions model
				getLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				getLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfigWithContext(ctx, getLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetLogForwardingConfigWithContext(ctx, getLogForwardingConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogForwardingConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": ["Sources"], "tags": ["Tags"], "log_server": {"type": "ibm-log-analysis"}, "enabled": true}`)
				}))
			})
			It(`Invoke GetLogForwardingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogForwardingConfigOptions model
				getLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				getLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogForwardingConfig with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLogForwardingConfigOptions model
				getLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				getLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLogForwardingConfigOptions model with no property values
				getLogForwardingConfigOptionsModelNew := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetLogForwardingConfig successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLogForwardingConfigOptions model
				getLogForwardingConfigOptionsModel := new(ibmanalyticsengineapiv3.GetLogForwardingConfigOptions)
				getLogForwardingConfigOptionsModel.InstanceID = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLogForwardingConfig(getLogForwardingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConfigurePlatformLogging(configurePlatformLoggingOptions *ConfigurePlatformLoggingOptions) - Operation response error`, func() {
		configurePlatformLoggingPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/logging"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(configurePlatformLoggingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ConfigurePlatformLogging with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ConfigurePlatformLoggingOptions model
				configurePlatformLoggingOptionsModel := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				configurePlatformLoggingOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.Enable = core.BoolPtr(true)
				configurePlatformLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConfigurePlatformLogging(configurePlatformLoggingOptions *ConfigurePlatformLoggingOptions)`, func() {
		configurePlatformLoggingPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/logging"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(configurePlatformLoggingPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"components": ["Components"], "log_server": {"type": "ibm-log-analysis"}, "enable": true}`)
				}))
			})
			It(`Invoke ConfigurePlatformLogging successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the ConfigurePlatformLoggingOptions model
				configurePlatformLoggingOptionsModel := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				configurePlatformLoggingOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.Enable = core.BoolPtr(true)
				configurePlatformLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLoggingWithContext(ctx, configurePlatformLoggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.ConfigurePlatformLoggingWithContext(ctx, configurePlatformLoggingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(configurePlatformLoggingPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"components": ["Components"], "log_server": {"type": "ibm-log-analysis"}, "enable": true}`)
				}))
			})
			It(`Invoke ConfigurePlatformLogging successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigurePlatformLoggingOptions model
				configurePlatformLoggingOptionsModel := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				configurePlatformLoggingOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.Enable = core.BoolPtr(true)
				configurePlatformLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ConfigurePlatformLogging with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ConfigurePlatformLoggingOptions model
				configurePlatformLoggingOptionsModel := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				configurePlatformLoggingOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.Enable = core.BoolPtr(true)
				configurePlatformLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ConfigurePlatformLoggingOptions model with no property values
				configurePlatformLoggingOptionsModelNew := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ConfigurePlatformLogging successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the ConfigurePlatformLoggingOptions model
				configurePlatformLoggingOptionsModel := new(ibmanalyticsengineapiv3.ConfigurePlatformLoggingOptions)
				configurePlatformLoggingOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.Enable = core.BoolPtr(true)
				configurePlatformLoggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.ConfigurePlatformLogging(configurePlatformLoggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoggingConfiguration(getLoggingConfigurationOptions *GetLoggingConfigurationOptions) - Operation response error`, func() {
		getLoggingConfigurationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/logging"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigurationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoggingConfiguration with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLoggingConfigurationOptions model
				getLoggingConfigurationOptionsModel := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				getLoggingConfigurationOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoggingConfiguration(getLoggingConfigurationOptions *GetLoggingConfigurationOptions)`, func() {
		getLoggingConfigurationPath := "/v3/analytics_engines/e64c907a-e82f-46fd-addc-ccfafbd28b09/logging"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"components": ["Components"], "log_server": {"type": "ibm-log-analysis"}, "enable": true}`)
				}))
			})
			It(`Invoke GetLoggingConfiguration successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLoggingConfigurationOptions model
				getLoggingConfigurationOptionsModel := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				getLoggingConfigurationOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfigurationWithContext(ctx, getLoggingConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfigurationWithContext(ctx, getLoggingConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLoggingConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"components": ["Components"], "log_server": {"type": "ibm-log-analysis"}, "enable": true}`)
				}))
			})
			It(`Invoke GetLoggingConfiguration successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoggingConfigurationOptions model
				getLoggingConfigurationOptionsModel := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				getLoggingConfigurationOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLoggingConfiguration with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLoggingConfigurationOptions model
				getLoggingConfigurationOptionsModel := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				getLoggingConfigurationOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoggingConfigurationOptions model with no property values
				getLoggingConfigurationOptionsModelNew := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetLoggingConfiguration successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetLoggingConfigurationOptions model
				getLoggingConfigurationOptionsModel := new(ibmanalyticsengineapiv3.GetLoggingConfigurationOptions)
				getLoggingConfigurationOptionsModel.InstanceGuid = core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetLoggingConfiguration(getLoggingConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmAnalyticsEngineApiService, _ := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
				URL:           "http://ibmanalyticsengineapiv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewConfigurePlatformLoggingOptions successfully`, func() {
				// Construct an instance of the ConfigurePlatformLoggingOptions model
				instanceGuid := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				configurePlatformLoggingOptionsModel := ibmAnalyticsEngineApiService.NewConfigurePlatformLoggingOptions(instanceGuid)
				configurePlatformLoggingOptionsModel.SetInstanceGuid("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				configurePlatformLoggingOptionsModel.SetEnable(true)
				configurePlatformLoggingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(configurePlatformLoggingOptionsModel).ToNot(BeNil())
				Expect(configurePlatformLoggingOptionsModel.InstanceGuid).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(configurePlatformLoggingOptionsModel.Enable).To(Equal(core.BoolPtr(true)))
				Expect(configurePlatformLoggingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateApplicationOptions successfully`, func() {
				// Construct an instance of the CreateApplicationOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				createApplicationOptionsModel := ibmAnalyticsEngineApiService.NewCreateApplicationOptions(instanceID)
				createApplicationOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				createApplicationOptionsModel.SetApplication("cos://bucket_name.my_cos/my_spark_app.py")
				createApplicationOptionsModel.SetJars("cos://cloud-object-storage/jars/tests.jar")
				createApplicationOptionsModel.SetPackages("testString")
				createApplicationOptionsModel.SetRepositories("testString")
				createApplicationOptionsModel.SetFiles("testString")
				createApplicationOptionsModel.SetArchives("testString")
				createApplicationOptionsModel.SetName("spark-app")
				createApplicationOptionsModel.SetClass("com.company.path.ClassName")
				createApplicationOptionsModel.SetArguments([]string{"[arg1, arg2, arg3]"})
				createApplicationOptionsModel.SetConf(make(map[string]interface{}))
				createApplicationOptionsModel.SetEnv(make(map[string]interface{}))
				createApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApplicationOptionsModel).ToNot(BeNil())
				Expect(createApplicationOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(createApplicationOptionsModel.Application).To(Equal(core.StringPtr("cos://bucket_name.my_cos/my_spark_app.py")))
				Expect(createApplicationOptionsModel.Jars).To(Equal(core.StringPtr("cos://cloud-object-storage/jars/tests.jar")))
				Expect(createApplicationOptionsModel.Packages).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.Repositories).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.Files).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.Archives).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.Name).To(Equal(core.StringPtr("spark-app")))
				Expect(createApplicationOptionsModel.Class).To(Equal(core.StringPtr("com.company.path.ClassName")))
				Expect(createApplicationOptionsModel.Arguments).To(Equal([]string{"[arg1, arg2, arg3]"}))
				Expect(createApplicationOptionsModel.Conf).To(Equal(make(map[string]interface{})))
				Expect(createApplicationOptionsModel.Env).To(Equal(make(map[string]interface{})))
				Expect(createApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteApplicationOptions successfully`, func() {
				// Construct an instance of the DeleteApplicationOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				applicationID := "ff48cc19-0e7e-4627-aac6-0b4ad080397b"
				deleteApplicationOptionsModel := ibmAnalyticsEngineApiService.NewDeleteApplicationOptions(instanceID, applicationID)
				deleteApplicationOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				deleteApplicationOptionsModel.SetApplicationID("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				deleteApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteApplicationOptionsModel).ToNot(BeNil())
				Expect(deleteApplicationOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(deleteApplicationOptionsModel.ApplicationID).To(Equal(core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")))
				Expect(deleteApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationOptions successfully`, func() {
				// Construct an instance of the GetApplicationOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				applicationID := "ff48cc19-0e7e-4627-aac6-0b4ad080397b"
				getApplicationOptionsModel := ibmAnalyticsEngineApiService.NewGetApplicationOptions(instanceID, applicationID)
				getApplicationOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationOptionsModel.SetApplicationID("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationOptionsModel).ToNot(BeNil())
				Expect(getApplicationOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getApplicationOptionsModel.ApplicationID).To(Equal(core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")))
				Expect(getApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationStateOptions successfully`, func() {
				// Construct an instance of the GetApplicationStateOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				applicationID := "ff48cc19-0e7e-4627-aac6-0b4ad080397b"
				getApplicationStateOptionsModel := ibmAnalyticsEngineApiService.NewGetApplicationStateOptions(instanceID, applicationID)
				getApplicationStateOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getApplicationStateOptionsModel.SetApplicationID("ff48cc19-0e7e-4627-aac6-0b4ad080397b")
				getApplicationStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationStateOptionsModel).ToNot(BeNil())
				Expect(getApplicationStateOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getApplicationStateOptionsModel.ApplicationID).To(Equal(core.StringPtr("ff48cc19-0e7e-4627-aac6-0b4ad080397b")))
				Expect(getApplicationStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCurrentResourceConsumptionOptions successfully`, func() {
				// Construct an instance of the GetCurrentResourceConsumptionOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getCurrentResourceConsumptionOptionsModel := ibmAnalyticsEngineApiService.NewGetCurrentResourceConsumptionOptions(instanceID)
				getCurrentResourceConsumptionOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getCurrentResourceConsumptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCurrentResourceConsumptionOptionsModel).ToNot(BeNil())
				Expect(getCurrentResourceConsumptionOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getCurrentResourceConsumptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceDefaultConfigsOptions successfully`, func() {
				// Construct an instance of the GetInstanceDefaultConfigsOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getInstanceDefaultConfigsOptionsModel := ibmAnalyticsEngineApiService.NewGetInstanceDefaultConfigsOptions(instanceID)
				getInstanceDefaultConfigsOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceDefaultConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceDefaultConfigsOptionsModel).ToNot(BeNil())
				Expect(getInstanceDefaultConfigsOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getInstanceDefaultConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceOptions successfully`, func() {
				// Construct an instance of the GetInstanceOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getInstanceOptionsModel := ibmAnalyticsEngineApiService.NewGetInstanceOptions(instanceID)
				getInstanceOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceOptionsModel).ToNot(BeNil())
				Expect(getInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceStateOptions successfully`, func() {
				// Construct an instance of the GetInstanceStateOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getInstanceStateOptionsModel := ibmAnalyticsEngineApiService.NewGetInstanceStateOptions(instanceID)
				getInstanceStateOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceStateOptionsModel).ToNot(BeNil())
				Expect(getInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogForwardingConfigOptions successfully`, func() {
				// Construct an instance of the GetLogForwardingConfigOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getLogForwardingConfigOptionsModel := ibmAnalyticsEngineApiService.NewGetLogForwardingConfigOptions(instanceID)
				getLogForwardingConfigOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLogForwardingConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogForwardingConfigOptionsModel).ToNot(BeNil())
				Expect(getLogForwardingConfigOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getLogForwardingConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoggingConfigurationOptions successfully`, func() {
				// Construct an instance of the GetLoggingConfigurationOptions model
				instanceGuid := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				getLoggingConfigurationOptionsModel := ibmAnalyticsEngineApiService.NewGetLoggingConfigurationOptions(instanceGuid)
				getLoggingConfigurationOptionsModel.SetInstanceGuid("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				getLoggingConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoggingConfigurationOptionsModel).ToNot(BeNil())
				Expect(getLoggingConfigurationOptionsModel.InstanceGuid).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(getLoggingConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListApplicationsOptions successfully`, func() {
				// Construct an instance of the ListApplicationsOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				listApplicationsOptionsModel := ibmAnalyticsEngineApiService.NewListApplicationsOptions(instanceID)
				listApplicationsOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				listApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listApplicationsOptionsModel).ToNot(BeNil())
				Expect(listApplicationsOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(listApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceInstanceDefaultConfigsOptions successfully`, func() {
				// Construct an instance of the ReplaceInstanceDefaultConfigsOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				body := make(map[string]string)
				replaceInstanceDefaultConfigsOptionsModel := ibmAnalyticsEngineApiService.NewReplaceInstanceDefaultConfigsOptions(instanceID, body)
				replaceInstanceDefaultConfigsOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceInstanceDefaultConfigsOptionsModel.SetBody(make(map[string]string))
				replaceInstanceDefaultConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceInstanceDefaultConfigsOptionsModel).ToNot(BeNil())
				Expect(replaceInstanceDefaultConfigsOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(replaceInstanceDefaultConfigsOptionsModel.Body).To(Equal(make(map[string]string)))
				Expect(replaceInstanceDefaultConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceLogForwardingConfigOptions successfully`, func() {
				// Construct an instance of the ReplaceLogForwardingConfigOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				replaceLogForwardingConfigOptionsModel := ibmAnalyticsEngineApiService.NewReplaceLogForwardingConfigOptions(instanceID)
				replaceLogForwardingConfigOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				replaceLogForwardingConfigOptionsModel.SetEnabled(true)
				replaceLogForwardingConfigOptionsModel.SetSources([]string{"spark-driver", "spark-executor"})
				replaceLogForwardingConfigOptionsModel.SetTags([]string{"<tag_1>", "<tag_2>", "<tag_n"})
				replaceLogForwardingConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceLogForwardingConfigOptionsModel).ToNot(BeNil())
				Expect(replaceLogForwardingConfigOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(replaceLogForwardingConfigOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceLogForwardingConfigOptionsModel.Sources).To(Equal([]string{"spark-driver", "spark-executor"}))
				Expect(replaceLogForwardingConfigOptionsModel.Tags).To(Equal([]string{"<tag_1>", "<tag_2>", "<tag_n"}))
				Expect(replaceLogForwardingConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetInstanceHomeOptions successfully`, func() {
				// Construct an instance of the SetInstanceHomeOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				setInstanceHomeOptionsModel := ibmAnalyticsEngineApiService.NewSetInstanceHomeOptions(instanceID)
				setInstanceHomeOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				setInstanceHomeOptionsModel.SetNewInstanceID("testString")
				setInstanceHomeOptionsModel.SetNewProvider("ibm-cos")
				setInstanceHomeOptionsModel.SetNewType("objectstore")
				setInstanceHomeOptionsModel.SetNewRegion("us-south")
				setInstanceHomeOptionsModel.SetNewEndpoint("s3.direct.us-south.cloud-object-storage.appdomain.cloud")
				setInstanceHomeOptionsModel.SetNewHmacAccessKey("821**********0ae")
				setInstanceHomeOptionsModel.SetNewHmacSecretKey("03e****************4fc3")
				setInstanceHomeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setInstanceHomeOptionsModel).ToNot(BeNil())
				Expect(setInstanceHomeOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(setInstanceHomeOptionsModel.NewInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(setInstanceHomeOptionsModel.NewProvider).To(Equal(core.StringPtr("ibm-cos")))
				Expect(setInstanceHomeOptionsModel.NewType).To(Equal(core.StringPtr("objectstore")))
				Expect(setInstanceHomeOptionsModel.NewRegion).To(Equal(core.StringPtr("us-south")))
				Expect(setInstanceHomeOptionsModel.NewEndpoint).To(Equal(core.StringPtr("s3.direct.us-south.cloud-object-storage.appdomain.cloud")))
				Expect(setInstanceHomeOptionsModel.NewHmacAccessKey).To(Equal(core.StringPtr("821**********0ae")))
				Expect(setInstanceHomeOptionsModel.NewHmacSecretKey).To(Equal(core.StringPtr("03e****************4fc3")))
				Expect(setInstanceHomeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceDefaultConfigsOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceDefaultConfigsOptions model
				instanceID := "e64c907a-e82f-46fd-addc-ccfafbd28b09"
				body := make(map[string]interface{})
				updateInstanceDefaultConfigsOptionsModel := ibmAnalyticsEngineApiService.NewUpdateInstanceDefaultConfigsOptions(instanceID, body)
				updateInstanceDefaultConfigsOptionsModel.SetInstanceID("e64c907a-e82f-46fd-addc-ccfafbd28b09")
				updateInstanceDefaultConfigsOptionsModel.SetBody(make(map[string]interface{}))
				updateInstanceDefaultConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceDefaultConfigsOptionsModel).ToNot(BeNil())
				Expect(updateInstanceDefaultConfigsOptionsModel.InstanceID).To(Equal(core.StringPtr("e64c907a-e82f-46fd-addc-ccfafbd28b09")))
				Expect(updateInstanceDefaultConfigsOptionsModel.Body).To(Equal(make(map[string]interface{})))
				Expect(updateInstanceDefaultConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
