# TransactionBenefitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BenefitTypes**](BenefitTypes.md) |  | 
**UnitType** | [**BenefitUnitTypes**](BenefitUnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | [**RangeBenefitAmountWithoutIncrement**](RangeBenefitAmountWithoutIncrement.md) |  | 
**AdditionalInformation** | **NullableString** |  | 

## Methods

### NewTransactionBenefitsInner

`func NewTransactionBenefitsInner(type_ BenefitTypes, unitType BenefitUnitTypes, unit string, amount RangeBenefitAmountWithoutIncrement, additionalInformation NullableString, ) *TransactionBenefitsInner`

NewTransactionBenefitsInner instantiates a new TransactionBenefitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBenefitsInnerWithDefaults

`func NewTransactionBenefitsInnerWithDefaults() *TransactionBenefitsInner`

NewTransactionBenefitsInnerWithDefaults instantiates a new TransactionBenefitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionBenefitsInner) GetType() BenefitTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionBenefitsInner) GetTypeOk() (*BenefitTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionBenefitsInner) SetType(v BenefitTypes)`

SetType sets Type field to given value.


### GetUnitType

`func (o *TransactionBenefitsInner) GetUnitType() BenefitUnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *TransactionBenefitsInner) GetUnitTypeOk() (*BenefitUnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *TransactionBenefitsInner) SetUnitType(v BenefitUnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *TransactionBenefitsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionBenefitsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionBenefitsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *TransactionBenefitsInner) GetAmount() RangeBenefitAmountWithoutIncrement`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionBenefitsInner) GetAmountOk() (*RangeBenefitAmountWithoutIncrement, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionBenefitsInner) SetAmount(v RangeBenefitAmountWithoutIncrement)`

SetAmount sets Amount field to given value.


### GetAdditionalInformation

`func (o *TransactionBenefitsInner) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *TransactionBenefitsInner) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *TransactionBenefitsInner) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### SetAdditionalInformationNil

`func (o *TransactionBenefitsInner) SetAdditionalInformationNil(b bool)`

 SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil

### UnsetAdditionalInformation
`func (o *TransactionBenefitsInner) UnsetAdditionalInformation()`

UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


