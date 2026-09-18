# GetSourceDevice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional]
**Enabled** | Pointer to **bool** |  | [optional]
**FcmToken** | Pointer to **string** |  | [optional]
**Id** | **string** |  |
**InstallId** | Pointer to **string** |  | [optional]
**Key** | Pointer to **string** |  | [optional]
**Locale** | Pointer to **string** |  | [optional]
**Model** | Pointer to **string** |  | [optional]
**Manufacturer** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**OsVersion** | Pointer to **string** |  | [optional]
**Place** | Pointer to **int32** |  | [optional]
**SerialNumber** | Pointer to **string** |  | [optional]
**DeviceId** | Pointer to **string** |  | [optional]
**Tos** | Pointer to **time.Time** |  | [optional]
**Type** | Pointer to **string** |  | [optional]
**Updated** | Pointer to **time.Time** |  | [optional]
**WaitlistEnabled** | Pointer to **bool** |  | [optional]
**WarpEnabled** | Pointer to **bool** |  | [optional]
**Account** | [**Account**](Account.md) |  |
**Config** | Pointer to [**Config**](Config.md) |  | [optional]
**Policy** | Pointer to [**Policy**](Policy.md) |  | [optional]
**OverrideCodes** | Pointer to [**OverrideCodes**](OverrideCodes.md) |  | [optional]
**AlternateNetworks** | Pointer to [**[]AlternateNetwork**](AlternateNetwork.md) |  | [optional]
**DexTests** | Pointer to [**[]DexTest**](DexTest.md) |  | [optional]
**KeyType** | **string** |  |
**TunnelType** | **string** |  |

## Methods

### NewGetSourceDevice200Response

`func NewGetSourceDevice200Response(id string, account Account, keyType string, tunnelType string, ) *GetSourceDevice200Response`

NewGetSourceDevice200Response instantiates a new GetSourceDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceDevice200ResponseWithDefaults

`func NewGetSourceDevice200ResponseWithDefaults() *GetSourceDevice200Response`

NewGetSourceDevice200ResponseWithDefaults instantiates a new GetSourceDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetSourceDevice200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetSourceDevice200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetSourceDevice200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetSourceDevice200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *GetSourceDevice200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSourceDevice200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSourceDevice200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetSourceDevice200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFcmToken

`func (o *GetSourceDevice200Response) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *GetSourceDevice200Response) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *GetSourceDevice200Response) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.

### HasFcmToken

`func (o *GetSourceDevice200Response) HasFcmToken() bool`

HasFcmToken returns a boolean if a field has been set.

### GetId

`func (o *GetSourceDevice200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSourceDevice200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSourceDevice200Response) SetId(v string)`

SetId sets Id field to given value.


### GetInstallId

`func (o *GetSourceDevice200Response) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *GetSourceDevice200Response) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *GetSourceDevice200Response) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *GetSourceDevice200Response) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### GetKey

`func (o *GetSourceDevice200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetSourceDevice200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetSourceDevice200Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetSourceDevice200Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocale

`func (o *GetSourceDevice200Response) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *GetSourceDevice200Response) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *GetSourceDevice200Response) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *GetSourceDevice200Response) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetModel

`func (o *GetSourceDevice200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetSourceDevice200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetSourceDevice200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetSourceDevice200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetManufacturer

`func (o *GetSourceDevice200Response) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *GetSourceDevice200Response) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *GetSourceDevice200Response) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *GetSourceDevice200Response) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetName

`func (o *GetSourceDevice200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSourceDevice200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSourceDevice200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSourceDevice200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetSourceDevice200Response) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetSourceDevice200Response) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetSourceDevice200Response) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetSourceDevice200Response) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlace

`func (o *GetSourceDevice200Response) GetPlace() int32`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *GetSourceDevice200Response) GetPlaceOk() (*int32, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *GetSourceDevice200Response) SetPlace(v int32)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *GetSourceDevice200Response) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetSerialNumber

`func (o *GetSourceDevice200Response) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetSourceDevice200Response) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetSourceDevice200Response) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetSourceDevice200Response) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDeviceId

`func (o *GetSourceDevice200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetSourceDevice200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetSourceDevice200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetSourceDevice200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetTos

`func (o *GetSourceDevice200Response) GetTos() time.Time`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *GetSourceDevice200Response) GetTosOk() (*time.Time, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *GetSourceDevice200Response) SetTos(v time.Time)`

SetTos sets Tos field to given value.

### HasTos

`func (o *GetSourceDevice200Response) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetType

`func (o *GetSourceDevice200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSourceDevice200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSourceDevice200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSourceDevice200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetSourceDevice200Response) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetSourceDevice200Response) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetSourceDevice200Response) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetSourceDevice200Response) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWaitlistEnabled

