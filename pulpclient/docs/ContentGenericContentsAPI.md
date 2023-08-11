# \ContentGenericContentsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebGenericContentsCreate**](ContentGenericContentsAPI.md#ContentDebGenericContentsCreate) | **Post** /pulp/api/v3/content/deb/generic_contents/ | Create a generic content
[**ContentDebGenericContentsList**](ContentGenericContentsAPI.md#ContentDebGenericContentsList) | **Get** /pulp/api/v3/content/deb/generic_contents/ | List generic contents
[**ContentDebGenericContentsRead**](ContentGenericContentsAPI.md#ContentDebGenericContentsRead) | **Get** {deb_generic_content_href} | Inspect a generic content



## ContentDebGenericContentsCreate

> AsyncOperationResponse ContentDebGenericContentsCreate(ctx).RelativePath(relativePath).Repository(repository).Artifact(artifact).File(file).Upload(upload).Execute()

Create a generic content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsAPI.ContentDebGenericContentsCreate(context.Background()).RelativePath(relativePath).Repository(repository).Artifact(artifact).File(file).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsAPI.ContentDebGenericContentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsAPI.ContentDebGenericContentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebGenericContentsList

> PaginateddebGenericContentResponseList ContentDebGenericContentsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List generic contents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `relative_path` - Relative path * `-relative_path` - Relative path (descending) * `sha256` - Sha256 * `-sha256` - Sha256 (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsAPI.ContentDebGenericContentsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsAPI.ContentDebGenericContentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsList`: PaginateddebGenericContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsAPI.ContentDebGenericContentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebGenericContentResponseList**](PaginateddebGenericContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebGenericContentsRead

> DebGenericContentResponse ContentDebGenericContentsRead(ctx, debGenericContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a generic content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func main() {
    debGenericContentHref := "debGenericContentHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentGenericContentsAPI.ContentDebGenericContentsRead(context.Background(), debGenericContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentGenericContentsAPI.ContentDebGenericContentsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebGenericContentsRead`: DebGenericContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentGenericContentsAPI.ContentDebGenericContentsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debGenericContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebGenericContentsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebGenericContentResponse**](DebGenericContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
