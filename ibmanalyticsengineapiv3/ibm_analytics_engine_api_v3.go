/**
 * (C) Copyright IBM Corp. 2023.
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
 * IBM OpenAPI SDK Code Generator Version: 3.54.1-1d9808a7-20220817-143039
 */

// Package ibmanalyticsengineapiv3 : Operations and models for the IbmAnalyticsEngineApiV3 service
package ibmanalyticsengineapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/ibm-iae-go-sdk/v2/common"
	"github.com/go-openapi/strfmt"
)

// IbmAnalyticsEngineApiV3 : Manage serverless Spark instances and run applications.
//
// API Version: 3.0.0
type IbmAnalyticsEngineApiV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.ae.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_analytics_engine_api"

// IbmAnalyticsEngineApiV3Options : Service options
type IbmAnalyticsEngineApiV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmAnalyticsEngineApiV3UsingExternalConfig : constructs an instance of IbmAnalyticsEngineApiV3 with passed in options and external configuration.
func NewIbmAnalyticsEngineApiV3UsingExternalConfig(options *IbmAnalyticsEngineApiV3Options) (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmAnalyticsEngineApi, err = NewIbmAnalyticsEngineApiV3(options)
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

// NewIbmAnalyticsEngineApiV3 : constructs an instance of IbmAnalyticsEngineApiV3 with passed in options.
func NewIbmAnalyticsEngineApiV3(options *IbmAnalyticsEngineApiV3Options) (service *IbmAnalyticsEngineApiV3, err error) {
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

	service = &IbmAnalyticsEngineApiV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://api.us-south.ae.cloud.ibm.com",
		"eu-de": "https://api.eu-de.ae.cloud.ibm.com",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "ibmAnalyticsEngineApi" suitable for processing requests.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) Clone() *IbmAnalyticsEngineApiV3 {
	if core.IsNil(ibmAnalyticsEngineApi) {
		return nil
	}
	clone := *ibmAnalyticsEngineApi
	clone.Service = ibmAnalyticsEngineApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) SetServiceURL(url string) error {
	return ibmAnalyticsEngineApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetServiceURL() string {
	return ibmAnalyticsEngineApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) SetDefaultHeaders(headers http.Header) {
	ibmAnalyticsEngineApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) SetEnableGzipCompression(enableGzip bool) {
	ibmAnalyticsEngineApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetEnableGzipCompression() bool {
	return ibmAnalyticsEngineApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmAnalyticsEngineApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DisableRetries() {
	ibmAnalyticsEngineApi.Service.DisableRetries()
}

// GetInstance : Find Analytics Engine by id
// Retrieve the details of a single Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstance(getInstanceOptions *GetInstanceOptions) (result *Instance, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetInstanceWithContext(context.Background(), getInstanceOptions)
}

// GetInstanceWithContext is an alternate form of the GetInstance method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceWithContext(ctx context.Context, getInstanceOptions *GetInstanceOptions) (result *Instance, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceOptions, "getInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceOptions, "getInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetInstance")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstance)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceState : Find Analytics Engine state by id
// Retrieve the state of a single Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceState(getInstanceStateOptions *GetInstanceStateOptions) (result *InstanceGetStateResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetInstanceStateWithContext(context.Background(), getInstanceStateOptions)
}

// GetInstanceStateWithContext is an alternate form of the GetInstanceState method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceStateWithContext(ctx context.Context, getInstanceStateOptions *GetInstanceStateOptions) (result *InstanceGetStateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceStateOptions, "getInstanceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceStateOptions, "getInstanceStateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceStateOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/state`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetInstanceState")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceGetStateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetInstanceHome : Set instance home
// Provide the details of the Cloud Object Storage instance to associate with the Analytics Engine instance and use as
// 'instance home' if 'instance home' has not already been set.
//
// **Note**: You can set 'instance home' again if the instance is in 'instance_home_creation_failure' state.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) SetInstanceHome(setInstanceHomeOptions *SetInstanceHomeOptions) (result *InstanceHomeResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.SetInstanceHomeWithContext(context.Background(), setInstanceHomeOptions)
}

// SetInstanceHomeWithContext is an alternate form of the SetInstanceHome method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) SetInstanceHomeWithContext(ctx context.Context, setInstanceHomeOptions *SetInstanceHomeOptions) (result *InstanceHomeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setInstanceHomeOptions, "setInstanceHomeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setInstanceHomeOptions, "setInstanceHomeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *setInstanceHomeOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/instance_home`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setInstanceHomeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "SetInstanceHome")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setInstanceHomeOptions.NewInstanceID != nil {
		body["instance_id"] = setInstanceHomeOptions.NewInstanceID
	}
	if setInstanceHomeOptions.NewProvider != nil {
		body["provider"] = setInstanceHomeOptions.NewProvider
	}
	if setInstanceHomeOptions.NewType != nil {
		body["type"] = setInstanceHomeOptions.NewType
	}
	if setInstanceHomeOptions.NewRegion != nil {
		body["region"] = setInstanceHomeOptions.NewRegion
	}
	if setInstanceHomeOptions.NewEndpoint != nil {
		body["endpoint"] = setInstanceHomeOptions.NewEndpoint
	}
	if setInstanceHomeOptions.NewHmacAccessKey != nil {
		body["hmac_access_key"] = setInstanceHomeOptions.NewHmacAccessKey
	}
	if setInstanceHomeOptions.NewHmacSecretKey != nil {
		body["hmac_secret_key"] = setInstanceHomeOptions.NewHmacSecretKey
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceHomeResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceDefaultConfigs : Get instance default Spark configurations
// Get the default Spark configuration properties that will be applied to all applications of the instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceDefaultConfigs(getInstanceDefaultConfigsOptions *GetInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetInstanceDefaultConfigsWithContext(context.Background(), getInstanceDefaultConfigsOptions)
}

// GetInstanceDefaultConfigsWithContext is an alternate form of the GetInstanceDefaultConfigs method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceDefaultConfigsWithContext(ctx context.Context, getInstanceDefaultConfigsOptions *GetInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceDefaultConfigsOptions, "getInstanceDefaultConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceDefaultConfigsOptions, "getInstanceDefaultConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceDefaultConfigsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/default_configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceDefaultConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetInstanceDefaultConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApi.Service.Request(request, &result)

	return
}

// ReplaceInstanceDefaultConfigs : Replace instance default Spark configurations
// Replace the default Spark configuration properties that will be applied to all applications of the instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceInstanceDefaultConfigs(replaceInstanceDefaultConfigsOptions *ReplaceInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ReplaceInstanceDefaultConfigsWithContext(context.Background(), replaceInstanceDefaultConfigsOptions)
}

// ReplaceInstanceDefaultConfigsWithContext is an alternate form of the ReplaceInstanceDefaultConfigs method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceInstanceDefaultConfigsWithContext(ctx context.Context, replaceInstanceDefaultConfigsOptions *ReplaceInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceInstanceDefaultConfigsOptions, "replaceInstanceDefaultConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceInstanceDefaultConfigsOptions, "replaceInstanceDefaultConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceInstanceDefaultConfigsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/default_configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceInstanceDefaultConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "ReplaceInstanceDefaultConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(replaceInstanceDefaultConfigsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApi.Service.Request(request, &result)

	return
}

// UpdateInstanceDefaultConfigs : Update instance default Spark configurations
// Update the default Spark configuration properties that will be applied to all applications of the instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) UpdateInstanceDefaultConfigs(updateInstanceDefaultConfigsOptions *UpdateInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.UpdateInstanceDefaultConfigsWithContext(context.Background(), updateInstanceDefaultConfigsOptions)
}

// UpdateInstanceDefaultConfigsWithContext is an alternate form of the UpdateInstanceDefaultConfigs method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) UpdateInstanceDefaultConfigsWithContext(ctx context.Context, updateInstanceDefaultConfigsOptions *UpdateInstanceDefaultConfigsOptions) (result map[string]string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceDefaultConfigsOptions, "updateInstanceDefaultConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateInstanceDefaultConfigsOptions, "updateInstanceDefaultConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceDefaultConfigsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/default_configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateInstanceDefaultConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "UpdateInstanceDefaultConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceDefaultConfigsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmAnalyticsEngineApi.Service.Request(request, &result)

	return
}

// GetInstanceDefaultRuntime : Get instance default runtime
// Get the default runtime environment on which all workloads of the instance will run.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceDefaultRuntime(getInstanceDefaultRuntimeOptions *GetInstanceDefaultRuntimeOptions) (result *Runtime, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetInstanceDefaultRuntimeWithContext(context.Background(), getInstanceDefaultRuntimeOptions)
}

