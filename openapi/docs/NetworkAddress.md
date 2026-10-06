# NetworkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V4** | **string** |  | 
**V6** | **string** |  | 

## Methods

### NewNetworkAddress

`func NewNetworkAddress(v4 string, v6 string, ) *NetworkAddress`

NewNetworkAddress instantiates a new NetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAddressWithDefaults

`func NewNetworkAddressWithDefaults() *NetworkAddress`

NewNetworkAddressWithDefaults instantiates a new NetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV4

`func (o *NetworkAddress) GetV4() string`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *NetworkAddress) GetV4Ok() (*string, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *NetworkAddress) SetV4(v string)`

SetV4 sets V4 field to given value.


### GetV6

`func (o *NetworkAddress) GetV6() string`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *NetworkAddress) GetV6Ok() (*string, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *NetworkAddress) SetV6(v string)`

SetV6 sets V6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


