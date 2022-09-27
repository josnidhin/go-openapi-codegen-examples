# FixedBenefit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BenefitTypes**](BenefitTypes.md) |  | 
**UnitType** | [**BenefitUnitTypes**](BenefitUnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | [**FixedBenefitAmount**](FixedBenefitAmount.md) |  | 
**AdditionalInformation** | **NullableString** |  | 

## Methods

### NewFixedBenefit

`func NewFixedBenefit(type_ BenefitTypes, unitType BenefitUnitTypes, unit string, amount FixedBenefitAmount, additionalInformation NullableString, ) *FixedBenefit`

NewFixedBenefit instantiates a new FixedBenefit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedBenefitWithDefaults

`func NewFixedBenefitWithDefaults() *FixedBenefit`

NewFixedBenefitWithDefaults instantiates a new FixedBenefit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedBenefit) GetType() BenefitTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedBenefit) GetTypeOk() (*BenefitTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedBenefit) SetType(v BenefitTypes)`

SetType sets Type field to given value.


### GetUnitType

`func (o *FixedBenefit) GetUnitType() BenefitUnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *FixedBenefit) GetUnitTypeOk() (*BenefitUnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *FixedBenefit) SetUnitType(v BenefitUnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *FixedBenefit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FixedBenefit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FixedBenefit) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *FixedBenefit) GetAmount() FixedBenefitAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FixedBenefit) GetAmountOk() (*FixedBenefitAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FixedBenefit) SetAmount(v FixedBenefitAmount)`

SetAmount sets Amount field to given value.


### GetAdditionalInformation

`func (o *FixedBenefit) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *FixedBenefit) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *FixedBenefit) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### SetAdditionalInformationNil

`func (o *FixedBenefit) SetAdditionalInformationNil(b bool)`

 SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil

### UnsetAdditionalInformation
`func (o *FixedBenefit) UnsetAdditionalInformation()`

UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


