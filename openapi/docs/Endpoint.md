# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**V4** | **string** |  | 
**V6** | **string** |  | 

## Methods

### NewEndpoint

`func NewEndpoint(host string, v4 string, v6 string, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Endpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Endpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Endpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetV4

`func (o *Endpoint) GetV4() string`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *Endpoint) GetV4Ok() (*string, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *Endpoint) SetV4(v string)`

SetV4 sets V4 field to given value.


### GetV6

`func (o *Endpoint) GetV6() string`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *Endpoint) GetV6Ok() (*string, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *Endpoint) SetV6(v string)`

SetV6 sets V6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


