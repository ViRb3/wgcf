# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** |  | [optional]
**Enable** | Pointer to **bool** |  | [optional]
**ServiceMode** | Pointer to [**AppMode**](AppMode.md) |  | [optional]
**ServiceModeV2** | Pointer to [**ServiceMode**](ServiceMode.md) |  | [optional]
**GatewayUniqueId** | Pointer to **string** |  | [optional]
**SupportUrl** | Pointer to **string** |  | [optional]
**Onboarding** | Pointer to **bool** |  | [optional]
**Exclude** | Pointer to [**[]PolicyExclude**](PolicyExclude.md) |  | [optional]
**Include** | Pointer to [**[]PolicyExclude**](PolicyExclude.md) |  | [optional]
**FallbackDomains** | Pointer to [**[]FallbackDomain**](FallbackDomain.md) |  | [optional]
**SwitchLocked** | Pointer to **bool** |  | [optional]
**AutoConnect** | Pointer to **int32** |  | [optional]
**ManagedMode** | Pointer to **bool** |  | [optional]
**AllowModeSwitch** | Pointer to **bool** |  | [optional]
**AuthClientId** | Pointer to **string** |  | [optional]
**AuthClientSecret** | Pointer to **string** |  | [optional]
**AllowedToLeave** | Pointer to **bool** |  | [optional]
**OverrideDohEndpoint** | Pointer to **string** |  | [optional]
**OverrideWarpEndpoint** | Pointer to **string** |  | [optional]
**OverrideApiEndpoint** | Pointer to **string** |  | [optional]
**UniqueClientId** | Pointer to **string** |  | [optional]
**TunnelProtocol** | Pointer to **string** |  | [optional]
**PolicyId** | Pointer to **string** |  | [optional]
**PostQuantum** | Pointer to [**PostQuantumSupport**](PostQuantumSupport.md) |  | [optional]
**AlwaysExclude** | Pointer to [**[]PolicyIP**](PolicyIP.md) |  | [optional]
**AlwaysInclude** | Pointer to [**[]PolicyIP**](PolicyIP.md) |  | [optional]

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *Policy) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Policy) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Policy) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Policy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEnable

