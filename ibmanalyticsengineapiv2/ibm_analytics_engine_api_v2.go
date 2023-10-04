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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.37.0-a85661cd-20210802-190136
 */

// Package ibmanalyticsengineapiv2 : Operations and models for the IbmAnalyticsEngineApiV2 service
package ibmanalyticsengineapiv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/ibm-iae-go-sdk/v2/common"
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
const DefaultServiceURL = "https://ibm-analytics-engine-api.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_analytics_engine_api" //#nosec G101

// IbmAnalyticsEngineApiV2Options : Service options
type IbmAnalyticsEngineApiV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmAnalyticsEngineApiV2UsingExternalConfig : constructs an instance of IbmAnalyticsEngineApiV2 with passed in options and external configuration.
func NewIbmAnalyticsEngineApiV2UsingExternalConfig(options *IbmAnalyticsEngineApiV2Options) (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmAnalyticsEngineApi, err = NewIbmAnalyticsEngineApiV2(options)
	if err != nil {
		return
	}

	err = ibmAnalyticsEngineApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmAnalyticsEngineApi.Service.SetServiceURL(options.URL)
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

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "ibmAnalyticsEngineApi" suitable for processing requests.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) Clone() *IbmAnalyticsEngineApiV2 {
	if core.IsNil(ibmAnalyticsEngineApi) {
		return nil
	}
	clone := *ibmAnalyticsEngineApi
	clone.Service = ibmAnalyticsEngineApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) SetServiceURL(url string) error {
	return ibmAnalyticsEngineApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetServiceURL() string {
	return ibmAnalyticsEngineApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) SetDefaultHeaders(headers http.Header) {
	ibmAnalyticsEngineApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) SetEnableGzipCompression(enableGzip bool) {
	ibmAnalyticsEngineApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetEnableGzipCompression() bool {
	return ibmAnalyticsEngineApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmAnalyticsEngineApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) DisableRetries() {
	ibmAnalyticsEngineApi.Service.DisableRetries()
}

// GetAllAnalyticsEngines : List all Analytics Engines
// Currently, you cannot fetch the list of all IBM Analytics Engine service instances through this REST API. You should
// use the IBM Cloud CLI instead.  For example, ```ibmcloud resource service-instances --service-name
// ibmanalyticsengine```.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAllAnalyticsEngines(getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetAllAnalyticsEnginesWithContext(context.Background(), getAllAnalyticsEnginesOptions)
}

// GetAllAnalyticsEnginesWithContext is an alternate form of the GetAllAnalyticsEngines method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAllAnalyticsEnginesWithContext(ctx context.Context, getAllAnalyticsEnginesOptions *GetAllAnalyticsEnginesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAllAnalyticsEnginesOptions, "getAllAnalyticsEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAllAnalyticsEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetAllAnalyticsEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApi.Service.Request(request, nil)

	return
}

// GetAnalyticsEngineByID : Get details of Analytics Engine
// Retrieves the following details of the IBM Analytics Engine service instance:
// * Hardware size and software package
//
//   - Timestamps at which the cluster was created, deleted or updated
//
//   - Service endpoint URLs
//
//     **NOTE:** No credentials are returned. You can get the IBM Analytics Engine service instance credentials by invoking
//
// the reset_password REST API.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAnalyticsEngineByID(getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions) (result *AnalyticsEngine, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetAnalyticsEngineByIDWithContext(context.Background(), getAnalyticsEngineByIdOptions)
}

// GetAnalyticsEngineByIDWithContext is an alternate form of the GetAnalyticsEngineByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAnalyticsEngineByIDWithContext(ctx context.Context, getAnalyticsEngineByIdOptions *GetAnalyticsEngineByIdOptions) (result *AnalyticsEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAnalyticsEngineByIdOptions, "getAnalyticsEngineByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAnalyticsEngineByIdOptions, "getAnalyticsEngineByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getAnalyticsEngineByIdOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnalyticsEngineByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetAnalyticsEngineByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngine)
		if err != nil {
			return
		}
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
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAnalyticsEngineStateByID(getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions) (result *AnalyticsEngineState, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetAnalyticsEngineStateByIDWithContext(context.Background(), getAnalyticsEngineStateByIdOptions)
}

// GetAnalyticsEngineStateByIDWithContext is an alternate form of the GetAnalyticsEngineStateByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAnalyticsEngineStateByIDWithContext(ctx context.Context, getAnalyticsEngineStateByIdOptions *GetAnalyticsEngineStateByIdOptions) (result *AnalyticsEngineState, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAnalyticsEngineStateByIdOptions, "getAnalyticsEngineStateByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAnalyticsEngineStateByIdOptions, "getAnalyticsEngineStateByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getAnalyticsEngineStateByIdOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/state`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnalyticsEngineStateByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetAnalyticsEngineStateByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineState)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomizationRequest : Create an adhoc customization request
// Creates a new adhoc customization request. Adhoc customization scripts can be run only once. They are not persisted
// with the cluster and are not run automatically when more nodes are added to the cluster.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) CreateCustomizationRequest(createCustomizationRequestOptions *CreateCustomizationRequestOptions) (result *AnalyticsEngineCreateCustomizationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.CreateCustomizationRequestWithContext(context.Background(), createCustomizationRequestOptions)
}

