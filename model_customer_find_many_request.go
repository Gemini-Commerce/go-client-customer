/*
CDP Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomerFindManyRequest struct for CustomerFindManyRequest
type CustomerFindManyRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	Filter *CustomerFindManyRequestFilter `json:"filter,omitempty"`
	FilterMask []string `json:"filterMask,omitempty"`
}

// NewCustomerFindManyRequest instantiates a new CustomerFindManyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFindManyRequest() *CustomerFindManyRequest {
	this := CustomerFindManyRequest{}
	return &this
}

// NewCustomerFindManyRequestWithDefaults instantiates a new CustomerFindManyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFindManyRequestWithDefaults() *CustomerFindManyRequest {
	this := CustomerFindManyRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerFindManyRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomerFindManyRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *CustomerFindManyRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *CustomerFindManyRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetFilter() CustomerFindManyRequestFilter {
	if o == nil || isNil(o.Filter) {
		var ret CustomerFindManyRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetFilterOk() (*CustomerFindManyRequestFilter, bool) {
	if o == nil || isNil(o.Filter) {
    return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given CustomerFindManyRequestFilter and assigns it to the Filter field.
func (o *CustomerFindManyRequest) SetFilter(v CustomerFindManyRequestFilter) {
	o.Filter = &v
}

// GetFilterMask returns the FilterMask field value if set, zero value otherwise.
func (o *CustomerFindManyRequest) GetFilterMask() []string {
	if o == nil || isNil(o.FilterMask) {
		var ret []string
		return ret
	}
	return o.FilterMask
}

// GetFilterMaskOk returns a tuple with the FilterMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindManyRequest) GetFilterMaskOk() ([]string, bool) {
	if o == nil || isNil(o.FilterMask) {
    return nil, false
	}
	return o.FilterMask, true
}

// HasFilterMask returns a boolean if a field has been set.
func (o *CustomerFindManyRequest) HasFilterMask() bool {
	if o != nil && !isNil(o.FilterMask) {
		return true
	}

	return false
}

// SetFilterMask gets a reference to the given []string and assigns it to the FilterMask field.
func (o *CustomerFindManyRequest) SetFilterMask(v []string) {
	o.FilterMask = v
}

func (o CustomerFindManyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.FilterMask) {
		toSerialize["filterMask"] = o.FilterMask
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFindManyRequest struct {
	value *CustomerFindManyRequest
	isSet bool
}

func (v NullableCustomerFindManyRequest) Get() *CustomerFindManyRequest {
	return v.value
}

func (v *NullableCustomerFindManyRequest) Set(val *CustomerFindManyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFindManyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFindManyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFindManyRequest(val *CustomerFindManyRequest) *NullableCustomerFindManyRequest {
	return &NullableCustomerFindManyRequest{value: val, isSet: true}
}

func (v NullableCustomerFindManyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFindManyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


