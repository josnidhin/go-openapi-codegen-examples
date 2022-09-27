# ProductPromotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Terms** | **NullableString** |  | 
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 

## Methods

### NewProductPromotion

`func NewProductPromotion(id int32, title string, description string, terms NullableString, startDate time.Time, endDate time.Time, ) *ProductPromotion`

NewProductPromotion instantiates a new ProductPromotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPromotionWithDefaults

`func NewProductPromotionWithDefaults() *ProductPromotion`

NewProductPromotionWithDefaults instantiates a new ProductPromotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductPromotion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductPromotion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductPromotion) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *ProductPromotion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductPromotion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductPromotion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ProductPromotion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductPromotion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductPromotion) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTerms

`func (o *ProductPromotion) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *ProductPromotion) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *ProductPromotion) SetTerms(v string)`

SetTerms sets Terms field to given value.


### SetTermsNil

`func (o *ProductPromotion) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *ProductPromotion) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetStartDate

`func (o *ProductPromotion) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProductPromotion) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProductPromotion) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *ProductPromotion) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProductPromotion) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProductPromotion) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