// GetInstanceDefaultRuntimeWithContext is an alternate form of the GetInstanceDefaultRuntime method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceDefaultRuntimeWithContext(ctx context.Context, getInstanceDefaultRuntimeOptions *GetInstanceDefaultRuntimeOptions) (result *Runtime, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceDefaultRuntimeOptions, "getInstanceDefaultRuntimeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceDefaultRuntimeOptions, "getInstanceDefaultRuntimeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceDefaultRuntimeOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/default_runtime`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceDefaultRuntimeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetInstanceDefaultRuntime")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuntime)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceInstanceDefaultRuntime : Replace instance default runtime
// Replace the default runtime environment on which all workloads of the instance will run.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceInstanceDefaultRuntime(replaceInstanceDefaultRuntimeOptions *ReplaceInstanceDefaultRuntimeOptions) (result *Runtime, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ReplaceInstanceDefaultRuntimeWithContext(context.Background(), replaceInstanceDefaultRuntimeOptions)
}

// ReplaceInstanceDefaultRuntimeWithContext is an alternate form of the ReplaceInstanceDefaultRuntime method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceInstanceDefaultRuntimeWithContext(ctx context.Context, replaceInstanceDefaultRuntimeOptions *ReplaceInstanceDefaultRuntimeOptions) (result *Runtime, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceInstanceDefaultRuntimeOptions, "replaceInstanceDefaultRuntimeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceInstanceDefaultRuntimeOptions, "replaceInstanceDefaultRuntimeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceInstanceDefaultRuntimeOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/default_runtime`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceInstanceDefaultRuntimeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "ReplaceInstanceDefaultRuntime")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceInstanceDefaultRuntimeOptions.SparkVersion != nil {
		body["spark_version"] = replaceInstanceDefaultRuntimeOptions.SparkVersion
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuntime)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateApplication : Deploy a Spark application
// Deploys a Spark application on a given serverless Spark instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) CreateApplication(createApplicationOptions *CreateApplicationOptions) (result *ApplicationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.CreateApplicationWithContext(context.Background(), createApplicationOptions)
}

// CreateApplicationWithContext is an alternate form of the CreateApplication method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) CreateApplicationWithContext(ctx context.Context, createApplicationOptions *CreateApplicationOptions) (result *ApplicationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createApplicationOptions, "createApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createApplicationOptions, "createApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createApplicationOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "CreateApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})

	// construct application_details property from exploded vars
	application_details := make(map[string]interface{})
	if createApplicationOptions.Application != nil {
		application_details["application"] = createApplicationOptions.Application
	}
	if createApplicationOptions.Runtime != nil {
		application_details["runtime"] = createApplicationOptions.Runtime
	}
	if createApplicationOptions.Jars != nil {
		application_details["jars"] = createApplicationOptions.Jars
	}
	if createApplicationOptions.Packages != nil {
		application_details["packages"] = createApplicationOptions.Packages
	}
	if createApplicationOptions.Repositories != nil {
		application_details["repositories"] = createApplicationOptions.Repositories
	}
	if createApplicationOptions.Files != nil {
		application_details["files"] = createApplicationOptions.Files
	}
	if createApplicationOptions.Archives != nil {
		application_details["archives"] = createApplicationOptions.Archives
	}
	if createApplicationOptions.Name != nil {
		application_details["name"] = createApplicationOptions.Name
	}
	if createApplicationOptions.Class != nil {
		application_details["class"] = createApplicationOptions.Class
	}
	if createApplicationOptions.Arguments != nil {
		application_details["arguments"] = createApplicationOptions.Arguments
	}
	if createApplicationOptions.Conf != nil {
		application_details["conf"] = createApplicationOptions.Conf
	}
	if createApplicationOptions.Env != nil {
		application_details["env"] = createApplicationOptions.Env
	}
	if len(application_details) > 0 {
		body["application_details"] = application_details
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListApplications : List all Spark applications
// Returns a list of all Spark applications submitted to the specified Analytics Engine instance. The result can be
// filtered by specifying query parameters.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ListApplications(listApplicationsOptions *ListApplicationsOptions) (result *ApplicationCollection, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ListApplicationsWithContext(context.Background(), listApplicationsOptions)
}

// ListApplicationsWithContext is an alternate form of the ListApplications method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ListApplicationsWithContext(ctx context.Context, listApplicationsOptions *ListApplicationsOptions) (result *ApplicationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listApplicationsOptions, "listApplicationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listApplicationsOptions, "listApplicationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listApplicationsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "ListApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listApplicationsOptions.State != nil {
		builder.AddQuery("state", strings.Join(listApplicationsOptions.State, ","))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetApplication : Retrieve the details of a given Spark application
// Gets the details of a given Spark application.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplication(getApplicationOptions *GetApplicationOptions) (result *ApplicationGetResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetApplicationWithContext(context.Background(), getApplicationOptions)
}

// GetApplicationWithContext is an alternate form of the GetApplication method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationWithContext(ctx context.Context, getApplicationOptions *GetApplicationOptions) (result *ApplicationGetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationOptions, "getApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationOptions, "getApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getApplicationOptions.InstanceID,
		"application_id": *getApplicationOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_applications/{application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetApplication")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationGetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteApplication : Stop application
// Stops a running application identified by the app_id identifier. This is an idempotent operation. Performs no action
// if the requested application is already stopped or completed.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.DeleteApplicationWithContext(context.Background(), deleteApplicationOptions)
}

// DeleteApplicationWithContext is an alternate form of the DeleteApplication method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteApplicationWithContext(ctx context.Context, deleteApplicationOptions *DeleteApplicationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteApplicationOptions, "deleteApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteApplicationOptions, "deleteApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteApplicationOptions.InstanceID,
		"application_id": *deleteApplicationOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_applications/{application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "DeleteApplication")
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

// GetApplicationState : Get the status of the application
// Returns the status of the given Spark application.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationState(getApplicationStateOptions *GetApplicationStateOptions) (result *ApplicationGetStateResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetApplicationStateWithContext(context.Background(), getApplicationStateOptions)
}

// GetApplicationStateWithContext is an alternate form of the GetApplicationState method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationStateWithContext(ctx context.Context, getApplicationStateOptions *GetApplicationStateOptions) (result *ApplicationGetStateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationStateOptions, "getApplicationStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationStateOptions, "getApplicationStateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getApplicationStateOptions.InstanceID,
		"application_id": *getApplicationStateOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_applications/{application_id}/state`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetApplicationState")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationGetStateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCurrentResourceConsumption : Get current resource consumption
// Gives the total memory and virtual processor cores allotted to all the applications running in the service instance
// at this point in time. When auto-scaled applications are running, the resources allotted will change over time, based
// on the applications's demands. Note: The consumption is not an indication of actual resource consumption by Spark
// processes. It is the sum of resources allocated to the currently running applications at the time of application
// submission.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetCurrentResourceConsumption(getCurrentResourceConsumptionOptions *GetCurrentResourceConsumptionOptions) (result *CurrentResourceConsumptionResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetCurrentResourceConsumptionWithContext(context.Background(), getCurrentResourceConsumptionOptions)
}

// GetCurrentResourceConsumptionWithContext is an alternate form of the GetCurrentResourceConsumption method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetCurrentResourceConsumptionWithContext(ctx context.Context, getCurrentResourceConsumptionOptions *GetCurrentResourceConsumptionOptions) (result *CurrentResourceConsumptionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCurrentResourceConsumptionOptions, "getCurrentResourceConsumptionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCurrentResourceConsumptionOptions, "getCurrentResourceConsumptionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getCurrentResourceConsumptionOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/current_resource_consumption`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCurrentResourceConsumptionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetCurrentResourceConsumption")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCurrentResourceConsumptionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetResourceConsumptionLimits : Get resource consumption limits
// Returns the maximum total memory and virtual processor cores that can be allotted across all the applications running
// in the service instance at any point in time.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetResourceConsumptionLimits(getResourceConsumptionLimitsOptions *GetResourceConsumptionLimitsOptions) (result *ResourceConsumptionLimitsResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetResourceConsumptionLimitsWithContext(context.Background(), getResourceConsumptionLimitsOptions)
}