// CreateCustomizationRequestWithContext is an alternate form of the CreateCustomizationRequest method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) CreateCustomizationRequestWithContext(ctx context.Context, createCustomizationRequestOptions *CreateCustomizationRequestOptions) (result *AnalyticsEngineCreateCustomizationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomizationRequestOptions, "createCustomizationRequestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomizationRequestOptions, "createCustomizationRequestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *createCustomizationRequestOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/customization_requests`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomizationRequestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "CreateCustomizationRequest")
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

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineCreateCustomizationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAllCustomizationRequests : Get all customization requests run on an Analytics Engine cluster
// Retrieves the request_id of all customization requests submitted to the specified Analytics Engine cluster.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAllCustomizationRequests(getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions) (result []AnalyticsEngineCustomizationRequestCollectionItem, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetAllCustomizationRequestsWithContext(context.Background(), getAllCustomizationRequestsOptions)
}

// GetAllCustomizationRequestsWithContext is an alternate form of the GetAllCustomizationRequests method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetAllCustomizationRequestsWithContext(ctx context.Context, getAllCustomizationRequestsOptions *GetAllCustomizationRequestsOptions) (result []AnalyticsEngineCustomizationRequestCollectionItem, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAllCustomizationRequestsOptions, "getAllCustomizationRequestsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAllCustomizationRequestsOptions, "getAllCustomizationRequestsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getAllCustomizationRequestsOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/customization_requests`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAllCustomizationRequestsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetAllCustomizationRequests")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse []json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineCustomizationRequestCollectionItem)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCustomizationRequestByID : Retrieve details of specified customization request ID
// Retrieves the status of the specified customization request, along with pointers to log files generated during the
// run.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetCustomizationRequestByID(getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions) (result *AnalyticsEngineCustomizationRunDetails, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetCustomizationRequestByIDWithContext(context.Background(), getCustomizationRequestByIdOptions)
}

// GetCustomizationRequestByIDWithContext is an alternate form of the GetCustomizationRequestByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetCustomizationRequestByIDWithContext(ctx context.Context, getCustomizationRequestByIdOptions *GetCustomizationRequestByIdOptions) (result *AnalyticsEngineCustomizationRunDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomizationRequestByIdOptions, "getCustomizationRequestByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomizationRequestByIdOptions, "getCustomizationRequestByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getCustomizationRequestByIdOptions.InstanceGuid,
		"request_id":    *getCustomizationRequestByIdOptions.RequestID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/customization_requests/{request_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomizationRequestByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetCustomizationRequestByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineCustomizationRunDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ResizeCluster : Resize the cluster
// Resizes the cluster by adjusting the number of compute and task nodes. Task nodes can be added and removed. Compute
// nodes, once added, can't be removed.
//
// **Note:**
//
//  1. You can't modify the number of compute nodes and tasks nodes in the same request.
//
// 2. You can't modify the number of task nodes if you enabled auto scaling when you created the cluster.
//
// 3. Task nodes are not supported on Lite plan clusters.
//
// 4. You can't resize the cluster if the software package on the cluster is deprecated or doesn't permit cluster
// resizing. See [here](https://cloud.ibm.com/docs/AnalyticsEngine?topic=AnalyticsEngine-unsupported-operations).
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ResizeCluster(resizeClusterOptions *ResizeClusterOptions) (result *AnalyticsEngineResizeClusterResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ResizeClusterWithContext(context.Background(), resizeClusterOptions)
}

// ResizeClusterWithContext is an alternate form of the ResizeCluster method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ResizeClusterWithContext(ctx context.Context, resizeClusterOptions *ResizeClusterOptions) (result *AnalyticsEngineResizeClusterResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resizeClusterOptions, "resizeClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resizeClusterOptions, "resizeClusterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *resizeClusterOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/resize`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resizeClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "ResizeCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(resizeClusterOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineResizeClusterResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ResetClusterPassword : Reset cluster password
// Resets the cluster's password to a new system-generated crytographically strong value.  The new password is included
// in the response and you should make a note of it.  This password is displayed only once here and cannot be retrieved
// later.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ResetClusterPassword(resetClusterPasswordOptions *ResetClusterPasswordOptions) (result *AnalyticsEngineResetClusterPasswordResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ResetClusterPasswordWithContext(context.Background(), resetClusterPasswordOptions)
}

