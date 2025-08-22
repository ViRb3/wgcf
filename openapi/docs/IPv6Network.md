# IPv6Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Prefix** | **float32** |  | 

## Methods

### NewIPv6Network

`func NewIPv6Network(address string, prefix float32, ) *IPv6Network`

NewIPv6Network instantiates a new IPv6Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6NetworkWithDefaults

`func NewIPv6NetworkWithDefaults() *IPv6Network`

NewIPv6NetworkWithDefaults instantiates a new IPv6Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPv6Network) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPv6Network) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPv6Network) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPrefix

`func (o *IPv6Network) GetPrefix() float32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IPv6Network) GetPrefixOk() (*float32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IPv6Network) SetPrefix(v float32)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


