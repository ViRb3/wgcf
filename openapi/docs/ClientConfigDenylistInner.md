# ClientConfigDenylistInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidPackages** | Pointer to **[]string** |  | [optional]
**Name** | **string** |  |
**Networks** | Pointer to [**ClientConfigDenylistInnerNetworks**](ClientConfigDenylistInnerNetworks.md) |  | [optional]
**Visible** | **bool** |  |

## Methods

### NewClientConfigDenylistInner

`func NewClientConfigDenylistInner(name string, visible bool, ) *ClientConfigDenylistInner`

NewClientConfigDenylistInner instantiates a new ClientConfigDenylistInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigDenylistInnerWithDefaults

`func NewClientConfigDenylistInnerWithDefaults() *ClientConfigDenylistInner`

NewClientConfigDenylistInnerWithDefaults instantiates a new ClientConfigDenylistInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidPackages

`func (o *ClientConfigDenylistInner) GetAndroidPackages() []string`

GetAndroidPackages returns the AndroidPackages field if non-nil, zero value otherwise.

### GetAndroidPackagesOk

`func (o *ClientConfigDenylistInner) GetAndroidPackagesOk() (*[]string, bool)`

GetAndroidPackagesOk returns a tuple with the AndroidPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackages

`func (o *ClientConfigDenylistInner) SetAndroidPackages(v []string)`

SetAndroidPackages sets AndroidPackages field to given value.

### HasAndroidPackages

`func (o *ClientConfigDenylistInner) HasAndroidPackages() bool`

HasAndroidPackages returns a boolean if a field has been set.

### GetName

`func (o *ClientConfigDenylistInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientConfigDenylistInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientConfigDenylistInner) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *ClientConfigDenylistInner) GetNetworks() ClientConfigDenylistInnerNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ClientConfigDenylistInner) GetNetworksOk() (*ClientConfigDenylistInnerNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ClientConfigDenylistInner) SetNetworks(v ClientConfigDenylistInnerNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ClientConfigDenylistInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetVisible

`func (o *ClientConfigDenylistInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ClientConfigDenylistInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ClientConfigDenylistInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


