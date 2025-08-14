# GetClientConfig200ResponseDenylistInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidPackages** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Networks** | Pointer to [**GetClientConfig200ResponseDenylistInnerNetworks**](GetClientConfig200ResponseDenylistInnerNetworks.md) |  | [optional] 
**Visible** | **bool** |  | 

## Methods

### NewGetClientConfig200ResponseDenylistInner

`func NewGetClientConfig200ResponseDenylistInner(name string, visible bool, ) *GetClientConfig200ResponseDenylistInner`

NewGetClientConfig200ResponseDenylistInner instantiates a new GetClientConfig200ResponseDenylistInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientConfig200ResponseDenylistInnerWithDefaults

`func NewGetClientConfig200ResponseDenylistInnerWithDefaults() *GetClientConfig200ResponseDenylistInner`

NewGetClientConfig200ResponseDenylistInnerWithDefaults instantiates a new GetClientConfig200ResponseDenylistInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidPackages

`func (o *GetClientConfig200ResponseDenylistInner) GetAndroidPackages() []string`

GetAndroidPackages returns the AndroidPackages field if non-nil, zero value otherwise.

### GetAndroidPackagesOk

`func (o *GetClientConfig200ResponseDenylistInner) GetAndroidPackagesOk() (*[]string, bool)`

GetAndroidPackagesOk returns a tuple with the AndroidPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackages

`func (o *GetClientConfig200ResponseDenylistInner) SetAndroidPackages(v []string)`

SetAndroidPackages sets AndroidPackages field to given value.

### HasAndroidPackages

`func (o *GetClientConfig200ResponseDenylistInner) HasAndroidPackages() bool`

HasAndroidPackages returns a boolean if a field has been set.

### GetName

`func (o *GetClientConfig200ResponseDenylistInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClientConfig200ResponseDenylistInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClientConfig200ResponseDenylistInner) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *GetClientConfig200ResponseDenylistInner) GetNetworks() GetClientConfig200ResponseDenylistInnerNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetClientConfig200ResponseDenylistInner) GetNetworksOk() (*GetClientConfig200ResponseDenylistInnerNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetClientConfig200ResponseDenylistInner) SetNetworks(v GetClientConfig200ResponseDenylistInnerNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetClientConfig200ResponseDenylistInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetVisible

`func (o *GetClientConfig200ResponseDenylistInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetClientConfig200ResponseDenylistInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetClientConfig200ResponseDenylistInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


