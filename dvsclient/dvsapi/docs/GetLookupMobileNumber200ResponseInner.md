# GetLookupMobileNumber200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Country** | [**Country**](Country.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 
**Identified** | **bool** | Indicates whether operator was identified as a direct match | 

## Methods

### NewGetLookupMobileNumber200ResponseInner

`func NewGetLookupMobileNumber200ResponseInner(id int32, name string, country Country, regions []ServiceRegion, identified bool, ) *GetLookupMobileNumber200ResponseInner`

NewGetLookupMobileNumber200ResponseInner instantiates a new GetLookupMobileNumber200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLookupMobileNumber200ResponseInnerWithDefaults

`func NewGetLookupMobileNumber200ResponseInnerWithDefaults() *GetLookupMobileNumber200ResponseInner`

NewGetLookupMobileNumber200ResponseInnerWithDefaults instantiates a new GetLookupMobileNumber200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLookupMobileNumber200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLookupMobileNumber200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLookupMobileNumber200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GetLookupMobileNumber200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLookupMobileNumber200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLookupMobileNumber200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *GetLookupMobileNumber200ResponseInner) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetLookupMobileNumber200ResponseInner) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetLookupMobileNumber200ResponseInner) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetRegions

`func (o *GetLookupMobileNumber200ResponseInner) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetLookupMobileNumber200ResponseInner) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetLookupMobileNumber200ResponseInner) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *GetLookupMobileNumber200ResponseInner) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *GetLookupMobileNumber200ResponseInner) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetIdentified

`func (o *GetLookupMobileNumber200ResponseInner) GetIdentified() bool`

GetIdentified returns the Identified field if non-nil, zero value otherwise.

### GetIdentifiedOk

`func (o *GetLookupMobileNumber200ResponseInner) GetIdentifiedOk() (*bool, bool)`

GetIdentifiedOk returns a tuple with the Identified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentified

`func (o *GetLookupMobileNumber200ResponseInner) SetIdentified(v bool)`

SetIdentified sets Identified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