`func (o *Policy) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Policy) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Policy) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Policy) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetServiceMode

`func (o *Policy) GetServiceMode() AppMode`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *Policy) GetServiceModeOk() (*AppMode, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *Policy) SetServiceMode(v AppMode)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *Policy) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServiceModeV2

`func (o *Policy) GetServiceModeV2() ServiceMode`

GetServiceModeV2 returns the ServiceModeV2 field if non-nil, zero value otherwise.

### GetServiceModeV2Ok

`func (o *Policy) GetServiceModeV2Ok() (*ServiceMode, bool)`

GetServiceModeV2Ok returns a tuple with the ServiceModeV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModeV2

`func (o *Policy) SetServiceModeV2(v ServiceMode)`

SetServiceModeV2 sets ServiceModeV2 field to given value.

### HasServiceModeV2

`func (o *Policy) HasServiceModeV2() bool`

HasServiceModeV2 returns a boolean if a field has been set.

### GetGatewayUniqueId

`func (o *Policy) GetGatewayUniqueId() string`

GetGatewayUniqueId returns the GatewayUniqueId field if non-nil, zero value otherwise.

### GetGatewayUniqueIdOk

`func (o *Policy) GetGatewayUniqueIdOk() (*string, bool)`

GetGatewayUniqueIdOk returns a tuple with the GatewayUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUniqueId

`func (o *Policy) SetGatewayUniqueId(v string)`

SetGatewayUniqueId sets GatewayUniqueId field to given value.

### HasGatewayUniqueId

`func (o *Policy) HasGatewayUniqueId() bool`

HasGatewayUniqueId returns a boolean if a field has been set.

### GetSupportUrl

`func (o *Policy) GetSupportUrl() string`

GetSupportUrl returns the SupportUrl field if non-nil, zero value otherwise.

### GetSupportUrlOk

`func (o *Policy) GetSupportUrlOk() (*string, bool)`

GetSupportUrlOk returns a tuple with the SupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUrl

`func (o *Policy) SetSupportUrl(v string)`

SetSupportUrl sets SupportUrl field to given value.

### HasSupportUrl

`func (o *Policy) HasSupportUrl() bool`

HasSupportUrl returns a boolean if a field has been set.

### GetOnboarding

`func (o *Policy) GetOnboarding() bool`

GetOnboarding returns the Onboarding field if non-nil, zero value otherwise.

### GetOnboardingOk

`func (o *Policy) GetOnboardingOk() (*bool, bool)`

GetOnboardingOk returns a tuple with the Onboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboarding

`func (o *Policy) SetOnboarding(v bool)`

SetOnboarding sets Onboarding field to given value.

### HasOnboarding

`func (o *Policy) HasOnboarding() bool`

HasOnboarding returns a boolean if a field has been set.

### GetExclude

`func (o *Policy) GetExclude() []PolicyExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Policy) GetExcludeOk() (*[]PolicyExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Policy) SetExclude(v []PolicyExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Policy) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *Policy) GetInclude() []PolicyExclude`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *Policy) GetIncludeOk() (*[]PolicyExclude, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *Policy) SetInclude(v []PolicyExclude)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *Policy) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetFallbackDomains

`func (o *Policy) GetFallbackDomains() []FallbackDomain`

GetFallbackDomains returns the FallbackDomains field if non-nil, zero value otherwise.

### GetFallbackDomainsOk

`func (o *Policy) GetFallbackDomainsOk() (*[]FallbackDomain, bool)`

GetFallbackDomainsOk returns a tuple with the FallbackDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDomains

`func (o *Policy) SetFallbackDomains(v []FallbackDomain)`

SetFallbackDomains sets FallbackDomains field to given value.

### HasFallbackDomains

`func (o *Policy) HasFallbackDomains() bool`

HasFallbackDomains returns a boolean if a field has been set.

### GetSwitchLocked

`func (o *Policy) GetSwitchLocked() bool`

GetSwitchLocked returns the SwitchLocked field if non-nil, zero value otherwise.

### GetSwitchLockedOk

`func (o *Policy) GetSwitchLockedOk() (*bool, bool)`

GetSwitchLockedOk returns a tuple with the SwitchLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchLocked

`func (o *Policy) SetSwitchLocked(v bool)`

SetSwitchLocked sets SwitchLocked field to given value.

### HasSwitchLocked

`func (o *Policy) HasSwitchLocked() bool`

HasSwitchLocked returns a boolean if a field has been set.

### GetAutoConnect

`func (o *Policy) GetAutoConnect() int32`

GetAutoConnect returns the AutoConnect field if non-nil, zero value otherwise.

### GetAutoConnectOk

`func (o *Policy) GetAutoConnectOk() (*int32, bool)`

GetAutoConnectOk returns a tuple with the AutoConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConnect

`func (o *Policy) SetAutoConnect(v int32)`

SetAutoConnect sets AutoConnect field to given value.

### HasAutoConnect

`func (o *Policy) HasAutoConnect() bool`

HasAutoConnect returns a boolean if a field has been set.

### GetManagedMode

`func (o *Policy) GetManagedMode() bool`

GetManagedMode returns the ManagedMode field if non-nil, zero value otherwise.

### GetManagedModeOk

`func (o *Policy) GetManagedModeOk() (*bool, bool)`

GetManagedModeOk returns a tuple with the ManagedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMode

`func (o *Policy) SetManagedMode(v bool)`

SetManagedMode sets ManagedMode field to given value.

### HasManagedMode

`func (o *Policy) HasManagedMode() bool`

HasManagedMode returns a boolean if a field has been set.

### GetAllowModeSwitch

`func (o *Policy) GetAllowModeSwitch() bool`

GetAllowModeSwitch returns the AllowModeSwitch field if non-nil, zero value otherwise.

### GetAllowModeSwitchOk

`func (o *Policy) GetAllowModeSwitchOk() (*bool, bool)`

GetAllowModeSwitchOk returns a tuple with the AllowModeSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowModeSwitch

`func (o *Policy) SetAllowModeSwitch(v bool)`

SetAllowModeSwitch sets AllowModeSwitch field to given value.

### HasAllowModeSwitch

`func (o *Policy) HasAllowModeSwitch() bool`

HasAllowModeSwitch returns a boolean if a field has been set.

### GetAuthClientId

`func (o *Policy) GetAuthClientId() string`

GetAuthClientId returns the AuthClientId field if non-nil, zero value otherwise.

### GetAuthClientIdOk

`func (o *Policy) GetAuthClientIdOk() (*string, bool)`

GetAuthClientIdOk returns a tuple with the AuthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthClientId

`func (o *Policy) SetAuthClientId(v string)`

SetAuthClientId sets AuthClientId field to given value.

### HasAuthClientId

`func (o *Policy) HasAuthClientId() bool`

HasAuthClientId returns a boolean if a field has been set.

### GetAuthClientSecret

`func (o *Policy) GetAuthClientSecret() string`

GetAuthClientSecret returns the AuthClientSecret field if non-nil, zero value otherwise.

### GetAuthClientSecretOk

`func (o *Policy) GetAuthClientSecretOk() (*string, bool)`

GetAuthClientSecretOk returns a tuple with the AuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthClientSecret

`func (o *Policy) SetAuthClientSecret(v string)`

SetAuthClientSecret sets AuthClientSecret field to given value.

### HasAuthClientSecret

`func (o *Policy) HasAuthClientSecret() bool`

HasAuthClientSecret returns a boolean if a field has been set.

### GetAllowedToLeave

`func (o *Policy) GetAllowedToLeave() bool`

GetAllowedToLeave returns the AllowedToLeave field if non-nil, zero value otherwise.

### GetAllowedToLeaveOk

`func (o *Policy) GetAllowedToLeaveOk() (*bool, bool)`

GetAllowedToLeaveOk returns a tuple with the AllowedToLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToLeave

`func (o *Policy) SetAllowedToLeave(v bool)`

SetAllowedToLeave sets AllowedToLeave field to given value.

### HasAllowedToLeave

`func (o *Policy) HasAllowedToLeave() bool`

HasAllowedToLeave returns a boolean if a field has been set.

### GetOverrideDohEndpoint

`func (o *Policy) GetOverrideDohEndpoint() string`

GetOverrideDohEndpoint returns the OverrideDohEndpoint field if non-nil, zero value otherwise.

### GetOverrideDohEndpointOk

`func (o *Policy) GetOverrideDohEndpointOk() (*string, bool)`

GetOverrideDohEndpointOk returns a tuple with the OverrideDohEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDohEndpoint

`func (o *Policy) SetOverrideDohEndpoint(v string)`

SetOverrideDohEndpoint sets OverrideDohEndpoint field to given value.

### HasOverrideDohEndpoint

`func (o *Policy) HasOverrideDohEndpoint() bool`

HasOverrideDohEndpoint returns a boolean if a field has been set.

### GetOverrideWarpEndpoint

`func (o *Policy) GetOverrideWarpEndpoint() string`

GetOverrideWarpEndpoint returns the OverrideWarpEndpoint field if non-nil, zero value otherwise.

### GetOverrideWarpEndpointOk

`func (o *Policy) GetOverrideWarpEndpointOk() (*string, bool)`

GetOverrideWarpEndpointOk returns a tuple with the OverrideWarpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideWarpEndpoint

`func (o *Policy) SetOverrideWarpEndpoint(v string)`

SetOverrideWarpEndpoint sets OverrideWarpEndpoint field to given value.

### HasOverrideWarpEndpoint

`func (o *Policy) HasOverrideWarpEndpoint() bool`

HasOverrideWarpEndpoint returns a boolean if a field has been set.

### GetOverrideApiEndpoint

`func (o *Policy) GetOverrideApiEndpoint() string`

GetOverrideApiEndpoint returns the OverrideApiEndpoint field if non-nil, zero value otherwise.

### GetOverrideApiEndpointOk

`func (o *Policy) GetOverrideApiEndpointOk() (*string, bool)`

GetOverrideApiEndpointOk returns a tuple with the OverrideApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideApiEndpoint

`func (o *Policy) SetOverrideApiEndpoint(v string)`

SetOverrideApiEndpoint sets OverrideApiEndpoint field to given value.

### HasOverrideApiEndpoint

`func (o *Policy) HasOverrideApiEndpoint() bool`

HasOverrideApiEndpoint returns a boolean if a field has been set.

### GetUniqueClientId

`func (o *Policy) GetUniqueClientId() string`

GetUniqueClientId returns the UniqueClientId field if non-nil, zero value otherwise.

### GetUniqueClientIdOk

`func (o *Policy) GetUniqueClientIdOk() (*string, bool)`

GetUniqueClientIdOk returns a tuple with the UniqueClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClientId

`func (o *Policy) SetUniqueClientId(v string)`

SetUniqueClientId sets UniqueClientId field to given value.

### HasUniqueClientId

`func (o *Policy) HasUniqueClientId() bool`

HasUniqueClientId returns a boolean if a field has been set.

### GetTunnelProtocol

`func (o *Policy) GetTunnelProtocol() string`

GetTunnelProtocol returns the TunnelProtocol field if non-nil, zero value otherwise.

### GetTunnelProtocolOk

`func (o *Policy) GetTunnelProtocolOk() (*string, bool)`

GetTunnelProtocolOk returns a tuple with the TunnelProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelProtocol

`func (o *Policy) SetTunnelProtocol(v string)`

SetTunnelProtocol sets TunnelProtocol field to given value.

### HasTunnelProtocol

`func (o *Policy) HasTunnelProtocol() bool`

HasTunnelProtocol returns a boolean if a field has been set.

### GetPolicyId

`func (o *Policy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Policy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Policy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Policy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPostQuantum

`func (o *Policy) GetPostQuantum() PostQuantumSupport`

GetPostQuantum returns the PostQuantum field if non-nil, zero value otherwise.

### GetPostQuantumOk

`func (o *Policy) GetPostQuantumOk() (*PostQuantumSupport, bool)`

GetPostQuantumOk returns a tuple with the PostQuantum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostQuantum

`func (o *Policy) SetPostQuantum(v PostQuantumSupport)`

SetPostQuantum sets PostQuantum field to given value.

### HasPostQuantum

`func (o *Policy) HasPostQuantum() bool`

HasPostQuantum returns a boolean if a field has been set.

### GetAlwaysExclude

`func (o *Policy) GetAlwaysExclude() []PolicyIP`

GetAlwaysExclude returns the AlwaysExclude field if non-nil, zero value otherwise.

### GetAlwaysExcludeOk

`func (o *Policy) GetAlwaysExcludeOk() (*[]PolicyIP, bool)`

GetAlwaysExcludeOk returns a tuple with the AlwaysExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysExclude

`func (o *Policy) SetAlwaysExclude(v []PolicyIP)`

SetAlwaysExclude sets AlwaysExclude field to given value.

### HasAlwaysExclude

`func (o *Policy) HasAlwaysExclude() bool`

HasAlwaysExclude returns a boolean if a field has been set.

### GetAlwaysInclude

`func (o *Policy) GetAlwaysInclude() []PolicyIP`

GetAlwaysInclude returns the AlwaysInclude field if non-nil, zero value otherwise.

### GetAlwaysIncludeOk

`func (o *Policy) GetAlwaysIncludeOk() (*[]PolicyIP, bool)`

GetAlwaysIncludeOk returns a tuple with the AlwaysInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysInclude

`func (o *Policy) SetAlwaysInclude(v []PolicyIP)`

SetAlwaysInclude sets AlwaysInclude field to given value.

### HasAlwaysInclude

`func (o *Policy) HasAlwaysInclude() bool`

HasAlwaysInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


