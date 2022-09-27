# TransactionSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**NationalityCountryIsoCode** | Pointer to **string** |  | [optional] 
**MobileNumber** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**AddressText** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountryIsoCode** | Pointer to **string** |  | [optional] 
**AddressPostalCode** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionSender

`func NewTransactionSender() *TransactionSender`

NewTransactionSender instantiates a new TransactionSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSenderWithDefaults

`func NewTransactionSenderWithDefaults() *TransactionSender`

NewTransactionSenderWithDefaults instantiates a new TransactionSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastName

`func (o *TransactionSender) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TransactionSender) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TransactionSender) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TransactionSender) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstName

`func (o *TransactionSender) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TransactionSender) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TransactionSender) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TransactionSender) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *TransactionSender) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *TransactionSender) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *TransactionSender) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *TransactionSender) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNationalityCountryIsoCode

`func (o *TransactionSender) GetNationalityCountryIsoCode() string`

GetNationalityCountryIsoCode returns the NationalityCountryIsoCode field if non-nil, zero value otherwise.

### GetNationalityCountryIsoCodeOk

`func (o *TransactionSender) GetNationalityCountryIsoCodeOk() (*string, bool)`

GetNationalityCountryIsoCodeOk returns a tuple with the NationalityCountryIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCountryIsoCode

`func (o *TransactionSender) SetNationalityCountryIsoCode(v string)`

SetNationalityCountryIsoCode sets NationalityCountryIsoCode field to given value.

### HasNationalityCountryIsoCode

`func (o *TransactionSender) HasNationalityCountryIsoCode() bool`

HasNationalityCountryIsoCode returns a boolean if a field has been set.

### GetMobileNumber

`func (o *TransactionSender) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *TransactionSender) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *TransactionSender) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *TransactionSender) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetEmail

`func (o *TransactionSender) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TransactionSender) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TransactionSender) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TransactionSender) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddressText

`func (o *TransactionSender) GetAddressText() string`

GetAddressText returns the AddressText field if non-nil, zero value otherwise.

### GetAddressTextOk

`func (o *TransactionSender) GetAddressTextOk() (*string, bool)`

GetAddressTextOk returns a tuple with the AddressText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressText

`func (o *TransactionSender) SetAddressText(v string)`

SetAddressText sets AddressText field to given value.

### HasAddressText

`func (o *TransactionSender) HasAddressText() bool`

HasAddressText returns a boolean if a field has been set.

### GetAddressCity

`func (o *TransactionSender) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *TransactionSender) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *TransactionSender) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *TransactionSender) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountryIsoCode

`func (o *TransactionSender) GetAddressCountryIsoCode() string`

GetAddressCountryIsoCode returns the AddressCountryIsoCode field if non-nil, zero value otherwise.

### GetAddressCountryIsoCodeOk

`func (o *TransactionSender) GetAddressCountryIsoCodeOk() (*string, bool)`

GetAddressCountryIsoCodeOk returns a tuple with the AddressCountryIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountryIsoCode

`func (o *TransactionSender) SetAddressCountryIsoCode(v string)`

SetAddressCountryIsoCode sets AddressCountryIsoCode field to given value.

### HasAddressCountryIsoCode

`func (o *TransactionSender) HasAddressCountryIsoCode() bool`

HasAddressCountryIsoCode returns a boolean if a field has been set.

### GetAddressPostalCode

`func (o *TransactionSender) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *TransactionSender) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *TransactionSender) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *TransactionSender) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


