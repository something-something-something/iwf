/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the TriggerContinueAsNewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerContinueAsNewRequest{}

// TriggerContinueAsNewRequest struct for TriggerContinueAsNewRequest
type TriggerContinueAsNewRequest struct {
	WorkflowId    string  `json:"workflowId"`
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
}

// NewTriggerContinueAsNewRequest instantiates a new TriggerContinueAsNewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerContinueAsNewRequest(workflowId string) *TriggerContinueAsNewRequest {
	this := TriggerContinueAsNewRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewTriggerContinueAsNewRequestWithDefaults instantiates a new TriggerContinueAsNewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerContinueAsNewRequestWithDefaults() *TriggerContinueAsNewRequest {
	this := TriggerContinueAsNewRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *TriggerContinueAsNewRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *TriggerContinueAsNewRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *TriggerContinueAsNewRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *TriggerContinueAsNewRequest) GetWorkflowRunId() string {
	if o == nil || IsNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerContinueAsNewRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *TriggerContinueAsNewRequest) HasWorkflowRunId() bool {
	if o != nil && !IsNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *TriggerContinueAsNewRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

func (o TriggerContinueAsNewRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerContinueAsNewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	if !IsNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	return toSerialize, nil
}

type NullableTriggerContinueAsNewRequest struct {
	value *TriggerContinueAsNewRequest
	isSet bool
}

func (v NullableTriggerContinueAsNewRequest) Get() *TriggerContinueAsNewRequest {
	return v.value
}

func (v *NullableTriggerContinueAsNewRequest) Set(val *TriggerContinueAsNewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerContinueAsNewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerContinueAsNewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerContinueAsNewRequest(val *TriggerContinueAsNewRequest) *NullableTriggerContinueAsNewRequest {
	return &NullableTriggerContinueAsNewRequest{value: val, isSet: true}
}

func (v NullableTriggerContinueAsNewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerContinueAsNewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
