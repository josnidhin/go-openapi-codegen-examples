# ProductBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**Service** | [**Service**](Service.md) |  | 
**Operator** | [**OperatorsGet200ResponseInner**](OperatorsGet200ResponseInner.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 

## Methods

### NewProductBase

`func NewProductBase(id int32, name string, description string, tags []string, service Service, operator OperatorsGet200ResponseInner, regions []ServiceRegion, ) *ProductBase`

NewProductBase instantiates a new ProductBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductBaseWithDefaults

`func NewProductBaseWithDefaults() *ProductBase`

NewProductBaseWithDefaults instantiates a new ProductBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductBase) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductBase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ProductBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductBase) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ProductBase) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProductBase) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetService

`func (o *ProductBase) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProductBase) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProductBase) SetService(v Service)`

SetService sets Service field to given value.


### GetOperator

`func (o *ProductBase) GetOperator() OperatorsGet200ResponseInner`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProductBase) GetOperatorOk() (*OperatorsGet200ResponseInner, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProductBase) SetOperator(v OperatorsGet200ResponseInner)`

SetOperator sets Operator field to given value.


### GetRegions

`func (o *ProductBase) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProductBase) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProductBase) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *ProductBase) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ProductBase) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


