# CaptivePortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**Networks** | [**[]CaptivePortalNetwork**](CaptivePortalNetwork.md) |  |

## Methods

### NewCaptivePortal

`func NewCaptivePortal(name string, networks []CaptivePortalNetwork, ) *CaptivePortal`

NewCaptivePortal instantiates a new CaptivePortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptivePortalWithDefaults

`func NewCaptivePortalWithDefaults() *CaptivePortal`

NewCaptivePortalWithDefaults instantiates a new CaptivePortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CaptivePortal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptivePortal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptivePortal) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *CaptivePortal) GetNetworks() []CaptivePortalNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CaptivePortal) GetNetworksOk() (*[]CaptivePortalNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CaptivePortal) SetNetworks(v []CaptivePortalNetwork)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


