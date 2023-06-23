# ProductService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Subservice** | Pointer to [**NullableProductServiceSubservice**](ProductServiceSubservice.md) |  | [optional] 

## Methods

### NewProductService

`func NewProductService(id int32, name string, ) *ProductService`

NewProductService instantiates a new ProductService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductServiceWithDefaults

`func NewProductServiceWithDefaults() *ProductService`

NewProductServiceWithDefaults instantiates a new ProductService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductService) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductService) SetName(v string)`

SetName sets Name field to given value.


### GetSubservice

`func (o *ProductService) GetSubservice() ProductServiceSubservice`

GetSubservice returns the Subservice field if non-nil, zero value otherwise.

### GetSubserviceOk

`func (o *ProductService) GetSubserviceOk() (*ProductServiceSubservice, bool)`

GetSubserviceOk returns a tuple with the Subservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubservice

`func (o *ProductService) SetSubservice(v ProductServiceSubservice)`

SetSubservice sets Subservice field to given value.

### HasSubservice

`func (o *ProductService) HasSubservice() bool`

HasSubservice returns a boolean if a field has been set.

### SetSubserviceNil

`func (o *ProductService) SetSubserviceNil(b bool)`

 SetSubserviceNil sets the value for Subservice to be an explicit nil

### UnsetSubservice
`func (o *ProductService) UnsetSubservice()`

UnsetSubservice ensures that no value is present for Subservice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