// ResetClusterPasswordWithContext is an alternate form of the ResetClusterPassword method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ResetClusterPasswordWithContext(ctx context.Context, resetClusterPasswordOptions *ResetClusterPasswordOptions) (result *AnalyticsEngineResetClusterPasswordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resetClusterPasswordOptions, "resetClusterPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resetClusterPasswordOptions, "resetClusterPasswordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *resetClusterPasswordOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/reset_password`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resetClusterPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "ResetClusterPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineResetClusterPasswordResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ConfigureLogging : Configure log aggregation
// Collects the logs for the following components in an IBM Analytics Engine cluster:
// * IBM Analytics Engine daemon logs, for example those for Spark, Hive, Yarn, and Knox on the management, data and
// task nodes
// * Yarn application job logs.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ConfigureLogging(configureLoggingOptions *ConfigureLoggingOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ConfigureLoggingWithContext(context.Background(), configureLoggingOptions)
}

// ConfigureLoggingWithContext is an alternate form of the ConfigureLogging method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) ConfigureLoggingWithContext(ctx context.Context, configureLoggingOptions *ConfigureLoggingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(configureLoggingOptions, "configureLoggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(configureLoggingOptions, "configureLoggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *configureLoggingOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/log_config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range configureLoggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "ConfigureLogging")
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

	response, err = ibmAnalyticsEngineApi.Service.Request(request, nil)

	return
}

// GetLoggingConfig : Retrieve the status of log configuration
// Retrieves the status and details of the log configuration for your cluster.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetLoggingConfig(getLoggingConfigOptions *GetLoggingConfigOptions) (result *AnalyticsEngineLoggingConfigDetails, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetLoggingConfigWithContext(context.Background(), getLoggingConfigOptions)
}

// GetLoggingConfigWithContext is an alternate form of the GetLoggingConfig method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) GetLoggingConfigWithContext(ctx context.Context, getLoggingConfigOptions *GetLoggingConfigOptions) (result *AnalyticsEngineLoggingConfigDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoggingConfigOptions, "getLoggingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoggingConfigOptions, "getLoggingConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getLoggingConfigOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/log_config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoggingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "GetLoggingConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineLoggingConfigDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteLoggingConfig : Delete the log configuration
// Deletes the log configuration. This operation stops sending logs to the centralized log server.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) DeleteLoggingConfig(deleteLoggingConfigOptions *DeleteLoggingConfigOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.DeleteLoggingConfigWithContext(context.Background(), deleteLoggingConfigOptions)
}

// DeleteLoggingConfigWithContext is an alternate form of the DeleteLoggingConfig method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) DeleteLoggingConfigWithContext(ctx context.Context, deleteLoggingConfigOptions *DeleteLoggingConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoggingConfigOptions, "deleteLoggingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoggingConfigOptions, "deleteLoggingConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *deleteLoggingConfigOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/log_config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoggingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "DeleteLoggingConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApi.Service.Request(request, nil)

	return
}

// UpdatePrivateEndpointWhitelist : Update private endpoint whitelist
// Updates the list of whitelisted private endpoints. This operation either adds ip ranges to the whitelist or deletes
// them.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) UpdatePrivateEndpointWhitelist(updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions) (result *AnalyticsEngineWhitelistResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.UpdatePrivateEndpointWhitelistWithContext(context.Background(), updatePrivateEndpointWhitelistOptions)
}

// UpdatePrivateEndpointWhitelistWithContext is an alternate form of the UpdatePrivateEndpointWhitelist method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV2) UpdatePrivateEndpointWhitelistWithContext(ctx context.Context, updatePrivateEndpointWhitelistOptions *UpdatePrivateEndpointWhitelistOptions) (result *AnalyticsEngineWhitelistResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePrivateEndpointWhitelistOptions, "updatePrivateEndpointWhitelistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePrivateEndpointWhitelistOptions, "updatePrivateEndpointWhitelistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *updatePrivateEndpointWhitelistOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v2/analytics_engines/{instance_guid}/private_endpoint_whitelist`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePrivateEndpointWhitelistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V2", "UpdatePrivateEndpointWhitelist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updatePrivateEndpointWhitelistOptions.IpRanges != nil {
		body["ip_ranges"] = updatePrivateEndpointWhitelistOptions.IpRanges
	}
	if updatePrivateEndpointWhitelistOptions.Action != nil {
		body["action"] = updatePrivateEndpointWhitelistOptions.Action
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmAnalyticsEngineApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticsEngineWhitelistResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

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

	// Whitelisted IP Ranges for Analytics Engine Service with private endpoints.
	PrivateEndpointWhitelist []string `json:"private_endpoint_whitelist,omitempty"`
}

