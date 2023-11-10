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

// CustomerSubscriberResponseWithNewsletterRequest struct for CustomerSubscriberResponseWithNewsletterRequest
type CustomerSubscriberResponseWithNewsletterRequest struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Country *string `json:"country,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Em *CustomerEMFields `json:"em,omitempty"`
	PreferredLocale *string `json:"preferredLocale,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Newsletters []CustomerNewsletterRequest `json:"newsletters,omitempty"`
}

// NewCustomerSubscriberResponseWithNewsletterRequest instantiates a new CustomerSubscriberResponseWithNewsletterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriberResponseWithNewsletterRequest() *CustomerSubscriberResponseWithNewsletterRequest {
	this := CustomerSubscriberResponseWithNewsletterRequest{}
	return &this
}

// NewCustomerSubscriberResponseWithNewsletterRequestWithDefaults instantiates a new CustomerSubscriberResponseWithNewsletterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriberResponseWithNewsletterRequestWithDefaults() *CustomerSubscriberResponseWithNewsletterRequest {
	this := CustomerSubscriberResponseWithNewsletterRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetName(v string) {
	o.Name = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
    return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetEmail(v string) {
	o.Email = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetCountry(v string) {
	o.Country = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetGender(v string) {
	o.Gender = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetBirthdate() time.Time {
	if o == nil || isNil(o.Birthdate) {
		var ret time.Time
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetBirthdateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Birthdate) {
    return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasBirthdate() bool {
	if o != nil && !isNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given time.Time and assigns it to the Birthdate field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetBirthdate(v time.Time) {
	o.Birthdate = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetNationality() string {
	if o == nil || isNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetNationalityOk() (*string, bool) {
	if o == nil || isNil(o.Nationality) {
    return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasNationality() bool {
	if o != nil && !isNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetNationality(v string) {
	o.Nationality = &v
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetPreferredLocale() string {
	if o == nil || isNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || isNil(o.PreferredLocale) {
    return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasPreferredLocale() bool {
	if o != nil && !isNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetNewsletters() []CustomerNewsletterRequest {
	if o == nil || isNil(o.Newsletters) {
		var ret []CustomerNewsletterRequest
		return ret
	}
	return o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) GetNewslettersOk() ([]CustomerNewsletterRequest, bool) {
	if o == nil || isNil(o.Newsletters) {
    return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *CustomerSubscriberResponseWithNewsletterRequest) HasNewsletters() bool {
	if o != nil && !isNil(o.Newsletters) {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given []CustomerNewsletterRequest and assigns it to the Newsletters field.
func (o *CustomerSubscriberResponseWithNewsletterRequest) SetNewsletters(v []CustomerNewsletterRequest) {
	o.Newsletters = v
}

func (o CustomerSubscriberResponseWithNewsletterRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !isNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !isNil(o.Em) {
		toSerialize["em"] = o.Em
	}
	if !isNil(o.PreferredLocale) {
		toSerialize["preferredLocale"] = o.PreferredLocale
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
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriberResponseWithNewsletterRequest struct {
	value *CustomerSubscriberResponseWithNewsletterRequest
	isSet bool
}

func (v NullableCustomerSubscriberResponseWithNewsletterRequest) Get() *CustomerSubscriberResponseWithNewsletterRequest {
	return v.value
}

func (v *NullableCustomerSubscriberResponseWithNewsletterRequest) Set(val *CustomerSubscriberResponseWithNewsletterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriberResponseWithNewsletterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriberResponseWithNewsletterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriberResponseWithNewsletterRequest(val *CustomerSubscriberResponseWithNewsletterRequest) *NullableCustomerSubscriberResponseWithNewsletterRequest {
	return &NullableCustomerSubscriberResponseWithNewsletterRequest{value: val, isSet: true}
}

func (v NullableCustomerSubscriberResponseWithNewsletterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriberResponseWithNewsletterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


