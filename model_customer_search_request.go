/*
CDP Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer

import (
	"encoding/json"
)

// checks if the CustomerSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSearchRequest{}

// CustomerSearchRequest struct for CustomerSearchRequest
type CustomerSearchRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Query *string `json:"query,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	Filter *CustomerSearchRequestFilter `json:"filter,omitempty"`
	FilterMask *string `json:"filterMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerSearchRequest CustomerSearchRequest

// NewCustomerSearchRequest instantiates a new CustomerSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSearchRequest() *CustomerSearchRequest {
	this := CustomerSearchRequest{}
	return &this
}

// NewCustomerSearchRequestWithDefaults instantiates a new CustomerSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSearchRequestWithDefaults() *CustomerSearchRequest {
	this := CustomerSearchRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerSearchRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CustomerSearchRequest) SetQuery(v string) {
	o.Query = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomerSearchRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *CustomerSearchRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *CustomerSearchRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetFilter() CustomerSearchRequestFilter {
	if o == nil || IsNil(o.Filter) {
		var ret CustomerSearchRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetFilterOk() (*CustomerSearchRequestFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given CustomerSearchRequestFilter and assigns it to the Filter field.
func (o *CustomerSearchRequest) SetFilter(v CustomerSearchRequestFilter) {
	o.Filter = &v
}

// GetFilterMask returns the FilterMask field value if set, zero value otherwise.
func (o *CustomerSearchRequest) GetFilterMask() string {
	if o == nil || IsNil(o.FilterMask) {
		var ret string
		return ret
	}
	return *o.FilterMask
}

// GetFilterMaskOk returns a tuple with the FilterMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequest) GetFilterMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FilterMask) {
		return nil, false
	}
	return o.FilterMask, true
}

// HasFilterMask returns a boolean if a field has been set.
func (o *CustomerSearchRequest) HasFilterMask() bool {
	if o != nil && !IsNil(o.FilterMask) {
		return true
	}

	return false
}

// SetFilterMask gets a reference to the given string and assigns it to the FilterMask field.
func (o *CustomerSearchRequest) SetFilterMask(v string) {
	o.FilterMask = &v
}

func (o CustomerSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.FilterMask) {
		toSerialize["filterMask"] = o.FilterMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerSearchRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomerSearchRequest := _CustomerSearchRequest{}

	err = json.Unmarshal(data, &varCustomerSearchRequest)

	if err != nil {
		return err
	}

	*o = CustomerSearchRequest(varCustomerSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "query")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "filterMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerSearchRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *CustomerSearchRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableCustomerSearchRequest struct {
	value *CustomerSearchRequest
	isSet bool
}

func (v NullableCustomerSearchRequest) Get() *CustomerSearchRequest {
	return v.value
}

func (v *NullableCustomerSearchRequest) Set(val *CustomerSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSearchRequest(val *CustomerSearchRequest) *NullableCustomerSearchRequest {
	return &NullableCustomerSearchRequest{value: val, isSet: true}
}

func (v NullableCustomerSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