`func (o *GetSourceDevice200Response) GetWaitlistEnabled() bool`

GetWaitlistEnabled returns the WaitlistEnabled field if non-nil, zero value otherwise.

### GetWaitlistEnabledOk

`func (o *GetSourceDevice200Response) GetWaitlistEnabledOk() (*bool, bool)`

GetWaitlistEnabledOk returns a tuple with the WaitlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistEnabled

`func (o *GetSourceDevice200Response) SetWaitlistEnabled(v bool)`

SetWaitlistEnabled sets WaitlistEnabled field to given value.

### HasWaitlistEnabled

`func (o *GetSourceDevice200Response) HasWaitlistEnabled() bool`

HasWaitlistEnabled returns a boolean if a field has been set.

### GetWarpEnabled

`func (o *GetSourceDevice200Response) GetWarpEnabled() bool`

GetWarpEnabled returns the WarpEnabled field if non-nil, zero value otherwise.

### GetWarpEnabledOk

`func (o *GetSourceDevice200Response) GetWarpEnabledOk() (*bool, bool)`

GetWarpEnabledOk returns a tuple with the WarpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarpEnabled

`func (o *GetSourceDevice200Response) SetWarpEnabled(v bool)`

SetWarpEnabled sets WarpEnabled field to given value.

### HasWarpEnabled

`func (o *GetSourceDevice200Response) HasWarpEnabled() bool`

HasWarpEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *GetSourceDevice200Response) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetSourceDevice200Response) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetSourceDevice200Response) SetAccount(v Account)`

SetAccount sets Account field to given value.


### GetConfig

`func (o *GetSourceDevice200Response) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSourceDevice200Response) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSourceDevice200Response) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetSourceDevice200Response) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPolicy

`func (o *GetSourceDevice200Response) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetSourceDevice200Response) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetSourceDevice200Response) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetSourceDevice200Response) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetOverrideCodes

`func (o *GetSourceDevice200Response) GetOverrideCodes() OverrideCodes`

GetOverrideCodes returns the OverrideCodes field if non-nil, zero value otherwise.

### GetOverrideCodesOk

`func (o *GetSourceDevice200Response) GetOverrideCodesOk() (*OverrideCodes, bool)`

GetOverrideCodesOk returns a tuple with the OverrideCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCodes

`func (o *GetSourceDevice200Response) SetOverrideCodes(v OverrideCodes)`

SetOverrideCodes sets OverrideCodes field to given value.

### HasOverrideCodes

`func (o *GetSourceDevice200Response) HasOverrideCodes() bool`

HasOverrideCodes returns a boolean if a field has been set.

### GetAlternateNetworks

`func (o *GetSourceDevice200Response) GetAlternateNetworks() []AlternateNetwork`

GetAlternateNetworks returns the AlternateNetworks field if non-nil, zero value otherwise.

### GetAlternateNetworksOk

`func (o *GetSourceDevice200Response) GetAlternateNetworksOk() (*[]AlternateNetwork, bool)`

GetAlternateNetworksOk returns a tuple with the AlternateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateNetworks

`func (o *GetSourceDevice200Response) SetAlternateNetworks(v []AlternateNetwork)`

SetAlternateNetworks sets AlternateNetworks field to given value.

### HasAlternateNetworks

`func (o *GetSourceDevice200Response) HasAlternateNetworks() bool`

HasAlternateNetworks returns a boolean if a field has been set.

### GetDexTests

`func (o *GetSourceDevice200Response) GetDexTests() []DexTest`

GetDexTests returns the DexTests field if non-nil, zero value otherwise.

### GetDexTestsOk

`func (o *GetSourceDevice200Response) GetDexTestsOk() (*[]DexTest, bool)`

GetDexTestsOk returns a tuple with the DexTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexTests

`func (o *GetSourceDevice200Response) SetDexTests(v []DexTest)`

SetDexTests sets DexTests field to given value.

### HasDexTests

`func (o *GetSourceDevice200Response) HasDexTests() bool`

HasDexTests returns a boolean if a field has been set.

### GetKeyType

`func (o *GetSourceDevice200Response) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GetSourceDevice200Response) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GetSourceDevice200Response) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetTunnelType

`func (o *GetSourceDevice200Response) GetTunnelType() string`

GetTunnelType returns the TunnelType field if non-nil, zero value otherwise.

### GetTunnelTypeOk

`func (o *GetSourceDevice200Response) GetTunnelTypeOk() (*string, bool)`

GetTunnelTypeOk returns a tuple with the TunnelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelType

`func (o *GetSourceDevice200Response) SetTunnelType(v string)`

SetTunnelType sets TunnelType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


