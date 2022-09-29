# TransactionProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**Service** | [**Service**](Service.md) |  | 
**Operator** | [**GetOperators200ResponseInner**](GetOperators200ResponseInner.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 
**Pin** | Pointer to [**TransactionProductAllOfPin**](TransactionProductAllOfPin.md) |  | [optional] 

## Methods

### NewTransactionProduct

`func NewTransactionProduct(id int32, name string, description string, tags []string, service Service, operator GetOperators200ResponseInner, regions []ServiceRegion, ) *TransactionProduct`

NewTransactionProduct instantiates a new TransactionProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionProductWithDefaults

`func NewTransactionProductWithDefaults() *TransactionProduct`

NewTransactionProductWithDefaults instantiates a new TransactionProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionProduct) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TransactionProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionProduct) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TransactionProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionProduct) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *TransactionProduct) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransactionProduct) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransactionProduct) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *TransactionProduct) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TransactionProduct) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetService

`func (o *TransactionProduct) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TransactionProduct) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TransactionProduct) SetService(v Service)`

SetService sets Service field to given value.


### GetOperator

`func (o *TransactionProduct) GetOperator() GetOperators200ResponseInner`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TransactionProduct) GetOperatorOk() (*GetOperators200ResponseInner, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TransactionProduct) SetOperator(v GetOperators200ResponseInner)`

SetOperator sets Operator field to given value.


### GetRegions

`func (o *TransactionProduct) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *TransactionProduct) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *TransactionProduct) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *TransactionProduct) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *TransactionProduct) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetPin

`func (o *TransactionProduct) GetPin() TransactionProductAllOfPin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *TransactionProduct) GetPinOk() (*TransactionProductAllOfPin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *TransactionProduct) SetPin(v TransactionProductAllOfPin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *TransactionProduct) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


