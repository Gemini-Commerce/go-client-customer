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

// checks if the CustomerConsent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerConsent{}

// CustomerConsent struct for CustomerConsent
type CustomerConsent struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Preferences *map[string]bool `json:"preferences,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Source *CustomerConsentSource `json:"source,omitempty"`
	Author *string `json:"author,omitempty"`
	SubjectId *string `json:"subjectId,omitempty"`
}

// NewCustomerConsent instantiates a new CustomerConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerConsent() *CustomerConsent {
	this := CustomerConsent{}
	var source CustomerConsentSource = UNKNOWN
	this.Source = &source
	return &this
}

// NewCustomerConsentWithDefaults instantiates a new CustomerConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerConsentWithDefaults() *CustomerConsent {
	this := CustomerConsent{}
	var source CustomerConsentSource = UNKNOWN
	this.Source = &source
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerConsent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerConsent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerConsent) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *CustomerConsent) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *CustomerConsent) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *CustomerConsent) SetGrn(v string) {
	o.Grn = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *CustomerConsent) GetPreferences() map[string]bool {
	if o == nil || IsNil(o.Preferences) {
		var ret map[string]bool
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetPreferencesOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Preferences) {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *CustomerConsent) HasPreferences() bool {
	if o != nil && !IsNil(o.Preferences) {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given map[string]bool and assigns it to the Preferences field.
func (o *CustomerConsent) SetPreferences(v map[string]bool) {
	o.Preferences = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerConsent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerConsent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerConsent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CustomerConsent) GetSource() CustomerConsentSource {
	if o == nil || IsNil(o.Source) {
		var ret CustomerConsentSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetSourceOk() (*CustomerConsentSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CustomerConsent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given CustomerConsentSource and assigns it to the Source field.
func (o *CustomerConsent) SetSource(v CustomerConsentSource) {
	o.Source = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CustomerConsent) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CustomerConsent) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *CustomerConsent) SetAuthor(v string) {
	o.Author = &v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *CustomerConsent) GetSubjectId() string {
	if o == nil || IsNil(o.SubjectId) {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerConsent) GetSubjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectId) {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *CustomerConsent) HasSubjectId() bool {
	if o != nil && !IsNil(o.SubjectId) {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *CustomerConsent) SetSubjectId(v string) {
	o.SubjectId = &v
}

func (o CustomerConsent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerConsent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Preferences) {
		toSerialize["preferences"] = o.Preferences
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.SubjectId) {
		toSerialize["subjectId"] = o.SubjectId
	}
	return toSerialize, nil
}

type NullableCustomerConsent struct {
	value *CustomerConsent
	isSet bool
}

func (v NullableCustomerConsent) Get() *CustomerConsent {
	return v.value
}

func (v *NullableCustomerConsent) Set(val *CustomerConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerConsent(val *CustomerConsent) *NullableCustomerConsent {
	return &NullableCustomerConsent{value: val, isSet: true}
}

func (v NullableCustomerConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


