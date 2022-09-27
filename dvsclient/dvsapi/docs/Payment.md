# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modes** | [**[]PaymentModes**](PaymentModes.md) |  | 
**PostingPeriod** | [**NullablePaymentPostingPeriod**](PaymentPostingPeriod.md) |  | 
**SupportsStatementInquiry** | **bool** |  | 
**RequiredInquiryFields** | Pointer to **[][]string** |  | [optional] 
**Notes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPayment

`func NewPayment(modes []PaymentModes, postingPeriod NullablePaymentPostingPeriod, supportsStatementInquiry bool, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModes

`func (o *Payment) GetModes() []PaymentModes`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *Payment) GetModesOk() (*[]PaymentModes, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *Payment) SetModes(v []PaymentModes)`

SetModes sets Modes field to given value.


### SetModesNil

`func (o *Payment) SetModesNil(b bool)`

 SetModesNil sets the value for Modes to be an explicit nil

### UnsetModes
`func (o *Payment) UnsetModes()`

UnsetModes ensures that no value is present for Modes, not even an explicit nil
### GetPostingPeriod

`func (o *Payment) GetPostingPeriod() PaymentPostingPeriod`

GetPostingPeriod returns the PostingPeriod field if non-nil, zero value otherwise.

### GetPostingPeriodOk

`func (o *Payment) GetPostingPeriodOk() (*PaymentPostingPeriod, bool)`

GetPostingPeriodOk returns a tuple with the PostingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingPeriod

`func (o *Payment) SetPostingPeriod(v PaymentPostingPeriod)`

SetPostingPeriod sets PostingPeriod field to given value.


### SetPostingPeriodNil

`func (o *Payment) SetPostingPeriodNil(b bool)`

 SetPostingPeriodNil sets the value for PostingPeriod to be an explicit nil

### UnsetPostingPeriod
`func (o *Payment) UnsetPostingPeriod()`

UnsetPostingPeriod ensures that no value is present for PostingPeriod, not even an explicit nil
### GetSupportsStatementInquiry

`func (o *Payment) GetSupportsStatementInquiry() bool`

GetSupportsStatementInquiry returns the SupportsStatementInquiry field if non-nil, zero value otherwise.

### GetSupportsStatementInquiryOk

`func (o *Payment) GetSupportsStatementInquiryOk() (*bool, bool)`

GetSupportsStatementInquiryOk returns a tuple with the SupportsStatementInquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStatementInquiry

`func (o *Payment) SetSupportsStatementInquiry(v bool)`

SetSupportsStatementInquiry sets SupportsStatementInquiry field to given value.


### GetRequiredInquiryFields

`func (o *Payment) GetRequiredInquiryFields() [][]string`

GetRequiredInquiryFields returns the RequiredInquiryFields field if non-nil, zero value otherwise.

### GetRequiredInquiryFieldsOk

`func (o *Payment) GetRequiredInquiryFieldsOk() (*[][]string, bool)`

GetRequiredInquiryFieldsOk returns a tuple with the RequiredInquiryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInquiryFields

`func (o *Payment) SetRequiredInquiryFields(v [][]string)`

SetRequiredInquiryFields sets RequiredInquiryFields field to given value.

### HasRequiredInquiryFields

`func (o *Payment) HasRequiredInquiryFields() bool`

HasRequiredInquiryFields returns a boolean if a field has been set.

### SetRequiredInquiryFieldsNil

`func (o *Payment) SetRequiredInquiryFieldsNil(b bool)`

 SetRequiredInquiryFieldsNil sets the value for RequiredInquiryFields to be an explicit nil

### UnsetRequiredInquiryFields
`func (o *Payment) UnsetRequiredInquiryFields()`

UnsetRequiredInquiryFields ensures that no value is present for RequiredInquiryFields, not even an explicit nil
### GetNotes

`func (o *Payment) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Payment) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Payment) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Payment) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Payment) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Payment) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


