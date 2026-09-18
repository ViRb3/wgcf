# ClientConfigDenylistInnerNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V4** | [**[]IPv4Network**](IPv4Network.md) |  |
**V6** | [**[]IPv6Network**](IPv6Network.md) |  |

## Methods

### NewClientConfigDenylistInnerNetworks

`func NewClientConfigDenylistInnerNetworks(v4 []IPv4Network, v6 []IPv6Network, ) *ClientConfigDenylistInnerNetworks`

NewClientConfigDenylistInnerNetworks instantiates a new ClientConfigDenylistInnerNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigDenylistInnerNetworksWithDefaults

`func NewClientConfigDenylistInnerNetworksWithDefaults() *ClientConfigDenylistInnerNetworks`

NewClientConfigDenylistInnerNetworksWithDefaults instantiates a new ClientConfigDenylistInnerNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV4

`func (o *ClientConfigDenylistInnerNetworks) GetV4() []IPv4Network`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *ClientConfigDenylistInnerNetworks) GetV4Ok() (*[]IPv4Network, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *ClientConfigDenylistInnerNetworks) SetV4(v []IPv4Network)`

SetV4 sets V4 field to given value.


### GetV6

`func (o *ClientConfigDenylistInnerNetworks) GetV6() []IPv6Network`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *ClientConfigDenylistInnerNetworks) GetV6Ok() (*[]IPv6Network, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *ClientConfigDenylistInnerNetworks) SetV6(v []IPv6Network)`

SetV6 sets V6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


