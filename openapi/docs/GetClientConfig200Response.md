# GetClientConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptivePortal** | [**[]GetClientConfig200ResponseCaptivePortal**](GetClientConfig200ResponseCaptivePortal.md) |  | 
**Denylist** | [**[]GetClientConfig200ResponseDenylist**](GetClientConfig200ResponseDenylist.md) |  | 
**PremiumDataBytes** | **float32** |  | 
**ReferralRewardBytes** | **float32** |  | 

## Methods

### NewGetClientConfig200Response

`func NewGetClientConfig200Response(captivePortal []GetClientConfig200ResponseCaptivePortal, denylist []GetClientConfig200ResponseDenylist, premiumDataBytes float32, referralRewardBytes float32, ) *GetClientConfig200Response`

NewGetClientConfig200Response instantiates a new GetClientConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientConfig200ResponseWithDefaults

`func NewGetClientConfig200ResponseWithDefaults() *GetClientConfig200Response`

NewGetClientConfig200ResponseWithDefaults instantiates a new GetClientConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptivePortal

`func (o *GetClientConfig200Response) GetCaptivePortal() []GetClientConfig200ResponseCaptivePortal`

GetCaptivePortal returns the CaptivePortal field if non-nil, zero value otherwise.

### GetCaptivePortalOk

`func (o *GetClientConfig200Response) GetCaptivePortalOk() (*[]GetClientConfig200ResponseCaptivePortal, bool)`

GetCaptivePortalOk returns a tuple with the CaptivePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptivePortal

`func (o *GetClientConfig200Response) SetCaptivePortal(v []GetClientConfig200ResponseCaptivePortal)`

SetCaptivePortal sets CaptivePortal field to given value.


### GetDenylist

`func (o *GetClientConfig200Response) GetDenylist() []GetClientConfig200ResponseDenylist`

GetDenylist returns the Denylist field if non-nil, zero value otherwise.

### GetDenylistOk

`func (o *GetClientConfig200Response) GetDenylistOk() (*[]GetClientConfig200ResponseDenylist, bool)`

GetDenylistOk returns a tuple with the Denylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenylist

`func (o *GetClientConfig200Response) SetDenylist(v []GetClientConfig200ResponseDenylist)`

SetDenylist sets Denylist field to given value.


### GetPremiumDataBytes

`func (o *GetClientConfig200Response) GetPremiumDataBytes() float32`

GetPremiumDataBytes returns the PremiumDataBytes field if non-nil, zero value otherwise.

### GetPremiumDataBytesOk

`func (o *GetClientConfig200Response) GetPremiumDataBytesOk() (*float32, bool)`

GetPremiumDataBytesOk returns a tuple with the PremiumDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumDataBytes

`func (o *GetClientConfig200Response) SetPremiumDataBytes(v float32)`

SetPremiumDataBytes sets PremiumDataBytes field to given value.


### GetReferralRewardBytes

`func (o *GetClientConfig200Response) GetReferralRewardBytes() float32`

GetReferralRewardBytes returns the ReferralRewardBytes field if non-nil, zero value otherwise.

### GetReferralRewardBytesOk

`func (o *GetClientConfig200Response) GetReferralRewardBytesOk() (*float32, bool)`

GetReferralRewardBytesOk returns a tuple with the ReferralRewardBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRewardBytes

`func (o *GetClientConfig200Response) SetReferralRewardBytes(v float32)`

SetReferralRewardBytes sets ReferralRewardBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


