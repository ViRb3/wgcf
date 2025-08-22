# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **string** |  | 
**Created** | **string** |  | 
**Id** | **string** |  | 
**License** | **string** |  | 
**PremiumData** | **float32** |  | 
**Quota** | **float32** |  | 
**ReferralCount** | **float32** |  | 
**ReferralRenewalCountdown** | **float32** |  | 
**Role** | **string** |  | 
**Updated** | **string** |  | 
**WarpPlus** | **bool** |  | 
**Usage** | Pointer to **float32** |  | [optional] 

## Methods

### NewAccount

`func NewAccount(accountType string, created string, id string, license string, premiumData float32, quota float32, referralCount float32, referralRenewalCountdown float32, role string, updated string, warpPlus bool, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Account) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Account) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Account) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetCreated

`func (o *Account) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Account) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.


### GetLicense

`func (o *Account) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Account) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Account) SetLicense(v string)`

SetLicense sets License field to given value.


### GetPremiumData

`func (o *Account) GetPremiumData() float32`

GetPremiumData returns the PremiumData field if non-nil, zero value otherwise.

### GetPremiumDataOk

`func (o *Account) GetPremiumDataOk() (*float32, bool)`

GetPremiumDataOk returns a tuple with the PremiumData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumData

`func (o *Account) SetPremiumData(v float32)`

SetPremiumData sets PremiumData field to given value.


### GetQuota

`func (o *Account) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Account) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Account) SetQuota(v float32)`

SetQuota sets Quota field to given value.


### GetReferralCount

`func (o *Account) GetReferralCount() float32`

GetReferralCount returns the ReferralCount field if non-nil, zero value otherwise.

### GetReferralCountOk

`func (o *Account) GetReferralCountOk() (*float32, bool)`

GetReferralCountOk returns a tuple with the ReferralCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCount

`func (o *Account) SetReferralCount(v float32)`

SetReferralCount sets ReferralCount field to given value.


### GetReferralRenewalCountdown

`func (o *Account) GetReferralRenewalCountdown() float32`

GetReferralRenewalCountdown returns the ReferralRenewalCountdown field if non-nil, zero value otherwise.

### GetReferralRenewalCountdownOk

`func (o *Account) GetReferralRenewalCountdownOk() (*float32, bool)`

GetReferralRenewalCountdownOk returns a tuple with the ReferralRenewalCountdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRenewalCountdown

`func (o *Account) SetReferralRenewalCountdown(v float32)`

SetReferralRenewalCountdown sets ReferralRenewalCountdown field to given value.


### GetRole

`func (o *Account) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Account) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Account) SetRole(v string)`

SetRole sets Role field to given value.


### GetUpdated

`func (o *Account) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Account) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Account) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetWarpPlus

`func (o *Account) GetWarpPlus() bool`

GetWarpPlus returns the WarpPlus field if non-nil, zero value otherwise.

### GetWarpPlusOk

`func (o *Account) GetWarpPlusOk() (*bool, bool)`

GetWarpPlusOk returns a tuple with the WarpPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarpPlus

`func (o *Account) SetWarpPlus(v bool)`

SetWarpPlus sets WarpPlus field to given value.


### GetUsage

`func (o *Account) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Account) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Account) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Account) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