// UnmarshalAnalyticsEngine unmarshals an instance of AnalyticsEngine from the specified map of raw messages.
func UnmarshalAnalyticsEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngine)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_plan", &obj.ServicePlan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hardware_size", &obj.HardwareSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "software_package", &obj.SoftwarePackage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creation_time", &obj.CreationTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "commision_time", &obj.CommisionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "decommision_time", &obj.DecommisionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deletion_time", &obj.DeletionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_change_time", &obj.StateChangeTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nodes", &obj.Nodes, UnmarshalAnalyticsEngineClusterNode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "user_credentials", &obj.UserCredentials, UnmarshalAnalyticsEngineUserCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "service_endpoints", &obj.ServiceEndpoints, UnmarshalServiceEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "service_endpoints_ip", &obj.ServiceEndpointsIp, UnmarshalServiceEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_endpoint_whitelist", &obj.PrivateEndpointWhitelist)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineClusterNode : Cluster node details.
type AnalyticsEngineClusterNode struct {
	// Node ID.
	ID *int64 `json:"id,omitempty"`

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

// UnmarshalAnalyticsEngineClusterNode unmarshals an instance of AnalyticsEngineClusterNode from the specified map of raw messages.
func UnmarshalAnalyticsEngineClusterNode(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineClusterNode)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fqdn", &obj.Fqdn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_ip", &obj.PublicIp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_ip", &obj.PrivateIp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_change_time", &obj.StateChangeTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "commission_time", &obj.CommissionTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineCreateCustomizationResponse : Create customization request response.
type AnalyticsEngineCreateCustomizationResponse struct {
	// Customization request ID.
	RequestID *int64 `json:"request_id,omitempty"`
}

// UnmarshalAnalyticsEngineCreateCustomizationResponse unmarshals an instance of AnalyticsEngineCreateCustomizationResponse from the specified map of raw messages.
func UnmarshalAnalyticsEngineCreateCustomizationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCreateCustomizationResponse)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineCustomAction(name string) (_model *AnalyticsEngineCustomAction, err error) {
	_model = &AnalyticsEngineCustomAction{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineCustomAction unmarshals an instance of AnalyticsEngineCustomAction from the specified map of raw messages.
func UnmarshalAnalyticsEngineCustomAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCustomAction)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "script", &obj.Script, UnmarshalAnalyticsEngineCustomActionScript)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "script_params", &obj.ScriptParams)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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
	AnalyticsEngineCustomActionScript_SourceType_Bluemixswift   = "BluemixSwift"
	AnalyticsEngineCustomActionScript_SourceType_Coss3          = "CosS3"
	AnalyticsEngineCustomActionScript_SourceType_Http           = "http"
	AnalyticsEngineCustomActionScript_SourceType_Https          = "https"
	AnalyticsEngineCustomActionScript_SourceType_Softlayerswift = "SoftLayerSwift"
)

// UnmarshalAnalyticsEngineCustomActionScript unmarshals an instance of AnalyticsEngineCustomActionScript from the specified map of raw messages.
func UnmarshalAnalyticsEngineCustomActionScript(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCustomActionScript)
	err = core.UnmarshalPrimitive(m, "source_type", &obj.SourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "script_path", &obj.ScriptPath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_props", &obj.SourceProps)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineCustomizationRequestCollectionItem : AnalyticsEngineCustomizationRequestCollectionItem struct
type AnalyticsEngineCustomizationRequestCollectionItem struct {
	// Customization request ID.
	ID *string `json:"id,omitempty"`
}

// UnmarshalAnalyticsEngineCustomizationRequestCollectionItem unmarshals an instance of AnalyticsEngineCustomizationRequestCollectionItem from the specified map of raw messages.
func UnmarshalAnalyticsEngineCustomizationRequestCollectionItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCustomizationRequestCollectionItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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

// UnmarshalAnalyticsEngineCustomizationRunDetails unmarshals an instance of AnalyticsEngineCustomizationRunDetails from the specified map of raw messages.
func UnmarshalAnalyticsEngineCustomizationRunDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCustomizationRunDetails)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_status", &obj.RunStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "run_details", &obj.RunDetails, UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineCustomizationRunDetailsRunDetails : Customization run details.
type AnalyticsEngineCustomizationRunDetailsRunDetails struct {
	// Customization run overall status.
	OverallStatus *string `json:"overall_status,omitempty"`

	// Customization run details for each node.
	Details []AnalyticsEngineNodeLevelCustomizationRunDetails `json:"details,omitempty"`
}

// UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails unmarshals an instance of AnalyticsEngineCustomizationRunDetailsRunDetails from the specified map of raw messages.
func UnmarshalAnalyticsEngineCustomizationRunDetailsRunDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineCustomizationRunDetailsRunDetails)
	err = core.UnmarshalPrimitive(m, "overall_status", &obj.OverallStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "details", &obj.Details, UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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

// UnmarshalAnalyticsEngineLoggingConfigDetails unmarshals an instance of AnalyticsEngineLoggingConfigDetails from the specified map of raw messages.
func UnmarshalAnalyticsEngineLoggingConfigDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineLoggingConfigDetails)
	err = core.UnmarshalModel(m, "log_specs", &obj.LogSpecs, UnmarshalAnalyticsEngineLoggingNodeSpec)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_server", &obj.LogServer, UnmarshalAnalyticsEngineLoggingServer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_config_status", &obj.LogConfigStatus, UnmarshalAnalyticsEngineLoggingConfigStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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
	AnalyticsEngineLoggingConfigStatus_NodeType_Data       = "data"
	AnalyticsEngineLoggingConfigStatus_NodeType_Management = "management"
	AnalyticsEngineLoggingConfigStatus_NodeType_Task       = "task"
)

// UnmarshalAnalyticsEngineLoggingConfigStatus unmarshals an instance of AnalyticsEngineLoggingConfigStatus from the specified map of raw messages.
func UnmarshalAnalyticsEngineLoggingConfigStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineLoggingConfigStatus)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "node_id", &obj.NodeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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
	AnalyticsEngineLoggingNodeSpec_NodeType_Data       = "data"
	AnalyticsEngineLoggingNodeSpec_NodeType_Management = "management"
	AnalyticsEngineLoggingNodeSpec_NodeType_Task       = "task"
)

// Constants associated with the AnalyticsEngineLoggingNodeSpec.Components property.
// Node components to be logged.
const (
	AnalyticsEngineLoggingNodeSpec_Components_AmbariServer    = "ambari-server"
	AnalyticsEngineLoggingNodeSpec_Components_HadoopMapreduce = "hadoop-mapreduce"
	AnalyticsEngineLoggingNodeSpec_Components_HadoopYarn      = "hadoop-yarn"
	AnalyticsEngineLoggingNodeSpec_Components_Hbase           = "hbase"
	AnalyticsEngineLoggingNodeSpec_Components_Hdfs            = "hdfs"
	AnalyticsEngineLoggingNodeSpec_Components_HdfsAudit       = "hdfs-audit"
	AnalyticsEngineLoggingNodeSpec_Components_Hive            = "hive"
	AnalyticsEngineLoggingNodeSpec_Components_Jnbg            = "jnbg"
	AnalyticsEngineLoggingNodeSpec_Components_Knox            = "knox"
	AnalyticsEngineLoggingNodeSpec_Components_KnoxAudit       = "knox-audit"
	AnalyticsEngineLoggingNodeSpec_Components_Livy2           = "livy2"
	AnalyticsEngineLoggingNodeSpec_Components_Spark2          = "spark2"
	AnalyticsEngineLoggingNodeSpec_Components_YarnApps        = "yarn-apps"
)

