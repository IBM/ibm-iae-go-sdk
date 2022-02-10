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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.43.5-e0ec19e2-20220124-172004
 */

// Package ibmanalyticsengineapiv3 : Operations and models for the IbmAnalyticsEngineApiV3 service
package ibmanalyticsengineapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/ibm-iae-go-sdk/common"
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

// CreateInstanceHome : Edit Instance Home details
// Instance details of the Object Storage instance that will be used as instance home.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) CreateInstanceHome(createInstanceHomeOptions *CreateInstanceHomeOptions) (result *InstanceHomeResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.CreateInstanceHomeWithContext(context.Background(), createInstanceHomeOptions)
}

// CreateInstanceHomeWithContext is an alternate form of the CreateInstanceHome method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) CreateInstanceHomeWithContext(ctx context.Context, createInstanceHomeOptions *CreateInstanceHomeOptions) (result *InstanceHomeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createInstanceHomeOptions, "createInstanceHomeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createInstanceHomeOptions, "createInstanceHomeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createInstanceHomeOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/instance_home`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createInstanceHomeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "CreateInstanceHome")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createInstanceHomeOptions.NewInstanceID != nil {
		body["instance_id"] = createInstanceHomeOptions.NewInstanceID
	}
	if createInstanceHomeOptions.NewProvider != nil {
		body["provider"] = createInstanceHomeOptions.NewProvider
	}
	if createInstanceHomeOptions.NewType != nil {
		body["type"] = createInstanceHomeOptions.NewType
	}
	if createInstanceHomeOptions.NewRegion != nil {
		body["region"] = createInstanceHomeOptions.NewRegion
	}
	if createInstanceHomeOptions.NewEndpoint != nil {
		body["endpoint"] = createInstanceHomeOptions.NewEndpoint
	}
	if createInstanceHomeOptions.NewHmacAccessKey != nil {
		body["hmac_access_key"] = createInstanceHomeOptions.NewHmacAccessKey
	}
	if createInstanceHomeOptions.NewHmacSecretKey != nil {
		body["hmac_secret_key"] = createInstanceHomeOptions.NewHmacSecretKey
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

// ListApplications : Retrieve all Spark applications
// Gets all applications submitted in an instance with a specified instance-id.
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
// Gets the details of the given Spark application.
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

// EnablePlatformLogging : Enable/disable platform logging
// Enable platform logging from IBM Analytics Engine.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) EnablePlatformLogging(enablePlatformLoggingOptions *EnablePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.EnablePlatformLoggingWithContext(context.Background(), enablePlatformLoggingOptions)
}

// EnablePlatformLoggingWithContext is an alternate form of the EnablePlatformLogging method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) EnablePlatformLoggingWithContext(ctx context.Context, enablePlatformLoggingOptions *EnablePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(enablePlatformLoggingOptions, "enablePlatformLoggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(enablePlatformLoggingOptions, "enablePlatformLoggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *enablePlatformLoggingOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_guid}/logging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range enablePlatformLoggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "EnablePlatformLogging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if enablePlatformLoggingOptions.Enable != nil {
		body["enable"] = enablePlatformLoggingOptions.Enable
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

// DisablePlatformLogging : Disable platform logging
// Disable platform logging from IBM Analytics Engine.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DisablePlatformLogging(disablePlatformLoggingOptions *DisablePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.DisablePlatformLoggingWithContext(context.Background(), disablePlatformLoggingOptions)
}

// DisablePlatformLoggingWithContext is an alternate form of the DisablePlatformLogging method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DisablePlatformLoggingWithContext(ctx context.Context, disablePlatformLoggingOptions *DisablePlatformLoggingOptions) (result *LoggingConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(disablePlatformLoggingOptions, "disablePlatformLoggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(disablePlatformLoggingOptions, "disablePlatformLoggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *disablePlatformLoggingOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_guid}/logging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range disablePlatformLoggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "DisablePlatformLogging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if disablePlatformLoggingOptions.Enable != nil {
		body["enable"] = disablePlatformLoggingOptions.Enable
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

// GetLoggingConfiguration : Find logging configuration by instance id
// Retrieve the logging configuration of a single Analytics Engine instance.
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

// DeleteLoggingConfiguration : Delete logging configuration by instance id
// Delete the logging configuration of a single Analytics Engine instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteLoggingConfiguration(deleteLoggingConfigurationOptions *DeleteLoggingConfigurationOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.DeleteLoggingConfigurationWithContext(context.Background(), deleteLoggingConfigurationOptions)
}

