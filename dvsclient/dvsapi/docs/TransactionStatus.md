# TransactionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Class** | Pointer to [**TransactionStatusClass**](TransactionStatusClass.md) |  | [optional] 

## Methods

### NewTransactionStatus

`func NewTransactionStatus() *TransactionStatus`

NewTransactionStatus instantiates a new TransactionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStatusWithDefaults

`func NewTransactionStatusWithDefaults() *TransactionStatus`

NewTransactionStatusWithDefaults instantiates a new TransactionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetClass

`func (o *TransactionStatus) GetClass() TransactionStatusClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *TransactionStatus) GetClassOk() (*TransactionStatusClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *TransactionStatus) SetClass(v TransactionStatusClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *TransactionStatus) HasClass() bool`

HasClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