// NewAnalyticsEngineLoggingNodeSpec : Instantiate AnalyticsEngineLoggingNodeSpec (Generic Model Constructor)
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineLoggingNodeSpec(nodeType string, components []string) (_model *AnalyticsEngineLoggingNodeSpec, err error) {
	_model = &AnalyticsEngineLoggingNodeSpec{
		NodeType:   core.StringPtr(nodeType),
		Components: components,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineLoggingNodeSpec unmarshals an instance of AnalyticsEngineLoggingNodeSpec from the specified map of raw messages.
func UnmarshalAnalyticsEngineLoggingNodeSpec(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineLoggingNodeSpec)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "components", &obj.Components)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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
func (*IbmAnalyticsEngineApiV2) NewAnalyticsEngineLoggingServer(typeVar string, credential string, apiHost string, logHost string) (_model *AnalyticsEngineLoggingServer, err error) {
	_model = &AnalyticsEngineLoggingServer{
		Type:       core.StringPtr(typeVar),
		Credential: core.StringPtr(credential),
		ApiHost:    core.StringPtr(apiHost),
		LogHost:    core.StringPtr(logHost),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAnalyticsEngineLoggingServer unmarshals an instance of AnalyticsEngineLoggingServer from the specified map of raw messages.
func UnmarshalAnalyticsEngineLoggingServer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineLoggingServer)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credential", &obj.Credential)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_host", &obj.ApiHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_host", &obj.LogHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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

// UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails unmarshals an instance of AnalyticsEngineNodeLevelCustomizationRunDetails from the specified map of raw messages.
func UnmarshalAnalyticsEngineNodeLevelCustomizationRunDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineNodeLevelCustomizationRunDetails)
	err = core.UnmarshalPrimitive(m, "node_name", &obj.NodeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time_taken", &obj.TimeTaken)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_file", &obj.LogFile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineResetClusterPasswordResponse : Response for resetting cluster password.
type AnalyticsEngineResetClusterPasswordResponse struct {
	// Instance guid.
	ID *string `json:"id,omitempty"`

	// User credentials.
	UserCredentials *AnalyticsEngineResetClusterPasswordResponseUserCredentials `json:"user_credentials,omitempty"`
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponse unmarshals an instance of AnalyticsEngineResetClusterPasswordResponse from the specified map of raw messages.
func UnmarshalAnalyticsEngineResetClusterPasswordResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineResetClusterPasswordResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "user_credentials", &obj.UserCredentials, UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineResetClusterPasswordResponseUserCredentials : User credentials.
type AnalyticsEngineResetClusterPasswordResponseUserCredentials struct {
	// Username.
	User *string `json:"user,omitempty"`

	// New password.
	Password *string `json:"password,omitempty"`
}

// UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials unmarshals an instance of AnalyticsEngineResetClusterPasswordResponseUserCredentials from the specified map of raw messages.
func UnmarshalAnalyticsEngineResetClusterPasswordResponseUserCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineResetClusterPasswordResponseUserCredentials)
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineResizeClusterResponse : Resize request response.
type AnalyticsEngineResizeClusterResponse struct {
	// Request ID.
	RequestID *string `json:"request_id,omitempty"`
}

// UnmarshalAnalyticsEngineResizeClusterResponse unmarshals an instance of AnalyticsEngineResizeClusterResponse from the specified map of raw messages.
func UnmarshalAnalyticsEngineResizeClusterResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineResizeClusterResponse)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineState : Cluster state.
type AnalyticsEngineState struct {
	// Cluster state.
	State *string `json:"state" validate:"required"`
}

// UnmarshalAnalyticsEngineState unmarshals an instance of AnalyticsEngineState from the specified map of raw messages.
func UnmarshalAnalyticsEngineState(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineState)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineUserCredentials : User credentials.
type AnalyticsEngineUserCredentials struct {
	// Username.
	User *string `json:"user,omitempty"`
}

// UnmarshalAnalyticsEngineUserCredentials unmarshals an instance of AnalyticsEngineUserCredentials from the specified map of raw messages.
func UnmarshalAnalyticsEngineUserCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineUserCredentials)
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticsEngineWhitelistResponse : Whitelisted IP Ranges.
type AnalyticsEngineWhitelistResponse struct {
	// Whitelisted IP Ranges.
	PrivateEndpointWhitelist []string `json:"private_endpoint_whitelist,omitempty"`
}

// UnmarshalAnalyticsEngineWhitelistResponse unmarshals an instance of AnalyticsEngineWhitelistResponse from the specified map of raw messages.
func UnmarshalAnalyticsEngineWhitelistResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngineWhitelistResponse)
	err = core.UnmarshalPrimitive(m, "private_endpoint_whitelist", &obj.PrivateEndpointWhitelist)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigureLoggingOptions : The ConfigureLogging options.
type ConfigureLoggingOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
		LogSpecs:     logSpecs,
		LogServer:    logServer,
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *ConfigureLoggingOptions) SetInstanceGuid(instanceGuid string) *ConfigureLoggingOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetLogSpecs : Allow user to set LogSpecs
func (_options *ConfigureLoggingOptions) SetLogSpecs(logSpecs []AnalyticsEngineLoggingNodeSpec) *ConfigureLoggingOptions {
	_options.LogSpecs = logSpecs
	return _options
}

// SetLogServer : Allow user to set LogServer
func (_options *ConfigureLoggingOptions) SetLogServer(logServer *AnalyticsEngineLoggingServer) *ConfigureLoggingOptions {
	_options.LogServer = logServer
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ConfigureLoggingOptions) SetHeaders(param map[string]string) *ConfigureLoggingOptions {
	options.Headers = param
	return options
}

// CreateCustomizationRequestOptions : The CreateCustomizationRequest options.
type CreateCustomizationRequestOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
	CreateCustomizationRequestOptions_Target_All              = "all"
	CreateCustomizationRequestOptions_Target_Data             = "data"
	CreateCustomizationRequestOptions_Target_MasterManagement = "master-management"
	CreateCustomizationRequestOptions_Target_Task             = "task"
)

// NewCreateCustomizationRequestOptions : Instantiate CreateCustomizationRequestOptions
func (*IbmAnalyticsEngineApiV2) NewCreateCustomizationRequestOptions(instanceGuid string, target string, customActions []AnalyticsEngineCustomAction) *CreateCustomizationRequestOptions {
	return &CreateCustomizationRequestOptions{
		InstanceGuid:  core.StringPtr(instanceGuid),
		Target:        core.StringPtr(target),
		CustomActions: customActions,
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *CreateCustomizationRequestOptions) SetInstanceGuid(instanceGuid string) *CreateCustomizationRequestOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *CreateCustomizationRequestOptions) SetTarget(target string) *CreateCustomizationRequestOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetCustomActions : Allow user to set CustomActions
func (_options *CreateCustomizationRequestOptions) SetCustomActions(customActions []AnalyticsEngineCustomAction) *CreateCustomizationRequestOptions {
	_options.CustomActions = customActions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomizationRequestOptions) SetHeaders(param map[string]string) *CreateCustomizationRequestOptions {
	options.Headers = param
	return options
}

