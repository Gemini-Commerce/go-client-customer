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
	"time"
)

// checks if the CustomerGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGroupResponse{}

// CustomerGroupResponse struct for CustomerGroupResponse
type CustomerGroupResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CustomerIds []string `json:"customerIds,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Code *string `json:"code,omitempty"`
	CustomerCount *string `json:"customerCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerGroupResponse CustomerGroupResponse

// NewCustomerGroupResponse instantiates a new CustomerGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupResponse() *CustomerGroupResponse {
	this := CustomerGroupResponse{}
	return &this
}

// NewCustomerGroupResponseWithDefaults instantiates a new CustomerGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupResponseWithDefaults() *CustomerGroupResponse {
	this := CustomerGroupResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerGroupResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerGroupResponse) SetName(v string) {
	o.Name = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetCustomerIds() []string {
	if o == nil || IsNil(o.CustomerIds) {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerIds) {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasCustomerIds() bool {
	if o != nil && !IsNil(o.CustomerIds) {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *CustomerGroupResponse) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerGroupResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerGroupResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *CustomerGroupResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomerGroupResponse) SetCode(v string) {
	o.Code = &v
}

// GetCustomerCount returns the CustomerCount field value if set, zero value otherwise.
func (o *CustomerGroupResponse) GetCustomerCount() string {
	if o == nil || IsNil(o.CustomerCount) {
		var ret string
		return ret
	}
	return *o.CustomerCount
}

// GetCustomerCountOk returns a tuple with the CustomerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponse) GetCustomerCountOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCount) {
		return nil, false
	}
	return o.CustomerCount, true
}

// HasCustomerCount returns a boolean if a field has been set.
func (o *CustomerGroupResponse) HasCustomerCount() bool {
	if o != nil && !IsNil(o.CustomerCount) {
		return true
	}

	return false
}

// SetCustomerCount gets a reference to the given string and assigns it to the CustomerCount field.
func (o *CustomerGroupResponse) SetCustomerCount(v string) {
	o.CustomerCount = &v
}

func (o CustomerGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CustomerIds) {
		toSerialize["customerIds"] = o.CustomerIds
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CustomerCount) {
		toSerialize["customerCount"] = o.CustomerCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerGroupResponse) UnmarshalJSON(data []byte) (err error) {
	varCustomerGroupResponse := _CustomerGroupResponse{}

	err = json.Unmarshal(data, &varCustomerGroupResponse)

	if err != nil {
		return err
	}

	*o = CustomerGroupResponse(varCustomerGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "customerIds")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "code")
		delete(additionalProperties, "customerCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerGroupResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *CustomerGroupResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableCustomerGroupResponse struct {
	value *CustomerGroupResponse
	isSet bool
}

func (v NullableCustomerGroupResponse) Get() *CustomerGroupResponse {
	return v.value
}

func (v *NullableCustomerGroupResponse) Set(val *CustomerGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupResponse(val *CustomerGroupResponse) *NullableCustomerGroupResponse {
	return &NullableCustomerGroupResponse{value: val, isSet: true}
}

func (v NullableCustomerGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


