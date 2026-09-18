# SourceDevice

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

## Methods

### NewSourceDevice

`func NewSourceDevice(id string, ) *SourceDevice`

NewSourceDevice instantiates a new SourceDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDeviceWithDefaults

`func NewSourceDeviceWithDefaults() *SourceDevice`

NewSourceDeviceWithDefaults instantiates a new SourceDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SourceDevice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SourceDevice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SourceDevice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SourceDevice) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *SourceDevice) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SourceDevice) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SourceDevice) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SourceDevice) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFcmToken

`func (o *SourceDevice) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *SourceDevice) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *SourceDevice) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.

### HasFcmToken

`func (o *SourceDevice) HasFcmToken() bool`

HasFcmToken returns a boolean if a field has been set.

### GetId

`func (o *SourceDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceDevice) SetId(v string)`

SetId sets Id field to given value.


### GetInstallId

`func (o *SourceDevice) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *SourceDevice) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *SourceDevice) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *SourceDevice) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### GetKey

`func (o *SourceDevice) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SourceDevice) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SourceDevice) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SourceDevice) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocale

`func (o *SourceDevice) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SourceDevice) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SourceDevice) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SourceDevice) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetModel

`func (o *SourceDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SourceDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SourceDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SourceDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetManufacturer

`func (o *SourceDevice) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *SourceDevice) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *SourceDevice) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *SourceDevice) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetName

`func (o *SourceDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVersion

`func (o *SourceDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *SourceDevice) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *SourceDevice) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *SourceDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlace

`func (o *SourceDevice) GetPlace() int32`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *SourceDevice) GetPlaceOk() (*int32, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *SourceDevice) SetPlace(v int32)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *SourceDevice) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SourceDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SourceDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SourceDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SourceDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDeviceId

`func (o *SourceDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SourceDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SourceDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SourceDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetTos

`func (o *SourceDevice) GetTos() time.Time`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *SourceDevice) GetTosOk() (*time.Time, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *SourceDevice) SetTos(v time.Time)`

SetTos sets Tos field to given value.

### HasTos

`func (o *SourceDevice) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetType

`func (o *SourceDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourceDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *SourceDevice) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SourceDevice) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SourceDevice) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SourceDevice) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWaitlistEnabled

`func (o *SourceDevice) GetWaitlistEnabled() bool`

GetWaitlistEnabled returns the WaitlistEnabled field if non-nil, zero value otherwise.

### GetWaitlistEnabledOk

`func (o *SourceDevice) GetWaitlistEnabledOk() (*bool, bool)`

GetWaitlistEnabledOk returns a tuple with the WaitlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistEnabled

`func (o *SourceDevice) SetWaitlistEnabled(v bool)`

SetWaitlistEnabled sets WaitlistEnabled field to given value.

### HasWaitlistEnabled

`func (o *SourceDevice) HasWaitlistEnabled() bool`

HasWaitlistEnabled returns a boolean if a field has been set.

### GetWarpEnabled

`func (o *SourceDevice) GetWarpEnabled() bool`

GetWarpEnabled returns the WarpEnabled field if non-nil, zero value otherwise.

### GetWarpEnabledOk

`func (o *SourceDevice) GetWarpEnabledOk() (*bool, bool)`

GetWarpEnabledOk returns a tuple with the WarpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarpEnabled

`func (o *SourceDevice) SetWarpEnabled(v bool)`

SetWarpEnabled sets WarpEnabled field to given value.

### HasWarpEnabled

`func (o *SourceDevice) HasWarpEnabled() bool`

HasWarpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


