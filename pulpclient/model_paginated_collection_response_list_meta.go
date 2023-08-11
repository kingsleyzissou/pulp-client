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
)

// checks if the PaginatedCollectionResponseListMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCollectionResponseListMeta{}

// PaginatedCollectionResponseListMeta struct for PaginatedCollectionResponseListMeta
type PaginatedCollectionResponseListMeta struct {
	Count *int32 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedCollectionResponseListMeta PaginatedCollectionResponseListMeta

// NewPaginatedCollectionResponseListMeta instantiates a new PaginatedCollectionResponseListMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCollectionResponseListMeta() *PaginatedCollectionResponseListMeta {
	this := PaginatedCollectionResponseListMeta{}
	return &this
}

// NewPaginatedCollectionResponseListMetaWithDefaults instantiates a new PaginatedCollectionResponseListMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCollectionResponseListMetaWithDefaults() *PaginatedCollectionResponseListMeta {
	this := PaginatedCollectionResponseListMeta{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedCollectionResponseListMeta) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCollectionResponseListMeta) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedCollectionResponseListMeta) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedCollectionResponseListMeta) SetCount(v int32) {
	o.Count = &v
}

func (o PaginatedCollectionResponseListMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedCollectionResponseListMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedCollectionResponseListMeta) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedCollectionResponseListMeta := _PaginatedCollectionResponseListMeta{}

	if err = json.Unmarshal(bytes, &varPaginatedCollectionResponseListMeta); err == nil {
		*o = PaginatedCollectionResponseListMeta(varPaginatedCollectionResponseListMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedCollectionResponseListMeta struct {
	value *PaginatedCollectionResponseListMeta
	isSet bool
}

func (v NullablePaginatedCollectionResponseListMeta) Get() *PaginatedCollectionResponseListMeta {
	return v.value
}

func (v *NullablePaginatedCollectionResponseListMeta) Set(val *PaginatedCollectionResponseListMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCollectionResponseListMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCollectionResponseListMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCollectionResponseListMeta(val *PaginatedCollectionResponseListMeta) *NullablePaginatedCollectionResponseListMeta {
	return &NullablePaginatedCollectionResponseListMeta{value: val, isSet: true}
}

func (v NullablePaginatedCollectionResponseListMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCollectionResponseListMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


