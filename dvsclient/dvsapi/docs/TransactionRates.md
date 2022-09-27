# TransactionRates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **float32** |  | 
**Wholesale** | **float32** |  | 
**Retail** | **NullableFloat32** |  | 

## Methods

### NewTransactionRates

`func NewTransactionRates(base float32, wholesale float32, retail NullableFloat32, ) *TransactionRates`

NewTransactionRates instantiates a new TransactionRates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRatesWithDefaults

`func NewTransactionRatesWithDefaults() *TransactionRates`

NewTransactionRatesWithDefaults instantiates a new TransactionRates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *TransactionRates) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *TransactionRates) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *TransactionRates) SetBase(v float32)`

SetBase sets Base field to given value.


### GetWholesale

`func (o *TransactionRates) GetWholesale() float32`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *TransactionRates) GetWholesaleOk() (*float32, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *TransactionRates) SetWholesale(v float32)`

SetWholesale sets Wholesale field to given value.


### GetRetail

`func (o *TransactionRates) GetRetail() float32`

GetRetail returns the Retail field if non-nil, zero value otherwise.

### GetRetailOk

`func (o *TransactionRates) GetRetailOk() (*float32, bool)`

GetRetailOk returns a tuple with the Retail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetail

`func (o *TransactionRates) SetRetail(v float32)`

SetRetail sets Retail field to given value.


### SetRetailNil

`func (o *TransactionRates) SetRetailNil(b bool)`

 SetRetailNil sets the value for Retail to be an explicit nil

### UnsetRetail
`func (o *TransactionRates) UnsetRetail()`

UnsetRetail ensures that no value is present for Retail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


