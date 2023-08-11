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
	"os"
)

// checks if the AnsibleAnsibleNamespaceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleAnsibleNamespaceMetadata{}

// AnsibleAnsibleNamespaceMetadata A serializer for Namespaces.
type AnsibleAnsibleNamespaceMetadata struct {
	// Required named, only accepts lowercase, numbers and underscores.
	Name string `json:"name"`
	// Optional namespace company owner.
	Company *string `json:"company,omitempty"`
	// Optional namespace contact email.
	Email *string `json:"email,omitempty"`
	// Optional short description.
	Description *string `json:"description,omitempty"`
	// Optional resource page in markdown format.
	Resources *string `json:"resources,omitempty"`
	// Labeled related links.
	Links []NamespaceLink `json:"links,omitempty"`
	// Optional avatar image for Namespace
	Avatar **os.File `json:"avatar,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnsibleAnsibleNamespaceMetadata AnsibleAnsibleNamespaceMetadata

// NewAnsibleAnsibleNamespaceMetadata instantiates a new AnsibleAnsibleNamespaceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleAnsibleNamespaceMetadata(name string) *AnsibleAnsibleNamespaceMetadata {
	this := AnsibleAnsibleNamespaceMetadata{}
	this.Name = name
	return &this
}

// NewAnsibleAnsibleNamespaceMetadataWithDefaults instantiates a new AnsibleAnsibleNamespaceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleAnsibleNamespaceMetadataWithDefaults() *AnsibleAnsibleNamespaceMetadata {
	this := AnsibleAnsibleNamespaceMetadata{}
	return &this
}

// GetName returns the Name field value
func (o *AnsibleAnsibleNamespaceMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleAnsibleNamespaceMetadata) SetName(v string) {
	o.Name = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AnsibleAnsibleNamespaceMetadata) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AnsibleAnsibleNamespaceMetadata) SetEmail(v string) {
	o.Email = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnsibleAnsibleNamespaceMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetResources() string {
	if o == nil || IsNil(o.Resources) {
		var ret string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetResourcesOk() (*string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given string and assigns it to the Resources field.
func (o *AnsibleAnsibleNamespaceMetadata) SetResources(v string) {
	o.Resources = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetLinks() []NamespaceLink {
	if o == nil || IsNil(o.Links) {
		var ret []NamespaceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetLinksOk() ([]NamespaceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []NamespaceLink and assigns it to the Links field.
func (o *AnsibleAnsibleNamespaceMetadata) SetLinks(v []NamespaceLink) {
	o.Links = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *AnsibleAnsibleNamespaceMetadata) GetAvatar() *os.File {
	if o == nil || IsNil(o.Avatar) {
		var ret *os.File
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleNamespaceMetadata) GetAvatarOk() (**os.File, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *AnsibleAnsibleNamespaceMetadata) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given *os.File and assigns it to the Avatar field.
func (o *AnsibleAnsibleNamespaceMetadata) SetAvatar(v *os.File) {
	o.Avatar = &v
}

func (o AnsibleAnsibleNamespaceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleAnsibleNamespaceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnsibleAnsibleNamespaceMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varAnsibleAnsibleNamespaceMetadata := _AnsibleAnsibleNamespaceMetadata{}

	if err = json.Unmarshal(bytes, &varAnsibleAnsibleNamespaceMetadata); err == nil {
		*o = AnsibleAnsibleNamespaceMetadata(varAnsibleAnsibleNamespaceMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "company")
		delete(additionalProperties, "email")
		delete(additionalProperties, "description")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "links")
		delete(additionalProperties, "avatar")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnsibleAnsibleNamespaceMetadata struct {
	value *AnsibleAnsibleNamespaceMetadata
	isSet bool
}

func (v NullableAnsibleAnsibleNamespaceMetadata) Get() *AnsibleAnsibleNamespaceMetadata {
	return v.value
}

func (v *NullableAnsibleAnsibleNamespaceMetadata) Set(val *AnsibleAnsibleNamespaceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleAnsibleNamespaceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleAnsibleNamespaceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleAnsibleNamespaceMetadata(val *AnsibleAnsibleNamespaceMetadata) *NullableAnsibleAnsibleNamespaceMetadata {
	return &NullableAnsibleAnsibleNamespaceMetadata{value: val, isSet: true}
}

func (v NullableAnsibleAnsibleNamespaceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleAnsibleNamespaceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


