# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UnitType** | [**UnitTypes**](UnitTypes.md) |  | 
**Unit** | **string** |  | 
**Available** | **float32** |  | 
**Holding** | **float32** |  | 
**CreditLimit** | **float32** |  | 

## Methods

### NewBalance

`func NewBalance(id int32, unitType UnitTypes, unit string, available float32, holding float32, creditLimit float32, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Balance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Balance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Balance) SetId(v int32)`

SetId sets Id field to given value.


### GetUnitType

`func (o *Balance) GetUnitType() UnitTypes`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *Balance) GetUnitTypeOk() (*UnitTypes, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *Balance) SetUnitType(v UnitTypes)`

SetUnitType sets UnitType field to given value.


### GetUnit

`func (o *Balance) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Balance) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Balance) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAvailable

`func (o *Balance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Balance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Balance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.


### GetHolding

`func (o *Balance) GetHolding() float32`

GetHolding returns the Holding field if non-nil, zero value otherwise.

### GetHoldingOk

`func (o *Balance) GetHoldingOk() (*float32, bool)`

GetHoldingOk returns a tuple with the Holding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolding

`func (o *Balance) SetHolding(v float32)`

SetHolding sets Holding field to given value.


### GetCreditLimit

`func (o *Balance) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *Balance) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *Balance) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


