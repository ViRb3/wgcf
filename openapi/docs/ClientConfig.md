# ClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denylist** | [**[]ClientConfigDenylistInner**](ClientConfigDenylistInner.md) |  |
**CaptivePortal** | [**[]CaptivePortal**](CaptivePortal.md) |  |
**PremiumDataBytes** | **int64** |  |
**ReferralRewardBytes** | **int64** |  |

## Methods

### NewClientConfig

`func NewClientConfig(denylist []ClientConfigDenylistInner, captivePortal []CaptivePortal, premiumDataBytes int64, referralRewardBytes int64, ) *ClientConfig`

NewClientConfig instantiates a new ClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigWithDefaults

`func NewClientConfigWithDefaults() *ClientConfig`

NewClientConfigWithDefaults instantiates a new ClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenylist

`func (o *ClientConfig) GetDenylist() []ClientConfigDenylistInner`

GetDenylist returns the Denylist field if non-nil, zero value otherwise.

### GetDenylistOk

`func (o *ClientConfig) GetDenylistOk() (*[]ClientConfigDenylistInner, bool)`

GetDenylistOk returns a tuple with the Denylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenylist

`func (o *ClientConfig) SetDenylist(v []ClientConfigDenylistInner)`

SetDenylist sets Denylist field to given value.


### GetCaptivePortal

`func (o *ClientConfig) GetCaptivePortal() []CaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *ClientConfig) GetCaptivePortalOk() (*[]CaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *ClientConfig) SetCaptivePortal(v []CaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.


### GetPremiumDataBytes

`func (o *ClientConfig) GetPremiumDataBytes() int64`

GetPremiumDataBytes returns the PremiumDataBytes field if non-nil, zero value otherwise.

### GetPremiumDataBytesOk

`func (o *ClientConfig) GetPremiumDataBytesOk() (*int64, bool)`

GetPremiumDataBytesOk returns a tuple with the PremiumDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumDataBytes

`func (o *ClientConfig) SetPremiumDataBytes(v int64)`

SetPremiumDataBytes sets PremiumDataBytes field to given value.


### GetReferralRewardBytes

`func (o *ClientConfig) GetReferralRewardBytes() int64`

GetReferralRewardBytes returns the ReferralRewardBytes field if non-nil, zero value otherwise.

### GetReferralRewardBytesOk

`func (o *ClientConfig) GetReferralRewardBytesOk() (*int64, bool)`

GetReferralRewardBytesOk returns a tuple with the ReferralRewardBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRewardBytes

`func (o *ClientConfig) SetReferralRewardBytes(v int64)`

SetReferralRewardBytes sets ReferralRewardBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


