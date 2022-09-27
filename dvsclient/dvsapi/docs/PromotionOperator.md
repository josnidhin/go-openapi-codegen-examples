# PromotionOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Country** | [**PromotionOperatorCountry**](PromotionOperatorCountry.md) |  | 

## Methods

### NewPromotionOperator

`func NewPromotionOperator(id int32, name string, country PromotionOperatorCountry, ) *PromotionOperator`

NewPromotionOperator instantiates a new PromotionOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionOperatorWithDefaults

`func NewPromotionOperatorWithDefaults() *PromotionOperator`

NewPromotionOperatorWithDefaults instantiates a new PromotionOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionOperator) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionOperator) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionOperator) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PromotionOperator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromotionOperator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromotionOperator) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *PromotionOperator) GetCountry() PromotionOperatorCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PromotionOperator) GetCountryOk() (*PromotionOperatorCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PromotionOperator) SetCountry(v PromotionOperatorCountry)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


