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
 * IBM OpenAPI SDK Code Generator Version: 3.29.0-cd9ba74f-20210305-183535
 */

// Package ibmanalyticsengineapiv3 : Operations and models for the IbmAnalyticsEngineApiV3 service
package ibmanalyticsengineapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/ibm-iae-go-sdk/common"
	"github.com/go-openapi/strfmt"
	"net/http"
	"reflect"
	"time"
)

// IbmAnalyticsEngineApiV3 : Create and manage serverless Spark instances and run applications.
//
// Version: 3.0.0
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

// GetInstanceByID : Find instance by id
// Retrieve the details of a single instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceByID(getInstanceByIdOptions *GetInstanceByIdOptions) (result *InstanceDetails, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetInstanceByIDWithContext(context.Background(), getInstanceByIdOptions)
}

// GetInstanceByIDWithContext is an alternate form of the GetInstanceByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetInstanceByIDWithContext(ctx context.Context, getInstanceByIdOptions *GetInstanceByIdOptions) (result *InstanceDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceByIdOptions, "getInstanceByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceByIdOptions, "getInstanceByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceByIdOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetInstanceByID")
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
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateApplication : Deploys a Spark application in the Analytics Engine instance
// Deploy a Spark application on a given serverless Spark instance.
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
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark/applications`, pathParamsMap)
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
	if createApplicationOptions.ApplicationArguments != nil {
		application_details["application_arguments"] = createApplicationOptions.ApplicationArguments
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
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetApplications : Gets all applications submitted in an instance with a specified inst_id
// Retrieve all Spark applications run on a given instance.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplications(getApplicationsOptions *GetApplicationsOptions) (result *ApplicationCollection, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetApplicationsWithContext(context.Background(), getApplicationsOptions)
}

// GetApplicationsWithContext is an alternate form of the GetApplications method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationsWithContext(ctx context.Context, getApplicationsOptions *GetApplicationsOptions) (result *ApplicationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationsOptions, "getApplicationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationsOptions, "getApplicationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getApplicationsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetApplications")
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
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetApplicationByID : Gets the details of the application identified by the app_id identifier
// Retrieve the details of a given Spark application.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationByID(getApplicationByIdOptions *GetApplicationByIdOptions) (result *ApplicationGetResponse, response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.GetApplicationByIDWithContext(context.Background(), getApplicationByIdOptions)
}

// GetApplicationByIDWithContext is an alternate form of the GetApplicationByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) GetApplicationByIDWithContext(ctx context.Context, getApplicationByIdOptions *GetApplicationByIdOptions) (result *ApplicationGetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationByIdOptions, "getApplicationByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationByIdOptions, "getApplicationByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getApplicationByIdOptions.InstanceID,
		"application_id": *getApplicationByIdOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark/applications/{application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "GetApplicationByID")
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
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationGetResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteApplicationByID : Stops the specified application
// Stops a running application identified by the app_id identifier. This is an idempotent operation. Performs no action
// if the requested application is already stopped or completed.
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteApplicationByID(deleteApplicationByIdOptions *DeleteApplicationByIdOptions) (response *core.DetailedResponse, err error) {
	return ibmAnalyticsEngineApi.DeleteApplicationByIDWithContext(context.Background(), deleteApplicationByIdOptions)
}

// DeleteApplicationByIDWithContext is an alternate form of the DeleteApplicationByID method which supports a Context parameter
func (ibmAnalyticsEngineApi *IbmAnalyticsEngineApiV3) DeleteApplicationByIDWithContext(ctx context.Context, deleteApplicationByIdOptions *DeleteApplicationByIdOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteApplicationByIdOptions, "deleteApplicationByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteApplicationByIdOptions, "deleteApplicationByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteApplicationByIdOptions.InstanceID,
		"application_id": *deleteApplicationByIdOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmAnalyticsEngineApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark/applications/{application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteApplicationByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_analytics_engine_api", "V3", "DeleteApplicationByID")
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

// GetApplicationState : Gets the status of the application identified by the app_id identifier
// Returns the status of the application identified by the app_id identifier.
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
	_, err = builder.ResolveRequestURL(ibmAnalyticsEngineApi.Service.Options.URL, `/v3/analytics_engines/{instance_id}/spark/applications/{application_id}/state`, pathParamsMap)
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
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationGetStateResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ApplicationCollection : An array of application details.
type ApplicationCollection struct {
	// List of applications.
	Applications []ApplicationDetails `json:"applications,omitempty"`
}

// UnmarshalApplicationCollection unmarshals an instance of ApplicationCollection from the specified map of raw messages.
func UnmarshalApplicationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationCollection)
	err = core.UnmarshalModel(m, "applications", &obj.Applications, UnmarshalApplicationDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationDetails : Details of an application.
type ApplicationDetails struct {
	// Identifier of the application.
	ApplicationID *string `json:"application_id,omitempty"`

	// Identifier of the Spark application.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Status of the application.
	State *string `json:"state,omitempty"`

	// Time when the application was started.
	StartTime *string `json:"start_time,omitempty"`

	// Time when the application was completed.
	FinishTime *string `json:"finish_time,omitempty"`
}

// UnmarshalApplicationDetails unmarshals an instance of ApplicationDetails from the specified map of raw messages.
func UnmarshalApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
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

// ApplicationGetResponse : Response of the Application Get API.
type ApplicationGetResponse struct {
	// Application request details.
	ApplicationDetails *ApplicationRequest `json:"application_details,omitempty"`

	// Application mode.
	Mode *string `json:"mode,omitempty"`

	// Application ID.
	ApplicationID *string `json:"application_id,omitempty"`

	// Application state.
	State *string `json:"state,omitempty"`

	// Application start time.
	StartTime *string `json:"start_time,omitempty"`

	// Application end time.
	FinishTime *string `json:"finish_time,omitempty"`
}

// UnmarshalApplicationGetResponse unmarshals an instance of ApplicationGetResponse from the specified map of raw messages.
func UnmarshalApplicationGetResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationGetResponse)
	err = core.UnmarshalModel(m, "application_details", &obj.ApplicationDetails, UnmarshalApplicationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
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
	ApplicationID *string `json:"application_id,omitempty"`

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
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
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

// ApplicationRequest : Application request details.
type ApplicationRequest struct {
	// Application details.
	ApplicationDetails *ApplicationRequestApplicationDetails `json:"application_details,omitempty"`
}

// UnmarshalApplicationRequest unmarshals an instance of ApplicationRequest from the specified map of raw messages.
func UnmarshalApplicationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationRequest)
	err = core.UnmarshalModel(m, "application_details", &obj.ApplicationDetails, UnmarshalApplicationRequestApplicationDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationRequestApplicationDetails : Application details.
type ApplicationRequestApplicationDetails struct {
	// Application name.
	Application *string `json:"application,omitempty"`

	// The entry point for your application.
	Class *string `json:"class,omitempty"`

	// Application arguments.
	ApplicationArguments []string `json:"application_arguments,omitempty"`

	// Application configurations to override. See [Spark environment variables](
	// https://spark.apache.org/docs/latest/configuration.html#available-properties) for a list of the supported variables.
	Conf map[string]interface{} `json:"conf,omitempty"`

	// Application environment configurations to override. See [Spark environment
	// variables](https://spark.apache.org/docs/latest/configuration.html#environment-variables) for a list of the
	// supported variables.
	Env map[string]interface{} `json:"env,omitempty"`
}

// UnmarshalApplicationRequestApplicationDetails unmarshals an instance of ApplicationRequestApplicationDetails from the specified map of raw messages.
func UnmarshalApplicationRequestApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationRequestApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "application_arguments", &obj.ApplicationArguments)
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

// ApplicationResponse : Application response details.
type ApplicationResponse struct {
	// Application ID.
	ApplicationID *string `json:"application_id,omitempty"`

	// Application state.
	State *string `json:"state,omitempty"`

	// Application start time.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`
}

