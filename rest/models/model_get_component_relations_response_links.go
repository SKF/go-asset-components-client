/*
Asset Components

An Asset Component is an internal part of an Asset which is defined by the Hierarchy API.  Currently the following components are supported, * Bearings * Shafts

API version: 0.1
Contact: team.gob@enlight.skf.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// GetComponentRelationsResponseLinks struct for GetComponentRelationsResponseLinks
type GetComponentRelationsResponseLinks struct {
	Self string `json:"self"`
	Next *string `json:"next,omitempty"`
}

// NewGetComponentRelationsResponseLinks instantiates a new GetComponentRelationsResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetComponentRelationsResponseLinks(self string) *GetComponentRelationsResponseLinks {
	this := GetComponentRelationsResponseLinks{}
	this.Self = self
	return &this
}

// NewGetComponentRelationsResponseLinksWithDefaults instantiates a new GetComponentRelationsResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetComponentRelationsResponseLinksWithDefaults() *GetComponentRelationsResponseLinks {
	this := GetComponentRelationsResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetComponentRelationsResponseLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetComponentRelationsResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetComponentRelationsResponseLinks) SetSelf(v string) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *GetComponentRelationsResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetComponentRelationsResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *GetComponentRelationsResponseLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *GetComponentRelationsResponseLinks) SetNext(v string) {
	o.Next = &v
}

func (o GetComponentRelationsResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullableGetComponentRelationsResponseLinks struct {
	value *GetComponentRelationsResponseLinks
	isSet bool
}

func (v NullableGetComponentRelationsResponseLinks) Get() *GetComponentRelationsResponseLinks {
	return v.value
}

func (v *NullableGetComponentRelationsResponseLinks) Set(val *GetComponentRelationsResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetComponentRelationsResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetComponentRelationsResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetComponentRelationsResponseLinks(val *GetComponentRelationsResponseLinks) *NullableGetComponentRelationsResponseLinks {
	return &NullableGetComponentRelationsResponseLinks{value: val, isSet: true}
}

func (v NullableGetComponentRelationsResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetComponentRelationsResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

