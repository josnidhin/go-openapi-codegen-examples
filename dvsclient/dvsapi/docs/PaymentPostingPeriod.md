# PaymentPostingPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | [**TimeUnits**](TimeUnits.md) |  | 
**Quantity** | **float32** |  | 
**Type** | [**TimeUnitTypes**](TimeUnitTypes.md) |  | 

## Methods

### NewPaymentPostingPeriod

`func NewPaymentPostingPeriod(unit TimeUnits, quantity float32, type_ TimeUnitTypes, ) *PaymentPostingPeriod`

NewPaymentPostingPeriod instantiates a new PaymentPostingPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPostingPeriodWithDefaults

`func NewPaymentPostingPeriodWithDefaults() *PaymentPostingPeriod`

NewPaymentPostingPeriodWithDefaults instantiates a new PaymentPostingPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *PaymentPostingPeriod) GetUnit() TimeUnits`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PaymentPostingPeriod) GetUnitOk() (*TimeUnits, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PaymentPostingPeriod) SetUnit(v TimeUnits)`

SetUnit sets Unit field to given value.


### GetQuantity

`func (o *PaymentPostingPeriod) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentPostingPeriod) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentPostingPeriod) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetType

`func (o *PaymentPostingPeriod) GetType() TimeUnitTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentPostingPeriod) GetTypeOk() (*TimeUnitTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentPostingPeriod) SetType(v TimeUnitTypes)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


