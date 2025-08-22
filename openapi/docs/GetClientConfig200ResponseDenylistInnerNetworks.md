# GetClientConfig200ResponseDenylistInnerNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V4** | [**[]IPv4Network**](IPv4Network.md) |  | 
**V6** | [**[]IPv6Network**](IPv6Network.md) |  | 

## Methods

### NewGetClientConfig200ResponseDenylistInnerNetworks

`func NewGetClientConfig200ResponseDenylistInnerNetworks(v4 []IPv4Network, v6 []IPv6Network, ) *GetClientConfig200ResponseDenylistInnerNetworks`

NewGetClientConfig200ResponseDenylistInnerNetworks instantiates a new GetClientConfig200ResponseDenylistInnerNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientConfig200ResponseDenylistInnerNetworksWithDefaults

`func NewGetClientConfig200ResponseDenylistInnerNetworksWithDefaults() *GetClientConfig200ResponseDenylistInnerNetworks`

NewGetClientConfig200ResponseDenylistInnerNetworksWithDefaults instantiates a new GetClientConfig200ResponseDenylistInnerNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV4

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) GetV4() []IPv4Network`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) GetV4Ok() (*[]IPv4Network, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) SetV4(v []IPv4Network)`

SetV4 sets V4 field to given value.


### GetV6

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) GetV6() []IPv6Network`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) GetV6Ok() (*[]IPv6Network, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *GetClientConfig200ResponseDenylistInnerNetworks) SetV6(v []IPv6Network)`

SetV6 sets V6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