// DeleteLoggingConfigOptions : The DeleteLoggingConfig options.
type DeleteLoggingConfigOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *DeleteLoggingConfigOptions) SetInstanceGuid(instanceGuid string) *DeleteLoggingConfigOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
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
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *GetAllCustomizationRequestsOptions) SetInstanceGuid(instanceGuid string) *GetAllCustomizationRequestsOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAllCustomizationRequestsOptions) SetHeaders(param map[string]string) *GetAllCustomizationRequestsOptions {
	options.Headers = param
	return options
}

// GetAnalyticsEngineByIdOptions : The GetAnalyticsEngineByID options.
type GetAnalyticsEngineByIdOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *GetAnalyticsEngineByIdOptions) SetInstanceGuid(instanceGuid string) *GetAnalyticsEngineByIdOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAnalyticsEngineByIdOptions) SetHeaders(param map[string]string) *GetAnalyticsEngineByIdOptions {
	options.Headers = param
	return options
}

// GetAnalyticsEngineStateByIdOptions : The GetAnalyticsEngineStateByID options.
type GetAnalyticsEngineStateByIdOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *GetAnalyticsEngineStateByIdOptions) SetInstanceGuid(instanceGuid string) *GetAnalyticsEngineStateByIdOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAnalyticsEngineStateByIdOptions) SetHeaders(param map[string]string) *GetAnalyticsEngineStateByIdOptions {
	options.Headers = param
	return options
}

// GetCustomizationRequestByIdOptions : The GetCustomizationRequestByID options.
type GetCustomizationRequestByIdOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"-" validate:"required,ne="`

	// customization request ID.
	RequestID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomizationRequestByIdOptions : Instantiate GetCustomizationRequestByIdOptions
func (*IbmAnalyticsEngineApiV2) NewGetCustomizationRequestByIdOptions(instanceGuid string, requestID string) *GetCustomizationRequestByIdOptions {
	return &GetCustomizationRequestByIdOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		RequestID:    core.StringPtr(requestID),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *GetCustomizationRequestByIdOptions) SetInstanceGuid(instanceGuid string) *GetCustomizationRequestByIdOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetRequestID : Allow user to set RequestID
func (_options *GetCustomizationRequestByIdOptions) SetRequestID(requestID string) *GetCustomizationRequestByIdOptions {
	_options.RequestID = core.StringPtr(requestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomizationRequestByIdOptions) SetHeaders(param map[string]string) *GetCustomizationRequestByIdOptions {
	options.Headers = param
	return options
}

// GetLoggingConfigOptions : The GetLoggingConfig options.
type GetLoggingConfigOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *GetLoggingConfigOptions) SetInstanceGuid(instanceGuid string) *GetLoggingConfigOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoggingConfigOptions) SetHeaders(param map[string]string) *GetLoggingConfigOptions {
	options.Headers = param
	return options
}

// ResetClusterPasswordOptions : The ResetClusterPassword options.
type ResetClusterPasswordOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"-" validate:"required,ne="`

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
func (_options *ResetClusterPasswordOptions) SetInstanceGuid(instanceGuid string) *ResetClusterPasswordOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResetClusterPasswordOptions) SetHeaders(param map[string]string) *ResetClusterPasswordOptions {
	options.Headers = param
	return options
}

