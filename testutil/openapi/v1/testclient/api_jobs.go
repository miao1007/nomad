/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// JobsApiService JobsApi service
type JobsApiService service

type ApiJobJobNamePlanPostRequest struct {
	ctx _context.Context
	ApiService *JobsApiService
	jobName string
	jobPlanRequest *JobPlanRequest
	region *string
	namespace *string
	index *int32
	wait *int32
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiJobJobNamePlanPostRequest) JobPlanRequest(jobPlanRequest JobPlanRequest) ApiJobJobNamePlanPostRequest {
	r.jobPlanRequest = &jobPlanRequest
	return r
}
func (r ApiJobJobNamePlanPostRequest) Region(region string) ApiJobJobNamePlanPostRequest {
	r.region = &region
	return r
}
func (r ApiJobJobNamePlanPostRequest) Namespace(namespace string) ApiJobJobNamePlanPostRequest {
	r.namespace = &namespace
	return r
}
func (r ApiJobJobNamePlanPostRequest) Index(index int32) ApiJobJobNamePlanPostRequest {
	r.index = &index
	return r
}
func (r ApiJobJobNamePlanPostRequest) Wait(wait int32) ApiJobJobNamePlanPostRequest {
	r.wait = &wait
	return r
}
func (r ApiJobJobNamePlanPostRequest) Stale(stale string) ApiJobJobNamePlanPostRequest {
	r.stale = &stale
	return r
}
func (r ApiJobJobNamePlanPostRequest) Prefix(prefix string) ApiJobJobNamePlanPostRequest {
	r.prefix = &prefix
	return r
}
func (r ApiJobJobNamePlanPostRequest) XNomadToken(xNomadToken string) ApiJobJobNamePlanPostRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiJobJobNamePlanPostRequest) PerPage(perPage int32) ApiJobJobNamePlanPostRequest {
	r.perPage = &perPage
	return r
}
func (r ApiJobJobNamePlanPostRequest) NextToken(nextToken string) ApiJobJobNamePlanPostRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiJobJobNamePlanPostRequest) Execute() (JobPlanResponse, *_nethttp.Response, error) {
	return r.ApiService.JobJobNamePlanPostExecute(r)
}

/*
 * JobJobNamePlanPost Method for JobJobNamePlanPost
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jobName The job identifier.
 * @return ApiJobJobNamePlanPostRequest
 */
func (a *JobsApiService) JobJobNamePlanPost(ctx _context.Context, jobName string) ApiJobJobNamePlanPostRequest {
	return ApiJobJobNamePlanPostRequest{
		ApiService: a,
		ctx: ctx,
		jobName: jobName,
	}
}

/*
 * Execute executes the request
 * @return JobPlanResponse
 */
func (a *JobsApiService) JobJobNamePlanPostExecute(r ApiJobJobNamePlanPostRequest) (JobPlanResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobPlanResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsApiService.JobJobNamePlanPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobName}/plan"
	localVarPath = strings.Replace(localVarPath, "{"+"jobName"+"}", _neturl.PathEscape(parameterToString(r.jobName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobPlanRequest == nil {
		return localVarReturnValue, nil, reportError("jobPlanRequest is required and must be specified")
	}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
	}
	if r.xNomadToken != nil {
		localVarHeaderParams["X-Nomad-Token"] = parameterToString(*r.xNomadToken, "")
	}
	// body params
	localVarPostBody = r.jobPlanRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiJobsJobNameGetRequest struct {
	ctx _context.Context
	ApiService *JobsApiService
	jobName string
	region *string
	namespace *string
	index *int32
	wait *int32
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiJobsJobNameGetRequest) Region(region string) ApiJobsJobNameGetRequest {
	r.region = &region
	return r
}
func (r ApiJobsJobNameGetRequest) Namespace(namespace string) ApiJobsJobNameGetRequest {
	r.namespace = &namespace
	return r
}
func (r ApiJobsJobNameGetRequest) Index(index int32) ApiJobsJobNameGetRequest {
	r.index = &index
	return r
}
func (r ApiJobsJobNameGetRequest) Wait(wait int32) ApiJobsJobNameGetRequest {
	r.wait = &wait
	return r
}
func (r ApiJobsJobNameGetRequest) Stale(stale string) ApiJobsJobNameGetRequest {
	r.stale = &stale
	return r
}
func (r ApiJobsJobNameGetRequest) Prefix(prefix string) ApiJobsJobNameGetRequest {
	r.prefix = &prefix
	return r
}
func (r ApiJobsJobNameGetRequest) XNomadToken(xNomadToken string) ApiJobsJobNameGetRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiJobsJobNameGetRequest) PerPage(perPage int32) ApiJobsJobNameGetRequest {
	r.perPage = &perPage
	return r
}
func (r ApiJobsJobNameGetRequest) NextToken(nextToken string) ApiJobsJobNameGetRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiJobsJobNameGetRequest) Execute() (Job, *_nethttp.Response, error) {
	return r.ApiService.JobsJobNameGetExecute(r)
}

/*
 * JobsJobNameGet Method for JobsJobNameGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jobName The job identifier.
 * @return ApiJobsJobNameGetRequest
 */
func (a *JobsApiService) JobsJobNameGet(ctx _context.Context, jobName string) ApiJobsJobNameGetRequest {
	return ApiJobsJobNameGetRequest{
		ApiService: a,
		ctx: ctx,
		jobName: jobName,
	}
}

/*
 * Execute executes the request
 * @return Job
 */
func (a *JobsApiService) JobsJobNameGetExecute(r ApiJobsJobNameGetRequest) (Job, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Job
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsApiService.JobsJobNameGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobs/{jobName}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobName"+"}", _neturl.PathEscape(parameterToString(r.jobName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
	}
	if r.xNomadToken != nil {
		localVarHeaderParams["X-Nomad-Token"] = parameterToString(*r.xNomadToken, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiJobsPostRequest struct {
	ctx _context.Context
	ApiService *JobsApiService
	jobRegisterRequest *JobRegisterRequest
}

func (r ApiJobsPostRequest) JobRegisterRequest(jobRegisterRequest JobRegisterRequest) ApiJobsPostRequest {
	r.jobRegisterRequest = &jobRegisterRequest
	return r
}

func (r ApiJobsPostRequest) Execute() (JobRegisterResponse, *_nethttp.Response, error) {
	return r.ApiService.JobsPostExecute(r)
}

/*
 * JobsPost Method for JobsPost
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiJobsPostRequest
 */
func (a *JobsApiService) JobsPost(ctx _context.Context) ApiJobsPostRequest {
	return ApiJobsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return JobRegisterResponse
 */
func (a *JobsApiService) JobsPostExecute(r ApiJobsPostRequest) (JobRegisterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobRegisterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsApiService.JobsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobRegisterRequest == nil {
		return localVarReturnValue, nil, reportError("jobRegisterRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.jobRegisterRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
