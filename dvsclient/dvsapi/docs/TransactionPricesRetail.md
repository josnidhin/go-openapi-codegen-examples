# TransactionPricesRetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitType** | [**UnitTypes**](UnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | **NullableFloat32** |  | 
**Fee** | **NullableFloat32** |  | 

## Methods

### NewTransactionPricesRetail

`func NewTransactionPricesRetail(unitType UnitTypes, unit string, amount NullableFloat32, fee NullableFloat32, ) *TransactionPricesRetail`

NewTransactionPricesRetail instantiates a new TransactionPricesRetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPricesRetailWithDefaults

`func NewTransactionPricesRetailWithDefaults() *TransactionPricesRetail`

NewTransactionPricesRetailWithDefaults instantiates a new TransactionPricesRetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitType

`func (o *TransactionPricesRetail) GetUnitType() UnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *TransactionPricesRetail) GetUnitTypeOk() (*UnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *TransactionPricesRetail) SetUnitType(v UnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *TransactionPricesRetail) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionPricesRetail) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionPricesRetail) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *TransactionPricesRetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionPricesRetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionPricesRetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *TransactionPricesRetail) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TransactionPricesRetail) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetFee

`func (o *TransactionPricesRetail) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionPricesRetail) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionPricesRetail) SetFee(v float32)`

SetFee sets Fee field to given value.


### SetFeeNil

`func (o *TransactionPricesRetail) SetFeeNil(b bool)`

 SetFeeNil sets the value for Fee to be an explicit nil

### UnsetFee
`func (o *TransactionPricesRetail) UnsetFee()`

UnsetFee ensures that no value is present for Fee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