// GetResourceConsumptionLimitsWithContext is an alternate form of the GetResourceConsumptionLimits method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetResourceConsumptionLimitsWithContext(ctx context.Context, getResourceConsumptionLimitsOptions *GetResourceConsumptionLimitsOptions) (result *ResourceConsumptionLimitsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceConsumptionLimitsOptions, "getResourceConsumptionLimitsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceConsumptionLimitsOptions, "getResourceConsumptionLimitsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getResourceConsumptionLimitsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/resource_consumption_limits`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceConsumptionLimitsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetResourceConsumptionLimits")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceConsumptionLimitsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceLogForwardingConfig : Replace log forwarding configuration
// Modify the configuration for forwarding logs from the Analytics Engine instance to IBM Log Analysis server. Use this
// endpoint to enable or disable log forwarding.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceLogForwardingConfig(replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions) (result *LogForwardingConfigResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ReplaceLogForwardingConfigWithContext(context.Background(), replaceLogForwardingConfigOptions)
}

// ReplaceLogForwardingConfigWithContext is an alternate form of the ReplaceLogForwardingConfig method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ReplaceLogForwardingConfigWithContext(ctx context.Context, replaceLogForwardingConfigOptions *ReplaceLogForwardingConfigOptions) (result *LogForwardingConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceLogForwardingConfigOptions, "replaceLogForwardingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceLogForwardingConfigOptions, "replaceLogForwardingConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceLogForwardingConfigOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/log_forwarding_config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceLogForwardingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "ReplaceLogForwardingConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceLogForwardingConfigOptions.Enabled != nil {
		body["enabled"] = replaceLogForwardingConfigOptions.Enabled
	}
	if replaceLogForwardingConfigOptions.Sources != nil {
		body["sources"] = replaceLogForwardingConfigOptions.Sources
	}
	if replaceLogForwardingConfigOptions.Tags != nil {
		body["tags"] = replaceLogForwardingConfigOptions.Tags
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogForwardingConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLogForwardingConfig : Get log forwarding configuration
// Retrieve the log forwarding configuration of the Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetLogForwardingConfig(getLogForwardingConfigOptions *GetLogForwardingConfigOptions) (result *LogForwardingConfigResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetLogForwardingConfigWithContext(context.Background(), getLogForwardingConfigOptions)
}

// GetLogForwardingConfigWithContext is an alternate form of the GetLogForwardingConfig method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetLogForwardingConfigWithContext(ctx context.Context, getLogForwardingConfigOptions *GetLogForwardingConfigOptions) (result *LogForwardingConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLogForwardingConfigOptions, "getLogForwardingConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLogForwardingConfigOptions, "getLogForwardingConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getLogForwardingConfigOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/log_forwarding_config`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLogForwardingConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetLogForwardingConfig")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogForwardingConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ConfigurePlatformLogging : Enable or disable log forwarding
// Enable or disable log forwarding from IBM Analytics Engine to IBM Log Analysis server.
// *Note:* Deprecated. Use the log forwarding config api instead.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ConfigurePlatformLogging(configurePlatformLoggingOptions *ConfigurePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.ConfigurePlatformLoggingWithContext(context.Background(), configurePlatformLoggingOptions)
}

// ConfigurePlatformLoggingWithContext is an alternate form of the ConfigurePlatformLogging method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) ConfigurePlatformLoggingWithContext(ctx context.Context, configurePlatformLoggingOptions *ConfigurePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(configurePlatformLoggingOptions, "configurePlatformLoggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(configurePlatformLoggingOptions, "configurePlatformLoggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *configurePlatformLoggingOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_guid}/logging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range configurePlatformLoggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "ConfigurePlatformLogging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if configurePlatformLoggingOptions.Enable != nil {
		body["enable"] = configurePlatformLoggingOptions.Enable
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoggingConfigurationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLoggingConfiguration : Retrieve the logging configuration for a given instance id
// Retrieve the logging configuration of a given Analytics Engine instance.
// *Note:* Deprecated. Use the log forwarding config api instead.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetLoggingConfiguration(getLoggingConfigurationOptions *GetLoggingConfigurationOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetLoggingConfigurationWithContext(context.Background(), getLoggingConfigurationOptions)
}

// GetLoggingConfigurationWithContext is an alternate form of the GetLoggingConfiguration method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetLoggingConfigurationWithContext(ctx context.Context, getLoggingConfigurationOptions *GetLoggingConfigurationOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoggingConfigurationOptions, "getLoggingConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoggingConfigurationOptions, "getLoggingConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *getLoggingConfigurationOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_guid}/logging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoggingConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetLoggingConfiguration")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoggingConfigurationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// StartSparkHistoryServer : Start Spark history server
// Start the Spark history server for the given Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) StartSparkHistoryServer(startSparkHistoryServerOptions *StartSparkHistoryServerOptions) (result *SparkHistoryServerResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.StartSparkHistoryServerWithContext(context.Background(), startSparkHistoryServerOptions)
}

// StartSparkHistoryServerWithContext is an alternate form of the StartSparkHistoryServer method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) StartSparkHistoryServerWithContext(ctx context.Context, startSparkHistoryServerOptions *StartSparkHistoryServerOptions) (result *SparkHistoryServerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startSparkHistoryServerOptions, "startSparkHistoryServerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(startSparkHistoryServerOptions, "startSparkHistoryServerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *startSparkHistoryServerOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_history_server`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range startSparkHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "StartSparkHistoryServer")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkHistoryServerResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSparkHistoryServer : Get Spark history server details
// Get the details of the Spark history server of the given Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetSparkHistoryServer(getSparkHistoryServerOptions *GetSparkHistoryServerOptions) (result *SparkHistoryServerResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetSparkHistoryServerWithContext(context.Background(), getSparkHistoryServerOptions)
}

// GetSparkHistoryServerWithContext is an alternate form of the GetSparkHistoryServer method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetSparkHistoryServerWithContext(ctx context.Context, getSparkHistoryServerOptions *GetSparkHistoryServerOptions) (result *SparkHistoryServerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkHistoryServerOptions, "getSparkHistoryServerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSparkHistoryServerOptions, "getSparkHistoryServerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getSparkHistoryServerOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_history_server`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSparkHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetSparkHistoryServer")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkHistoryServerResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// StopSparkHistoryServer : Stop Spark history server
// Stop the Spark history server of the given Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) StopSparkHistoryServer(stopSparkHistoryServerOptions *StopSparkHistoryServerOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.StopSparkHistoryServerWithContext(context.Background(), stopSparkHistoryServerOptions)
}

// StopSparkHistoryServerWithContext is an alternate form of the StopSparkHistoryServer method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) StopSparkHistoryServerWithContext(ctx context.Context, stopSparkHistoryServerOptions *StopSparkHistoryServerOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(stopSparkHistoryServerOptions, "stopSparkHistoryServerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(stopSparkHistoryServerOptions, "stopSparkHistoryServerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *stopSparkHistoryServerOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark_history_server`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range stopSparkHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "StopSparkHistoryServer")
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

// Application : Details of a Spark application.
type Application struct {
	// Identifier provided by Analytics Engine service for the Spark application.
	ID *string `json:"id,omitempty"`

	// Full URL of the resource.
	Href *string `json:"href,omitempty"`

	// Runtime enviroment for applications and other workloads.
	Runtime *Runtime `json:"runtime,omitempty"`

	// Identifier provided by Apache Spark for the application.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Name of the Spark application.
	SparkApplicationName *string `json:"spark_application_name,omitempty"`

	// State of the Spark application.
	State *string `json:"state,omitempty"`

	// URL of the Apache Spark web UI that is available when the application is running.
	SparkUi *string `json:"spark_ui,omitempty"`

	// Time when the application was submitted.
	SubmissionTime *strfmt.DateTime `json:"submission_time,omitempty"`

	// Time when the application was started.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Time when the application run ended in success, failure or was stopped.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// (deprecated) Time when the application was completed.
	FinishTime *strfmt.DateTime `json:"finish_time,omitempty"`

	// Time when the application will be automatically stopped by the service.
	AutoTerminationTime *strfmt.DateTime `json:"auto_termination_time,omitempty"`
}

// Constants associated with the Application.State property.
// State of the Spark application.
const (
	Application_State_Accepted = "accepted"
	Application_State_AutoTerminated = "auto_terminated"
	Application_State_Failed = "failed"
	Application_State_Finished = "finished"
	Application_State_OpsTerminated = "ops_terminated"
	Application_State_Running = "running"
	Application_State_Stopped = "stopped"
)

// UnmarshalApplication unmarshals an instance of Application from the specified map of raw messages.
func UnmarshalApplication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Application)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "runtime", &obj.Runtime, UnmarshalRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_id", &obj.SparkApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_name", &obj.SparkApplicationName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_ui", &obj.SparkUi)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "submission_time", &obj.SubmissionTime)
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
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationCollection : An array of application details.
type ApplicationCollection struct {
	// List of applications.
	Applications []Application `json:"applications,omitempty"`
}

