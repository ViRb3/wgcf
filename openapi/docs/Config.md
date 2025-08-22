# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Interface** | [**ConfigInterface**](ConfigInterface.md) |  | 
**Peers** | [**[]Peer**](Peer.md) |  | 
**Services** | [**ConfigServices**](ConfigServices.md) |  | 

## Methods

### NewConfig

`func NewConfig(clientId string, interface_ ConfigInterface, peers []Peer, services ConfigServices, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Config) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Config) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Config) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetInterface

`func (o *Config) GetInterface() ConfigInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Config) GetInterfaceOk() (*ConfigInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Config) SetInterface(v ConfigInterface)`

SetInterface sets Interface field to given value.


### GetPeers

`func (o *Config) GetPeers() []Peer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *Config) GetPeersOk() (*[]Peer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *Config) SetPeers(v []Peer)`

SetPeers sets Peers field to given value.


### GetServices

`func (o *Config) GetServices() ConfigServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Config) GetServicesOk() (*ConfigServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Config) SetServices(v ConfigServices)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


