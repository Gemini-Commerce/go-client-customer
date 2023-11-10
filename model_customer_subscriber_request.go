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

// CustomerSubscriberRequest struct for CustomerSubscriberRequest
type CustomerSubscriberRequest struct {
	Name *string `json:"name,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Country *string `json:"country,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Em *CustomerEMFields `json:"em,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Newsletters []CustomerNewsletterRequest `json:"newsletters,omitempty"`
	Market *string `json:"market,omitempty"`
	PreferredLocale *string `json:"preferredLocale,omitempty"`
}

// NewCustomerSubscriberRequest instantiates a new CustomerSubscriberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriberRequest() *CustomerSubscriberRequest {
	this := CustomerSubscriberRequest{}
	return &this
}

// NewCustomerSubscriberRequestWithDefaults instantiates a new CustomerSubscriberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriberRequestWithDefaults() *CustomerSubscriberRequest {
	this := CustomerSubscriberRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerSubscriberRequest) SetName(v string) {
	o.Name = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
    return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *CustomerSubscriberRequest) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerSubscriberRequest) SetEmail(v string) {
	o.Email = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerSubscriberRequest) SetCountry(v string) {
	o.Country = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CustomerSubscriberRequest) SetGender(v string) {
	o.Gender = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetBirthdate() time.Time {
	if o == nil || isNil(o.Birthdate) {
		var ret time.Time
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetBirthdateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Birthdate) {
    return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasBirthdate() bool {
	if o != nil && !isNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given time.Time and assigns it to the Birthdate field.
func (o *CustomerSubscriberRequest) SetBirthdate(v time.Time) {
	o.Birthdate = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetNationality() string {
	if o == nil || isNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetNationalityOk() (*string, bool) {
	if o == nil || isNil(o.Nationality) {
    return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasNationality() bool {
	if o != nil && !isNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *CustomerSubscriberRequest) SetNationality(v string) {
	o.Nationality = &v
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerSubscriberRequest) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerSubscriberRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerSubscriberRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetNewsletters() []CustomerNewsletterRequest {
	if o == nil || isNil(o.Newsletters) {
		var ret []CustomerNewsletterRequest
		return ret
	}
	return o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetNewslettersOk() ([]CustomerNewsletterRequest, bool) {
	if o == nil || isNil(o.Newsletters) {
    return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasNewsletters() bool {
	if o != nil && !isNil(o.Newsletters) {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given []CustomerNewsletterRequest and assigns it to the Newsletters field.
func (o *CustomerSubscriberRequest) SetNewsletters(v []CustomerNewsletterRequest) {
	o.Newsletters = v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetMarket() string {
	if o == nil || isNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetMarketOk() (*string, bool) {
	if o == nil || isNil(o.Market) {
    return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasMarket() bool {
	if o != nil && !isNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CustomerSubscriberRequest) SetMarket(v string) {
	o.Market = &v
}

// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *CustomerSubscriberRequest) GetPreferredLocale() string {
	if o == nil || isNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriberRequest) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || isNil(o.PreferredLocale) {
    return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *CustomerSubscriberRequest) HasPreferredLocale() bool {
	if o != nil && !isNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *CustomerSubscriberRequest) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}

func (o CustomerSubscriberRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriberRequest struct {
	value *CustomerSubscriberRequest
	isSet bool
}

func (v NullableCustomerSubscriberRequest) Get() *CustomerSubscriberRequest {
	return v.value
}

func (v *NullableCustomerSubscriberRequest) Set(val *CustomerSubscriberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriberRequest(val *CustomerSubscriberRequest) *NullableCustomerSubscriberRequest {
	return &NullableCustomerSubscriberRequest{value: val, isSet: true}
}

func (v NullableCustomerSubscriberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


