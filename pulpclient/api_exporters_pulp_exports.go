/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ExportersPulpExportsAPIService ExportersPulpExportsAPI service
type ExportersPulpExportsAPIService service

type ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpExportsAPIService
	pulpExporterHref string
	pulpExport *PulpExport
}

func (r ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest) PulpExport(pulpExport PulpExport) ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest {
	r.pulpExport = &pulpExport
	return r
}

func (r ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpExportsCreateExecute(r)
}

/*
ExportersCorePulpExportsCreate Create a pulp export

Trigger an asynchronous task to export a set of repositories

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest
*/
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsCreate(ctx context.Context, pulpExporterHref string) ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest {
	return ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsCreateExecute(r ExportersPulpExportsAPIExportersCorePulpExportsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpExportsAPIService.ExportersCorePulpExportsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{pulp_exporter_href}exports/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", parameterValueToString(r.pulpExporterHref, "pulpExporterHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpExport == nil {
		return localVarReturnValue, nil, reportError("pulpExport is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	localVarPostBody = r.pulpExport
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest struct {
	ctx context.Context
	ApiService *ExportersPulpExportsAPIService
	pulpPulpExportHref string
}

func (r ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExportersCorePulpExportsDeleteExecute(r)
}

/*
ExportersCorePulpExportsDelete Delete a pulp export

ViewSet for viewing exports from a PulpExporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpPulpExportHref
 @return ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest
*/
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsDelete(ctx context.Context, pulpPulpExportHref string) ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest {
	return ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpPulpExportHref: pulpPulpExportHref,
	}
}

// Execute executes the request
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsDeleteExecute(r ExportersPulpExportsAPIExportersCorePulpExportsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpExportsAPIService.ExportersCorePulpExportsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{pulp_pulp_export_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_pulp_export_href"+"}", parameterValueToString(r.pulpPulpExportHref, "pulpPulpExportHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ExportersPulpExportsAPIExportersCorePulpExportsListRequest struct {
	ctx context.Context
	ApiService *ExportersPulpExportsAPIService
	pulpExporterHref string
	limit *int32
	offset *int32
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) Limit(limit int32) ExportersPulpExportsAPIExportersCorePulpExportsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) Offset(offset int32) ExportersPulpExportsAPIExportersCorePulpExportsListRequest {
	r.offset = &offset
	return r
}

// A list of fields to include in the response.
func (r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) Fields(fields []string) ExportersPulpExportsAPIExportersCorePulpExportsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) ExcludeFields(excludeFields []string) ExportersPulpExportsAPIExportersCorePulpExportsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) Execute() (*PaginatedPulpExportResponseList, *http.Response, error) {
	return r.ApiService.ExportersCorePulpExportsListExecute(r)
}

/*
ExportersCorePulpExportsList List pulp exports

ViewSet for viewing exports from a PulpExporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpExportsAPIExportersCorePulpExportsListRequest
*/
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsList(ctx context.Context, pulpExporterHref string) ExportersPulpExportsAPIExportersCorePulpExportsListRequest {
	return ExportersPulpExportsAPIExportersCorePulpExportsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return PaginatedPulpExportResponseList
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsListExecute(r ExportersPulpExportsAPIExportersCorePulpExportsListRequest) (*PaginatedPulpExportResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPulpExportResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpExportsAPIService.ExportersCorePulpExportsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{pulp_exporter_href}exports/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", parameterValueToString(r.pulpExporterHref, "pulpExporterHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExportersPulpExportsAPIExportersCorePulpExportsReadRequest struct {
	ctx context.Context
	ApiService *ExportersPulpExportsAPIService
	pulpPulpExportHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ExportersPulpExportsAPIExportersCorePulpExportsReadRequest) Fields(fields []string) ExportersPulpExportsAPIExportersCorePulpExportsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersPulpExportsAPIExportersCorePulpExportsReadRequest) ExcludeFields(excludeFields []string) ExportersPulpExportsAPIExportersCorePulpExportsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersPulpExportsAPIExportersCorePulpExportsReadRequest) Execute() (*PulpExportResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpExportsReadExecute(r)
}

/*
ExportersCorePulpExportsRead Inspect a pulp export

ViewSet for viewing exports from a PulpExporter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpPulpExportHref
 @return ExportersPulpExportsAPIExportersCorePulpExportsReadRequest
*/
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsRead(ctx context.Context, pulpPulpExportHref string) ExportersPulpExportsAPIExportersCorePulpExportsReadRequest {
	return ExportersPulpExportsAPIExportersCorePulpExportsReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpPulpExportHref: pulpPulpExportHref,
	}
}

// Execute executes the request
//  @return PulpExportResponse
func (a *ExportersPulpExportsAPIService) ExportersCorePulpExportsReadExecute(r ExportersPulpExportsAPIExportersCorePulpExportsReadRequest) (*PulpExportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpExportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpExportsAPIService.ExportersCorePulpExportsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{pulp_pulp_export_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_pulp_export_href"+"}", parameterValueToString(r.pulpPulpExportHref, "pulpPulpExportHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