// UnmarshalApplicationCollection unmarshals an instance of ApplicationCollection from the specified map of raw messages.
func UnmarshalApplicationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationCollection)
	err = core.UnmarshalModel(m, "applications", &obj.Applications, UnmarshalApplication)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationDetails : Application details.
type ApplicationDetails struct {
	// Path of the application to run.
	Application *string `json:"application,omitempty"`

	// Runtime enviroment for applications and other workloads.
	Runtime *Runtime `json:"runtime,omitempty"`

	// Path of the jar files containing the application.
	Jars *string `json:"jars,omitempty"`

	// Package names.
	Packages *string `json:"packages,omitempty"`

	// Repositories names.
	Repositories *string `json:"repositories,omitempty"`

	// File names.
	Files *string `json:"files,omitempty"`

	// Archive Names.
	Archives *string `json:"archives,omitempty"`

	// Name of the application.
	Name *string `json:"name,omitempty"`

	// Entry point for a Spark application bundled as a '.jar' file. This is applicable only for Java or Scala
	// applications.
	Class *string `json:"class,omitempty"`

	// An array of arguments to be passed to the application.
	Arguments []string `json:"arguments,omitempty"`

	// Application configurations to override the value specified at instance level. See [Spark environment variables](
	// https://spark.apache.org/docs/latest/configuration.html#available-properties) for a list of the supported variables.
	Conf map[string]interface{} `json:"conf,omitempty"`

	// Application environment configurations to use. See [Spark environment
	// variables](https://spark.apache.org/docs/latest/configuration.html#environment-variables) for a list of the
	// supported variables.
	Env map[string]interface{} `json:"env,omitempty"`
}

// UnmarshalApplicationDetails unmarshals an instance of ApplicationDetails from the specified map of raw messages.
func UnmarshalApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "runtime", &obj.Runtime, UnmarshalRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "jars", &obj.Jars)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "packages", &obj.Packages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repositories", &obj.Repositories)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "files", &obj.Files)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archives", &obj.Archives)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conf", &obj.Conf)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "env", &obj.Env)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationGetResponse : Response of the Application Get API.
type ApplicationGetResponse struct {
	// Application details.
	ApplicationDetails *ApplicationDetails `json:"application_details,omitempty"`

	// Application ID.
	ID *string `json:"id,omitempty"`

	// Identifier provided by Apache Spark for the application.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Name of the Spark application.
	SparkApplicationName *string `json:"spark_application_name,omitempty"`

	// State of the Spark application.
	State *string `json:"state,omitempty"`

	// URL of the Apache Spark web UI that is available when the application is running.
	SparkUi *string `json:"spark_ui,omitempty"`

	// List of additional information messages on the current state of the application.
	StateDetails []ApplicationGetResponseStateDetailsItem `json:"state_details,omitempty"`

	// Time when the application was submitted.
	SubmissionTime *strfmt.DateTime `json:"submission_time,omitempty"`

	// Time when the application started, in the format YYYY-MM-DDTHH:mm:ssZ.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Time when the application ended either in success or failure, in the format YYYY-MM-DDTHH:mm:ssZ.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// (deprecated) Time when the application completed successfully, in the format YYYY-MM-DDTHH:mm:ssZ.
	FinishTime *strfmt.DateTime `json:"finish_time,omitempty"`

	// Time when the application will be automatically stopped by the service.
	AutoTerminationTime *strfmt.DateTime `json:"auto_termination_time,omitempty"`
}

// Constants associated with the ApplicationGetResponse.State property.
// State of the Spark application.
const (
	ApplicationGetResponse_State_Accepted = "accepted"
	ApplicationGetResponse_State_AutoTerminated = "auto_terminated"
	ApplicationGetResponse_State_Failed = "failed"
	ApplicationGetResponse_State_Finished = "finished"
	ApplicationGetResponse_State_OpsTerminated = "ops_terminated"
	ApplicationGetResponse_State_Running = "running"
	ApplicationGetResponse_State_Stopped = "stopped"
)

// UnmarshalApplicationGetResponse unmarshals an instance of ApplicationGetResponse from the specified map of raw messages.
func UnmarshalApplicationGetResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationGetResponse)
	err = core.UnmarshalModel(m, "application_details", &obj.ApplicationDetails, UnmarshalApplicationDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_id", &obj.SparkApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_name", &obj.SparkApplicationName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_ui", &obj.SparkUi)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "state_details", &obj.StateDetails, UnmarshalApplicationGetResponseStateDetailsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "submission_time", &obj.SubmissionTime)
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
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationGetResponseStateDetailsItem : Additional information message on the current state of the application.
type ApplicationGetResponseStateDetailsItem struct {
	// Type of the message.
	Type *string `json:"type,omitempty"`

	// Fixed code for the message.
	Code *string `json:"code,omitempty"`

	// A descriptive message providing additional information on the current application state.
	Message *string `json:"message,omitempty"`
}

// Constants associated with the ApplicationGetResponseStateDetailsItem.Type property.
// Type of the message.
const (
	ApplicationGetResponseStateDetailsItem_Type_Info = "info"
	ApplicationGetResponseStateDetailsItem_Type_ServerError = "server_error"
	ApplicationGetResponseStateDetailsItem_Type_UserError = "user_error"
)

// UnmarshalApplicationGetResponseStateDetailsItem unmarshals an instance of ApplicationGetResponseStateDetailsItem from the specified map of raw messages.
func UnmarshalApplicationGetResponseStateDetailsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationGetResponseStateDetailsItem)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationGetStateResponse : State of a given application.
type ApplicationGetStateResponse struct {
	// Identifier of the application.
	ID *string `json:"id,omitempty"`

	// State of the Spark application.
	State *string `json:"state,omitempty"`

	// Time when the application was started.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Time when the application run ended in success, failure or was stopped.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// (deprecated) Time when the application was completed.
	FinishTime *strfmt.DateTime `json:"finish_time,omitempty"`

	// Time when the application will be automatically stopped by the service.
	AutoTerminationTime *strfmt.DateTime `json:"auto_termination_time,omitempty"`
}

// Constants associated with the ApplicationGetStateResponse.State property.
// State of the Spark application.
const (
	ApplicationGetStateResponse_State_Accepted = "accepted"
	ApplicationGetStateResponse_State_AutoTerminated = "auto_terminated"
	ApplicationGetStateResponse_State_Failed = "failed"
	ApplicationGetStateResponse_State_Finished = "finished"
	ApplicationGetStateResponse_State_OpsTerminated = "ops_terminated"
	ApplicationGetStateResponse_State_Running = "running"
	ApplicationGetStateResponse_State_Stopped = "stopped"
)

// UnmarshalApplicationGetStateResponse unmarshals an instance of ApplicationGetStateResponse from the specified map of raw messages.
func UnmarshalApplicationGetStateResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationGetStateResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
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
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationResponse : Application response details.
type ApplicationResponse struct {
	// Identifier of the application that was submitted.
	ID *string `json:"id,omitempty"`

	// State of the Spark application.
	State *string `json:"state,omitempty"`
}

// Constants associated with the ApplicationResponse.State property.
// State of the Spark application.
const (
	ApplicationResponse_State_Accepted = "accepted"
	ApplicationResponse_State_AutoTerminated = "auto_terminated"
	ApplicationResponse_State_Failed = "failed"
	ApplicationResponse_State_Finished = "finished"
	ApplicationResponse_State_OpsTerminated = "ops_terminated"
	ApplicationResponse_State_Running = "running"
	ApplicationResponse_State_Stopped = "stopped"
)