// DeleteLoggingConfigurationWithContext is an alternate form of the DeleteLoggingConfiguration method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteLoggingConfigurationWithContext(ctx context.Context, deleteLoggingConfigurationOptions *DeleteLoggingConfigurationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoggingConfigurationOptions, "deleteLoggingConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoggingConfigurationOptions, "deleteLoggingConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_guid": *deleteLoggingConfigurationOptions.InstanceGuid,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_guid}/logging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoggingConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "DeleteLoggingConfiguration")
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

	// Identifier provided by Apache Spark for the application.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Status of the application.
	State *string `json:"state,omitempty"`

	// Time when the application was started.
	StartTime *string `json:"start_time,omitempty"`

	// Time when the application was completed.
	FinishTime *string `json:"finish_time,omitempty"`
}

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
	err = core.UnmarshalPrimitive(m, "spark_application_id", &obj.SparkApplicationID)
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
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
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

	// Application state.
	State *string `json:"state,omitempty"`

	// Application start time in the format YYYY-MM-DDTHH:mm:ssZ.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Application end time in the format YYYY-MM-DDTHH:mm:ssZ.
	FinishTime *strfmt.DateTime `json:"finish_time,omitempty"`
}

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
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
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

	// Status of the application.
	State *string `json:"state,omitempty"`

	// Time when the application was started.
	StartTime *string `json:"start_time,omitempty"`

	// Time when the application was completed.
	FinishTime *string `json:"finish_time,omitempty"`
}

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
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
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

	// State of the submitted application.
	State *string `json:"state,omitempty"`
}

