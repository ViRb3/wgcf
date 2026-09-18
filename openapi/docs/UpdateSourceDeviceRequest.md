# UpdateSourceDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional]
**KeyType** | Pointer to **string** |  | [optional]
**TunnelType** | Pointer to **string** |  | [optional]
**FcmToken** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**OsVersion** | Pointer to **string** |  | [optional]
**Manufacturer** | Pointer to **string** |  | [optional]
**Model** | Pointer to **string** |  | [optional]
**DeviceId** | Pointer to **string** |  | [optional]

## Methods

### NewUpdateSourceDeviceRequest

`func NewUpdateSourceDeviceRequest() *UpdateSourceDeviceRequest`

NewUpdateSourceDeviceRequest instantiates a new UpdateSourceDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourceDeviceRequestWithDefaults

`func NewUpdateSourceDeviceRequestWithDefaults() *UpdateSourceDeviceRequest`

NewUpdateSourceDeviceRequestWithDefaults instantiates a new UpdateSourceDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateSourceDeviceRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSourceDeviceRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSourceDeviceRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateSourceDeviceRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyType

`func (o *UpdateSourceDeviceRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *UpdateSourceDeviceRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *UpdateSourceDeviceRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *UpdateSourceDeviceRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetTunnelType

`func (o *UpdateSourceDeviceRequest) GetTunnelType() string`

GetTunnelType returns the TunnelType field if non-nil, zero value otherwise.

### GetTunnelTypeOk

`func (o *UpdateSourceDeviceRequest) GetTunnelTypeOk() (*string, bool)`

GetTunnelTypeOk returns a tuple with the TunnelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelType

`func (o *UpdateSourceDeviceRequest) SetTunnelType(v string)`

SetTunnelType sets TunnelType field to given value.

### HasTunnelType

`func (o *UpdateSourceDeviceRequest) HasTunnelType() bool`

HasTunnelType returns a boolean if a field has been set.

### GetFcmToken

`func (o *UpdateSourceDeviceRequest) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *UpdateSourceDeviceRequest) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *UpdateSourceDeviceRequest) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.

### HasFcmToken

`func (o *UpdateSourceDeviceRequest) HasFcmToken() bool`

HasFcmToken returns a boolean if a field has been set.

### GetName

`func (o *UpdateSourceDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSourceDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSourceDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSourceDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVersion

`func (o *UpdateSourceDeviceRequest) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *UpdateSourceDeviceRequest) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *UpdateSourceDeviceRequest) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *UpdateSourceDeviceRequest) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetManufacturer

`func (o *UpdateSourceDeviceRequest) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *UpdateSourceDeviceRequest) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *UpdateSourceDeviceRequest) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *UpdateSourceDeviceRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *UpdateSourceDeviceRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateSourceDeviceRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateSourceDeviceRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpdateSourceDeviceRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDeviceId

`func (o *UpdateSourceDeviceRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpdateSourceDeviceRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpdateSourceDeviceRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpdateSourceDeviceRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


