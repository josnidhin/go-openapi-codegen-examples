# FixedBenefitAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **float32** |  | 
**PromotionBonus** | **float32** |  | [default to 0]
**TotalExcludingTax** | **float32** |  | 
**TotalIncludingTax** | **float32** |  | 

## Methods

### NewFixedBenefitAmount

`func NewFixedBenefitAmount(base float32, promotionBonus float32, totalExcludingTax float32, totalIncludingTax float32, ) *FixedBenefitAmount`

NewFixedBenefitAmount instantiates a new FixedBenefitAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedBenefitAmountWithDefaults

`func NewFixedBenefitAmountWithDefaults() *FixedBenefitAmount`

NewFixedBenefitAmountWithDefaults instantiates a new FixedBenefitAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *FixedBenefitAmount) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FixedBenefitAmount) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FixedBenefitAmount) SetBase(v float32)`

SetBase sets Base field to given value.


### GetPromotionBonus

`func (o *FixedBenefitAmount) GetPromotionBonus() float32`

GetPromotionBonus returns the PromotionBonus field if non-nil, zero value otherwise.

### GetPromotionBonusOk

`func (o *FixedBenefitAmount) GetPromotionBonusOk() (*float32, bool)`

GetPromotionBonusOk returns a tuple with the PromotionBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionBonus

`func (o *FixedBenefitAmount) SetPromotionBonus(v float32)`

SetPromotionBonus sets PromotionBonus field to given value.


### GetTotalExcludingTax

`func (o *FixedBenefitAmount) GetTotalExcludingTax() float32`

GetTotalExcludingTax returns the TotalExcludingTax field if non-nil, zero value otherwise.

### GetTotalExcludingTaxOk

`func (o *FixedBenefitAmount) GetTotalExcludingTaxOk() (*float32, bool)`

GetTotalExcludingTaxOk returns a tuple with the TotalExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExcludingTax

`func (o *FixedBenefitAmount) SetTotalExcludingTax(v float32)`

SetTotalExcludingTax sets TotalExcludingTax field to given value.


### GetTotalIncludingTax

`func (o *FixedBenefitAmount) GetTotalIncludingTax() float32`

GetTotalIncludingTax returns the TotalIncludingTax field if non-nil, zero value otherwise.

### GetTotalIncludingTaxOk

`func (o *FixedBenefitAmount) GetTotalIncludingTaxOk() (*float32, bool)`

GetTotalIncludingTaxOk returns a tuple with the TotalIncludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncludingTax

`func (o *FixedBenefitAmount) SetTotalIncludingTax(v float32)`

SetTotalIncludingTax sets TotalIncludingTax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


