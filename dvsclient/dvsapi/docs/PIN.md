# PIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageInfo** | **[]string** |  | 
**Validity** | [**NullablePINValidity**](PINValidity.md) |  | 
**Terms** | **NullableString** |  | 

## Methods

### NewPIN

`func NewPIN(usageInfo []string, validity NullablePINValidity, terms NullableString, ) *PIN`

NewPIN instantiates a new PIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPINWithDefaults

`func NewPINWithDefaults() *PIN`

NewPINWithDefaults instantiates a new PIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageInfo

`func (o *PIN) GetUsageInfo() []string`

GetUsageInfo returns the UsageInfo field if non-nil, zero value otherwise.

### GetUsageInfoOk

`func (o *PIN) GetUsageInfoOk() (*[]string, bool)`

GetUsageInfoOk returns a tuple with the UsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInfo

`func (o *PIN) SetUsageInfo(v []string)`

SetUsageInfo sets UsageInfo field to given value.


### SetUsageInfoNil

`func (o *PIN) SetUsageInfoNil(b bool)`

 SetUsageInfoNil sets the value for UsageInfo to be an explicit nil

### UnsetUsageInfo
`func (o *PIN) UnsetUsageInfo()`

UnsetUsageInfo ensures that no value is present for UsageInfo, not even an explicit nil
### GetValidity

`func (o *PIN) GetValidity() PINValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *PIN) GetValidityOk() (*PINValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *PIN) SetValidity(v PINValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *PIN) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *PIN) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetTerms

`func (o *PIN) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PIN) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PIN) SetTerms(v string)`

SetTerms sets Terms field to given value.


### SetTermsNil

`func (o *PIN) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *PIN) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


