# GetClientConfig200ResponseDenylist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidPackages** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Networks** | Pointer to [**GetClientConfig200ResponseNetworks1**](GetClientConfig_200_Response_networks_1.md) |  | [optional] 
**Visible** | **bool** |  | 

## Methods

### NewGetClientConfig200ResponseDenylist

`func NewGetClientConfig200ResponseDenylist(name string, visible bool, ) *GetClientConfig200ResponseDenylist`

NewGetClientConfig200ResponseDenylist instantiates a new GetClientConfig200ResponseDenylist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientConfig200ResponseDenylistWithDefaults

`func NewGetClientConfig200ResponseDenylistWithDefaults() *GetClientConfig200ResponseDenylist`

NewGetClientConfig200ResponseDenylistWithDefaults instantiates a new GetClientConfig200ResponseDenylist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidPackages

`func (o *GetClientConfig200ResponseDenylist) GetAndroidPackages() []string`

GetAndroidPackages returns the AndroidPackages field if non-nil, zero value otherwise.

### GetAndroidPackagesOk

`func (o *GetClientConfig200ResponseDenylist) GetAndroidPackagesOk() (*[]string, bool)`

GetAndroidPackagesOk returns a tuple with the AndroidPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackages

`func (o *GetClientConfig200ResponseDenylist) SetAndroidPackages(v []string)`

SetAndroidPackages sets AndroidPackages field to given value.

### HasAndroidPackages

`func (o *GetClientConfig200ResponseDenylist) HasAndroidPackages() bool`

HasAndroidPackages returns a boolean if a field has been set.

### GetName

`func (o *GetClientConfig200ResponseDenylist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClientConfig200ResponseDenylist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClientConfig200ResponseDenylist) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *GetClientConfig200ResponseDenylist) GetNetworks() GetClientConfig200ResponseNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetClientConfig200ResponseDenylist) GetNetworksOk() (*GetClientConfig200ResponseNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetClientConfig200ResponseDenylist) SetNetworks(v GetClientConfig200ResponseNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetClientConfig200ResponseDenylist) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetVisible

`func (o *GetClientConfig200ResponseDenylist) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetClientConfig200ResponseDenylist) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetClientConfig200ResponseDenylist) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