// Constants associated with the ApplicationResponse.State property.
// State of the submitted application.
const (
	ApplicationResponse_State_Accepted = "accepted"
	ApplicationResponse_State_Error = "error"
	ApplicationResponse_State_Failed = "failed"
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

// CreateApplicationOptions : The CreateApplication options.
type CreateApplicationOptions struct {
	// The identifier of the instance where the Spark application is submitted.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Path of the application to run.
	Application *string `json:"application,omitempty"`

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

// CreateInstanceHomeOptions : The CreateInstanceHome options.
type CreateInstanceHomeOptions struct {
	// The identifier of the instance details to be added.
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

// NewCreateInstanceHomeOptions : Instantiate CreateInstanceHomeOptions
func (*IbmAnalyticsEngineApiV3) NewCreateInstanceHomeOptions(instanceID string) *CreateInstanceHomeOptions {
	return &CreateInstanceHomeOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateInstanceHomeOptions) SetInstanceID(instanceID string) *CreateInstanceHomeOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetNewInstanceID : Allow user to set NewInstanceID
func (_options *CreateInstanceHomeOptions) SetNewInstanceID(newInstanceID string) *CreateInstanceHomeOptions {
	_options.NewInstanceID = core.StringPtr(newInstanceID)
	return _options
}

// SetNewProvider : Allow user to set NewProvider
func (_options *CreateInstanceHomeOptions) SetNewProvider(newProvider string) *CreateInstanceHomeOptions {
	_options.NewProvider = core.StringPtr(newProvider)
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *CreateInstanceHomeOptions) SetNewType(newType string) *CreateInstanceHomeOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetNewRegion : Allow user to set NewRegion
func (_options *CreateInstanceHomeOptions) SetNewRegion(newRegion string) *CreateInstanceHomeOptions {
	_options.NewRegion = core.StringPtr(newRegion)
	return _options
}

// SetNewEndpoint : Allow user to set NewEndpoint
func (_options *CreateInstanceHomeOptions) SetNewEndpoint(newEndpoint string) *CreateInstanceHomeOptions {
	_options.NewEndpoint = core.StringPtr(newEndpoint)
	return _options
}

// SetNewHmacAccessKey : Allow user to set NewHmacAccessKey
func (_options *CreateInstanceHomeOptions) SetNewHmacAccessKey(newHmacAccessKey string) *CreateInstanceHomeOptions {
	_options.NewHmacAccessKey = core.StringPtr(newHmacAccessKey)
	return _options
}

// SetNewHmacSecretKey : Allow user to set NewHmacSecretKey
func (_options *CreateInstanceHomeOptions) SetNewHmacSecretKey(newHmacSecretKey string) *CreateInstanceHomeOptions {
	_options.NewHmacSecretKey = core.StringPtr(newHmacSecretKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateInstanceHomeOptions) SetHeaders(param map[string]string) *CreateInstanceHomeOptions {
	options.Headers = param
	return options
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

// DeleteLoggingConfigurationOptions : The DeleteLoggingConfiguration options.
type DeleteLoggingConfigurationOptions struct {
	// Identifier of the instance to which the application belongs.
	InstanceGuid *string `json:"instance_guid" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLoggingConfigurationOptions : Instantiate DeleteLoggingConfigurationOptions
func (*IbmAnalyticsEngineApiV3) NewDeleteLoggingConfigurationOptions(instanceGuid string) *DeleteLoggingConfigurationOptions {
	return &DeleteLoggingConfigurationOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *DeleteLoggingConfigurationOptions) SetInstanceGuid(instanceGuid string) *DeleteLoggingConfigurationOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLoggingConfigurationOptions) SetHeaders(param map[string]string) *DeleteLoggingConfigurationOptions {
	options.Headers = param
	return options
}

// DisablePlatformLoggingOptions : The DisablePlatformLogging options.
type DisablePlatformLoggingOptions struct {
	// The identifier of the instance details to be added.
	InstanceGuid *string `json:"instance_guid" validate:"required,ne="`

	// enable platform logging.
	Enable *bool `json:"enable,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDisablePlatformLoggingOptions : Instantiate DisablePlatformLoggingOptions
func (*IbmAnalyticsEngineApiV3) NewDisablePlatformLoggingOptions(instanceGuid string) *DisablePlatformLoggingOptions {
	return &DisablePlatformLoggingOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *DisablePlatformLoggingOptions) SetInstanceGuid(instanceGuid string) *DisablePlatformLoggingOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetEnable : Allow user to set Enable
func (_options *DisablePlatformLoggingOptions) SetEnable(enable bool) *DisablePlatformLoggingOptions {
	_options.Enable = core.BoolPtr(enable)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DisablePlatformLoggingOptions) SetHeaders(param map[string]string) *DisablePlatformLoggingOptions {
	options.Headers = param
	return options
}

// EnablePlatformLoggingOptions : The EnablePlatformLogging options.
type EnablePlatformLoggingOptions struct {
	// The identifier of the instance details to be added.
	InstanceGuid *string `json:"instance_guid" validate:"required,ne="`

	// enable platform logging.
	Enable *bool `json:"enable,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEnablePlatformLoggingOptions : Instantiate EnablePlatformLoggingOptions
func (*IbmAnalyticsEngineApiV3) NewEnablePlatformLoggingOptions(instanceGuid string) *EnablePlatformLoggingOptions {
	return &EnablePlatformLoggingOptions{
		InstanceGuid: core.StringPtr(instanceGuid),
	}
}

// SetInstanceGuid : Allow user to set InstanceGuid
func (_options *EnablePlatformLoggingOptions) SetInstanceGuid(instanceGuid string) *EnablePlatformLoggingOptions {
	_options.InstanceGuid = core.StringPtr(instanceGuid)
	return _options
}

// SetEnable : Allow user to set Enable
func (_options *EnablePlatformLoggingOptions) SetEnable(enable bool) *EnablePlatformLoggingOptions {
	_options.Enable = core.BoolPtr(enable)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *EnablePlatformLoggingOptions) SetHeaders(param map[string]string) *EnablePlatformLoggingOptions {
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

// GetLoggingConfigurationOptions : The GetLoggingConfiguration options.
type GetLoggingConfigurationOptions struct {
	// GUID of the Analytics Engine service instance to retrieve.
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

// Instance : Details of Analytics Engine instance.
type Instance struct {
	// GUID of the Analytics Engine instance.
	ID *string `json:"id,omitempty"`

	// Full URL of the resource.
	Href *string `json:"href,omitempty"`

	// Instance state.
	State *string `json:"state,omitempty"`

	// Timestamp when the state of the instance was changed, in the format YYYY-MM-DDTHH:mm:ssZ.
	StateChangeTime *strfmt.DateTime `json:"state_change_time,omitempty"`

	// Specifies the default runtime to use for all workloads that run in this instance.
	DefaultRuntime *InstanceDefaultRuntime `json:"default_runtime,omitempty"`

	// Object storage instance that acts as the home for custom libraries and Spark events.
	InstanceHome *InstanceHome `json:"instance_home,omitempty"`

	// Instance level default configuration for Spark workloads.
	DefaultConfig *InstanceDefaultConfig `json:"default_config,omitempty"`
}

// Constants associated with the Instance.State property.
// Instance state.
const (
	Instance_State_Created = "created"
	Instance_State_Deleted = "deleted"
	Instance_State_Failed = "failed"
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
	err = core.UnmarshalModel(m, "default_runtime", &obj.DefaultRuntime, UnmarshalInstanceDefaultRuntime)
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

// InstanceDefaultRuntime : Specifies the default runtime to use for all workloads that run in this instance.
type InstanceDefaultRuntime struct {
	// Version of Spark runtime to use. Currently, only 3.1 is supported.
	SparkVersion *string `json:"spark_version,omitempty"`
}

// UnmarshalInstanceDefaultRuntime unmarshals an instance of InstanceDefaultRuntime from the specified map of raw messages.
func UnmarshalInstanceDefaultRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDefaultRuntime)
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
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

	// Instance state.
	State *string `json:"state,omitempty"`
}

// Constants associated with the InstanceGetStateResponse.State property.
// Instance state.
const (
	InstanceGetStateResponse_State_Created = "created"
	InstanceGetStateResponse_State_Deleted = "deleted"
	InstanceGetStateResponse_State_Failed = "failed"
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
	// Identifier of the instance where the applications run.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

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

// SetHeaders : Allow user to set Headers
func (options *ListApplicationsOptions) SetHeaders(param map[string]string) *ListApplicationsOptions {
	options.Headers = param
	return options
}

// LoggingConfigurationResponse : Response of logging API.
type LoggingConfigurationResponse struct {
	// component array.
	Components []string `json:"components,omitempty"`

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

// LoggingConfigurationResponseLogServer : LoggingConfigurationResponseLogServer struct
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