// Constants associated with the ApplicationResponse.State property.
// Application state.
const (
	ApplicationResponse_State_Accepted = "accepted"
	ApplicationResponse_State_Error = "error"
	ApplicationResponse_State_Failed = "failed"
	ApplicationResponse_State_Finished = "finished"
	ApplicationResponse_State_Killed = "killed"
	ApplicationResponse_State_Running = "running"
	ApplicationResponse_State_Stopped = "stopped"
	ApplicationResponse_State_Submitted = "submitted"
	ApplicationResponse_State_Waiting = "waiting"
)

// UnmarshalApplicationResponse unmarshals an instance of ApplicationResponse from the specified map of raw messages.
func UnmarshalApplicationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationResponse)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateApplicationOptions : The CreateApplication options.
type CreateApplicationOptions struct {
	// The identifier of the instance where the Spark application is submitted.
	InstanceID *string `validate:"required,ne="`

	// Application name.
	Application *string

	// The entry point for your application.
	Class *string

	// Application arguments.
	ApplicationArguments []string

	// Application configurations to override. See [Spark environment variables](
	// https://spark.apache.org/docs/latest/configuration.html#available-properties) for a list of the supported variables.
	Conf map[string]interface{}

	// Application environment configurations to override. See [Spark environment
	// variables](https://spark.apache.org/docs/latest/configuration.html#environment-variables) for a list of the
	// supported variables.
	Env map[string]interface{}

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
func (options *CreateApplicationOptions) SetInstanceID(instanceID string) *CreateApplicationOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetApplication : Allow user to set Application
func (options *CreateApplicationOptions) SetApplication(application string) *CreateApplicationOptions {
	options.Application = core.StringPtr(application)
	return options
}

// SetClass : Allow user to set Class
func (options *CreateApplicationOptions) SetClass(class string) *CreateApplicationOptions {
	options.Class = core.StringPtr(class)
	return options
}

// SetApplicationArguments : Allow user to set ApplicationArguments
func (options *CreateApplicationOptions) SetApplicationArguments(applicationArguments []string) *CreateApplicationOptions {
	options.ApplicationArguments = applicationArguments
	return options
}

// SetConf : Allow user to set Conf
func (options *CreateApplicationOptions) SetConf(conf map[string]interface{}) *CreateApplicationOptions {
	options.Conf = conf
	return options
}

// SetEnv : Allow user to set Env
func (options *CreateApplicationOptions) SetEnv(env map[string]interface{}) *CreateApplicationOptions {
	options.Env = env
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateApplicationOptions) SetHeaders(param map[string]string) *CreateApplicationOptions {
	options.Headers = param
	return options
}

// DeleteApplicationByIdOptions : The DeleteApplicationByID options.
type DeleteApplicationByIdOptions struct {
	// Identifier of the instance to which the application belongs.
	InstanceID *string `validate:"required,ne="`

	// Identifier of the application that needs to be stopped.
	ApplicationID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteApplicationByIdOptions : Instantiate DeleteApplicationByIdOptions
func (*IbmAnalyticsEngineApiV3) NewDeleteApplicationByIdOptions(instanceID string, applicationID string) *DeleteApplicationByIdOptions {
	return &DeleteApplicationByIdOptions{
		InstanceID: core.StringPtr(instanceID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteApplicationByIdOptions) SetInstanceID(instanceID string) *DeleteApplicationByIdOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetApplicationID : Allow user to set ApplicationID
func (options *DeleteApplicationByIdOptions) SetApplicationID(applicationID string) *DeleteApplicationByIdOptions {
	options.ApplicationID = core.StringPtr(applicationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteApplicationByIdOptions) SetHeaders(param map[string]string) *DeleteApplicationByIdOptions {
	options.Headers = param
	return options
}

// GetApplicationByIdOptions : The GetApplicationByID options.
type GetApplicationByIdOptions struct {
	// Identifier of the instance to which the application belongs.
	InstanceID *string `validate:"required,ne="`

	// Identifier of the application for which details are requested.
	ApplicationID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationByIdOptions : Instantiate GetApplicationByIdOptions
func (*IbmAnalyticsEngineApiV3) NewGetApplicationByIdOptions(instanceID string, applicationID string) *GetApplicationByIdOptions {
	return &GetApplicationByIdOptions{
		InstanceID: core.StringPtr(instanceID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetApplicationByIdOptions) SetInstanceID(instanceID string) *GetApplicationByIdOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetApplicationID : Allow user to set ApplicationID
func (options *GetApplicationByIdOptions) SetApplicationID(applicationID string) *GetApplicationByIdOptions {
	options.ApplicationID = core.StringPtr(applicationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationByIdOptions) SetHeaders(param map[string]string) *GetApplicationByIdOptions {
	options.Headers = param
	return options
}

// GetApplicationStateOptions : The GetApplicationState options.
type GetApplicationStateOptions struct {
	// Identifier of the instance to which the applications belongs.
	InstanceID *string `validate:"required,ne="`

	// Identifier of the application for which details are requested.
	ApplicationID *string `validate:"required,ne="`

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
func (options *GetApplicationStateOptions) SetInstanceID(instanceID string) *GetApplicationStateOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetApplicationID : Allow user to set ApplicationID
func (options *GetApplicationStateOptions) SetApplicationID(applicationID string) *GetApplicationStateOptions {
	options.ApplicationID = core.StringPtr(applicationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationStateOptions) SetHeaders(param map[string]string) *GetApplicationStateOptions {
	options.Headers = param
	return options
}

// GetApplicationsOptions : The GetApplications options.
type GetApplicationsOptions struct {
	// Identifier of the instance where the applications run.
	InstanceID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationsOptions : Instantiate GetApplicationsOptions
func (*IbmAnalyticsEngineApiV3) NewGetApplicationsOptions(instanceID string) *GetApplicationsOptions {
	return &GetApplicationsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetApplicationsOptions) SetInstanceID(instanceID string) *GetApplicationsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationsOptions) SetHeaders(param map[string]string) *GetApplicationsOptions {
	options.Headers = param
	return options
}

// GetInstanceByIdOptions : The GetInstanceByID options.
type GetInstanceByIdOptions struct {
	// Identifier of the instance to retrieve.
	InstanceID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceByIdOptions : Instantiate GetInstanceByIdOptions
func (*IbmAnalyticsEngineApiV3) NewGetInstanceByIdOptions(instanceID string) *GetInstanceByIdOptions {
	return &GetInstanceByIdOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetInstanceByIdOptions) SetInstanceID(instanceID string) *GetInstanceByIdOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceByIdOptions) SetHeaders(param map[string]string) *GetInstanceByIdOptions {
	options.Headers = param
	return options
}

// InstanceDetails : Instance details.
type InstanceDetails struct {
	// Instance id.
	InstanceID *string `json:"instance_id,omitempty"`

	// Instance state.
	State *string `json:"state,omitempty"`

	// Timestamp when the state of the instance was changed.
	StateChangeTime *strfmt.DateTime `json:"state_change_time,omitempty"`

	// Specifies the default runtime to use for all workloads that run in this instance.
	DefaultRuntime *InstanceDetailsDefaultRuntime `json:"default_runtime,omitempty"`

	// Instance home storage associated with the instance.
	InstanceHome *InstanceDetailsInstanceHome `json:"instance_home,omitempty"`

	// Instance level default configuration for Spark workloads.
	DefaultConfig *InstanceDetailsDefaultConfig `json:"default_config,omitempty"`
}

// Constants associated with the InstanceDetails.State property.
// Instance state.
const (
	InstanceDetails_State_Created = "created"
	InstanceDetails_State_Deleted = "deleted"
	InstanceDetails_State_Failed = "failed"
)

// UnmarshalInstanceDetails unmarshals an instance of InstanceDetails from the specified map of raw messages.
func UnmarshalInstanceDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDetails)
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
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
	err = core.UnmarshalModel(m, "default_runtime", &obj.DefaultRuntime, UnmarshalInstanceDetailsDefaultRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "instance_home", &obj.InstanceHome, UnmarshalInstanceDetailsInstanceHome)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_config", &obj.DefaultConfig, UnmarshalInstanceDetailsDefaultConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceDetailsDefaultConfig : Instance level default configuration for Spark workloads.
type InstanceDetailsDefaultConfig struct {
	// Value of the Spark configuration key.
	Key *string `json:"key,omitempty"`
}

// UnmarshalInstanceDetailsDefaultConfig unmarshals an instance of InstanceDetailsDefaultConfig from the specified map of raw messages.
func UnmarshalInstanceDetailsDefaultConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDetailsDefaultConfig)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceDetailsDefaultRuntime : Specifies the default runtime to use for all workloads that run in this instance.
type InstanceDetailsDefaultRuntime struct {
	// Version of Spark runtime to use. Currently, only 3.0 is supported.
	SparkVersion *string `json:"spark_version,omitempty"`

	// Add-on packages.
	AdditionalPackages []string `json:"additional_packages,omitempty"`
}

// UnmarshalInstanceDetailsDefaultRuntime unmarshals an instance of InstanceDetailsDefaultRuntime from the specified map of raw messages.
func UnmarshalInstanceDetailsDefaultRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDetailsDefaultRuntime)
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "additional_packages", &obj.AdditionalPackages)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceDetailsInstanceHome : Instance home storage associated with the instance.
type InstanceDetailsInstanceHome struct {
	// UUID of the instance home storage instance.
	Guid *string `json:"guid,omitempty"`

	// Currently only ibm-cos (IBM Cloud Object Storage) is supported.
	Provider *string `json:"provider,omitempty"`

	// Type of the instance home storage. Currently, only objectstore (Cloud Object Storage)is supported.
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

// UnmarshalInstanceDetailsInstanceHome unmarshals an instance of InstanceDetailsInstanceHome from the specified map of raw messages.
func UnmarshalInstanceDetailsInstanceHome(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceDetailsInstanceHome)
	err = core.UnmarshalPrimitive(m, "guid", &obj.Guid)
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
