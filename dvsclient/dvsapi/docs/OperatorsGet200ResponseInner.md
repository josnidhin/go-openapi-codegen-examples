# OperatorsGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Country** | [**Country**](Country.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 

## Methods

### NewOperatorsGet200ResponseInner

`func NewOperatorsGet200ResponseInner(id int32, name string, country Country, regions []ServiceRegion, ) *OperatorsGet200ResponseInner`

NewOperatorsGet200ResponseInner instantiates a new OperatorsGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorsGet200ResponseInnerWithDefaults

`func NewOperatorsGet200ResponseInnerWithDefaults() *OperatorsGet200ResponseInner`

NewOperatorsGet200ResponseInnerWithDefaults instantiates a new OperatorsGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatorsGet200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatorsGet200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatorsGet200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OperatorsGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatorsGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatorsGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *OperatorsGet200ResponseInner) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OperatorsGet200ResponseInner) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OperatorsGet200ResponseInner) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetRegions

`func (o *OperatorsGet200ResponseInner) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OperatorsGet200ResponseInner) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OperatorsGet200ResponseInner) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *OperatorsGet200ResponseInner) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OperatorsGet200ResponseInner) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