// UnmarshalApplicationResponse unmarshals an instance of ApplicationResponse from the specified map of raw messages.
func UnmarshalApplicationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurePlatformLoggingOptions : The ConfigurePlatformLogging options.
type ConfigurePlatformLoggingOptions struct {
	// GUID of the instance details for which log forwarding is to be configured.
	InstanceGuid *string `json:"instance_guid" validate:"required,ne="`

	// Enable or disable log forwarding.
	Enable *bool `json:"enable,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewConfigurePlatformLoggingOptions : Instantiate ConfigurePlatformLoggingOptions
func (*IbmAnalyticsEngineApiV3) NewConfigurePlatformLoggingOptions(instanceGuid string) *ConfigurePlatformLoggingOptions {
	return &ConfigurePlatformLoggingOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *ConfigurePlatformLoggingOptions) SetInstanceGuid(instanceGuid string) *ConfigurePlatformLoggingOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetEnable : Allow user to set Enable
func (_options *ConfigurePlatformLoggingOptions) SetEnable(enable bool) *ConfigurePlatformLoggingOptions {
	_options.Enable = core.BoolPtr(enable)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ConfigurePlatformLoggingOptions) SetHeaders(param map[string]string) *ConfigurePlatformLoggingOptions {
	options.Headers = param
	return options
}

// CreateApplicationOptions : The CreateApplication options.
type CreateApplicationOptions struct {
	// The identifier of the Analytics Engine instance associated with the Spark application(s).
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Path of the application to run.
	Application *string `json:"application,omitempty"`

	// Runtime enviroment for applications and other workloads.
	Runtime *Runtime `json:"runtime,omitempty"`

	// Path of the jar files containing the application.
	Jars *string `json:"jars,omitempty"`

	// Package names.
	Packages *string `json:"packages,omitempty"`

	// Repositories names.
	Repositories *string `json:"repositories,omitempty"`

	// File names.
	Files *string `json:"files,omitempty"`

	// Archive Names.
	Archives *string `json:"archives,omitempty"`

	// Name of the application.
	Name *string `json:"name,omitempty"`

	// Entry point for a Spark application bundled as a '.jar' file. This is applicable only for Java or Scala
	// applications.
	Class *string `json:"class,omitempty"`

	// An array of arguments to be passed to the application.
	Arguments []string `json:"arguments,omitempty"`

	// Application configurations to override the value specified at instance level. See [Spark environment variables](
	// https://spark.apache.org/docs/latest/configuration.html#available-properties) for a list of the supported variables.
	Conf map[string]interface{} `json:"conf,omitempty"`

	// Application environment configurations to use. See [Spark environment
	// variables](https://spark.apache.org/docs/latest/configuration.html#environment-variables) for a list of the
	// supported variables.
	Env map[string]interface{} `json:"env,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateApplicationOptions : Instantiate CreateApplicationOptions
func (*IbmAnalyticsEngineApiV3) NewCreateApplicationOptions(instanceID string) *CreateApplicationOptions {
	return &CreateApplicationOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateApplicationOptions) SetInstanceID(instanceID string) *CreateApplicationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetApplication : Allow user to set Application
func (_options *CreateApplicationOptions) SetApplication(application string) *CreateApplicationOptions {
	_options.Application = core.StringPtr(application)
	return _options
}

// SetRuntime : Allow user to set Runtime
func (_options *CreateApplicationOptions) SetRuntime(runtime *Runtime) *CreateApplicationOptions {
	_options.Runtime = runtime
	return _options
}

// SetJars : Allow user to set Jars
func (_options *CreateApplicationOptions) SetJars(jars string) *CreateApplicationOptions {
	_options.Jars = core.StringPtr(jars)
	return _options
}

// SetPackages : Allow user to set Packages
func (_options *CreateApplicationOptions) SetPackages(packages string) *CreateApplicationOptions {
	_options.Packages = core.StringPtr(packages)
	return _options
}

// SetRepositories : Allow user to set Repositories
func (_options *CreateApplicationOptions) SetRepositories(repositories string) *CreateApplicationOptions {
	_options.Repositories = core.StringPtr(repositories)
	return _options
}

// SetFiles : Allow user to set Files
func (_options *CreateApplicationOptions) SetFiles(files string) *CreateApplicationOptions {
	_options.Files = core.StringPtr(files)
	return _options
}

// SetArchives : Allow user to set Archives
func (_options *CreateApplicationOptions) SetArchives(archives string) *CreateApplicationOptions {
	_options.Archives = core.StringPtr(archives)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateApplicationOptions) SetName(name string) *CreateApplicationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetClass : Allow user to set Class
func (_options *CreateApplicationOptions) SetClass(class string) *CreateApplicationOptions {
	_options.Class = core.StringPtr(class)
	return _options
}

// SetArguments : Allow user to set Arguments
func (_options *CreateApplicationOptions) SetArguments(arguments []string) *CreateApplicationOptions {
	_options.Arguments = arguments
	return _options
}

// SetConf : Allow user to set Conf
func (_options *CreateApplicationOptions) SetConf(conf map[string]interface{}) *CreateApplicationOptions {
	_options.Conf = conf
	return _options
}

// SetEnv : Allow user to set Env
func (_options *CreateApplicationOptions) SetEnv(env map[string]interface{}) *CreateApplicationOptions {
	_options.Env = env
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateApplicationOptions) SetHeaders(param map[string]string) *CreateApplicationOptions {
	options.Headers = param
	return options
}

// CurrentResourceConsumptionResponse : Current resource consumption of the instance.
type CurrentResourceConsumptionResponse struct {
	// Number of virtual processor cores used.
	Cores *string `json:"cores,omitempty"`

	// Amount of memory used.
	Memory *string `json:"memory,omitempty"`
}

// UnmarshalCurrentResourceConsumptionResponse unmarshals an instance of CurrentResourceConsumptionResponse from the specified map of raw messages.
func UnmarshalCurrentResourceConsumptionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CurrentResourceConsumptionResponse)
	err = core.UnmarshalPrimitive(m, "cores", &obj.Cores)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteApplicationOptions : The DeleteApplication options.
