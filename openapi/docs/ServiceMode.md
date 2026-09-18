# ServiceMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**AppMode**](AppMode.md) |  | [optional]
**Port** | Pointer to **int32** |  | [optional]

## Methods

### NewServiceMode

`func NewServiceMode() *ServiceMode`

NewServiceMode instantiates a new ServiceMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceModeWithDefaults

`func NewServiceModeWithDefaults() *ServiceMode`

NewServiceModeWithDefaults instantiates a new ServiceMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ServiceMode) GetMode() AppMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ServiceMode) GetModeOk() (*AppMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ServiceMode) SetMode(v AppMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ServiceMode) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPort

`func (o *ServiceMode) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServiceMode) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServiceMode) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServiceMode) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


