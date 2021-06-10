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
	Describe(`GetInstanceByID(getInstanceByIdOptions *GetInstanceByIdOptions) - Operation response error`, func() {
		getInstanceByIDPath := "/v3/analytics_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceByID with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceByIdOptions model
				getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				getInstanceByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceByID(getInstanceByIdOptions *GetInstanceByIdOptions)`, func() {
		getInstanceByIDPath := "/v3/analytics_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance_id": "InstanceID", "state": "created", "state_change_time": "2019-01-01T12:00:00.000Z", "default_runtime": {"spark_version": "SparkVersion", "additional_packages": ["AdditionalPackages"]}, "instance_home": {"guid": "Guid", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "bucket": "Bucket", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}, "default_config": {"key": "Key"}}`)
				}))
			})
			It(`Invoke GetInstanceByID successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceByIdOptions model
				getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				getInstanceByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetInstanceByIDWithContext(ctx, getInstanceByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetInstanceByIDWithContext(ctx, getInstanceByIdOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"instance_id": "InstanceID", "state": "created", "state_change_time": "2019-01-01T12:00:00.000Z", "default_runtime": {"spark_version": "SparkVersion", "additional_packages": ["AdditionalPackages"]}, "instance_home": {"guid": "Guid", "provider": "Provider", "type": "Type", "region": "Region", "endpoint": "Endpoint", "bucket": "Bucket", "hmac_access_key": "HmacAccessKey", "hmac_secret_key": "HmacSecretKey"}, "default_config": {"key": "Key"}}`)
				}))
			})
			It(`Invoke GetInstanceByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceByIdOptions model
				getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				getInstanceByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceByIdOptions model
				getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				getInstanceByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceByIdOptions model with no property values
				getInstanceByIdOptionsModelNew := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModelNew)
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
			It(`Invoke GetInstanceByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetInstanceByIdOptions model
				getInstanceByIdOptionsModel := new(ibmanalyticsengineapiv3.GetInstanceByIdOptions)
				getInstanceByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetInstanceByID(getInstanceByIdOptionsModel)
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
		createApplicationPath := "/v3/analytics_engines/testString/spark/applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApplicationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				createApplicationOptionsModel.InstanceID = core.StringPtr("testString")
				createApplicationOptionsModel.Application = core.StringPtr("testString")
				createApplicationOptionsModel.ApplicationArguments = []string{"testString"}
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
		createApplicationPath := "/v3/analytics_engines/testString/spark/applications"
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"application_id": "ApplicationID", "state": "accepted", "start_time": "2019-01-01T12:00:00.000Z"}`)
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
				createApplicationOptionsModel.InstanceID = core.StringPtr("testString")
				createApplicationOptionsModel.Application = core.StringPtr("testString")
				createApplicationOptionsModel.ApplicationArguments = []string{"testString"}
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"application_id": "ApplicationID", "state": "accepted", "start_time": "2019-01-01T12:00:00.000Z"}`)
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
				createApplicationOptionsModel.InstanceID = core.StringPtr("testString")
				createApplicationOptionsModel.Application = core.StringPtr("testString")
				createApplicationOptionsModel.ApplicationArguments = []string{"testString"}
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
				createApplicationOptionsModel.InstanceID = core.StringPtr("testString")
				createApplicationOptionsModel.Application = core.StringPtr("testString")
				createApplicationOptionsModel.ApplicationArguments = []string{"testString"}
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
					res.WriteHeader(201)
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
				createApplicationOptionsModel.InstanceID = core.StringPtr("testString")
				createApplicationOptionsModel.Application = core.StringPtr("testString")
				createApplicationOptionsModel.ApplicationArguments = []string{"testString"}
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
	Describe(`GetApplications(getApplicationsOptions *GetApplicationsOptions) - Operation response error`, func() {
		getApplicationsPath := "/v3/analytics_engines/testString/spark/applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplications with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationsOptions model
				getApplicationsOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				getApplicationsOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplications(getApplicationsOptions *GetApplicationsOptions)`, func() {
		getApplicationsPath := "/v3/analytics_engines/testString/spark/applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"application_id": "ApplicationID", "spark_application_id": "SparkApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}]}`)
				}))
			})
			It(`Invoke GetApplications successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationsOptions model
				getApplicationsOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				getApplicationsOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetApplicationsWithContext(ctx, getApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetApplicationsWithContext(ctx, getApplicationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"application_id": "ApplicationID", "spark_application_id": "SparkApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}]}`)
				}))
			})
			It(`Invoke GetApplications successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationsOptions model
				getApplicationsOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				getApplicationsOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplications with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationsOptions model
				getApplicationsOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				getApplicationsOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationsOptions model with no property values
				getApplicationsOptionsModelNew := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModelNew)
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
			It(`Invoke GetApplications successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationsOptions model
				getApplicationsOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationsOptions)
				getApplicationsOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplications(getApplicationsOptionsModel)
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
	Describe(`GetApplicationByID(getApplicationByIdOptions *GetApplicationByIdOptions) - Operation response error`, func() {
		getApplicationByIDPath := "/v3/analytics_engines/testString/spark/applications/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationByIDPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApplicationByID with error: Operation response processing error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationByIdOptions model
				getApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				getApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationByID(getApplicationByIdOptions *GetApplicationByIdOptions)`, func() {
		getApplicationByIDPath := "/v3/analytics_engines/testString/spark/applications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application_details": {"application_details": {"application": "Application", "application_arguments": ["ApplicationArguments"], "conf": {"mapKey": "anyValue"}, "env": {"mapKey": "anyValue"}}}, "mode": "Mode", "application_id": "ApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}`)
				}))
			})
			It(`Invoke GetApplicationByID successfully with retries`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())
				ibmAnalyticsEngineApiService.EnableRetries(0, 0)

				// Construct an instance of the GetApplicationByIdOptions model
				getApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				getApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmAnalyticsEngineApiService.GetApplicationByIDWithContext(ctx, getApplicationByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmAnalyticsEngineApiService.DisableRetries()
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmAnalyticsEngineApiService.GetApplicationByIDWithContext(ctx, getApplicationByIdOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationByIDPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application_details": {"application_details": {"application": "Application", "application_arguments": ["ApplicationArguments"], "conf": {"mapKey": "anyValue"}, "env": {"mapKey": "anyValue"}}}, "mode": "Mode", "application_id": "ApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}`)
				}))
			})
			It(`Invoke GetApplicationByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApplicationByIdOptions model
				getApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				getApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApplicationByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationByIdOptions model
				getApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				getApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApplicationByIdOptions model with no property values
				getApplicationByIdOptionsModelNew := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModelNew)
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
			It(`Invoke GetApplicationByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the GetApplicationByIdOptions model
				getApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.GetApplicationByIdOptions)
				getApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				getApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmAnalyticsEngineApiService.GetApplicationByID(getApplicationByIdOptionsModel)
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
	Describe(`DeleteApplicationByID(deleteApplicationByIdOptions *DeleteApplicationByIdOptions)`, func() {
		deleteApplicationByIDPath := "/v3/analytics_engines/testString/spark/applications/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteApplicationByIDPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteApplicationByID successfully`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmAnalyticsEngineApiService.DeleteApplicationByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApplicationByIdOptions model
				deleteApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.DeleteApplicationByIdOptions)
				deleteApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				deleteApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApplicationByID with error: Operation validation and request error`, func() {
				ibmAnalyticsEngineApiService, serviceErr := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmAnalyticsEngineApiService).ToNot(BeNil())

				// Construct an instance of the DeleteApplicationByIdOptions model
				deleteApplicationByIdOptionsModel := new(ibmanalyticsengineapiv3.DeleteApplicationByIdOptions)
				deleteApplicationByIdOptionsModel.InstanceID = core.StringPtr("testString")
				deleteApplicationByIdOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteApplicationByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmAnalyticsEngineApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteApplicationByIdOptions model with no property values
				deleteApplicationByIdOptionsModelNew := new(ibmanalyticsengineapiv3.DeleteApplicationByIdOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmAnalyticsEngineApiService.DeleteApplicationByID(deleteApplicationByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApplicationState(getApplicationStateOptions *GetApplicationStateOptions) - Operation response error`, func() {
		getApplicationStatePath := "/v3/analytics_engines/testString/spark/applications/testString/state"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApplicationStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("testString")
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
		getApplicationStatePath := "/v3/analytics_engines/testString/spark/applications/testString/state"
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
					fmt.Fprintf(res, "%s", `{"application_id": "ApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}`)
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
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"application_id": "ApplicationID", "state": "State", "start_time": "StartTime", "finish_time": "FinishTime"}`)
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
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("testString")
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
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("testString")
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
				getApplicationStateOptionsModel.InstanceID = core.StringPtr("testString")
				getApplicationStateOptionsModel.ApplicationID = core.StringPtr("testString")
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmAnalyticsEngineApiService, _ := ibmanalyticsengineapiv3.NewIbmAnalyticsEngineApiV3(&ibmanalyticsengineapiv3.IbmAnalyticsEngineApiV3Options{
				URL:           "http://ibmanalyticsengineapiv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateApplicationOptions successfully`, func() {
				// Construct an instance of the CreateApplicationOptions model
				instanceID := "testString"
				createApplicationOptionsModel := ibmAnalyticsEngineApiService.NewCreateApplicationOptions(instanceID)
				createApplicationOptionsModel.SetInstanceID("testString")
				createApplicationOptionsModel.SetApplication("testString")
				createApplicationOptionsModel.SetApplicationArguments([]string{"testString"})
				createApplicationOptionsModel.SetConf(make(map[string]interface{}))
				createApplicationOptionsModel.SetEnv(make(map[string]interface{}))
				createApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApplicationOptionsModel).ToNot(BeNil())
				Expect(createApplicationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.Application).To(Equal(core.StringPtr("testString")))
				Expect(createApplicationOptionsModel.ApplicationArguments).To(Equal([]string{"testString"}))
				Expect(createApplicationOptionsModel.Conf).To(Equal(make(map[string]interface{})))
				Expect(createApplicationOptionsModel.Env).To(Equal(make(map[string]interface{})))
				Expect(createApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteApplicationByIdOptions successfully`, func() {
				// Construct an instance of the DeleteApplicationByIdOptions model
				instanceID := "testString"
				applicationID := "testString"
				deleteApplicationByIdOptionsModel := ibmAnalyticsEngineApiService.NewDeleteApplicationByIdOptions(instanceID, applicationID)
				deleteApplicationByIdOptionsModel.SetInstanceID("testString")
				deleteApplicationByIdOptionsModel.SetApplicationID("testString")
				deleteApplicationByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteApplicationByIdOptionsModel).ToNot(BeNil())
				Expect(deleteApplicationByIdOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteApplicationByIdOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteApplicationByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationByIdOptions successfully`, func() {
				// Construct an instance of the GetApplicationByIdOptions model
				instanceID := "testString"
				applicationID := "testString"
				getApplicationByIdOptionsModel := ibmAnalyticsEngineApiService.NewGetApplicationByIdOptions(instanceID, applicationID)
				getApplicationByIdOptionsModel.SetInstanceID("testString")
				getApplicationByIdOptionsModel.SetApplicationID("testString")
				getApplicationByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationByIdOptionsModel).ToNot(BeNil())
				Expect(getApplicationByIdOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationByIdOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationStateOptions successfully`, func() {
				// Construct an instance of the GetApplicationStateOptions model
				instanceID := "testString"
				applicationID := "testString"
				getApplicationStateOptionsModel := ibmAnalyticsEngineApiService.NewGetApplicationStateOptions(instanceID, applicationID)
				getApplicationStateOptionsModel.SetInstanceID("testString")
				getApplicationStateOptionsModel.SetApplicationID("testString")
				getApplicationStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationStateOptionsModel).ToNot(BeNil())
				Expect(getApplicationStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationStateOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApplicationsOptions successfully`, func() {
				// Construct an instance of the GetApplicationsOptions model
				instanceID := "testString"
				getApplicationsOptionsModel := ibmAnalyticsEngineApiService.NewGetApplicationsOptions(instanceID)
				getApplicationsOptionsModel.SetInstanceID("testString")
				getApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApplicationsOptionsModel).ToNot(BeNil())
				Expect(getApplicationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceByIdOptions successfully`, func() {
				// Construct an instance of the GetInstanceByIdOptions model
				instanceID := "testString"
				getInstanceByIdOptionsModel := ibmAnalyticsEngineApiService.NewGetInstanceByIdOptions(instanceID)
				getInstanceByIdOptionsModel.SetInstanceID("testString")
				getInstanceByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceByIdOptionsModel).ToNot(BeNil())
				Expect(getInstanceByIdOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getInstanceByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
