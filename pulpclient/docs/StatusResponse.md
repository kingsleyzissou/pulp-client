# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]VersionResponse**](VersionResponse.md) | Version information of Pulp components | 
**OnlineWorkers** | [**[]WorkerResponse**](WorkerResponse.md) | List of online workers known to the application. An online worker is actively heartbeating and can respond to new work | 
**OnlineContentApps** | [**[]ContentAppStatusResponse**](ContentAppStatusResponse.md) | List of online content apps known to the application. An online content app is actively heartbeating and can serve data to clients | 
**DatabaseConnection** | [**DatabaseConnectionResponse**](DatabaseConnectionResponse.md) |  | 
**RedisConnection** | Pointer to [**RedisConnectionResponse**](RedisConnectionResponse.md) |  | [optional] 
**Storage** | Pointer to [**StorageResponse**](StorageResponse.md) |  | [optional] 
**ContentSettings** | [**ContentSettingsResponse**](ContentSettingsResponse.md) |  | 
**DomainEnabled** | **bool** | Is Domains enabled | 

## Methods

### NewStatusResponse

`func NewStatusResponse(versions []VersionResponse, onlineWorkers []WorkerResponse, onlineContentApps []ContentAppStatusResponse, databaseConnection DatabaseConnectionResponse, contentSettings ContentSettingsResponse, domainEnabled bool, ) *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *StatusResponse) GetVersions() []VersionResponse`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StatusResponse) GetVersionsOk() (*[]VersionResponse, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StatusResponse) SetVersions(v []VersionResponse)`

SetVersions sets Versions field to given value.


### GetOnlineWorkers

`func (o *StatusResponse) GetOnlineWorkers() []WorkerResponse`

GetOnlineWorkers returns the OnlineWorkers field if non-nil, zero value otherwise.

### GetOnlineWorkersOk

`func (o *StatusResponse) GetOnlineWorkersOk() (*[]WorkerResponse, bool)`

GetOnlineWorkersOk returns a tuple with the OnlineWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineWorkers

`func (o *StatusResponse) SetOnlineWorkers(v []WorkerResponse)`

SetOnlineWorkers sets OnlineWorkers field to given value.


### GetOnlineContentApps

`func (o *StatusResponse) GetOnlineContentApps() []ContentAppStatusResponse`

GetOnlineContentApps returns the OnlineContentApps field if non-nil, zero value otherwise.

### GetOnlineContentAppsOk

`func (o *StatusResponse) GetOnlineContentAppsOk() (*[]ContentAppStatusResponse, bool)`

GetOnlineContentAppsOk returns a tuple with the OnlineContentApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineContentApps

`func (o *StatusResponse) SetOnlineContentApps(v []ContentAppStatusResponse)`

SetOnlineContentApps sets OnlineContentApps field to given value.


### GetDatabaseConnection

`func (o *StatusResponse) GetDatabaseConnection() DatabaseConnectionResponse`

GetDatabaseConnection returns the DatabaseConnection field if non-nil, zero value otherwise.

### GetDatabaseConnectionOk

`func (o *StatusResponse) GetDatabaseConnectionOk() (*DatabaseConnectionResponse, bool)`

GetDatabaseConnectionOk returns a tuple with the DatabaseConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseConnection

`func (o *StatusResponse) SetDatabaseConnection(v DatabaseConnectionResponse)`

SetDatabaseConnection sets DatabaseConnection field to given value.


### GetRedisConnection

`func (o *StatusResponse) GetRedisConnection() RedisConnectionResponse`

GetRedisConnection returns the RedisConnection field if non-nil, zero value otherwise.

### GetRedisConnectionOk

`func (o *StatusResponse) GetRedisConnectionOk() (*RedisConnectionResponse, bool)`

GetRedisConnectionOk returns a tuple with the RedisConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisConnection

`func (o *StatusResponse) SetRedisConnection(v RedisConnectionResponse)`

SetRedisConnection sets RedisConnection field to given value.

### HasRedisConnection

`func (o *StatusResponse) HasRedisConnection() bool`

HasRedisConnection returns a boolean if a field has been set.

### GetStorage

`func (o *StatusResponse) GetStorage() StorageResponse`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StatusResponse) GetStorageOk() (*StorageResponse, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StatusResponse) SetStorage(v StorageResponse)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *StatusResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetContentSettings

`func (o *StatusResponse) GetContentSettings() ContentSettingsResponse`

GetContentSettings returns the ContentSettings field if non-nil, zero value otherwise.

### GetContentSettingsOk

`func (o *StatusResponse) GetContentSettingsOk() (*ContentSettingsResponse, bool)`

GetContentSettingsOk returns a tuple with the ContentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSettings

`func (o *StatusResponse) SetContentSettings(v ContentSettingsResponse)`

SetContentSettings sets ContentSettings field to given value.


### GetDomainEnabled

`func (o *StatusResponse) GetDomainEnabled() bool`

GetDomainEnabled returns the DomainEnabled field if non-nil, zero value otherwise.

### GetDomainEnabledOk

`func (o *StatusResponse) GetDomainEnabledOk() (*bool, bool)`

GetDomainEnabledOk returns a tuple with the DomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEnabled

`func (o *StatusResponse) SetDomainEnabled(v bool)`

SetDomainEnabled sets DomainEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