type DeleteApplicationOptions struct {
	// Identifier of the instance to which the application belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Identifier of the application that needs to be stopped.
	ApplicationID *string `json:"application_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteApplicationOptions : Instantiate DeleteApplicationOptions
func (*IbmAnalyticsEngineApiV3) NewDeleteApplicationOptions(instanceID string, applicationID string) *DeleteApplicationOptions {
	return &DeleteApplicationOptions{
		InstanceID: core.StringPtr(instanceID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteApplicationOptions) SetInstanceID(instanceID string) *DeleteApplicationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *DeleteApplicationOptions) SetApplicationID(applicationID string) *DeleteApplicationOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteApplicationOptions) SetHeaders(param map[string]string) *DeleteApplicationOptions {
	options.Headers = param
	return options
}

// GetApplicationOptions : The GetApplication options.
type GetApplicationOptions struct {
	// Identifier of the instance to which the application belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Identifier of the application for which details are requested.
	ApplicationID *string `json:"application_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationOptions : Instantiate GetApplicationOptions
func (*IbmAnalyticsEngineApiV3) NewGetApplicationOptions(instanceID string, applicationID string) *GetApplicationOptions {
	return &GetApplicationOptions{
		InstanceID: core.StringPtr(instanceID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetApplicationOptions) SetInstanceID(instanceID string) *GetApplicationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *GetApplicationOptions) SetApplicationID(applicationID string) *GetApplicationOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationOptions) SetHeaders(param map[string]string) *GetApplicationOptions {
	options.Headers = param
	return options
}

// GetApplicationStateOptions : The GetApplicationState options.
type GetApplicationStateOptions struct {
	// Identifier of the instance to which the applications belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Identifier of the application for which details are requested.
	ApplicationID *string `json:"application_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationStateOptions : Instantiate GetApplicationStateOptions
func (*IbmAnalyticsEngineApiV3) NewGetApplicationStateOptions(instanceID string, applicationID string) *GetApplicationStateOptions {
	return &GetApplicationStateOptions{
		InstanceID: core.StringPtr(instanceID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetApplicationStateOptions) SetInstanceID(instanceID string) *GetApplicationStateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *GetApplicationStateOptions) SetApplicationID(applicationID string) *GetApplicationStateOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationStateOptions) SetHeaders(param map[string]string) *GetApplicationStateOptions {
	options.Headers = param
	return options
}

// GetCurrentResourceConsumptionOptions : The GetCurrentResourceConsumption options.
type GetCurrentResourceConsumptionOptions struct {
	// ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCurrentResourceConsumptionOptions : Instantiate GetCurrentResourceConsumptionOptions
func (*IbmAnalyticsEngineApiV3) NewGetCurrentResourceConsumptionOptions(instanceID string) *GetCurrentResourceConsumptionOptions {
	return &GetCurrentResourceConsumptionOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetCurrentResourceConsumptionOptions) SetInstanceID(instanceID string) *GetCurrentResourceConsumptionOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCurrentResourceConsumptionOptions) SetHeaders(param map[string]string) *GetCurrentResourceConsumptionOptions {
	options.Headers = param
	return options
}

// GetInstanceDefaultConfigsOptions : The GetInstanceDefaultConfigs options.
type GetInstanceDefaultConfigsOptions struct {
	// The ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceDefaultConfigsOptions : Instantiate GetInstanceDefaultConfigsOptions
func (*IbmAnalyticsEngineApiV3) NewGetInstanceDefaultConfigsOptions(instanceID string) *GetInstanceDefaultConfigsOptions {
	return &GetInstanceDefaultConfigsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceDefaultConfigsOptions) SetInstanceID(instanceID string) *GetInstanceDefaultConfigsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceDefaultConfigsOptions) SetHeaders(param map[string]string) *GetInstanceDefaultConfigsOptions {
	options.Headers = param
	return options
}

// GetInstanceDefaultRuntimeOptions : The GetInstanceDefaultRuntime options.
type GetInstanceDefaultRuntimeOptions struct {
	// The ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceDefaultRuntimeOptions : Instantiate GetInstanceDefaultRuntimeOptions
func (*IbmAnalyticsEngineApiV3) NewGetInstanceDefaultRuntimeOptions(instanceID string) *GetInstanceDefaultRuntimeOptions {
	return &GetInstanceDefaultRuntimeOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceDefaultRuntimeOptions) SetInstanceID(instanceID string) *GetInstanceDefaultRuntimeOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceDefaultRuntimeOptions) SetHeaders(param map[string]string) *GetInstanceDefaultRuntimeOptions {
	options.Headers = param
	return options
}

// GetInstanceOptions : The GetInstance options.
type GetInstanceOptions struct {
	// GUID of the Analytics Engine service instance to retrieve.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceOptions : Instantiate GetInstanceOptions
func (*IbmAnalyticsEngineApiV3) NewGetInstanceOptions(instanceID string) *GetInstanceOptions {
	return &GetInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceOptions) SetInstanceID(instanceID string) *GetInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceOptions) SetHeaders(param map[string]string) *GetInstanceOptions {
	options.Headers = param
	return options
}

// GetInstanceStateOptions : The GetInstanceState options.
type GetInstanceStateOptions struct {
	// GUID of the Analytics Engine service instance to retrieve state.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceStateOptions : Instantiate GetInstanceStateOptions
func (*IbmAnalyticsEngineApiV3) NewGetInstanceStateOptions(instanceID string) *GetInstanceStateOptions {
	return &GetInstanceStateOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceStateOptions) SetInstanceID(instanceID string) *GetInstanceStateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceStateOptions) SetHeaders(param map[string]string) *GetInstanceStateOptions {
	options.Headers = param
	return options
}

// GetLogForwardingConfigOptions : The GetLogForwardingConfig options.
type GetLogForwardingConfigOptions struct {
	// ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLogForwardingConfigOptions : Instantiate GetLogForwardingConfigOptions
func (*IbmAnalyticsEngineApiV3) NewGetLogForwardingConfigOptions(instanceID string) *GetLogForwardingConfigOptions {
	return &GetLogForwardingConfigOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetLogForwardingConfigOptions) SetInstanceID(instanceID string) *GetLogForwardingConfigOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLogForwardingConfigOptions) SetHeaders(param map[string]string) *GetLogForwardingConfigOptions {
	options.Headers = param
	return options
}

// GetLoggingConfigurationOptions : The GetLoggingConfiguration options.
type GetLoggingConfigurationOptions struct {
	// GUID of the Analytics Engine service instance to retrieve log configuration.
	InstanceGuid *string `json:"instance_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoggingConfigurationOptions : Instantiate GetLoggingConfigurationOptions
func (*IbmAnalyticsEngineApiV3) NewGetLoggingConfigurationOptions(instanceGuid string) *GetLoggingConfigurationOptions {
	return &GetLoggingConfigurationOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *GetLoggingConfigurationOptions) SetInstanceGuid(instanceGuid string) *GetLoggingConfigurationOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoggingConfigurationOptions) SetHeaders(param map[string]string) *GetLoggingConfigurationOptions {
	options.Headers = param
	return options
}

// GetResourceConsumptionLimitsOptions : The GetResourceConsumptionLimits options.
type GetResourceConsumptionLimitsOptions struct {
	// ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceConsumptionLimitsOptions : Instantiate GetResourceConsumptionLimitsOptions
func (*IbmAnalyticsEngineApiV3) NewGetResourceConsumptionLimitsOptions(instanceID string) *GetResourceConsumptionLimitsOptions {
	return &GetResourceConsumptionLimitsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetResourceConsumptionLimitsOptions) SetInstanceID(instanceID string) *GetResourceConsumptionLimitsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceConsumptionLimitsOptions) SetHeaders(param map[string]string) *GetResourceConsumptionLimitsOptions {
	options.Headers = param
	return options
}

// GetSparkHistoryServerOptions : The GetSparkHistoryServer options.
type GetSparkHistoryServerOptions struct {
	// The ID of the Analytics Engine instance to which the Spark history server belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSparkHistoryServerOptions : Instantiate GetSparkHistoryServerOptions
func (*IbmAnalyticsEngineApiV3) NewGetSparkHistoryServerOptions(instanceID string) *GetSparkHistoryServerOptions {
	return &GetSparkHistoryServerOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetSparkHistoryServerOptions) SetInstanceID(instanceID string) *GetSparkHistoryServerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkHistoryServerOptions) SetHeaders(param map[string]string) *GetSparkHistoryServerOptions {
	options.Headers = param
	return options
}

// Instance : Details of Analytics Engine instance.
type Instance struct {
	// GUID of the Analytics Engine instance.
	ID *string `json:"id,omitempty"`

	// Full URL of the resource.
	Href *string `json:"href,omitempty"`

	// State of the Analytics Engine instance.
	State *string `json:"state,omitempty"`

	// Timestamp when the state of the instance was changed, in the format YYYY-MM-DDTHH:mm:ssZ.
	StateChangeTime *strfmt.DateTime `json:"state_change_time,omitempty"`

	// Runtime enviroment for applications and other workloads.
	DefaultRuntime *Runtime `json:"default_runtime,omitempty"`

	// Object storage instance that acts as the home for custom libraries and Spark events.
	InstanceHome *InstanceHome `json:"instance_home,omitempty"`

	// Instance level default configuration for Spark workloads.
	DefaultConfig *InstanceDefaultConfig `json:"default_config,omitempty"`
}

// Constants associated with the Instance.State property.
// State of the Analytics Engine instance.
const (
	Instance_State_Active = "active"
	Instance_State_CreationAccepted = "creation_accepted"
	Instance_State_CreationFailed = "creation_failed"
	Instance_State_Deleted = "deleted"
	Instance_State_Disabled = "disabled"
	Instance_State_Initialized = "initialized"
	Instance_State_Preparing = "preparing"
)

// UnmarshalInstance unmarshals an instance of Instance from the specified map of raw messages.
func UnmarshalInstance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Instance)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_change_time", &obj.StateChangeTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_runtime", &obj.DefaultRuntime, UnmarshalRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "instance_home", &obj.InstanceHome, UnmarshalInstanceHome)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_config", &obj.DefaultConfig, UnmarshalInstanceDefaultConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceDefaultConfig : Instance level default configuration for Spark workloads.
type InstanceDefaultConfig struct {
	// Value of the Spark configuration key.
	Key *string `json:"key,omitempty"`
}

// UnmarshalInstanceDefaultConfig unmarshals an instance of InstanceDefaultConfig from the specified map of raw messages.
func UnmarshalInstanceDefaultConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDefaultConfig)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceGetStateResponse : State details of Analytics Engine instance.
type InstanceGetStateResponse struct {
	// GUID of the Analytics Engine instance.
	ID *string `json:"id,omitempty"`

	// State of the Analytics Engine instance.
	State *string `json:"state,omitempty"`
}

// Constants associated with the InstanceGetStateResponse.State property.
// State of the Analytics Engine instance.
const (
	InstanceGetStateResponse_State_Active = "active"
	InstanceGetStateResponse_State_CreationAccepted = "creation_accepted"
	InstanceGetStateResponse_State_CreationFailed = "creation_failed"
	InstanceGetStateResponse_State_Deleted = "deleted"
	InstanceGetStateResponse_State_Disabled = "disabled"
	InstanceGetStateResponse_State_Initialized = "initialized"
	InstanceGetStateResponse_State_Preparing = "preparing"
)

// UnmarshalInstanceGetStateResponse unmarshals an instance of InstanceGetStateResponse from the specified map of raw messages.
func UnmarshalInstanceGetStateResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceGetStateResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceHome : Object storage instance that acts as the home for custom libraries and Spark events.
type InstanceHome struct {
	// UUID of the instance home storage instance.
	ID *string `json:"id,omitempty"`

	// Currently only ibm-cos (IBM Cloud Object Storage) is supported.
	Provider *string `json:"provider,omitempty"`

	// Type of the instance home storage. Currently, only objectstore (Cloud Object Storage) is supported.
	Type *string `json:"type,omitempty"`

	// Region of the Cloud Object Storage instance.
	Region *string `json:"region,omitempty"`

	// Endpoint to access the Cloud Object Storage instance.
	Endpoint *string `json:"endpoint,omitempty"`

	// Cloud Object Storage bucket used as instance home.
	Bucket *string `json:"bucket,omitempty"`

	// Cloud Object Storage access key. Masked for security reasons.
	HmacAccessKey *string `json:"hmac_access_key,omitempty"`

	// Cloud Object Storage secret key. Masked for security reasons.
	HmacSecretKey *string `json:"hmac_secret_key,omitempty"`
}

