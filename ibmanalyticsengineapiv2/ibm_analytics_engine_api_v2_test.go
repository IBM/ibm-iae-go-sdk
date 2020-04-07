/**
 * (C) Copyright IBM Corp. 2020.
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
	"fmt"
	"github.com/IBM/cloud-go-sdk/ibmanalyticsengineapiv2"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`IbmAnalyticsEngineApiV2`, func() {
	Describe(`GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions)`, func() {
		getAllAnalyticsEnginesPath := "/v2/analytics_engines"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAllAnalyticsEnginesPath))
				Expect(req.Method).To(Equal("GET"))
				res.WriteHeader(200)
			}))
			It(`Invoke GetAllAnalyticsEngines successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetAllAnalyticsEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetAllAnalyticsEnginesOptions model
				getAllAnalyticsEnginesOptionsModel := new(ibmanalyticsengineapiv2.GetAllAnalyticsEnginesOptions)

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetAllAnalyticsEngines(getAllAnalyticsEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions)`, func() {
		getAnalyticsEngineByIDPath := "/v2/analytics_engines/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAnalyticsEngineByIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "service_plan": "ServicePlan", "hardware_size": "HardwareSize", "software_package": "SoftwarePackage", "domain": "Domain", "creation_time": "2019-01-01T12:00:00", "commision_time": "2019-01-01T12:00:00", "decommision_time": "2019-01-01T12:00:00", "deletion_time": "2019-01-01T12:00:00", "state_change_time": "2019-01-01T12:00:00", "state": "State", "nodes": [{"id": "ID", "fqdn": "Fqdn", "type": "Type", "state": "State", "public_ip": "PublicIp", "private_ip": "PrivateIp", "state_change_time": "2019-01-01T12:00:00", "commission_time": "2019-01-01T12:00:00"}], "user_credentials": {"user": "User"}, "service_endpoints": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}, "service_endpoints_ip": {"phoenix_jdbc": "PhoenixJdbc", "ambari_console": "AmbariConsole", "livy": "Livy", "spark_history_server": "SparkHistoryServer", "oozie_rest": "OozieRest", "hive_jdbc": "HiveJdbc", "notebook_gateway_websocket": "NotebookGatewayWebsocket", "notebook_gateway": "NotebookGateway", "webhdfs": "Webhdfs", "ssh": "Ssh", "spark_sql": "SparkSql"}}`)
			}))
			It(`Invoke GetAnalyticsEngineByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAnalyticsEngineByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAnalyticsEngineByIdOptions model
				getAnalyticsEngineByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineByIdOptions)
				getAnalyticsEngineByIdOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAnalyticsEngineByID(getAnalyticsEngineByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions)`, func() {
		getAnalyticsEngineStateByIDPath := "/v2/analytics_engines/testString/state"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAnalyticsEngineStateByIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"state": "State"}`)
			}))
			It(`Invoke GetAnalyticsEngineStateByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAnalyticsEngineStateByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAnalyticsEngineStateByIdOptions model
				getAnalyticsEngineStateByIdOptionsModel := new(ibmanalyticsengineapiv2.GetAnalyticsEngineStateByIdOptions)
				getAnalyticsEngineStateByIdOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions)`, func() {
		createCustomizationRequestPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCustomizationRequestPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"request_id": "RequestID"}`)
			}))
			It(`Invoke CreateCustomizationRequest successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateCustomizationRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnalyticsEngineCustomActionScript model
				analyticsEngineCustomActionScriptModel := new(ibmanalyticsengineapiv2.AnalyticsEngineCustomActionScript)
				analyticsEngineCustomActionScriptModel.SourceType = core.StringPtr("http")
				analyticsEngineCustomActionScriptModel.ScriptPath = core.StringPtr("testString")
				analyticsEngineCustomActionScriptModel.SourceProps = CreateMockMap()

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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateCustomizationRequest(createCustomizationRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions)`, func() {
		getAllCustomizationRequestsPath := "/v2/analytics_engines/testString/customization_requests"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAllCustomizationRequestsPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke GetAllCustomizationRequests successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAllCustomizationRequests(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllCustomizationRequestsOptions model
				getAllCustomizationRequestsOptionsModel := new(ibmanalyticsengineapiv2.GetAllCustomizationRequestsOptions)
				getAllCustomizationRequestsOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAllCustomizationRequests(getAllCustomizationRequestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions)`, func() {
		getCustomizationRequestByIDPath := "/v2/analytics_engines/testString/customization_requests/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCustomizationRequestByIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "run_status": "RunStatus", "run_details": {"overall_status": "OverallStatus", "details": [{"node_name": "NodeName", "node_type": "NodeType", "start_time": "StartTime", "end_time": "EndTime", "time_taken": "TimeTaken", "status": "Status", "log_file": "LogFile"}]}}`)
			}))
			It(`Invoke GetCustomizationRequestByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCustomizationRequestByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomizationRequestByIdOptions model
				getCustomizationRequestByIdOptionsModel := new(ibmanalyticsengineapiv2.GetCustomizationRequestByIdOptions)
				getCustomizationRequestByIdOptionsModel.InstanceGuid = core.StringPtr("testString")
				getCustomizationRequestByIdOptionsModel.RequestID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCustomizationRequestByID(getCustomizationRequestByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ResizeCluster(resizeClusterOptions *ResizeClusterOptions)`, func() {
		resizeClusterPath := "/v2/analytics_engines/testString/resize"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(resizeClusterPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"request_id": "RequestID"}`)
			}))
			It(`Invoke ResizeCluster successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ResizeCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResizeClusterOptions model
				resizeClusterOptionsModel := new(ibmanalyticsengineapiv2.ResizeClusterOptions)
				resizeClusterOptionsModel.InstanceGuid = core.StringPtr("testString")
				resizeClusterOptionsModel.ComputeNodesCount = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ResizeCluster(resizeClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions)`, func() {
		resetClusterPasswordPath := "/v2/analytics_engines/testString/reset_password"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(resetClusterPasswordPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "user_credentials": {"user": "User", "password": "Password"}}`)
			}))
			It(`Invoke ResetClusterPassword successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ResetClusterPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResetClusterPasswordOptions model
				resetClusterPasswordOptionsModel := new(ibmanalyticsengineapiv2.ResetClusterPasswordOptions)
				resetClusterPasswordOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ResetClusterPassword(resetClusterPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions)`, func() {
		configureLoggingPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(configureLoggingPath))
				Expect(req.Method).To(Equal("PUT"))
				res.WriteHeader(202)
			}))
			It(`Invoke ConfigureLogging successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ConfigureLogging(nil)
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

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ConfigureLogging(configureLoggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions)`, func() {
		getLoggingConfigPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLoggingConfigPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"log_specs": [{"node_type": "management", "components": ["ambari-server"]}], "log_server": {"type": "logdna", "credential": "Credential", "api_host": "ApiHost", "log_host": "LogHost", "owner": "Owner"}, "log_config_status": [{"node_type": "management", "node_id": "NodeID", "action": "Action", "status": "Status"}]}`)
			}))
			It(`Invoke GetLoggingConfig successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoggingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoggingConfigOptions model
				getLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.GetLoggingConfigOptions)
				getLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoggingConfig(getLoggingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions)`, func() {
		deleteLoggingConfigPath := "/v2/analytics_engines/testString/log_config"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLoggingConfigPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(202)
			}))
			It(`Invoke DeleteLoggingConfig successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteLoggingConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoggingConfigOptions model
				deleteLoggingConfigOptionsModel := new(ibmanalyticsengineapiv2.DeleteLoggingConfigOptions)
				deleteLoggingConfigOptionsModel.InstanceGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteLoggingConfig(deleteLoggingConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := ibmanalyticsengineapiv2.NewIbmAnalyticsEngineApiV2(&ibmanalyticsengineapiv2.IbmAnalyticsEngineApiV2Options{
				URL:           "http://ibmanalyticsengineapiv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAnalyticsEngineCustomAction successfully`, func() {
				name := "testString"
				model, err := testService.NewAnalyticsEngineCustomAction(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAnalyticsEngineLoggingNodeSpec successfully`, func() {
				nodeType := "management"
				components := []string{"ambari-server"}
				model, err := testService.NewAnalyticsEngineLoggingNodeSpec(nodeType, components)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAnalyticsEngineLoggingServer successfully`, func() {
				typeVar := "logdna"
				credential := "testString"
				apiHost := "testString"
				logHost := "testString"
				model, err := testService.NewAnalyticsEngineLoggingServer(typeVar, credential, apiHost, logHost)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
