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

// checks if the PatchedostreeOstreeRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedostreeOstreeRemote{}

// PatchedostreeOstreeRemote A Serializer class for a remote OSTree repository.
type PatchedostreeOstreeRemote struct {
	// A unique name for this remote.
	Name *string `json:"name,omitempty"`
	// The URL of an external content source.
	Url *string `json:"url,omitempty"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// A PEM encoded private key used for authentication.
	ClientKey NullableString `json:"client_key,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	// The username to authenticte to the proxy.
	ProxyUsername NullableString `json:"proxy_username,omitempty"`
	// The password to authenticate to the proxy. Extra leading and trailing whitespace characters are not trimmed.
	ProxyPassword NullableString `json:"proxy_password,omitempty"`
	// The username to be used for authentication when syncing.
	Username NullableString `json:"username,omitempty"`
	// The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed.
	Password NullableString `json:"password,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt64 `json:"download_concurrency,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt64 `json:"max_retries,omitempty"`
	Policy *Policy462Enum `json:"policy,omitempty"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt64 `json:"rate_limit,omitempty"`
	// An option to specify how many commits to traverse.
	Depth *int64 `json:"depth,omitempty"`
	//              A list of refs to include during a sync.             The wildcards *, ? are recognized.             'include_refs' is evaluated before 'exclude_refs'.             
	IncludeRefs []string `json:"include_refs,omitempty"`
	//              A list of tags to exclude during a sync.             The wildcards *, ? are recognized.             'exclude_refs' is evaluated after 'include_refs'.             
	ExcludeRefs []string `json:"exclude_refs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedostreeOstreeRemote PatchedostreeOstreeRemote

// NewPatchedostreeOstreeRemote instantiates a new PatchedostreeOstreeRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedostreeOstreeRemote() *PatchedostreeOstreeRemote {
	this := PatchedostreeOstreeRemote{}
	var policy Policy462Enum = POLICY462ENUM_IMMEDIATE
	this.Policy = &policy
	var depth int64 = 0
	this.Depth = &depth
	return &this
}

// NewPatchedostreeOstreeRemoteWithDefaults instantiates a new PatchedostreeOstreeRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedostreeOstreeRemoteWithDefaults() *PatchedostreeOstreeRemote {
	this := PatchedostreeOstreeRemote{}
	var policy Policy462Enum = POLICY462ENUM_IMMEDIATE
	this.Policy = &policy
	var depth int64 = 0
	this.Depth = &depth
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedostreeOstreeRemote) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedostreeOstreeRemote) SetUrl(v string) {
	o.Url = &v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *PatchedostreeOstreeRemote) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *PatchedostreeOstreeRemote) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetClientKey() string {
	if o == nil || IsNil(o.ClientKey.Get()) {
		var ret string
		return ret
	}
	return *o.ClientKey.Get()
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientKey.Get(), o.ClientKey.IsSet()
}

// HasClientKey returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasClientKey() bool {
	if o != nil && o.ClientKey.IsSet() {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given NullableString and assigns it to the ClientKey field.
func (o *PatchedostreeOstreeRemote) SetClientKey(v string) {
	o.ClientKey.Set(&v)
}
// SetClientKeyNil sets the value for ClientKey to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetClientKeyNil() {
	o.ClientKey.Set(nil)
}

// UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetClientKey() {
	o.ClientKey.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *PatchedostreeOstreeRemote) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *PatchedostreeOstreeRemote) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetProxyUsername returns the ProxyUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetProxyUsername() string {
	if o == nil || IsNil(o.ProxyUsername.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUsername.Get()
}

// GetProxyUsernameOk returns a tuple with the ProxyUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetProxyUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUsername.Get(), o.ProxyUsername.IsSet()
}

// HasProxyUsername returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasProxyUsername() bool {
	if o != nil && o.ProxyUsername.IsSet() {
		return true
	}

	return false
}

// SetProxyUsername gets a reference to the given NullableString and assigns it to the ProxyUsername field.
func (o *PatchedostreeOstreeRemote) SetProxyUsername(v string) {
	o.ProxyUsername.Set(&v)
}
// SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetProxyUsernameNil() {
	o.ProxyUsername.Set(nil)
}

// UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetProxyUsername() {
	o.ProxyUsername.Unset()
}

// GetProxyPassword returns the ProxyPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetProxyPassword() string {
	if o == nil || IsNil(o.ProxyPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyPassword.Get()
}

// GetProxyPasswordOk returns a tuple with the ProxyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetProxyPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyPassword.Get(), o.ProxyPassword.IsSet()
}

// HasProxyPassword returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasProxyPassword() bool {
	if o != nil && o.ProxyPassword.IsSet() {
		return true
	}

	return false
}

// SetProxyPassword gets a reference to the given NullableString and assigns it to the ProxyPassword field.
func (o *PatchedostreeOstreeRemote) SetProxyPassword(v string) {
	o.ProxyPassword.Set(&v)
}
// SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetProxyPasswordNil() {
	o.ProxyPassword.Set(nil)
}

// UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetProxyPassword() {
	o.ProxyPassword.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *PatchedostreeOstreeRemote) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PatchedostreeOstreeRemote) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetPassword() {
	o.Password.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedostreeOstreeRemote) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetDownloadConcurrency() int64 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int64
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetDownloadConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt64 and assigns it to the DownloadConcurrency field.
func (o *PatchedostreeOstreeRemote) SetDownloadConcurrency(v int64) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetMaxRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt64 and assigns it to the MaxRetries field.
func (o *PatchedostreeOstreeRemote) SetMaxRetries(v int64) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetPolicy() Policy462Enum {
	if o == nil || IsNil(o.Policy) {
		var ret Policy462Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetPolicyOk() (*Policy462Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given Policy462Enum and assigns it to the Policy field.
func (o *PatchedostreeOstreeRemote) SetPolicy(v Policy462Enum) {
	o.Policy = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *PatchedostreeOstreeRemote) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *PatchedostreeOstreeRemote) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *PatchedostreeOstreeRemote) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *PatchedostreeOstreeRemote) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *PatchedostreeOstreeRemote) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetRateLimit() int64 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetRateLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt64 and assigns it to the RateLimit field.
func (o *PatchedostreeOstreeRemote) SetRateLimit(v int64) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *PatchedostreeOstreeRemote) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *PatchedostreeOstreeRemote) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRemote) GetDepth() int64 {
	if o == nil || IsNil(o.Depth) {
		var ret int64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRemote) GetDepthOk() (*int64, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int64 and assigns it to the Depth field.
func (o *PatchedostreeOstreeRemote) SetDepth(v int64) {
	o.Depth = &v
}

// GetIncludeRefs returns the IncludeRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetIncludeRefs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeRefs
}

// GetIncludeRefsOk returns a tuple with the IncludeRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetIncludeRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeRefs) {
		return nil, false
	}
	return o.IncludeRefs, true
}

// HasIncludeRefs returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasIncludeRefs() bool {
	if o != nil && IsNil(o.IncludeRefs) {
		return true
	}

	return false
}

// SetIncludeRefs gets a reference to the given []string and assigns it to the IncludeRefs field.
func (o *PatchedostreeOstreeRemote) SetIncludeRefs(v []string) {
	o.IncludeRefs = v
}

// GetExcludeRefs returns the ExcludeRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRemote) GetExcludeRefs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeRefs
}

// GetExcludeRefsOk returns a tuple with the ExcludeRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRemote) GetExcludeRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeRefs) {
		return nil, false
	}
	return o.ExcludeRefs, true
}

// HasExcludeRefs returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRemote) HasExcludeRefs() bool {
	if o != nil && IsNil(o.ExcludeRefs) {
		return true
	}

	return false
}

// SetExcludeRefs gets a reference to the given []string and assigns it to the ExcludeRefs field.
func (o *PatchedostreeOstreeRemote) SetExcludeRefs(v []string) {
	o.ExcludeRefs = v
}

func (o PatchedostreeOstreeRemote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedostreeOstreeRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.ClientKey.IsSet() {
		toSerialize["client_key"] = o.ClientKey.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if o.ProxyUsername.IsSet() {
		toSerialize["proxy_username"] = o.ProxyUsername.Get()
	}
	if o.ProxyPassword.IsSet() {
		toSerialize["proxy_password"] = o.ProxyPassword.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}
	if o.IncludeRefs != nil {
		toSerialize["include_refs"] = o.IncludeRefs
	}
	if o.ExcludeRefs != nil {
		toSerialize["exclude_refs"] = o.ExcludeRefs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedostreeOstreeRemote) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedostreeOstreeRemote := _PatchedostreeOstreeRemote{}

	if err = json.Unmarshal(bytes, &varPatchedostreeOstreeRemote); err == nil {
		*o = PatchedostreeOstreeRemote(varPatchedostreeOstreeRemote)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "ca_cert")
		delete(additionalProperties, "client_cert")
		delete(additionalProperties, "client_key")
		delete(additionalProperties, "tls_validation")
		delete(additionalProperties, "proxy_url")
		delete(additionalProperties, "proxy_username")
		delete(additionalProperties, "proxy_password")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "download_concurrency")
		delete(additionalProperties, "max_retries")
		delete(additionalProperties, "policy")
		delete(additionalProperties, "total_timeout")
		delete(additionalProperties, "connect_timeout")
		delete(additionalProperties, "sock_connect_timeout")
		delete(additionalProperties, "sock_read_timeout")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "rate_limit")
		delete(additionalProperties, "depth")
		delete(additionalProperties, "include_refs")
		delete(additionalProperties, "exclude_refs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedostreeOstreeRemote struct {
	value *PatchedostreeOstreeRemote
	isSet bool
}

func (v NullablePatchedostreeOstreeRemote) Get() *PatchedostreeOstreeRemote {
	return v.value
}

func (v *NullablePatchedostreeOstreeRemote) Set(val *PatchedostreeOstreeRemote) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedostreeOstreeRemote) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedostreeOstreeRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedostreeOstreeRemote(val *PatchedostreeOstreeRemote) *NullablePatchedostreeOstreeRemote {
	return &NullablePatchedostreeOstreeRemote{value: val, isSet: true}
}

func (v NullablePatchedostreeOstreeRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedostreeOstreeRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


