# TransactionRequestedValuesSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitType** | [**UnitTypes**](UnitTypes.md) |  | 
**Unit** | **string** |  | 
**Amount** | **float32** |  | 

## Methods

### NewTransactionRequestedValuesSource

`func NewTransactionRequestedValuesSource(unitType UnitTypes, unit string, amount float32, ) *TransactionRequestedValuesSource`

NewTransactionRequestedValuesSource instantiates a new TransactionRequestedValuesSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestedValuesSourceWithDefaults

`func NewTransactionRequestedValuesSourceWithDefaults() *TransactionRequestedValuesSource`

NewTransactionRequestedValuesSourceWithDefaults instantiates a new TransactionRequestedValuesSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitType

`func (o *TransactionRequestedValuesSource) GetUnitType() UnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *TransactionRequestedValuesSource) GetUnitTypeOk() (*UnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *TransactionRequestedValuesSource) SetUnitType(v UnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *TransactionRequestedValuesSource) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionRequestedValuesSource) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionRequestedValuesSource) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAmount

`func (o *TransactionRequestedValuesSource) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionRequestedValuesSource) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionRequestedValuesSource) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


