# DexData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**DexDataKind**](DexDataKind.md) |  |
**Method** | Pointer to [**HTTPRequestMethod**](HTTPRequestMethod.md) |  | [optional]
**Host** | **string** |  |

## Methods

### NewDexData

`func NewDexData(kind DexDataKind, host string, ) *DexData`

NewDexData instantiates a new DexData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexDataWithDefaults

`func NewDexDataWithDefaults() *DexData`

NewDexDataWithDefaults instantiates a new DexData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DexData) GetKind() DexDataKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DexData) GetKindOk() (*DexDataKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DexData) SetKind(v DexDataKind)`

SetKind sets Kind field to given value.


### GetMethod

`func (o *DexData) GetMethod() HTTPRequestMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DexData) GetMethodOk() (*HTTPRequestMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DexData) SetMethod(v HTTPRequestMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *DexData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetHost

`func (o *DexData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DexData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DexData) SetHost(v string)`

SetHost sets Host field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


