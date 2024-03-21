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

// checks if the CustomerBulkUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerBulkUpdateRequest{}

// CustomerBulkUpdateRequest struct for CustomerBulkUpdateRequest
type CustomerBulkUpdateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CustomerIds []string `json:"customerIds,omitempty"`
	Action *BulkUpdateRequestAction `json:"action,omitempty"`
}

// NewCustomerBulkUpdateRequest instantiates a new CustomerBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerBulkUpdateRequest() *CustomerBulkUpdateRequest {
	this := CustomerBulkUpdateRequest{}
	var action BulkUpdateRequestAction = BULKUPDATEREQUESTACTION_ENABLE
	this.Action = &action
	return &this
}

// NewCustomerBulkUpdateRequestWithDefaults instantiates a new CustomerBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerBulkUpdateRequestWithDefaults() *CustomerBulkUpdateRequest {
	this := CustomerBulkUpdateRequest{}
	var action BulkUpdateRequestAction = BULKUPDATEREQUESTACTION_ENABLE
	this.Action = &action
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerBulkUpdateRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerBulkUpdateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerBulkUpdateRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerBulkUpdateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *CustomerBulkUpdateRequest) GetCustomerIds() []string {
	if o == nil || IsNil(o.CustomerIds) {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerBulkUpdateRequest) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerIds) {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *CustomerBulkUpdateRequest) HasCustomerIds() bool {
	if o != nil && !IsNil(o.CustomerIds) {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *CustomerBulkUpdateRequest) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CustomerBulkUpdateRequest) GetAction() BulkUpdateRequestAction {
	if o == nil || IsNil(o.Action) {
		var ret BulkUpdateRequestAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerBulkUpdateRequest) GetActionOk() (*BulkUpdateRequestAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CustomerBulkUpdateRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given BulkUpdateRequestAction and assigns it to the Action field.
func (o *CustomerBulkUpdateRequest) SetAction(v BulkUpdateRequestAction) {
	o.Action = &v
}

func (o CustomerBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerBulkUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CustomerIds) {
		toSerialize["customerIds"] = o.CustomerIds
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableCustomerBulkUpdateRequest struct {
	value *CustomerBulkUpdateRequest
	isSet bool
}

func (v NullableCustomerBulkUpdateRequest) Get() *CustomerBulkUpdateRequest {
	return v.value
}

func (v *NullableCustomerBulkUpdateRequest) Set(val *CustomerBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerBulkUpdateRequest(val *CustomerBulkUpdateRequest) *NullableCustomerBulkUpdateRequest {
	return &NullableCustomerBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableCustomerBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


