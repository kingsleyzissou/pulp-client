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


// ContentPackagegroupsAPIService ContentPackagegroupsAPI service
type ContentPackagegroupsAPIService service

type ContentPackagegroupsAPIContentRpmPackagegroupsListRequest struct {
	ctx context.Context
	ApiService *ContentPackagegroupsAPIService
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
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) Limit(limit int32) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) Offset(offset int32) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) Ordering(ordering []string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) PulpIdIn(pulpIdIn []string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) RepositoryVersion(repositoryVersion string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) Fields(fields []string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) ExcludeFields(excludeFields []string) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) Execute() (*PaginatedrpmPackageGroupResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagegroupsListExecute(r)
}

/*
ContentRpmPackagegroupsList List package groups

PackageGroup ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackagegroupsAPIContentRpmPackagegroupsListRequest
*/
func (a *ContentPackagegroupsAPIService) ContentRpmPackagegroupsList(ctx context.Context) ContentPackagegroupsAPIContentRpmPackagegroupsListRequest {
	return ContentPackagegroupsAPIContentRpmPackagegroupsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageGroupResponseList
func (a *ContentPackagegroupsAPIService) ContentRpmPackagegroupsListExecute(r ContentPackagegroupsAPIContentRpmPackagegroupsListRequest) (*PaginatedrpmPackageGroupResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageGroupResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagegroupsAPIService.ContentRpmPackagegroupsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/packagegroups/"

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

type ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagegroupsAPIService
	rpmPackageGroupHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest) Fields(fields []string) ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest) ExcludeFields(excludeFields []string) ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest) Execute() (*RpmPackageGroupResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagegroupsReadExecute(r)
}

/*
ContentRpmPackagegroupsRead Inspect a package group

PackageGroup ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageGroupHref
 @return ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest
*/
func (a *ContentPackagegroupsAPIService) ContentRpmPackagegroupsRead(ctx context.Context, rpmPackageGroupHref string) ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest {
	return ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageGroupHref: rpmPackageGroupHref,
	}
}

// Execute executes the request
//  @return RpmPackageGroupResponse
func (a *ContentPackagegroupsAPIService) ContentRpmPackagegroupsReadExecute(r ContentPackagegroupsAPIContentRpmPackagegroupsReadRequest) (*RpmPackageGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagegroupsAPIService.ContentRpmPackagegroupsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{rpm_package_group_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_group_href"+"}", parameterValueToString(r.rpmPackageGroupHref, "rpmPackageGroupHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