// UnmarshalInstanceHome unmarshals an instance of InstanceHome from the specified map of raw messages.
func UnmarshalInstanceHome(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceHome)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hmac_access_key", &obj.HmacAccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hmac_secret_key", &obj.HmacSecretKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceHomeResponse : Response of Instance home API.
type InstanceHomeResponse struct {
	// UUID of the instance home storage instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// Currently only ibm-cos (IBM Cloud Object Storage) is supported.
	Provider *string `json:"provider,omitempty"`

	// Type of the instance home storage. Currently, only objectstore (Cloud Object Storage) is supported.
	Type *string `json:"type,omitempty"`

	// Region of the Cloud Object Storage instance.
	Region *string `json:"region,omitempty"`

	// Endpoint to access the Cloud Object Storage instance.
	Endpoint *string `json:"endpoint,omitempty"`

	// Cloud Object Storage access key.
	HmacAccessKey *string `json:"hmac_access_key,omitempty"`

	// Cloud Object Storage secret key.
	HmacSecretKey *string `json:"hmac_secret_key,omitempty"`
}

// UnmarshalInstanceHomeResponse unmarshals an instance of InstanceHomeResponse from the specified map of raw messages.
func UnmarshalInstanceHomeResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceHomeResponse)
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hmac_access_key", &obj.HmacAccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hmac_secret_key", &obj.HmacSecretKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListApplicationsOptions : The ListApplications options.
type ListApplicationsOptions struct {
	// The identifier of the Analytics Engine instance associated with the Spark application(s).
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// List of Spark application states that will be used to filter the response.
	State []string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListApplicationsOptions.State property.
// State of the Spark application.
const (
	ListApplicationsOptions_State_Accepted = "accepted"
	ListApplicationsOptions_State_AutoTerminated = "auto_terminated"
	ListApplicationsOptions_State_Failed = "failed"
	ListApplicationsOptions_State_Finished = "finished"
	ListApplicationsOptions_State_OpsTerminated = "ops_terminated"
	ListApplicationsOptions_State_Running = "running"
	ListApplicationsOptions_State_Stopped = "stopped"
)

// NewListApplicationsOptions : Instantiate ListApplicationsOptions
func (*IbmAnalyticsEngineApiV3) NewListApplicationsOptions(instanceID string) *ListApplicationsOptions {
	return &ListApplicationsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListApplicationsOptions) SetInstanceID(instanceID string) *ListApplicationsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetState : Allow user to set State
func (_options *ListApplicationsOptions) SetState(state []string) *ListApplicationsOptions {
	_options.State = state
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListApplicationsOptions) SetHeaders(param map[string]string) *ListApplicationsOptions {
	options.Headers = param
	return options
}

// LogForwardingConfigResponse : Log forwarding configuration details.
type LogForwardingConfigResponse struct {
	// List of sources of logs that are being forwarded.
	Sources []string `json:"sources,omitempty"`

	// List of tags that are applied to the logs being forwarded.
	Tags []string `json:"tags,omitempty"`

	// Log server properties.
	LogServer *LogForwardingConfigResponseLogServer `json:"log_server,omitempty"`

	// Indicates whether log forwarding is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// UnmarshalLogForwardingConfigResponse unmarshals an instance of LogForwardingConfigResponse from the specified map of raw messages.
func UnmarshalLogForwardingConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogForwardingConfigResponse)
	err = core.UnmarshalPrimitive(m, "sources", &obj.Sources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_server", &obj.LogServer, UnmarshalLogForwardingConfigResponseLogServer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogForwardingConfigResponseLogServer : Log server properties.
type LogForwardingConfigResponseLogServer struct {
	// Type of the log server.
	Type *string `json:"type,omitempty"`
}

// UnmarshalLogForwardingConfigResponseLogServer unmarshals an instance of LogForwardingConfigResponseLogServer from the specified map of raw messages.
func UnmarshalLogForwardingConfigResponseLogServer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogForwardingConfigResponseLogServer)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LoggingConfigurationResponse : (deprecated) Response of logging API.
type LoggingConfigurationResponse struct {
	// component array.
	Components []string `json:"components,omitempty"`

	// log server properties.
	LogServer *LoggingConfigurationResponseLogServer `json:"log_server,omitempty"`

	// enable.
	Enable *bool `json:"enable,omitempty"`
}

// UnmarshalLoggingConfigurationResponse unmarshals an instance of LoggingConfigurationResponse from the specified map of raw messages.
func UnmarshalLoggingConfigurationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoggingConfigurationResponse)
	err = core.UnmarshalPrimitive(m, "components", &obj.Components)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_server", &obj.LogServer, UnmarshalLoggingConfigurationResponseLogServer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable", &obj.Enable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LoggingConfigurationResponseLogServer : log server properties.
type LoggingConfigurationResponseLogServer struct {
	// type of log server.
	Type *string `json:"type,omitempty"`
}

