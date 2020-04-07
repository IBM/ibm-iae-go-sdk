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

// Package ibmanalyticsengineapiv2 : Operations and models for the IbmAnalyticsEngineApiV2 service
package ibmanalyticsengineapiv2

import (
	"fmt"
	common "github.com/IBM/cloud-go-sdk/common"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
)

// IbmAnalyticsEngineApiV2 : With IBM Analytics Engine you can create Apache Spark and Apache Hadoop clusters and
// customize these clusters by using scripts. You can work with data in IBM Cloud Object Storage, as well as integrate
// other Watson Data Platform services like IBM Watson Studio and Machine Learning.
//
// Version: 2.0.5
type IbmAnalyticsEngineApiV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://ibm-analytics-engine-api-docs.cloud.ibm.com/"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_analytics_engine_api_docs"

// IbmAnalyticsEngineApiV2Options : Service options
type IbmAnalyticsEngineApiV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmAnalyticsEngineApiV2UsingExternalConfig : constructs an instance of IbmAnalyticsEngineApiV2 with passed in options and external configuration.
func NewIbmAnalyticsEngineApiV2UsingExternalConfig(options *IbmAnalyticsEngineApiV2Options) (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmAnalyticsEngineApiDocs, err = NewIbmAnalyticsEngineApiV2(options)
	if err != nil {
		return
	}

	err = ibmAnalyticsEngineApiDocs.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmAnalyticsEngineApiDocs.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmAnalyticsEngineApiV2 : constructs an instance of IbmAnalyticsEngineApiV2 with passed in options.
func NewIbmAnalyticsEngineApiV2(options *IbmAnalyticsEngineApiV2Options) (service *IbmAnalyticsEngineApiV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &IbmAnalyticsEngineApiV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) SetServiceURL(url string) error {
	return ibmAnalyticsEngineApiDocs.Service.SetServiceURL(url)
}

// GetAllAnalyticsEngines : List all Analytics Engines
// Currently, you cannot fetch the list of all IBM Analytics Engine service instances through this REST API. You should
// use the IBM Cloud CLI instead.  For example, ```ibmcloud resource service-instances --service-name
// ibmanalyticsengine```.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAllAnalyticsEnginesOptions, "getAllAnalyticsEnginesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAllAnalyticsEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetAllAnalyticsEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, nil)

	return
}

// GetAnalyticsEngineByID : Get details of Analytics Engine
// Retrieves the following details of the IBM Analytics Engine service instance:
// * Hardware size and software package
//  * Timestamps at which the cluster was created, deleted or updated
//  * Service endpoint URLs
//
//  **NOTE:** No credentials are returned. You can get the IBM Analytics Engine service instance credentials by invoking
// the reset_password REST API.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions) (result *AnalyticsEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAnalyticsEngineByIdOptions, "getAnalyticsEngineByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAnalyticsEngineByIdOptions, "getAnalyticsEngineByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines"}
	pathParameters := []string{*getAnalyticsEngineByIdOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnalyticsEngineByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetAnalyticsEngineByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngine(m)
		response.Result = result
	}

	return
}

// GetAnalyticsEngineStateByID : Get state of Analytics Engine
// Returns the state of the Analytics Engine cluster. The following states exist:
// * Preparing : A cluster is being created.
// * Active : The cluster is created and running.
// * Deleted : The cluster was deleted.
// * Failed : A cluster couldn't be created.
// * Expired : The service instance has expired. The cluster has been deleted.
// * ResizeFailed : The cluster couldn't be resized. The cluster will be reactivated based on the old settings.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions) (result *AnalyticsEngineState, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAnalyticsEngineStateByIdOptions, "getAnalyticsEngineStateByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAnalyticsEngineStateByIdOptions, "getAnalyticsEngineStateByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "state"}
	pathParameters := []string{*getAnalyticsEngineStateByIdOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnalyticsEngineStateByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetAnalyticsEngineStateByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineState(m)
		response.Result = result
	}

	return
}

// CreateCustomizationRequest : Create an adhoc customization request
// Creates a new adhoc customization request. Adhoc customization scripts can be run only once. They are not persisted
// with the cluster and are not run automatically when more nodes are added to the cluster.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions) (result *AnalyticsEngineCreateCustomizationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomizationRequestOptions, "createCustomizationRequestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomizationRequestOptions, "createCustomizationRequestOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "customization_requests"}
	pathParameters := []string{*createCustomizationRequestOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomizationRequestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "CreateCustomizationRequest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCustomizationRequestOptions.Target != nil {
		body["target"] = createCustomizationRequestOptions.Target
	}
	if createCustomizationRequestOptions.CustomActions != nil {
		body["custom_actions"] = createCustomizationRequestOptions.CustomActions
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineCreateCustomizationResponse(m)
		response.Result = result
	}

	return
}

// GetAllCustomizationRequests : Get all customization requests run on an Analytics Engine cluster
// Retrieves the request_id of all customization requests submitted to the specified Analytics Engine cluster.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions) (result *[]interface{}, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAllCustomizationRequestsOptions, "getAllCustomizationRequestsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAllCustomizationRequestsOptions, "getAllCustomizationRequestsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "customization_requests"}
	pathParameters := []string{*getAllCustomizationRequestsOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAllCustomizationRequestsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetAllCustomizationRequests")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make([]map[string]interface{}, 1))
	if err == nil {
		s, ok := response.Result.([]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		slice, e := Unmarshalinterface{}Slice(s)
		result = &slice
		err = e
		response.Result = result
	}

	return
}

// GetCustomizationRequestByID : Retrieve details of specified customization request ID
// Retrieves the status of the specified customization request, along with pointers to log files generated during the
// run.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions) (result *AnalyticsEngineCustomizationRunDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomizationRequestByIdOptions, "getCustomizationRequestByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomizationRequestByIdOptions, "getCustomizationRequestByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "customization_requests"}
	pathParameters := []string{*getCustomizationRequestByIdOptions.InstanceGuid, *getCustomizationRequestByIdOptions.RequestID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomizationRequestByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetCustomizationRequestByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineCustomizationRunDetails(m)
		response.Result = result
	}

	return
}

// ResizeCluster : Add nodes to the cluster
// Resizes the cluster by adding compute nodes.
//
// **Note:** You can't resize the cluster if the software package on the cluster is deprecated or if the software
// package doesn't permit cluster resizing. See
// [here](https://cloud.ibm.com/docs/AnalyticsEngine?topic=AnalyticsEngine-unsupported-operations).
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) ResizeCluster(resizeClusterOptions *ResizeClusterOptions) (result *AnalyticsEngineResizeClusterResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resizeClusterOptions, "resizeClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resizeClusterOptions, "resizeClusterOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "resize"}
	pathParameters := []string{*resizeClusterOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range resizeClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "ResizeCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if resizeClusterOptions.ComputeNodesCount != nil {
		body["compute_nodes_count"] = resizeClusterOptions.ComputeNodesCount
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineResizeClusterResponse(m)
		response.Result = result
	}

	return
}

// ResetClusterPassword : Reset cluster password
// Resets the cluster's password to a new system-generated crytographically strong value.  The new password is included
// in the response and you should make a note of it.  This password is displayed only once here and cannot be retrieved
// later.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions) (result *AnalyticsEngineResetClusterPasswordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resetClusterPasswordOptions, "resetClusterPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resetClusterPasswordOptions, "resetClusterPasswordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "reset_password"}
	pathParameters := []string{*resetClusterPasswordOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range resetClusterPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "ResetClusterPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineResetClusterPasswordResponse(m)
		response.Result = result
	}

	return
}

