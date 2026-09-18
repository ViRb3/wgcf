# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FcmToken** | **string** |  |
**InstallId** | **string** |  |
**Key** | **string** |  |
**Locale** | **string** |  |
**Model** | Pointer to **string** |  | [optional]
**Referrer** | Pointer to **string** |  | [optional]
**Tos** | Pointer to **string** |  | [optional]
**Expired** | Pointer to **string** |  | [optional]
**SerialNumber** | **string** |  |
**OsVersion** | Pointer to **string** |  | [optional]
**KeyType** | Pointer to **string** |  | [optional]
**TunnelType** | Pointer to **string** |  | [optional]

## Methods

### NewRegisterRequest

`func NewRegisterRequest(fcmToken string, installId string, key string, locale string, serialNumber string, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFcmToken

`func (o *RegisterRequest) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *RegisterRequest) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *RegisterRequest) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.


### GetInstallId

`func (o *RegisterRequest) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *RegisterRequest) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *RegisterRequest) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.


### GetKey

`func (o *RegisterRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RegisterRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RegisterRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetLocale

`func (o *RegisterRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *RegisterRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *RegisterRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetModel

`func (o *RegisterRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RegisterRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RegisterRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RegisterRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetReferrer

`func (o *RegisterRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *RegisterRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *RegisterRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *RegisterRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetTos

`func (o *RegisterRequest) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *RegisterRequest) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *RegisterRequest) SetTos(v string)`

SetTos sets Tos field to given value.

### HasTos

`func (o *RegisterRequest) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetExpired

`func (o *RegisterRequest) GetExpired() string`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *RegisterRequest) GetExpiredOk() (*string, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *RegisterRequest) SetExpired(v string)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *RegisterRequest) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RegisterRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RegisterRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RegisterRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetOsVersion

`func (o *RegisterRequest) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *RegisterRequest) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *RegisterRequest) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *RegisterRequest) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetKeyType

`func (o *RegisterRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *RegisterRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *RegisterRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *RegisterRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetTunnelType

`func (o *RegisterRequest) GetTunnelType() string`

GetTunnelType returns the TunnelType field if non-nil, zero value otherwise.

### GetTunnelTypeOk

`func (o *RegisterRequest) GetTunnelTypeOk() (*string, bool)`

GetTunnelTypeOk returns a tuple with the TunnelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelType

`func (o *RegisterRequest) SetTunnelType(v string)`

SetTunnelType sets TunnelType field to given value.

### HasTunnelType

`func (o *RegisterRequest) HasTunnelType() bool`

HasTunnelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


