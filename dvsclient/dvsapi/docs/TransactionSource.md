# TransactionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitType** | [**UnitTypes**](UnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | **float32** |  | 

## Methods

### NewTransactionSource

`func NewTransactionSource(unitType UnitTypes, unit string, amount float32, ) *TransactionSource`

NewTransactionSource instantiates a new TransactionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSourceWithDefaults

`func NewTransactionSourceWithDefaults() *TransactionSource`

NewTransactionSourceWithDefaults instantiates a new TransactionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitType

`func (o *TransactionSource) GetUnitType() UnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *TransactionSource) GetUnitTypeOk() (*UnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *TransactionSource) SetUnitType(v UnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *TransactionSource) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionSource) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionSource) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *TransactionSource) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSource) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSource) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


