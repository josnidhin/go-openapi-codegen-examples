# TransactionPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wholesale** | [**TransactionPricesWholesale**](TransactionPricesWholesale.md) |  | 
**Retail** | [**NullableTransactionPricesRetail**](TransactionPricesRetail.md) |  | 

## Methods

### NewTransactionPrices

`func NewTransactionPrices(wholesale TransactionPricesWholesale, retail NullableTransactionPricesRetail, ) *TransactionPrices`

NewTransactionPrices instantiates a new TransactionPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPricesWithDefaults

`func NewTransactionPricesWithDefaults() *TransactionPrices`

NewTransactionPricesWithDefaults instantiates a new TransactionPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWholesale

`func (o *TransactionPrices) GetWholesale() TransactionPricesWholesale`

GetWholesale returns the Wholesale field if non-nil, zero value otherwise.

### GetWholesaleOk

`func (o *TransactionPrices) GetWholesaleOk() (*TransactionPricesWholesale, bool)`

GetWholesaleOk returns a tuple with the Wholesale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesale

`func (o *TransactionPrices) SetWholesale(v TransactionPricesWholesale)`

SetWholesale sets Wholesale field to given value.


### GetRetail

`func (o *TransactionPrices) GetRetail() TransactionPricesRetail`

GetRetail returns the Retail field if non-nil, zero value otherwise.

### GetRetailOk

`func (o *TransactionPrices) GetRetailOk() (*TransactionPricesRetail, bool)`

GetRetailOk returns a tuple with the Retail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetail

`func (o *TransactionPrices) SetRetail(v TransactionPricesRetail)`

SetRetail sets Retail field to given value.


### SetRetailNil

`func (o *TransactionPrices) SetRetailNil(b bool)`

 SetRetailNil sets the value for Retail to be an explicit nil

### UnsetRetail
`func (o *TransactionPrices) UnsetRetail()`

UnsetRetail ensures that no value is present for Retail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


