# IPv4Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Netmask** | **string** |  | 

## Methods

### NewIPv4Network

`func NewIPv4Network(address string, netmask string, ) *IPv4Network`

NewIPv4Network instantiates a new IPv4Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4NetworkWithDefaults

`func NewIPv4NetworkWithDefaults() *IPv4Network`

NewIPv4NetworkWithDefaults instantiates a new IPv4Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPv4Network) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPv4Network) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPv4Network) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetmask

`func (o *IPv4Network) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IPv4Network) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IPv4Network) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


