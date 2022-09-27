# RangedBenefit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BenefitTypes**](BenefitTypes.md) |  | 
**UnitType** | [**BenefitUnitTypes**](BenefitUnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | [**RangeBenefitAmountWithoutIncrement**](RangeBenefitAmountWithoutIncrement.md) |  | 
**AdditionalInformation** | **NullableString** |  | 

## Methods

### NewRangedBenefit

`func NewRangedBenefit(type_ BenefitTypes, unitType BenefitUnitTypes, unit string, amount RangeBenefitAmountWithoutIncrement, additionalInformation NullableString, ) *RangedBenefit`

NewRangedBenefit instantiates a new RangedBenefit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangedBenefitWithDefaults

`func NewRangedBenefitWithDefaults() *RangedBenefit`

NewRangedBenefitWithDefaults instantiates a new RangedBenefit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RangedBenefit) GetType() BenefitTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RangedBenefit) GetTypeOk() (*BenefitTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RangedBenefit) SetType(v BenefitTypes)`

SetType sets Type field to given value.


### GetUnitType

`func (o *RangedBenefit) GetUnitType() BenefitUnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *RangedBenefit) GetUnitTypeOk() (*BenefitUnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *RangedBenefit) SetUnitType(v BenefitUnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *RangedBenefit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RangedBenefit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RangedBenefit) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *RangedBenefit) GetAmount() RangeBenefitAmountWithoutIncrement`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RangedBenefit) GetAmountOk() (*RangeBenefitAmountWithoutIncrement, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RangedBenefit) SetAmount(v RangeBenefitAmountWithoutIncrement)`

SetAmount sets Amount field to given value.


### GetAdditionalInformation

`func (o *RangedBenefit) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *RangedBenefit) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *RangedBenefit) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### SetAdditionalInformationNil

`func (o *RangedBenefit) SetAdditionalInformationNil(b bool)`

 SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil

### UnsetAdditionalInformation
`func (o *RangedBenefit) UnsetAdditionalInformation()`

UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


