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


// PulpAnsibleApiV3CollectionsAPIService PulpAnsibleApiV3CollectionsAPI service
type PulpAnsibleApiV3CollectionsAPIService service

type PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsAPIService
	name string
	namespace string
	path string
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsDeleteExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsDelete Method for PulpAnsibleGalaxyApiV3CollectionsDelete

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsDelete(ctx context.Context, name string, namespace string, path string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest {
	return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsDeleteExecute(r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsAPIService.PulpAnsibleGalaxyApiV3CollectionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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

type PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsAPIService
	path string
	deprecated *bool
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Deprecated(deprecated bool) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.deprecated = &deprecated
	return r
}

// Number of results to return per page.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Limit(limit int32) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Name(name string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.name = &name
	return r
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Namespace(namespace string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Offset(offset int32) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Ordering(ordering []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) PulpHrefIn(pulpHrefIn []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) PulpIdIn(pulpIdIn []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) Execute() (*PaginatedCollectionResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsListExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsList Method for PulpAnsibleGalaxyApiV3CollectionsList

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsList(ctx context.Context, path string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest {
	return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return PaginatedCollectionResponseList
// Deprecated
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsListExecute(r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsListRequest) (*PaginatedCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsAPIService.PulpAnsibleGalaxyApiV3CollectionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deprecated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deprecated", r.deprecated, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
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

type PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsAPIService
	name string
	namespace string
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest) Execute() (*CollectionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsRead Method for PulpAnsibleGalaxyApiV3CollectionsRead

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsRead(ctx context.Context, name string, namespace string, path string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest {
	return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return CollectionResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsReadExecute(r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsReadRequest) (*CollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsAPIService.PulpAnsibleGalaxyApiV3CollectionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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

type PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsAPIService
	name string
	namespace string
	path string
	patchedCollection *PatchedCollection
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest) PatchedCollection(patchedCollection PatchedCollection) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest {
	r.patchedCollection = &patchedCollection
	return r
}

func (r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsUpdateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsUpdate Method for PulpAnsibleGalaxyApiV3CollectionsUpdate

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsUpdate(ctx context.Context, name string, namespace string, path string) PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest {
	return PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsAPIService) PulpAnsibleGalaxyApiV3CollectionsUpdateExecute(r PulpAnsibleApiV3CollectionsAPIPulpAnsibleGalaxyApiV3CollectionsUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsAPIService.PulpAnsibleGalaxyApiV3CollectionsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedCollection == nil {
		return localVarReturnValue, nil, reportError("patchedCollection is required and must be specified")
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
	localVarPostBody = r.patchedCollection
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
