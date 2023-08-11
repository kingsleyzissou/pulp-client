/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
	"time"
)

// checks if the PulpImporterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpImporterResponse{}

// PulpImporterResponse Serializer for PulpImporters.
type PulpImporterResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Unique name of the Importer.
	Name string `json:"name"`
	// Mapping of repo names in an export file to the repo names in Pulp. For example, if the export has a repo named 'foo' and the repo to import content into was 'bar', the mapping would be \"{'foo': 'bar'}\".
	RepoMapping *map[string]string `json:"repo_mapping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PulpImporterResponse PulpImporterResponse

// NewPulpImporterResponse instantiates a new PulpImporterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpImporterResponse(name string) *PulpImporterResponse {
	this := PulpImporterResponse{}
	this.Name = name
	return &this
}

// NewPulpImporterResponseWithDefaults instantiates a new PulpImporterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpImporterResponseWithDefaults() *PulpImporterResponse {
	this := PulpImporterResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *PulpImporterResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImporterResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *PulpImporterResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *PulpImporterResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *PulpImporterResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImporterResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *PulpImporterResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *PulpImporterResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *PulpImporterResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PulpImporterResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PulpImporterResponse) SetName(v string) {
	o.Name = v
}

// GetRepoMapping returns the RepoMapping field value if set, zero value otherwise.
func (o *PulpImporterResponse) GetRepoMapping() map[string]string {
	if o == nil || IsNil(o.RepoMapping) {
		var ret map[string]string
		return ret
	}
	return *o.RepoMapping
}

// GetRepoMappingOk returns a tuple with the RepoMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImporterResponse) GetRepoMappingOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.RepoMapping) {
		return nil, false
	}
	return o.RepoMapping, true
}

// HasRepoMapping returns a boolean if a field has been set.
func (o *PulpImporterResponse) HasRepoMapping() bool {
	if o != nil && !IsNil(o.RepoMapping) {
		return true
	}

	return false
}

// SetRepoMapping gets a reference to the given map[string]string and assigns it to the RepoMapping field.
func (o *PulpImporterResponse) SetRepoMapping(v map[string]string) {
	o.RepoMapping = &v
}

func (o PulpImporterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpImporterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.RepoMapping) {
		toSerialize["repo_mapping"] = o.RepoMapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PulpImporterResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPulpImporterResponse := _PulpImporterResponse{}

	if err = json.Unmarshal(bytes, &varPulpImporterResponse); err == nil {
		*o = PulpImporterResponse(varPulpImporterResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repo_mapping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePulpImporterResponse struct {
	value *PulpImporterResponse
	isSet bool
}

func (v NullablePulpImporterResponse) Get() *PulpImporterResponse {
	return v.value
}

func (v *NullablePulpImporterResponse) Set(val *PulpImporterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpImporterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpImporterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpImporterResponse(val *PulpImporterResponse) *NullablePulpImporterResponse {
	return &NullablePulpImporterResponse{value: val, isSet: true}
}

func (v NullablePulpImporterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpImporterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


