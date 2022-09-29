# LookupMobileNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Country** | [**Country**](Country.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 
**Identified** | **bool** | Indicates whether operator was identified as a direct match | 

## Methods

### NewLookupMobileNumber

`func NewLookupMobileNumber(id int32, name string, country Country, regions []ServiceRegion, identified bool, ) *LookupMobileNumber`

NewLookupMobileNumber instantiates a new LookupMobileNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupMobileNumberWithDefaults

`func NewLookupMobileNumberWithDefaults() *LookupMobileNumber`

NewLookupMobileNumberWithDefaults instantiates a new LookupMobileNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LookupMobileNumber) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LookupMobileNumber) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LookupMobileNumber) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *LookupMobileNumber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LookupMobileNumber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LookupMobileNumber) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *LookupMobileNumber) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LookupMobileNumber) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LookupMobileNumber) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetRegions

`func (o *LookupMobileNumber) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *LookupMobileNumber) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *LookupMobileNumber) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *LookupMobileNumber) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *LookupMobileNumber) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetIdentified

`func (o *LookupMobileNumber) GetIdentified() bool`

GetIdentified returns the Identified field if non-nil, zero value otherwise.

### GetIdentifiedOk

`func (o *LookupMobileNumber) GetIdentifiedOk() (*bool, bool)`

GetIdentifiedOk returns a tuple with the Identified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentified

`func (o *LookupMobileNumber) SetIdentified(v bool)`

SetIdentified sets Identified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


