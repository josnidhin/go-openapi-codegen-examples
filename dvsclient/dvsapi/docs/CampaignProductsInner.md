# CampaignProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**ProductTypes**](ProductTypes.md) |  | 
**Operator** | [**Operator**](Operator.md) |  | 

## Methods

### NewCampaignProductsInner

`func NewCampaignProductsInner(id int32, name string, description string, type_ ProductTypes, operator Operator, ) *CampaignProductsInner`

NewCampaignProductsInner instantiates a new CampaignProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignProductsInnerWithDefaults

`func NewCampaignProductsInnerWithDefaults() *CampaignProductsInner`

NewCampaignProductsInnerWithDefaults instantiates a new CampaignProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignProductsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignProductsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignProductsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignProductsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CampaignProductsInner) GetType() ProductTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignProductsInner) GetTypeOk() (*ProductTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignProductsInner) SetType(v ProductTypes)`

SetType sets Type field to given value.


### GetOperator

`func (o *CampaignProductsInner) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignProductsInner) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CampaignProductsInner) SetOperator(v Operator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