// ConfigureLogging : Configure log aggregation
// Collects the logs for the following components in an IBM Analytics Engine cluster:
// * IBM Analytics Engine daemon logs, for example those for Spark, Hive, Yarn, and Knox on the management and data
// nodes
// * Yarn application job logs.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(configureLoggingOptions, "configureLoggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(configureLoggingOptions, "configureLoggingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "log_config"}
	pathParameters := []string{*configureLoggingOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range configureLoggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "ConfigureLogging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if configureLoggingOptions.LogSpecs != nil {
		body["log_specs"] = configureLoggingOptions.LogSpecs
	}
	if configureLoggingOptions.LogServer != nil {
		body["log_server"] = configureLoggingOptions.LogServer
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, nil)

	return
}

// GetLoggingConfig : Retrieve the status of log configuration
// Retrieves the status and details of the log configuration for your cluster.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions) (result *AnalyticsEngineLoggingConfigDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoggingConfigOptions, "getLoggingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoggingConfigOptions, "getLoggingConfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "log_config"}
	pathParameters := []string{*getLoggingConfigOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoggingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "GetLoggingConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAnalyticsEngineLoggingConfigDetails(m)
		response.Result = result
	}

	return
}

// DeleteLoggingConfig : Delete the log configuration
// Deletes the log configuration. This operation stops sending logs to the centralized log server.
func (ibmAnalyticsEngineApiDocs *IbmAnalyticsEngineApiV2) DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoggingConfigOptions, "deleteLoggingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoggingConfigOptions, "deleteLoggingConfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/analytics_engines", "log_config"}
	pathParameters := []string{*deleteLoggingConfigOptions.InstanceGuid}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(ibmAnalyticsEngineApiDocs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoggingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api_docs", "V2", "DeleteLoggingConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApiDocs.Service.Request(request, nil)

	return
}

// AnalyticsEngine : Analytics Engine cluster details.
type AnalyticsEngine struct {
	// Instance GUID.
	ID *string `json:"id" validate:"required"`

	// Analytics Engine.
	Name *string `json:"name" validate:"required"`

	// ID of Analytics Engine service plan.
	ServicePlan *string `json:"service_plan" validate:"required"`

	// Hardware size.
	HardwareSize *string `json:"hardware_size" validate:"required"`

	// Software package.
	SoftwarePackage *string `json:"software_package" validate:"required"`

	// Domain.
	Domain *string `json:"domain" validate:"required"`

	// Cluster creation time.
	CreationTime *strfmt.DateTime `json:"creation_time" validate:"required"`

	// Cluster commision time.
	CommisionTime *strfmt.DateTime `json:"commision_time" validate:"required"`

	// Cluster decommision time.
	DecommisionTime *strfmt.DateTime `json:"decommision_time" validate:"required"`

	// Cluster deletion time.
	DeletionTime *strfmt.DateTime `json:"deletion_time" validate:"required"`

	// Cluster state change time.
	StateChangeTime *strfmt.DateTime `json:"state_change_time" validate:"required"`

	// Cluster state.
	State *string `json:"state" validate:"required"`

	// List of nodes in the cluster.
	Nodes []AnalyticsEngineClusterNode `json:"nodes,omitempty"`

	// User credentials.
	UserCredentials *AnalyticsEngineUserCredentials `json:"user_credentials" validate:"required"`

	// Service endpoint URLs with host names. Endpoints will vary based on software package chosen for the cluster.
	ServiceEndpoints *ServiceEndpoints `json:"service_endpoints,omitempty"`

	// Service endpoint URLs with host IPS. Endpoints will vary based on software package chosen for the cluster.
	ServiceEndpointsIp *ServiceEndpoints `json:"service_endpoints_ip,omitempty"`
}


// UnmarshalAnalyticsEngine constructs an instance of AnalyticsEngine from the specified map.
func UnmarshalAnalyticsEngine(m map[string]interface{}) (result *AnalyticsEngine, err error) {
	obj := new(AnalyticsEngine)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.ServicePlan, err = core.UnmarshalString(m, "service_plan")
	if err != nil {
		return
	}
	obj.HardwareSize, err = core.UnmarshalString(m, "hardware_size")
	if err != nil {
		return
	}
	obj.SoftwarePackage, err = core.UnmarshalString(m, "software_package")
	if err != nil {
		return
	}
	obj.Domain, err = core.UnmarshalString(m, "domain")
	if err != nil {
		return
	}
	obj.CreationTime, err = core.UnmarshalDateTime(m, "creation_time")
	if err != nil {
		return
	}
	obj.CommisionTime, err = core.UnmarshalDateTime(m, "commision_time")
	if err != nil {
		return
	}
	obj.DecommisionTime, err = core.UnmarshalDateTime(m, "decommision_time")
	if err != nil {
		return
	}
	obj.DeletionTime, err = core.UnmarshalDateTime(m, "deletion_time")
	if err != nil {
		return
	}
	obj.StateChangeTime, err = core.UnmarshalDateTime(m, "state_change_time")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.Nodes, err = UnmarshalAnalyticsEngineClusterNodeSliceAsProperty(m, "nodes")
	if err != nil {
		return
	}
	obj.UserCredentials, err = UnmarshalAnalyticsEngineUserCredentialsAsProperty(m, "user_credentials")
	if err != nil {
		return
	}
	obj.ServiceEndpoints, err = UnmarshalServiceEndpointsAsProperty(m, "service_endpoints")
	if err != nil {
		return
	}
	obj.ServiceEndpointsIp, err = UnmarshalServiceEndpointsAsProperty(m, "service_endpoints_ip")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineSlice unmarshals a slice of AnalyticsEngine instances from the specified list of maps.
func UnmarshalAnalyticsEngineSlice(s []interface{}) (slice []AnalyticsEngine, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngine'")
			return
		}
		obj, e := UnmarshalAnalyticsEngine(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineAsProperty unmarshals an instance of AnalyticsEngine that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngine, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngine'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngine(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineSliceAsProperty unmarshals a slice of AnalyticsEngine instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngine, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngine'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineSlice(vSlice)
	}
	return
}

// AnalyticsEngineClusterNode : Cluster node details.
type AnalyticsEngineClusterNode struct {
	// Node ID.
	ID *string `json:"id,omitempty"`

	// Fully qualified domain name.
	Fqdn *string `json:"fqdn,omitempty"`

	// Node type.
	Type *string `json:"type,omitempty"`

	// State of node.
	State *string `json:"state,omitempty"`

	// Public IP address.
	PublicIp *string `json:"public_ip,omitempty"`

	// Private IP address.
	PrivateIp *string `json:"private_ip,omitempty"`

	// State change time.
	StateChangeTime *strfmt.DateTime `json:"state_change_time,omitempty"`

	// Commission time.
	CommissionTime *strfmt.DateTime `json:"commission_time,omitempty"`
}


