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
	"time"
)

// CustomerSubscriberResponse struct for CustomerSubscriberResponse
type CustomerSubscriberResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Country *string `json:"country,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Em *CustomerEMFields `json:"em,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Newsletters []CustomerNewsletterResponse `json:"newsletters,omitempty"`
	Market *string `json:"market,omitempty"`
	PreferredLocale *string `json:"preferredLocale,omitempty"`
	CustomerGroups []string `json:"customerGroups,omitempty"`
}

// NewCustomerSubscriberResponse instantiates a new CustomerSubscriberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriberResponse() *CustomerSubscriberResponse {
	this := CustomerSubscriberResponse{}
	return &this
}

// NewCustomerSubscriberResponseWithDefaults instantiates a new CustomerSubscriberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriberResponseWithDefaults() *CustomerSubscriberResponse {
	this := CustomerSubscriberResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerSubscriberResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerSubscriberResponse) SetName(v string) {
	o.Name = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
    return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *CustomerSubscriberResponse) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerSubscriberResponse) SetEmail(v string) {
	o.Email = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerSubscriberResponse) SetCountry(v string) {
	o.Country = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CustomerSubscriberResponse) SetGender(v string) {
	o.Gender = &v
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerSubscriberResponse) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerSubscriberResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerSubscriberResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetNewsletters() []CustomerNewsletterResponse {
	if o == nil || isNil(o.Newsletters) {
		var ret []CustomerNewsletterResponse
		return ret
	}
	return o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetNewslettersOk() ([]CustomerNewsletterResponse, bool) {
	if o == nil || isNil(o.Newsletters) {
    return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasNewsletters() bool {
	if o != nil && !isNil(o.Newsletters) {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given []CustomerNewsletterResponse and assigns it to the Newsletters field.
func (o *CustomerSubscriberResponse) SetNewsletters(v []CustomerNewsletterResponse) {
	o.Newsletters = v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetMarket() string {
	if o == nil || isNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetMarketOk() (*string, bool) {
	if o == nil || isNil(o.Market) {
    return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasMarket() bool {
	if o != nil && !isNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CustomerSubscriberResponse) SetMarket(v string) {
	o.Market = &v
}

// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetPreferredLocale() string {
	if o == nil || isNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || isNil(o.PreferredLocale) {
    return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasPreferredLocale() bool {
	if o != nil && !isNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *CustomerSubscriberResponse) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}

// GetCustomerGroups returns the CustomerGroups field value if set, zero value otherwise.
func (o *CustomerSubscriberResponse) GetCustomerGroups() []string {
	if o == nil || isNil(o.CustomerGroups) {
		var ret []string
		return ret
	}
	return o.CustomerGroups
}

// GetCustomerGroupsOk returns a tuple with the CustomerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponse) GetCustomerGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.CustomerGroups) {
    return nil, false
	}
	return o.CustomerGroups, true
}

// HasCustomerGroups returns a boolean if a field has been set.
func (o *CustomerSubscriberResponse) HasCustomerGroups() bool {
	if o != nil && !isNil(o.CustomerGroups) {
		return true
	}

	return false
}

// SetCustomerGroups gets a reference to the given []string and assigns it to the CustomerGroups field.
func (o *CustomerSubscriberResponse) SetCustomerGroups(v []string) {
	o.CustomerGroups = v
}

func (o CustomerSubscriberResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Em) {
		toSerialize["em"] = o.Em
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Newsletters) {
		toSerialize["newsletters"] = o.Newsletters
	}
	if !isNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !isNil(o.PreferredLocale) {
		toSerialize["preferredLocale"] = o.PreferredLocale
	}
	if !isNil(o.CustomerGroups) {
		toSerialize["customerGroups"] = o.CustomerGroups
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriberResponse struct {
	value *CustomerSubscriberResponse
	isSet bool
}

func (v NullableCustomerSubscriberResponse) Get() *CustomerSubscriberResponse {
	return v.value
}

func (v *NullableCustomerSubscriberResponse) Set(val *CustomerSubscriberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriberResponse(val *CustomerSubscriberResponse) *NullableCustomerSubscriberResponse {
	return &NullableCustomerSubscriberResponse{value: val, isSet: true}
}

func (v NullableCustomerSubscriberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


