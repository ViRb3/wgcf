# GetSourceDevice200ResponseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Interface** | [**GetSourceDevice200ResponseConfigInterface**](GetSourceDevice_200_Response_config_interface.md) |  | 
**Peers** | [**[]GetSourceDevice200ResponseConfigPeers**](GetSourceDevice200ResponseConfigPeers.md) |  | 
**Services** | [**GetSourceDevice200ResponseConfigServices**](GetSourceDevice_200_Response_config_services.md) |  | 

## Methods

### NewGetSourceDevice200ResponseConfig

`func NewGetSourceDevice200ResponseConfig(clientId string, interface_ GetSourceDevice200ResponseConfigInterface, peers []GetSourceDevice200ResponseConfigPeers, services GetSourceDevice200ResponseConfigServices, ) *GetSourceDevice200ResponseConfig`

NewGetSourceDevice200ResponseConfig instantiates a new GetSourceDevice200ResponseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceDevice200ResponseConfigWithDefaults

`func NewGetSourceDevice200ResponseConfigWithDefaults() *GetSourceDevice200ResponseConfig`

NewGetSourceDevice200ResponseConfigWithDefaults instantiates a new GetSourceDevice200ResponseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GetSourceDevice200ResponseConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetSourceDevice200ResponseConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetSourceDevice200ResponseConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetInterface

`func (o *GetSourceDevice200ResponseConfig) GetInterface() GetSourceDevice200ResponseConfigInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetSourceDevice200ResponseConfig) GetInterfaceOk() (*GetSourceDevice200ResponseConfigInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetSourceDevice200ResponseConfig) SetInterface(v GetSourceDevice200ResponseConfigInterface)`

SetInterface sets Interface field to given value.


### GetPeers

`func (o *GetSourceDevice200ResponseConfig) GetPeers() []GetSourceDevice200ResponseConfigPeers`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *GetSourceDevice200ResponseConfig) GetPeersOk() (*[]GetSourceDevice200ResponseConfigPeers, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *GetSourceDevice200ResponseConfig) SetPeers(v []GetSourceDevice200ResponseConfigPeers)`

SetPeers sets Peers field to given value.


### GetServices

`func (o *GetSourceDevice200ResponseConfig) GetServices() GetSourceDevice200ResponseConfigServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetSourceDevice200ResponseConfig) GetServicesOk() (*GetSourceDevice200ResponseConfigServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetSourceDevice200ResponseConfig) SetServices(v GetSourceDevice200ResponseConfigServices)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