// ResizeClusterOptions : The ResizeCluster options.
type ResizeClusterOptions struct {
	// Service instance GUID.
	InstanceGuid *string `json:"-" validate:"required,ne="`

	// Expected size of the cluster after the resize operation. If the number of nodes in the cluster is 5 and you want to
	// add 2 nodes, specify 7.
	Body ResizeClusterRequestIntf `json:"body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResizeClusterOptions : Instantiate ResizeClusterOptions
func (*IbmAnalyticsEngineApiV2) NewResizeClusterOptions(instanceGuid string, body ResizeClusterRequestIntf) *ResizeClusterOptions {
	return &ResizeClusterOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		Body:         body,
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *ResizeClusterOptions) SetInstanceGuid(instanceGuid string) *ResizeClusterOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetBody : Allow user to set Body
func (_options *ResizeClusterOptions) SetBody(body ResizeClusterRequestIntf) *ResizeClusterOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResizeClusterOptions) SetHeaders(param map[string]string) *ResizeClusterOptions {
	options.Headers = param
	return options
}

// ResizeClusterRequest : ResizeClusterRequest struct
// Models which "extend" this model:
// - ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest
// - ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest
type ResizeClusterRequest struct {
	// Expected number of compute nodes in the cluster after the resize operation.
	ComputeNodesCount *int64 `json:"compute_nodes_count,omitempty"`

	// Expected number of task nodes in the cluster after the resize operation.
	TaskNodesCount *int64 `json:"task_nodes_count,omitempty"`
}

func (*ResizeClusterRequest) isaResizeClusterRequest() bool {
	return true
}

type ResizeClusterRequestIntf interface {
	isaResizeClusterRequest() bool
}

// UnmarshalResizeClusterRequest unmarshals an instance of ResizeClusterRequest from the specified map of raw messages.
func UnmarshalResizeClusterRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResizeClusterRequest)
	err = core.UnmarshalPrimitive(m, "compute_nodes_count", &obj.ComputeNodesCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "task_nodes_count", &obj.TaskNodesCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

// UnmarshalServiceEndpoints unmarshals an instance of ServiceEndpoints from the specified map of raw messages.
func UnmarshalServiceEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceEndpoints)
	err = core.UnmarshalPrimitive(m, "phoenix_jdbc", &obj.PhoenixJdbc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ambari_console", &obj.AmbariConsole)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "livy", &obj.Livy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_history_server", &obj.SparkHistoryServer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "oozie_rest", &obj.OozieRest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hive_jdbc", &obj.HiveJdbc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notebook_gateway_websocket", &obj.NotebookGatewayWebsocket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notebook_gateway", &obj.NotebookGateway)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "webhdfs", &obj.Webhdfs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssh", &obj.Ssh)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql", &obj.SparkSql)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdatePrivateEndpointWhitelistOptions : The UpdatePrivateEndpointWhitelist options.
type UpdatePrivateEndpointWhitelistOptions struct {
	// GUID of the service instance.
	InstanceGuid *string `json:"-" validate:"required,ne="`

	// List of IP ranges to add to or remove from the whitelist.
	IpRanges []string `json:"ip_ranges" validate:"required"`

	// Update Whitelist IP ranges. Add (or) Delete.
	Action *string `json:"action" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdatePrivateEndpointWhitelistOptions.Action property.
// Update Whitelist IP ranges. Add (or) Delete.
const (
	UpdatePrivateEndpointWhitelistOptions_Action_Add    = "add"
	UpdatePrivateEndpointWhitelistOptions_Action_Delete = "delete"
)

// NewUpdatePrivateEndpointWhitelistOptions : Instantiate UpdatePrivateEndpointWhitelistOptions
func (*IbmAnalyticsEngineApiV2) NewUpdatePrivateEndpointWhitelistOptions(instanceGuid string, ipRanges []string, action string) *UpdatePrivateEndpointWhitelistOptions {
	return &UpdatePrivateEndpointWhitelistOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
		IpRanges:     ipRanges,
		Action:       core.StringPtr(action),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *UpdatePrivateEndpointWhitelistOptions) SetInstanceGuid(instanceGuid string) *UpdatePrivateEndpointWhitelistOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetIpRanges : Allow user to set IpRanges
func (_options *UpdatePrivateEndpointWhitelistOptions) SetIpRanges(ipRanges []string) *UpdatePrivateEndpointWhitelistOptions {
	_options.IpRanges = ipRanges
	return _options
}

// SetAction : Allow user to set Action
func (_options *UpdatePrivateEndpointWhitelistOptions) SetAction(action string) *UpdatePrivateEndpointWhitelistOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrivateEndpointWhitelistOptions) SetHeaders(param map[string]string) *UpdatePrivateEndpointWhitelistOptions {
	options.Headers = param
	return options
}

// ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest : Resize cluster request.
// This model "extends" ResizeClusterRequest
type ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest struct {
	// Expected number of compute nodes in the cluster after the resize operation.
	ComputeNodesCount *int64 `json:"compute_nodes_count,omitempty"`
}

func (*ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest) isaResizeClusterRequest() bool {
	return true
}

// UnmarshalResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest unmarshals an instance of ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest from the specified map of raw messages.
func UnmarshalResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResizeClusterRequestAnalyticsEngineResizeClusterByComputeNodesRequest)
	err = core.UnmarshalPrimitive(m, "compute_nodes_count", &obj.ComputeNodesCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest : ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest struct
// This model "extends" ResizeClusterRequest
type ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest struct {
	// Expected number of task nodes in the cluster after the resize operation.
	TaskNodesCount *int64 `json:"task_nodes_count,omitempty"`
}

func (*ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest) isaResizeClusterRequest() bool {
	return true
}

// UnmarshalResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest unmarshals an instance of ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest from the specified map of raw messages.
func UnmarshalResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResizeClusterRequestAnalyticsEngineResizeClusterByTaskNodesRequest)
	err = core.UnmarshalPrimitive(m, "task_nodes_count", &obj.TaskNodesCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
