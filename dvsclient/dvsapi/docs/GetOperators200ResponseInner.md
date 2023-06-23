# GetOperators200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Country** | [**Country**](Country.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 

## Methods

### NewGetOperators200ResponseInner

`func NewGetOperators200ResponseInner(id int32, name string, country Country, regions []ServiceRegion, ) *GetOperators200ResponseInner`

NewGetOperators200ResponseInner instantiates a new GetOperators200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOperators200ResponseInnerWithDefaults

`func NewGetOperators200ResponseInnerWithDefaults() *GetOperators200ResponseInner`

NewGetOperators200ResponseInnerWithDefaults instantiates a new GetOperators200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOperators200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOperators200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOperators200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GetOperators200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOperators200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOperators200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *GetOperators200ResponseInner) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetOperators200ResponseInner) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetOperators200ResponseInner) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetRegions

`func (o *GetOperators200ResponseInner) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetOperators200ResponseInner) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetOperators200ResponseInner) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *GetOperators200ResponseInner) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *GetOperators200ResponseInner) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


