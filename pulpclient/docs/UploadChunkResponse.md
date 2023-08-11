# UploadChunkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** |  | [optional] [readonly] 
**Size** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewUploadChunkResponse

`func NewUploadChunkResponse() *UploadChunkResponse`

NewUploadChunkResponse instantiates a new UploadChunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadChunkResponseWithDefaults

`func NewUploadChunkResponseWithDefaults() *UploadChunkResponse`

NewUploadChunkResponseWithDefaults instantiates a new UploadChunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *UploadChunkResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UploadChunkResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UploadChunkResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UploadChunkResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSize

`func (o *UploadChunkResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadChunkResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadChunkResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UploadChunkResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


