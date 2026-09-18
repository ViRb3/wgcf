# FallbackDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suffix** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**DnsServer** | Pointer to **[]string** |  | [optional]

## Methods

### NewFallbackDomain

`func NewFallbackDomain() *FallbackDomain`

NewFallbackDomain instantiates a new FallbackDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackDomainWithDefaults

`func NewFallbackDomainWithDefaults() *FallbackDomain`

NewFallbackDomainWithDefaults instantiates a new FallbackDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuffix

`func (o *FallbackDomain) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *FallbackDomain) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *FallbackDomain) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *FallbackDomain) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetDescription

`func (o *FallbackDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FallbackDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FallbackDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FallbackDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsServer

`func (o *FallbackDomain) GetDnsServer() []string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *FallbackDomain) GetDnsServerOk() (*[]string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *FallbackDomain) SetDnsServer(v []string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *FallbackDomain) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


