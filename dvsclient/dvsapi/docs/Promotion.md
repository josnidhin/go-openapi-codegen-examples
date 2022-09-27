# Promotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Terms** | **NullableString** |  | 
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 
**Operator** | [**PromotionOperator**](PromotionOperator.md) |  | 
**Products** | [**[]PromotionProductsInner**](PromotionProductsInner.md) |  | 

## Methods

### NewPromotion

`func NewPromotion(id int32, title string, description string, terms NullableString, startDate time.Time, endDate time.Time, operator PromotionOperator, products []PromotionProductsInner, ) *Promotion`

NewPromotion instantiates a new Promotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionWithDefaults

`func NewPromotionWithDefaults() *Promotion`

NewPromotionWithDefaults instantiates a new Promotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Promotion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Promotion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Promotion) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *Promotion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Promotion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Promotion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Promotion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Promotion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Promotion) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTerms

`func (o *Promotion) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Promotion) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Promotion) SetTerms(v string)`

SetTerms sets Terms field to given value.


### SetTermsNil

`func (o *Promotion) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *Promotion) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetStartDate

`func (o *Promotion) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Promotion) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Promotion) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *Promotion) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Promotion) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Promotion) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetOperator

`func (o *Promotion) GetOperator() PromotionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Promotion) GetOperatorOk() (*PromotionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Promotion) SetOperator(v PromotionOperator)`

SetOperator sets Operator field to given value.


### GetProducts

`func (o *Promotion) GetProducts() []PromotionProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Promotion) GetProductsOk() (*[]PromotionProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Promotion) SetProducts(v []PromotionProductsInner)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


