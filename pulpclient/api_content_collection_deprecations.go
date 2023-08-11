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


// ContentCollectionDeprecationsAPIService ContentCollectionDeprecationsAPI service
type ContentCollectionDeprecationsAPIService service

type ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest struct {
	ctx context.Context
	ApiService *ContentCollectionDeprecationsAPIService
	ansibleCollection *AnsibleCollection
}

func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest) AnsibleCollection(ansibleCollection AnsibleCollection) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest {
	r.ansibleCollection = &ansibleCollection
	return r
}

func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest) Execute() (*AnsibleCollectionResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionDeprecationsCreateExecute(r)
}

/*
ContentAnsibleCollectionDeprecationsCreate Create an ansible collection deprecated

ViewSet for AnsibleCollectionDeprecated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest
*/
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsCreate(ctx context.Context) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest {
	return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AnsibleCollectionResponse
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsCreateExecute(r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsCreateRequest) (*AnsibleCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionDeprecationsAPIService.ContentAnsibleCollectionDeprecationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_deprecations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ansibleCollection == nil {
		return localVarReturnValue, nil, reportError("ansibleCollection is required and must be specified")
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
	localVarPostBody = r.ansibleCollection
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

type ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest struct {
	ctx context.Context
	ApiService *ContentCollectionDeprecationsAPIService
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) Limit(limit int32) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) Offset(offset int32) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) Ordering(ordering []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) PulpHrefIn(pulpHrefIn []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) PulpIdIn(pulpIdIn []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) RepositoryVersion(repositoryVersion string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) Fields(fields []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) ExcludeFields(excludeFields []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) Execute() (*PaginatedansibleCollectionResponseList, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionDeprecationsListExecute(r)
}

/*
ContentAnsibleCollectionDeprecationsList List ansible collection deprecateds

ViewSet for AnsibleCollectionDeprecated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest
*/
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsList(ctx context.Context) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest {
	return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionResponseList
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsListExecute(r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsListRequest) (*PaginatedansibleCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionDeprecationsAPIService.ContentAnsibleCollectionDeprecationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_deprecations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
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

type ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest struct {
	ctx context.Context
	ApiService *ContentCollectionDeprecationsAPIService
	ansibleAnsibleCollectionDeprecatedHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest) Fields(fields []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest) ExcludeFields(excludeFields []string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest) Execute() (*AnsibleCollectionResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionDeprecationsReadExecute(r)
}

/*
ContentAnsibleCollectionDeprecationsRead Inspect an ansible collection deprecated

ViewSet for AnsibleCollectionDeprecated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleAnsibleCollectionDeprecatedHref
 @return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest
*/
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsRead(ctx context.Context, ansibleAnsibleCollectionDeprecatedHref string) ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest {
	return ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleAnsibleCollectionDeprecatedHref: ansibleAnsibleCollectionDeprecatedHref,
	}
}

// Execute executes the request
//  @return AnsibleCollectionResponse
func (a *ContentCollectionDeprecationsAPIService) ContentAnsibleCollectionDeprecationsReadExecute(r ContentCollectionDeprecationsAPIContentAnsibleCollectionDeprecationsReadRequest) (*AnsibleCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionDeprecationsAPIService.ContentAnsibleCollectionDeprecationsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{ansible_ansible_collection_deprecated_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_ansible_collection_deprecated_href"+"}", parameterValueToString(r.ansibleAnsibleCollectionDeprecatedHref, "ansibleAnsibleCollectionDeprecatedHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