// UnmarshalAnalyticsEngineClusterNode constructs an instance of AnalyticsEngineClusterNode from the specified map.
func UnmarshalAnalyticsEngineClusterNode(m map[string]interface{}) (result *AnalyticsEngineClusterNode, err error) {
	obj := new(AnalyticsEngineClusterNode)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Fqdn, err = core.UnmarshalString(m, "fqdn")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.PublicIp, err = core.UnmarshalString(m, "public_ip")
	if err != nil {
		return
	}
	obj.PrivateIp, err = core.UnmarshalString(m, "private_ip")
	if err != nil {
		return
	}
	obj.StateChangeTime, err = core.UnmarshalDateTime(m, "state_change_time")
	if err != nil {
		return
	}
	obj.CommissionTime, err = core.UnmarshalDateTime(m, "commission_time")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineClusterNodeSlice unmarshals a slice of AnalyticsEngineClusterNode instances from the specified list of maps.
func UnmarshalAnalyticsEngineClusterNodeSlice(s []interface{}) (slice []AnalyticsEngineClusterNode, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineClusterNode'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineClusterNode(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineClusterNodeAsProperty unmarshals an instance of AnalyticsEngineClusterNode that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineClusterNodeAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineClusterNode, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineClusterNode'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineClusterNode(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineClusterNodeSliceAsProperty unmarshals a slice of AnalyticsEngineClusterNode instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineClusterNodeSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineClusterNode, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineClusterNode'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineClusterNodeSlice(vSlice)
	}
	return
}

// AnalyticsEngineCreateCustomizationResponse : Create customization request response.
type AnalyticsEngineCreateCustomizationResponse struct {
	// Customization request ID.
	RequestID *string `json:"request_id,omitempty"`
}


// UnmarshalAnalyticsEngineCreateCustomizationResponse constructs an instance of AnalyticsEngineCreateCustomizationResponse from the specified map.
func UnmarshalAnalyticsEngineCreateCustomizationResponse(m map[string]interface{}) (result *AnalyticsEngineCreateCustomizationResponse, err error) {
	obj := new(AnalyticsEngineCreateCustomizationResponse)
	obj.RequestID, err = core.UnmarshalString(m, "request_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineCreateCustomizationResponseSlice unmarshals a slice of AnalyticsEngineCreateCustomizationResponse instances from the specified list of maps.
func UnmarshalAnalyticsEngineCreateCustomizationResponseSlice(s []interface{}) (slice []AnalyticsEngineCreateCustomizationResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineCreateCustomizationResponse'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineCreateCustomizationResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineCreateCustomizationResponseAsProperty unmarshals an instance of AnalyticsEngineCreateCustomizationResponse that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCreateCustomizationResponseAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineCreateCustomizationResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineCreateCustomizationResponse'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineCreateCustomizationResponse(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineCreateCustomizationResponseSliceAsProperty unmarshals a slice of AnalyticsEngineCreateCustomizationResponse instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCreateCustomizationResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineCreateCustomizationResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineCreateCustomizationResponse'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineCreateCustomizationResponseSlice(vSlice)
	}
	return
}

// AnalyticsEngineCustomAction : Custom action details for customization.
type AnalyticsEngineCustomAction struct {
	// Custom action name.
	Name *string `json:"name" validate:"required"`

	// Customization type.
	Type *string `json:"type,omitempty"`

	// Customization script details.
	Script *AnalyticsEngineCustomActionScript `json:"script,omitempty"`

	// Customization script parameters.
	ScriptParams []string `json:"script_params,omitempty"`
}

// Constants associated with the AnalyticsEngineCustomAction.Type property.
// Customization type.
const (
	AnalyticsEngineCustomAction_Type_Bootstrap = "bootstrap"
)


// NewAnalyticsEngineCustomAction : Instantiate AnalyticsEngineCustomAction (Generic Model Constructor)
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineCustomAction(name string) (model *AnalyticsEngineCustomAction, err error) {
	model = &AnalyticsEngineCustomAction{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineCustomAction constructs an instance of AnalyticsEngineCustomAction from the specified map.
func UnmarshalAnalyticsEngineCustomAction(m map[string]interface{}) (result *AnalyticsEngineCustomAction, err error) {
	obj := new(AnalyticsEngineCustomAction)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Script, err = UnmarshalAnalyticsEngineCustomActionScriptAsProperty(m, "script")
	if err != nil {
		return
	}
	obj.ScriptParams, err = core.UnmarshalStringSlice(m, "script_params")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineCustomActionSlice unmarshals a slice of AnalyticsEngineCustomAction instances from the specified list of maps.
func UnmarshalAnalyticsEngineCustomActionSlice(s []interface{}) (slice []AnalyticsEngineCustomAction, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineCustomAction'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineCustomAction(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineCustomActionAsProperty unmarshals an instance of AnalyticsEngineCustomAction that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomActionAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineCustomAction, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineCustomAction'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineCustomAction(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineCustomActionSliceAsProperty unmarshals a slice of AnalyticsEngineCustomAction instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomActionSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineCustomAction, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineCustomAction'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineCustomActionSlice(vSlice)
	}
	return
}

// AnalyticsEngineCustomActionScript : Customization script details.
type AnalyticsEngineCustomActionScript struct {
	// Defines where to access the customization script.
	SourceType *string `json:"source_type,omitempty"`

	// Path to the customization script.
	ScriptPath *string `json:"script_path,omitempty"`

	// Customization script properties.
	SourceProps interface{} `json:"source_props,omitempty"`
}

// Constants associated with the AnalyticsEngineCustomActionScript.SourceType property.
// Defines where to access the customization script.
const (
	AnalyticsEngineCustomActionScript_SourceType_Bluemixswift = "BluemixSwift"
	AnalyticsEngineCustomActionScript_SourceType_Coss3 = "CosS3"
	AnalyticsEngineCustomActionScript_SourceType_Http = "http"
	AnalyticsEngineCustomActionScript_SourceType_Https = "https"
	AnalyticsEngineCustomActionScript_SourceType_Softlayerswift = "SoftLayerSwift"
)


// UnmarshalAnalyticsEngineCustomActionScript constructs an instance of AnalyticsEngineCustomActionScript from the specified map.
func UnmarshalAnalyticsEngineCustomActionScript(m map[string]interface{}) (result *AnalyticsEngineCustomActionScript, err error) {
	obj := new(AnalyticsEngineCustomActionScript)
	obj.SourceType, err = core.UnmarshalString(m, "source_type")
	if err != nil {
		return
	}
	obj.ScriptPath, err = core.UnmarshalString(m, "script_path")
	if err != nil {
		return
	}
	obj.SourceProps, err = core.UnmarshalAny(m, "source_props")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineCustomActionScriptSlice unmarshals a slice of AnalyticsEngineCustomActionScript instances from the specified list of maps.
func UnmarshalAnalyticsEngineCustomActionScriptSlice(s []interface{}) (slice []AnalyticsEngineCustomActionScript, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineCustomActionScript'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineCustomActionScript(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineCustomActionScriptAsProperty unmarshals an instance of AnalyticsEngineCustomActionScript that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomActionScriptAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineCustomActionScript, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineCustomActionScript'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineCustomActionScript(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineCustomActionScriptSliceAsProperty unmarshals a slice of AnalyticsEngineCustomActionScript instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomActionScriptSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineCustomActionScript, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineCustomActionScript'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineCustomActionScriptSlice(vSlice)
	}
	return
}

// AnalyticsEngineCustomizationRunDetails : Customization run details for the cluster.
type AnalyticsEngineCustomizationRunDetails struct {
	// Instance GUID.
	ID *string `json:"id,omitempty"`

	// Customization run status.
	RunStatus *string `json:"run_status,omitempty"`

	// Customization run details.
	RunDetails *AnalyticsEngineCustomizationRunDetailsRunDetails `json:"run_details,omitempty"`
}


// UnmarshalAnalyticsEngineCustomizationRunDetails constructs an instance of AnalyticsEngineCustomizationRunDetails from the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetails(m map[string]interface{}) (result *AnalyticsEngineCustomizationRunDetails, err error) {
	obj := new(AnalyticsEngineCustomizationRunDetails)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.RunStatus, err = core.UnmarshalString(m, "run_status")
	if err != nil {
		return
	}
	obj.RunDetails, err = UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsAsProperty(m, "run_details")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsSlice unmarshals a slice of AnalyticsEngineCustomizationRunDetails instances from the specified list of maps.
func UnmarshalAnalyticsEngineCustomizationRunDetailsSlice(s []interface{}) (slice []AnalyticsEngineCustomizationRunDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineCustomizationRunDetails'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineCustomizationRunDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsAsProperty unmarshals an instance of AnalyticsEngineCustomizationRunDetails that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetailsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineCustomizationRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineCustomizationRunDetails'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineCustomizationRunDetails(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsSliceAsProperty unmarshals a slice of AnalyticsEngineCustomizationRunDetails instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineCustomizationRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineCustomizationRunDetails'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineCustomizationRunDetailsSlice(vSlice)
	}
	return
}

// AnalyticsEngineCustomizationRunDetailsRunDetails : Customization run details.
type AnalyticsEngineCustomizationRunDetailsRunDetails struct {
	// Customization run overall status.
	OverallStatus *string `json:"overall_status,omitempty"`

	// Customization run details for each node.
	Details []AnalyticsEngineNodeLevelCustomizationRunDetails `json:"details,omitempty"`
}


// UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails constructs an instance of AnalyticsEngineCustomizationRunDetailsRunDetails from the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails(m map[string]interface{}) (result *AnalyticsEngineCustomizationRunDetailsRunDetails, err error) {
	obj := new(AnalyticsEngineCustomizationRunDetailsRunDetails)
	obj.OverallStatus, err = core.UnmarshalString(m, "overall_status")
	if err != nil {
		return
	}
	obj.Details, err = UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSliceAsProperty(m, "details")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsSlice unmarshals a slice of AnalyticsEngineCustomizationRunDetailsRunDetails instances from the specified list of maps.
func UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsSlice(s []interface{}) (slice []AnalyticsEngineCustomizationRunDetailsRunDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineCustomizationRunDetailsRunDetails'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsAsProperty unmarshals an instance of AnalyticsEngineCustomizationRunDetailsRunDetails that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineCustomizationRunDetailsRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineCustomizationRunDetailsRunDetails'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsSliceAsProperty unmarshals a slice of AnalyticsEngineCustomizationRunDetailsRunDetails instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineCustomizationRunDetailsRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineCustomizationRunDetailsRunDetails'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetailsSlice(vSlice)
	}
	return
}

// AnalyticsEngineLoggingConfigDetails : Logging configuration.
type AnalyticsEngineLoggingConfigDetails struct {
	// Log specifications for nodes.
	LogSpecs []AnalyticsEngineLoggingNodeSpec `json:"log_specs" validate:"required"`

	// Logging server configuration.
	LogServer *AnalyticsEngineLoggingServer `json:"log_server" validate:"required"`

	// Log configuration status.
	LogConfigStatus []AnalyticsEngineLoggingConfigStatus `json:"log_config_status" validate:"required"`
}


// UnmarshalAnalyticsEngineLoggingConfigDetails constructs an instance of AnalyticsEngineLoggingConfigDetails from the specified map.
func UnmarshalAnalyticsEngineLoggingConfigDetails(m map[string]interface{}) (result *AnalyticsEngineLoggingConfigDetails, err error) {
	obj := new(AnalyticsEngineLoggingConfigDetails)
	obj.LogSpecs, err = UnmarshalAnalyticsEngineLoggingNodeSpecSliceAsProperty(m, "log_specs")
	if err != nil {
		return
	}
	obj.LogServer, err = UnmarshalAnalyticsEngineLoggingServerAsProperty(m, "log_server")
	if err != nil {
		return
	}
	obj.LogConfigStatus, err = UnmarshalAnalyticsEngineLoggingConfigStatusSliceAsProperty(m, "log_config_status")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineLoggingConfigDetailsSlice unmarshals a slice of AnalyticsEngineLoggingConfigDetails instances from the specified list of maps.
func UnmarshalAnalyticsEngineLoggingConfigDetailsSlice(s []interface{}) (slice []AnalyticsEngineLoggingConfigDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineLoggingConfigDetails'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineLoggingConfigDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingConfigDetailsAsProperty unmarshals an instance of AnalyticsEngineLoggingConfigDetails that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingConfigDetailsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineLoggingConfigDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineLoggingConfigDetails'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineLoggingConfigDetails(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingConfigDetailsSliceAsProperty unmarshals a slice of AnalyticsEngineLoggingConfigDetails instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingConfigDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineLoggingConfigDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineLoggingConfigDetails'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineLoggingConfigDetailsSlice(vSlice)
	}
	return
}

// AnalyticsEngineLoggingConfigStatus : Log configuration status.
type AnalyticsEngineLoggingConfigStatus struct {
	// Node type.
	NodeType *string `json:"node_type" validate:"required"`

	// Node ID.
	NodeID *string `json:"node_id" validate:"required"`

	// Action.
	Action *string `json:"action" validate:"required"`

	// Log configuration status.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the AnalyticsEngineLoggingConfigStatus.NodeType property.
// Node type.
const (
	AnalyticsEngineLoggingConfigStatus_NodeType_Data = "data"
	AnalyticsEngineLoggingConfigStatus_NodeType_Management = "management"
)


// UnmarshalAnalyticsEngineLoggingConfigStatus constructs an instance of AnalyticsEngineLoggingConfigStatus from the specified map.
func UnmarshalAnalyticsEngineLoggingConfigStatus(m map[string]interface{}) (result *AnalyticsEngineLoggingConfigStatus, err error) {
	obj := new(AnalyticsEngineLoggingConfigStatus)
	obj.NodeType, err = core.UnmarshalString(m, "node_type")
	if err != nil {
		return
	}
	obj.NodeID, err = core.UnmarshalString(m, "node_id")
	if err != nil {
		return
	}
	obj.Action, err = core.UnmarshalString(m, "action")
	if err != nil {
		return
	}
	obj.Status, err = core.UnmarshalString(m, "status")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineLoggingConfigStatusSlice unmarshals a slice of AnalyticsEngineLoggingConfigStatus instances from the specified list of maps.
func UnmarshalAnalyticsEngineLoggingConfigStatusSlice(s []interface{}) (slice []AnalyticsEngineLoggingConfigStatus, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineLoggingConfigStatus'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineLoggingConfigStatus(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingConfigStatusAsProperty unmarshals an instance of AnalyticsEngineLoggingConfigStatus that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingConfigStatusAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineLoggingConfigStatus, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineLoggingConfigStatus'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineLoggingConfigStatus(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingConfigStatusSliceAsProperty unmarshals a slice of AnalyticsEngineLoggingConfigStatus instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingConfigStatusSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineLoggingConfigStatus, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineLoggingConfigStatus'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineLoggingConfigStatusSlice(vSlice)
	}
	return
}

// AnalyticsEngineLoggingNodeSpec : Log specifications for node.
type AnalyticsEngineLoggingNodeSpec struct {
	// Node type.
	NodeType *string `json:"node_type" validate:"required"`

	// Node components to be monitored.
	Components []string `json:"components" validate:"required"`
}

// Constants associated with the AnalyticsEngineLoggingNodeSpec.NodeType property.
// Node type.
const (
	AnalyticsEngineLoggingNodeSpec_NodeType_Data = "data"
	AnalyticsEngineLoggingNodeSpec_NodeType_Management = "management"
)

// Constants associated with the AnalyticsEngineLoggingNodeSpec.Components property.
// Node components to be logged.
const (
	AnalyticsEngineLoggingNodeSpec_Components_AmbariServer = "ambari-server"
	AnalyticsEngineLoggingNodeSpec_Components_HadoopMapreduce = "hadoop-mapreduce"
	AnalyticsEngineLoggingNodeSpec_Components_HadoopYarn = "hadoop-yarn"
	AnalyticsEngineLoggingNodeSpec_Components_Hbase = "hbase"
	AnalyticsEngineLoggingNodeSpec_Components_Hive = "hive"
	AnalyticsEngineLoggingNodeSpec_Components_Jnbg = "jnbg"
	AnalyticsEngineLoggingNodeSpec_Components_Knox = "knox"
	AnalyticsEngineLoggingNodeSpec_Components_Livy2 = "livy2"
	AnalyticsEngineLoggingNodeSpec_Components_Spark2 = "spark2"
	AnalyticsEngineLoggingNodeSpec_Components_YarnApps = "yarn-apps"
)


// NewAnalyticsEngineLoggingNodeSpec : Instantiate AnalyticsEngineLoggingNodeSpec (Generic Model Constructor)
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineLoggingNodeSpec(nodeType string, components []string) (model *AnalyticsEngineLoggingNodeSpec, err error) {
	model = &AnalyticsEngineLoggingNodeSpec{
		NodeType: core.StringPtr(nodeType),
		Components: components,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineLoggingNodeSpec constructs an instance of AnalyticsEngineLoggingNodeSpec from the specified map.
func UnmarshalAnalyticsEngineLoggingNodeSpec(m map[string]interface{}) (result *AnalyticsEngineLoggingNodeSpec, err error) {
	obj := new(AnalyticsEngineLoggingNodeSpec)
	obj.NodeType, err = core.UnmarshalString(m, "node_type")
	if err != nil {
		return
	}
	obj.Components, err = core.UnmarshalStringSlice(m, "components")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineLoggingNodeSpecSlice unmarshals a slice of AnalyticsEngineLoggingNodeSpec instances from the specified list of maps.
func UnmarshalAnalyticsEngineLoggingNodeSpecSlice(s []interface{}) (slice []AnalyticsEngineLoggingNodeSpec, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineLoggingNodeSpec'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineLoggingNodeSpec(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingNodeSpecAsProperty unmarshals an instance of AnalyticsEngineLoggingNodeSpec that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingNodeSpecAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineLoggingNodeSpec, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineLoggingNodeSpec'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineLoggingNodeSpec(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingNodeSpecSliceAsProperty unmarshals a slice of AnalyticsEngineLoggingNodeSpec instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingNodeSpecSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineLoggingNodeSpec, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineLoggingNodeSpec'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineLoggingNodeSpecSlice(vSlice)
	}
	return
}

// AnalyticsEngineLoggingServer : Logging server configuration.
type AnalyticsEngineLoggingServer struct {
	// Logging server type.
	Type *string `json:"type" validate:"required"`

	// Logging server credential.
	Credential *string `json:"credential" validate:"required"`

	// Logging server API host.
	ApiHost *string `json:"api_host" validate:"required"`

	// Logging server host.
	LogHost *string `json:"log_host" validate:"required"`

	// Logging server owner.
	Owner *string `json:"owner,omitempty"`
}

// Constants associated with the AnalyticsEngineLoggingServer.Type property.
// Logging server type.
const (
	AnalyticsEngineLoggingServer_Type_Logdna = "logdna"
)


// NewAnalyticsEngineLoggingServer : Instantiate AnalyticsEngineLoggingServer (Generic Model Constructor)
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineLoggingServer(typeVar string, credential string, apiHost string, logHost string) (model *AnalyticsEngineLoggingServer, err error) {
	model = &AnalyticsEngineLoggingServer{
		Type: core.StringPtr(typeVar),
		Credential: core.StringPtr(credential),
		ApiHost: core.StringPtr(apiHost),
		LogHost: core.StringPtr(logHost),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineLoggingServer constructs an instance of AnalyticsEngineLoggingServer from the specified map.
func UnmarshalAnalyticsEngineLoggingServer(m map[string]interface{}) (result *AnalyticsEngineLoggingServer, err error) {
	obj := new(AnalyticsEngineLoggingServer)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Credential, err = core.UnmarshalString(m, "credential")
	if err != nil {
		return
	}
	obj.ApiHost, err = core.UnmarshalString(m, "api_host")
	if err != nil {
		return
	}
	obj.LogHost, err = core.UnmarshalString(m, "log_host")
	if err != nil {
		return
	}
	obj.Owner, err = core.UnmarshalString(m, "owner")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineLoggingServerSlice unmarshals a slice of AnalyticsEngineLoggingServer instances from the specified list of maps.
func UnmarshalAnalyticsEngineLoggingServerSlice(s []interface{}) (slice []AnalyticsEngineLoggingServer, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineLoggingServer'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineLoggingServer(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingServerAsProperty unmarshals an instance of AnalyticsEngineLoggingServer that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingServerAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineLoggingServer, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineLoggingServer'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineLoggingServer(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineLoggingServerSliceAsProperty unmarshals a slice of AnalyticsEngineLoggingServer instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineLoggingServerSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineLoggingServer, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineLoggingServer'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineLoggingServerSlice(vSlice)
	}
	return
}

// AnalyticsEngineNodeLevelCustomizationRunDetails : Customization run details for the node.
type AnalyticsEngineNodeLevelCustomizationRunDetails struct {
	// Node name.
	NodeName *string `json:"node_name,omitempty"`

	// Node type.
	NodeType *string `json:"node_type,omitempty"`

	// Customization request start time.
	StartTime *string `json:"start_time,omitempty"`

	// Customization request end time.
	EndTime *string `json:"end_time,omitempty"`

	// Total time taken for customization request.
	TimeTaken *string `json:"time_taken,omitempty"`

	// Status of customization request.
	Status *string `json:"status,omitempty"`

	// Log file to track for customization run information.
	LogFile *string `json:"log_file,omitempty"`
}


// UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails constructs an instance of AnalyticsEngineNodeLevelCustomizationRunDetails from the specified map.
func UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails(m map[string]interface{}) (result *AnalyticsEngineNodeLevelCustomizationRunDetails, err error) {
	obj := new(AnalyticsEngineNodeLevelCustomizationRunDetails)
	obj.NodeName, err = core.UnmarshalString(m, "node_name")
	if err != nil {
		return
	}
	obj.NodeType, err = core.UnmarshalString(m, "node_type")
	if err != nil {
		return
	}
	obj.StartTime, err = core.UnmarshalString(m, "start_time")
	if err != nil {
		return
	}
	obj.EndTime, err = core.UnmarshalString(m, "end_time")
	if err != nil {
		return
	}
	obj.TimeTaken, err = core.UnmarshalString(m, "time_taken")
	if err != nil {
		return
	}
	obj.Status, err = core.UnmarshalString(m, "status")
	if err != nil {
		return
	}
	obj.LogFile, err = core.UnmarshalString(m, "log_file")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSlice unmarshals a slice of AnalyticsEngineNodeLevelCustomizationRunDetails instances from the specified list of maps.
func UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSlice(s []interface{}) (slice []AnalyticsEngineNodeLevelCustomizationRunDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineNodeLevelCustomizationRunDetails'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsAsProperty unmarshals an instance of AnalyticsEngineNodeLevelCustomizationRunDetails that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineNodeLevelCustomizationRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineNodeLevelCustomizationRunDetails'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSliceAsProperty unmarshals a slice of AnalyticsEngineNodeLevelCustomizationRunDetails instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineNodeLevelCustomizationRunDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineNodeLevelCustomizationRunDetails'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetailsSlice(vSlice)
	}
	return
}

// AnalyticsEngineResetClusterPasswordResponse : Response for resetting cluster password.
type AnalyticsEngineResetClusterPasswordResponse struct {
	// Instance guid.
	ID *string `json:"id,omitempty"`

	// User credentials.
	UserCredentials *AnalyticsEngineResetClusterPasswordResponseUserCredentials `json:"user_credentials,omitempty"`
}


// UnmarshalAnalyticsEngineResetClusterPasswordResponse constructs an instance of AnalyticsEngineResetClusterPasswordResponse from the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponse(m map[string]interface{}) (result *AnalyticsEngineResetClusterPasswordResponse, err error) {
	obj := new(AnalyticsEngineResetClusterPasswordResponse)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.UserCredentials, err = UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsAsProperty(m, "user_credentials")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseSlice unmarshals a slice of AnalyticsEngineResetClusterPasswordResponse instances from the specified list of maps.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseSlice(s []interface{}) (slice []AnalyticsEngineResetClusterPasswordResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineResetClusterPasswordResponse'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineResetClusterPasswordResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseAsProperty unmarshals an instance of AnalyticsEngineResetClusterPasswordResponse that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineResetClusterPasswordResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineResetClusterPasswordResponse'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineResetClusterPasswordResponse(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseSliceAsProperty unmarshals a slice of AnalyticsEngineResetClusterPasswordResponse instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineResetClusterPasswordResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineResetClusterPasswordResponse'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineResetClusterPasswordResponseSlice(vSlice)
	}
	return
}

// AnalyticsEngineResetClusterPasswordResponseUserCredentials : User credentials.
type AnalyticsEngineResetClusterPasswordResponseUserCredentials struct {
	// Username.
	User *string `json:"user,omitempty"`

	// New password.
	Password *string `json:"password,omitempty"`
}


// UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials constructs an instance of AnalyticsEngineResetClusterPasswordResponseUserCredentials from the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials(m map[string]interface{}) (result *AnalyticsEngineResetClusterPasswordResponseUserCredentials, err error) {
	obj := new(AnalyticsEngineResetClusterPasswordResponseUserCredentials)
	obj.User, err = core.UnmarshalString(m, "user")
	if err != nil {
		return
	}
	obj.Password, err = core.UnmarshalString(m, "password")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsSlice unmarshals a slice of AnalyticsEngineResetClusterPasswordResponseUserCredentials instances from the specified list of maps.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsSlice(s []interface{}) (slice []AnalyticsEngineResetClusterPasswordResponseUserCredentials, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineResetClusterPasswordResponseUserCredentials'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsAsProperty unmarshals an instance of AnalyticsEngineResetClusterPasswordResponseUserCredentials that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineResetClusterPasswordResponseUserCredentials, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineResetClusterPasswordResponseUserCredentials'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsSliceAsProperty unmarshals a slice of AnalyticsEngineResetClusterPasswordResponseUserCredentials instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineResetClusterPasswordResponseUserCredentials, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineResetClusterPasswordResponseUserCredentials'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentialsSlice(vSlice)
	}
	return
}

// AnalyticsEngineResizeClusterResponse : Resize request response.
type AnalyticsEngineResizeClusterResponse struct {
	// Request ID.
	RequestID *string `json:"request_id,omitempty"`
}


// UnmarshalAnalyticsEngineResizeClusterResponse constructs an instance of AnalyticsEngineResizeClusterResponse from the specified map.
func UnmarshalAnalyticsEngineResizeClusterResponse(m map[string]interface{}) (result *AnalyticsEngineResizeClusterResponse, err error) {
	obj := new(AnalyticsEngineResizeClusterResponse)
	obj.RequestID, err = core.UnmarshalString(m, "request_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineResizeClusterResponseSlice unmarshals a slice of AnalyticsEngineResizeClusterResponse instances from the specified list of maps.
func UnmarshalAnalyticsEngineResizeClusterResponseSlice(s []interface{}) (slice []AnalyticsEngineResizeClusterResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineResizeClusterResponse'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineResizeClusterResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineResizeClusterResponseAsProperty unmarshals an instance of AnalyticsEngineResizeClusterResponse that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResizeClusterResponseAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineResizeClusterResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineResizeClusterResponse'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineResizeClusterResponse(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineResizeClusterResponseSliceAsProperty unmarshals a slice of AnalyticsEngineResizeClusterResponse instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineResizeClusterResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineResizeClusterResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineResizeClusterResponse'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineResizeClusterResponseSlice(vSlice)
	}
	return
}

// AnalyticsEngineState : Cluster state.
type AnalyticsEngineState struct {
	// Cluster state.
	State *string `json:"state" validate:"required"`
}


// UnmarshalAnalyticsEngineState constructs an instance of AnalyticsEngineState from the specified map.
func UnmarshalAnalyticsEngineState(m map[string]interface{}) (result *AnalyticsEngineState, err error) {
	obj := new(AnalyticsEngineState)
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineStateSlice unmarshals a slice of AnalyticsEngineState instances from the specified list of maps.
func UnmarshalAnalyticsEngineStateSlice(s []interface{}) (slice []AnalyticsEngineState, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineState'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineState(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineStateAsProperty unmarshals an instance of AnalyticsEngineState that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineStateAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineState, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineState'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineState(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineStateSliceAsProperty unmarshals a slice of AnalyticsEngineState instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineStateSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineState, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineState'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineStateSlice(vSlice)
	}
	return
}

// AnalyticsEngineUserCredentials : User credentials.
type AnalyticsEngineUserCredentials struct {
	// Username.
	User *string `json:"user,omitempty"`
}


// UnmarshalAnalyticsEngineUserCredentials constructs an instance of AnalyticsEngineUserCredentials from the specified map.
func UnmarshalAnalyticsEngineUserCredentials(m map[string]interface{}) (result *AnalyticsEngineUserCredentials, err error) {
	obj := new(AnalyticsEngineUserCredentials)
	obj.User, err = core.UnmarshalString(m, "user")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAnalyticsEngineUserCredentialsSlice unmarshals a slice of AnalyticsEngineUserCredentials instances from the specified list of maps.
func UnmarshalAnalyticsEngineUserCredentialsSlice(s []interface{}) (slice []AnalyticsEngineUserCredentials, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AnalyticsEngineUserCredentials'")
			return
		}
		obj, e := UnmarshalAnalyticsEngineUserCredentials(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAnalyticsEngineUserCredentialsAsProperty unmarshals an instance of AnalyticsEngineUserCredentials that is stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineUserCredentialsAsProperty(m map[string]interface{}, propertyName string) (result *AnalyticsEngineUserCredentials, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AnalyticsEngineUserCredentials'", propertyName)
			return
		}
		result, err = UnmarshalAnalyticsEngineUserCredentials(objMap)
	}
	return
}

// UnmarshalAnalyticsEngineUserCredentialsSliceAsProperty unmarshals a slice of AnalyticsEngineUserCredentials instances that are stored as a property
// within the specified map.
func UnmarshalAnalyticsEngineUserCredentialsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AnalyticsEngineUserCredentials, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AnalyticsEngineUserCredentials'", propertyName)
			return
		}
		slice, err = UnmarshalAnalyticsEngineUserCredentialsSlice(vSlice)
	}
	return
}

// ConfigureLoggingOptions : The ConfigureLogging options.
type ConfigureLoggingOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Logging specifications on each node.
	LogSpecs []AnalyticsEngineLoggingNodeSpec `json:"log_specs" validate:"required"`

	// Logging server configuration.
	LogServer *AnalyticsEngineLoggingServer `json:"log_server" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewConfigureLoggingOptions : Instantiate ConfigureLoggingOptions
func (*IbmAnalyticsEngineApiV2) NewConfigureLoggingOptions(instanceGuid string, logSpecs []AnalyticsEngineLoggingNodeSpec, logServer *AnalyticsEngineLoggingServer) *ConfigureLoggingOptions {
	return &ConfigureLoggingOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		LogSpecs: logSpecs,
		LogServer: logServer,
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *ConfigureLoggingOptions) SetInstanceGuid(instanceGuid string) *ConfigureLoggingOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetLogSpecs : Allow user to set LogSpecs
func (options *ConfigureLoggingOptions) SetLogSpecs(logSpecs []AnalyticsEngineLoggingNodeSpec) *ConfigureLoggingOptions {
	options.LogSpecs = logSpecs
	return options
}

// SetLogServer : Allow user to set LogServer
func (options *ConfigureLoggingOptions) SetLogServer(logServer *AnalyticsEngineLoggingServer) *ConfigureLoggingOptions {
	options.LogServer = logServer
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ConfigureLoggingOptions) SetHeaders(param map[string]string) *ConfigureLoggingOptions {
	options.Headers = param
	return options
}

// CreateCustomizationRequestOptions : The CreateCustomizationRequest options.
type CreateCustomizationRequestOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Type of nodes to target for this customization.
	Target *string `json:"target" validate:"required"`

	// List of custom actions.
	CustomActions []AnalyticsEngineCustomAction `json:"custom_actions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCustomizationRequestOptions.Target property.
// Type of nodes to target for this customization.
const (
	CreateCustomizationRequestOptions_Target_All = "all"
	CreateCustomizationRequestOptions_Target_Data = "data"
	CreateCustomizationRequestOptions_Target_MasterManagement = "master-management"
)

// NewCreateCustomizationRequestOptions : Instantiate CreateCustomizationRequestOptions
func (*IbmAnalyticsEngineApiV2) NewCreateCustomizationRequestOptions(instanceGuid string, target string, customActions []AnalyticsEngineCustomAction) *CreateCustomizationRequestOptions {
	return &CreateCustomizationRequestOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		Target: core.StringPtr(target),
		CustomActions: customActions,
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *CreateCustomizationRequestOptions) SetInstanceGuid(instanceGuid string) *CreateCustomizationRequestOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetTarget : Allow user to set Target
func (options *CreateCustomizationRequestOptions) SetTarget(target string) *CreateCustomizationRequestOptions {
	options.Target = core.StringPtr(target)
	return options
}

// SetCustomActions : Allow user to set CustomActions
func (options *CreateCustomizationRequestOptions) SetCustomActions(customActions []AnalyticsEngineCustomAction) *CreateCustomizationRequestOptions {
	options.CustomActions = customActions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomizationRequestOptions) SetHeaders(param map[string]string) *CreateCustomizationRequestOptions {
	options.Headers = param
	return options
}

// DeleteLoggingConfigOptions : The DeleteLoggingConfig options.
type DeleteLoggingConfigOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLoggingConfigOptions : Instantiate DeleteLoggingConfigOptions
func (*IbmAnalyticsEngineApiV2) NewDeleteLoggingConfigOptions(instanceGuid string) *DeleteLoggingConfigOptions {
	return &DeleteLoggingConfigOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *DeleteLoggingConfigOptions) SetInstanceGuid(instanceGuid string) *DeleteLoggingConfigOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLoggingConfigOptions) SetHeaders(param map[string]string) *DeleteLoggingConfigOptions {
	options.Headers = param
	return options
}

// GetAllAnalyticsEnginesOptions : The GetAllAnalyticsEngines options.
type GetAllAnalyticsEnginesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAllAnalyticsEnginesOptions : Instantiate GetAllAnalyticsEnginesOptions
func (*IbmAnalyticsEngineApiV2) NewGetAllAnalyticsEnginesOptions() *GetAllAnalyticsEnginesOptions {
	return &GetAllAnalyticsEnginesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAllAnalyticsEnginesOptions) SetHeaders(param map[string]string) *GetAllAnalyticsEnginesOptions {
	options.Headers = param
	return options
}

// GetAllCustomizationRequestsOptions : The GetAllCustomizationRequests options.
type GetAllCustomizationRequestsOptions struct {
	// service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAllCustomizationRequestsOptions : Instantiate GetAllCustomizationRequestsOptions
func (*IbmAnalyticsEngineApiV2) NewGetAllCustomizationRequestsOptions(instanceGuid string) *GetAllCustomizationRequestsOptions {
	return &GetAllCustomizationRequestsOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *GetAllCustomizationRequestsOptions) SetInstanceGuid(instanceGuid string) *GetAllCustomizationRequestsOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAllCustomizationRequestsOptions) SetHeaders(param map[string]string) *GetAllCustomizationRequestsOptions {
	options.Headers = param
	return options
}

// GetAnalyticsEngineByIdOptions : The GetAnalyticsEngineByID options.
type GetAnalyticsEngineByIdOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAnalyticsEngineByIdOptions : Instantiate GetAnalyticsEngineByIdOptions
func (*IbmAnalyticsEngineApiV2) NewGetAnalyticsEngineByIdOptions(instanceGuid string) *GetAnalyticsEngineByIdOptions {
	return &GetAnalyticsEngineByIdOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *GetAnalyticsEngineByIdOptions) SetInstanceGuid(instanceGuid string) *GetAnalyticsEngineByIdOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAnalyticsEngineByIdOptions) SetHeaders(param map[string]string) *GetAnalyticsEngineByIdOptions {
	options.Headers = param
	return options
}

// GetAnalyticsEngineStateByIdOptions : The GetAnalyticsEngineStateByID options.
type GetAnalyticsEngineStateByIdOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAnalyticsEngineStateByIdOptions : Instantiate GetAnalyticsEngineStateByIdOptions
func (*IbmAnalyticsEngineApiV2) NewGetAnalyticsEngineStateByIdOptions(instanceGuid string) *GetAnalyticsEngineStateByIdOptions {
	return &GetAnalyticsEngineStateByIdOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *GetAnalyticsEngineStateByIdOptions) SetInstanceGuid(instanceGuid string) *GetAnalyticsEngineStateByIdOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAnalyticsEngineStateByIdOptions) SetHeaders(param map[string]string) *GetAnalyticsEngineStateByIdOptions {
	options.Headers = param
	return options
}

// GetCustomizationRequestByIdOptions : The GetCustomizationRequestByID options.
type GetCustomizationRequestByIdOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// customization request ID.
	RequestID *string `json:"request_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomizationRequestByIdOptions : Instantiate GetCustomizationRequestByIdOptions
func (*IbmAnalyticsEngineApiV2) NewGetCustomizationRequestByIdOptions(instanceGuid string, requestID string) *GetCustomizationRequestByIdOptions {
	return &GetCustomizationRequestByIdOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		RequestID: core.StringPtr(requestID),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *GetCustomizationRequestByIdOptions) SetInstanceGuid(instanceGuid string) *GetCustomizationRequestByIdOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetRequestID : Allow user to set RequestID
func (options *GetCustomizationRequestByIdOptions) SetRequestID(requestID string) *GetCustomizationRequestByIdOptions {
	options.RequestID = core.StringPtr(requestID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomizationRequestByIdOptions) SetHeaders(param map[string]string) *GetCustomizationRequestByIdOptions {
	options.Headers = param
	return options
}

// GetLoggingConfigOptions : The GetLoggingConfig options.
type GetLoggingConfigOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoggingConfigOptions : Instantiate GetLoggingConfigOptions
func (*IbmAnalyticsEngineApiV2) NewGetLoggingConfigOptions(instanceGuid string) *GetLoggingConfigOptions {
	return &GetLoggingConfigOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *GetLoggingConfigOptions) SetInstanceGuid(instanceGuid string) *GetLoggingConfigOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoggingConfigOptions) SetHeaders(param map[string]string) *GetLoggingConfigOptions {
	options.Headers = param
	return options
}

// ResetClusterPasswordOptions : The ResetClusterPassword options.
type ResetClusterPasswordOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResetClusterPasswordOptions : Instantiate ResetClusterPasswordOptions
func (*IbmAnalyticsEngineApiV2) NewResetClusterPasswordOptions(instanceGuid string) *ResetClusterPasswordOptions {
	return &ResetClusterPasswordOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *ResetClusterPasswordOptions) SetInstanceGuid(instanceGuid string) *ResetClusterPasswordOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ResetClusterPasswordOptions) SetHeaders(param map[string]string) *ResetClusterPasswordOptions {
	options.Headers = param
	return options
}

// ResizeClusterOptions : The ResizeCluster options.
type ResizeClusterOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"instance_guid" validate:"required"`

	// Expected number of nodes in the cluster after the resize operation.
	ComputeNodesCount *int64 `json:"compute_nodes_count,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResizeClusterOptions : Instantiate ResizeClusterOptions
func (*IbmAnalyticsEngineApiV2) NewResizeClusterOptions(instanceGuid string) *ResizeClusterOptions {
	return &ResizeClusterOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (options *ResizeClusterOptions) SetInstanceGuid(instanceGuid string) *ResizeClusterOptions {
	options.InstanceGuid = core.StringPtr(instanceGuid)
	return options
}

// SetComputeNodesCount : Allow user to set ComputeNodesCount
func (options *ResizeClusterOptions) SetComputeNodesCount(computeNodesCount int64) *ResizeClusterOptions {
	options.ComputeNodesCount = core.Int64Ptr(computeNodesCount)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ResizeClusterOptions) SetHeaders(param map[string]string) *ResizeClusterOptions {
	options.Headers = param
	return options
}

// ServiceEndpoints : Service Endpoints.
type ServiceEndpoints struct {
	// Phoenix JDBC service endpoint.
	PhoenixJdbc *string `json:"phoenix_jdbc,omitempty"`

	// Amabri console service endpoint.
	AmbariConsole *string `json:"ambari_console,omitempty"`

	// Livy service endpoint.
	Livy *string `json:"livy,omitempty"`

	// Spark history server serivce endpoint.
	SparkHistoryServer *string `json:"spark_history_server,omitempty"`

	// Oozie REST service endpi'.
	OozieRest *string `json:"oozie_rest,omitempty"`

	// Hive JDBC service endpoint.
	HiveJdbc *string `json:"hive_jdbc,omitempty"`

	// Notebook gateway websocket service endpoint.
	NotebookGatewayWebsocket *string `json:"notebook_gateway_websocket,omitempty"`

	// Notebook gateway service endpoint.
	NotebookGateway *string `json:"notebook_gateway,omitempty"`

	// WebHDFS service endpoint.
	Webhdfs *string `json:"webhdfs,omitempty"`

	// SSH service endpoint.
	Ssh *string `json:"ssh,omitempty"`

	// Spark SQL service endpoint.
	SparkSql *string `json:"spark_sql,omitempty"`
}


// UnmarshalServiceEndpoints constructs an instance of ServiceEndpoints from the specified map.
func UnmarshalServiceEndpoints(m map[string]interface{}) (result *ServiceEndpoints, err error) {
	obj := new(ServiceEndpoints)
	obj.PhoenixJdbc, err = core.UnmarshalString(m, "phoenix_jdbc")
	if err != nil {
		return
	}
	obj.AmbariConsole, err = core.UnmarshalString(m, "ambari_console")
	if err != nil {
		return
	}
	obj.Livy, err = core.UnmarshalString(m, "livy")
	if err != nil {
		return
	}
	obj.SparkHistoryServer, err = core.UnmarshalString(m, "spark_history_server")
	if err != nil {
		return
	}
	obj.OozieRest, err = core.UnmarshalString(m, "oozie_rest")
	if err != nil {
		return
	}
	obj.HiveJdbc, err = core.UnmarshalString(m, "hive_jdbc")
	if err != nil {
		return
	}
	obj.NotebookGatewayWebsocket, err = core.UnmarshalString(m, "notebook_gateway_websocket")
	if err != nil {
		return
	}
	obj.NotebookGateway, err = core.UnmarshalString(m, "notebook_gateway")
	if err != nil {
		return
	}
	obj.Webhdfs, err = core.UnmarshalString(m, "webhdfs")
	if err != nil {
		return
	}
	obj.Ssh, err = core.UnmarshalString(m, "ssh")
	if err != nil {
		return
	}
	obj.SparkSql, err = core.UnmarshalString(m, "spark_sql")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalServiceEndpointsSlice unmarshals a slice of ServiceEndpoints instances from the specified list of maps.
func UnmarshalServiceEndpointsSlice(s []interface{}) (slice []ServiceEndpoints, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ServiceEndpoints'")
			return
		}
		obj, e := UnmarshalServiceEndpoints(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalServiceEndpointsAsProperty unmarshals an instance of ServiceEndpoints that is stored as a property
// within the specified map.
func UnmarshalServiceEndpointsAsProperty(m map[string]interface{}, propertyName string) (result *ServiceEndpoints, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ServiceEndpoints'", propertyName)
			return
		}
		result, err = UnmarshalServiceEndpoints(objMap)
	}
	return
}

// UnmarshalServiceEndpointsSliceAsProperty unmarshals a slice of ServiceEndpoints instances that are stored as a property
// within the specified map.
func UnmarshalServiceEndpointsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ServiceEndpoints, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ServiceEndpoints'", propertyName)
			return
		}
		slice, err = UnmarshalServiceEndpointsSlice(vSlice)
	}
	return
}
