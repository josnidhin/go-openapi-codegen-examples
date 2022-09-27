# StatementBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitType** | Pointer to [**UnitTypes**](UnitTypes.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 

## Methods

### NewStatementBalance

`func NewStatementBalance() *StatementBalance`

NewStatementBalance instantiates a new StatementBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementBalanceWithDefaults

`func NewStatementBalanceWithDefaults() *StatementBalance`

NewStatementBalanceWithDefaults instantiates a new StatementBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitType

`func (o *StatementBalance) GetUnitType() UnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *StatementBalance) GetUnitTypeOk() (*UnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *StatementBalance) SetUnitType(v UnitTypes)`

SetUnitType sets UnitType field to given value.

### HasUnitType

`func (o *StatementBalance) HasUnitType() bool`

HasUnitType returns a boolean if a field has been set.

### GetUnit

`func (o *StatementBalance) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *StatementBalance) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *StatementBalance) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *StatementBalance) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetAmount

`func (o *StatementBalance) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StatementBalance) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StatementBalance) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StatementBalance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