// UnmarshalLoggingConfigurationResponseLogServer unmarshals an instance of LoggingConfigurationResponseLogServer from the specified map of raw messages.
func UnmarshalLoggingConfigurationResponseLogServer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoggingConfigurationResponseLogServer)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceInstanceDefaultConfigsOptions : The ReplaceInstanceDefaultConfigs options.
type ReplaceInstanceDefaultConfigsOptions struct {
	// The ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Spark configuration properties to replace existing instance default Spark configurations.
	Body map[string]string `json:"body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceInstanceDefaultConfigsOptions : Instantiate ReplaceInstanceDefaultConfigsOptions
func (*IbmAnalyticsEngineApiV3) NewReplaceInstanceDefaultConfigsOptions(instanceID string, body map[string]string) *ReplaceInstanceDefaultConfigsOptions {
	return &ReplaceInstanceDefaultConfigsOptions{
		InstanceID: core.StringPtr(instanceID),
		Body: body,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceInstanceDefaultConfigsOptions) SetInstanceID(instanceID string) *ReplaceInstanceDefaultConfigsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *ReplaceInstanceDefaultConfigsOptions) SetBody(body map[string]string) *ReplaceInstanceDefaultConfigsOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceInstanceDefaultConfigsOptions) SetHeaders(param map[string]string) *ReplaceInstanceDefaultConfigsOptions {
	options.Headers = param
	return options
}

// ReplaceInstanceDefaultRuntimeOptions : The ReplaceInstanceDefaultRuntime options.
type ReplaceInstanceDefaultRuntimeOptions struct {
	// The ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Spark version of the runtime environment.
	SparkVersion *string `json:"spark_version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceInstanceDefaultRuntimeOptions : Instantiate ReplaceInstanceDefaultRuntimeOptions
func (*IbmAnalyticsEngineApiV3) NewReplaceInstanceDefaultRuntimeOptions(instanceID string) *ReplaceInstanceDefaultRuntimeOptions {
	return &ReplaceInstanceDefaultRuntimeOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceInstanceDefaultRuntimeOptions) SetInstanceID(instanceID string) *ReplaceInstanceDefaultRuntimeOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetSparkVersion : Allow user to set SparkVersion
func (_options *ReplaceInstanceDefaultRuntimeOptions) SetSparkVersion(sparkVersion string) *ReplaceInstanceDefaultRuntimeOptions {
	_options.SparkVersion = core.StringPtr(sparkVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceInstanceDefaultRuntimeOptions) SetHeaders(param map[string]string) *ReplaceInstanceDefaultRuntimeOptions {
	options.Headers = param
	return options
}

// ReplaceLogForwardingConfigOptions : The ReplaceLogForwardingConfig options.
type ReplaceLogForwardingConfigOptions struct {
	// ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Enable or disable log forwarding.
	Enabled *bool `json:"enabled,omitempty"`

	// List of sources of logs that will be forwarded. By default, only 'spark-driver' logs are forwarded.
	Sources []string `json:"sources,omitempty"`

	// List of tags to be applied to the logs being forwarded. They can be used to filter the logs in the IBM Log Analysis
	// server.
	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceLogForwardingConfigOptions : Instantiate ReplaceLogForwardingConfigOptions
func (*IbmAnalyticsEngineApiV3) NewReplaceLogForwardingConfigOptions(instanceID string) *ReplaceLogForwardingConfigOptions {
	return &ReplaceLogForwardingConfigOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceLogForwardingConfigOptions) SetInstanceID(instanceID string) *ReplaceLogForwardingConfigOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *ReplaceLogForwardingConfigOptions) SetEnabled(enabled bool) *ReplaceLogForwardingConfigOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetSources : Allow user to set Sources
func (_options *ReplaceLogForwardingConfigOptions) SetSources(sources []string) *ReplaceLogForwardingConfigOptions {
	_options.Sources = sources
	return _options
}

// SetTags : Allow user to set Tags
func (_options *ReplaceLogForwardingConfigOptions) SetTags(tags []string) *ReplaceLogForwardingConfigOptions {
	_options.Tags = tags
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceLogForwardingConfigOptions) SetHeaders(param map[string]string) *ReplaceLogForwardingConfigOptions {
	options.Headers = param
	return options
}

// ResourceConsumptionLimitsResponse : Resource consumption limits for the instance.
type ResourceConsumptionLimitsResponse struct {
	// Maximum number of virtual processor cores that be used in the instance.
	MaxCores *string `json:"max_cores,omitempty"`

	// Maximum memory that can be used in the instance.
	MaxMemory *string `json:"max_memory,omitempty"`
}

// UnmarshalResourceConsumptionLimitsResponse unmarshals an instance of ResourceConsumptionLimitsResponse from the specified map of raw messages.
func UnmarshalResourceConsumptionLimitsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceConsumptionLimitsResponse)
	err = core.UnmarshalPrimitive(m, "max_cores", &obj.MaxCores)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_memory", &obj.MaxMemory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Runtime : Runtime enviroment for applications and other workloads.
type Runtime struct {
	// Spark version of the runtime environment.
	SparkVersion *string `json:"spark_version,omitempty"`
}

// UnmarshalRuntime unmarshals an instance of Runtime from the specified map of raw messages.
func UnmarshalRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Runtime)
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetInstanceHomeOptions : The SetInstanceHome options.
type SetInstanceHomeOptions struct {
	// The ID of the Analytics Engine instance for which 'instance home' is to be set.
	InstanceID *string `json:"-" validate:"required,ne="`

	// UUID of the instance home storage instance.
	NewInstanceID *string `json:"instance_id,omitempty"`

	// Currently only ibm-cos (IBM Cloud Object Storage) is supported.
	NewProvider *string `json:"provider,omitempty"`

	// Type of the instance home storage. Currently, only objectstore (Cloud Object Storage) is supported.
	NewType *string `json:"type,omitempty"`

	// Region of the Cloud Object Storage instance.
	NewRegion *string `json:"region,omitempty"`

	// Endpoint to access the Cloud Object Storage instance.
	NewEndpoint *string `json:"endpoint,omitempty"`

	// Cloud Object Storage access key.
	NewHmacAccessKey *string `json:"hmac_access_key,omitempty"`

	// Cloud Object Storage secret key.
	NewHmacSecretKey *string `json:"hmac_secret_key,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetInstanceHomeOptions : Instantiate SetInstanceHomeOptions
func (*IbmAnalyticsEngineApiV3) NewSetInstanceHomeOptions(instanceID string) *SetInstanceHomeOptions {
	return &SetInstanceHomeOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *SetInstanceHomeOptions) SetInstanceID(instanceID string) *SetInstanceHomeOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetNewInstanceID : Allow user to set NewInstanceID
func (_options *SetInstanceHomeOptions) SetNewInstanceID(newInstanceID string) *SetInstanceHomeOptions {
	_options.NewInstanceID = core.StringPtr(newInstanceID)
	return _options
}

// SetNewProvider : Allow user to set NewProvider
func (_options *SetInstanceHomeOptions) SetNewProvider(newProvider string) *SetInstanceHomeOptions {
	_options.NewProvider = core.StringPtr(newProvider)
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *SetInstanceHomeOptions) SetNewType(newType string) *SetInstanceHomeOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetNewRegion : Allow user to set NewRegion
func (_options *SetInstanceHomeOptions) SetNewRegion(newRegion string) *SetInstanceHomeOptions {
	_options.NewRegion = core.StringPtr(newRegion)
	return _options
}

// SetNewEndpoint : Allow user to set NewEndpoint
func (_options *SetInstanceHomeOptions) SetNewEndpoint(newEndpoint string) *SetInstanceHomeOptions {
	_options.NewEndpoint = core.StringPtr(newEndpoint)
	return _options
}

// SetNewHmacAccessKey : Allow user to set NewHmacAccessKey
func (_options *SetInstanceHomeOptions) SetNewHmacAccessKey(newHmacAccessKey string) *SetInstanceHomeOptions {
	_options.NewHmacAccessKey = core.StringPtr(newHmacAccessKey)
	return _options
}

// SetNewHmacSecretKey : Allow user to set NewHmacSecretKey
func (_options *SetInstanceHomeOptions) SetNewHmacSecretKey(newHmacSecretKey string) *SetInstanceHomeOptions {
	_options.NewHmacSecretKey = core.StringPtr(newHmacSecretKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetInstanceHomeOptions) SetHeaders(param map[string]string) *SetInstanceHomeOptions {
	options.Headers = param
	return options
}

// SparkHistoryServerResponse : Status of the Spark history server.
type SparkHistoryServerResponse struct {
	// State of the Spark history server.
	State *string `json:"state,omitempty"`

	// Number of cpu cores used by the Spark history server.
	Cores *string `json:"cores,omitempty"`

	// Amount of memory used by the Spark history server.
	Memory *string `json:"memory,omitempty"`

	// Time when the Spark history server was started.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Time when the Spark history server was stopped.
	StopTime *strfmt.DateTime `json:"stop_time,omitempty"`

	// Time when the Spark history server will be stopped automatically.
	AutoTerminationTime *strfmt.DateTime `json:"auto_termination_time,omitempty"`
}

// Constants associated with the SparkHistoryServerResponse.State property.
// State of the Spark history server.
const (
	SparkHistoryServerResponse_State_Started = "started"
	SparkHistoryServerResponse_State_Stopped = "stopped"
)

// UnmarshalSparkHistoryServerResponse unmarshals an instance of SparkHistoryServerResponse from the specified map of raw messages.
func UnmarshalSparkHistoryServerResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkHistoryServerResponse)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cores", &obj.Cores)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stop_time", &obj.StopTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StartSparkHistoryServerOptions : The StartSparkHistoryServer options.
type StartSparkHistoryServerOptions struct {
	// The ID of the Analytics Engine instance to which the Spark history server belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStartSparkHistoryServerOptions : Instantiate StartSparkHistoryServerOptions
func (*IbmAnalyticsEngineApiV3) NewStartSparkHistoryServerOptions(instanceID string) *StartSparkHistoryServerOptions {
	return &StartSparkHistoryServerOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *StartSparkHistoryServerOptions) SetInstanceID(instanceID string) *StartSparkHistoryServerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *StartSparkHistoryServerOptions) SetHeaders(param map[string]string) *StartSparkHistoryServerOptions {
	options.Headers = param
	return options
}

// StopSparkHistoryServerOptions : The StopSparkHistoryServer options.
type StopSparkHistoryServerOptions struct {
	// The ID of the Analytics Engine instance to which the Spark history server belongs.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStopSparkHistoryServerOptions : Instantiate StopSparkHistoryServerOptions
func (*IbmAnalyticsEngineApiV3) NewStopSparkHistoryServerOptions(instanceID string) *StopSparkHistoryServerOptions {
	return &StopSparkHistoryServerOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *StopSparkHistoryServerOptions) SetInstanceID(instanceID string) *StopSparkHistoryServerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *StopSparkHistoryServerOptions) SetHeaders(param map[string]string) *StopSparkHistoryServerOptions {
	options.Headers = param
	return options
}

// UpdateInstanceDefaultConfigsOptions : The UpdateInstanceDefaultConfigs options.
type UpdateInstanceDefaultConfigsOptions struct {
	// The ID of the Analytics Engine instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Spark configuration properties to be updated. Properties will be merged with existing configuration properties. Set
	// a property value to `null` in order to unset it.
	Body map[string]interface{} `json:"body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateInstanceDefaultConfigsOptions : Instantiate UpdateInstanceDefaultConfigsOptions
func (*IbmAnalyticsEngineApiV3) NewUpdateInstanceDefaultConfigsOptions(instanceID string, body map[string]interface{}) *UpdateInstanceDefaultConfigsOptions {
	return &UpdateInstanceDefaultConfigsOptions{
		InstanceID: core.StringPtr(instanceID),
		Body: body,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceDefaultConfigsOptions) SetInstanceID(instanceID string) *UpdateInstanceDefaultConfigsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateInstanceDefaultConfigsOptions) SetBody(body map[string]interface{}) *UpdateInstanceDefaultConfigsOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceDefaultConfigsOptions) SetHeaders(param map[string]string) *UpdateInstanceDefaultConfigsOptions {
	options.Headers = param
	return options
}
