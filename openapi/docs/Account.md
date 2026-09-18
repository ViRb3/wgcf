# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **string** |  |
**Created** | Pointer to **time.Time** |  | [optional]
**Id** | **string** |  |
**License** | Pointer to **string** |  | [optional]
**PlusExpiresAt** | Pointer to **time.Time** |  | [optional]
**Managed** | Pointer to **string** |  | [optional]
**PremiumData** | Pointer to **int64** |  | [optional]
**Quota** | Pointer to **int64** |  | [optional]
**ReferralCount** | Pointer to **int32** |  | [optional]
**ReferralRenewalCountdown** | Pointer to **int32** |  | [optional]
**Role** | Pointer to **string** |  | [optional]
**Ttl** | Pointer to **time.Time** |  | [optional]
**Updated** | Pointer to **time.Time** |  | [optional]
**Usage** | Pointer to **int64** |  | [optional]
**WarpPlus** | Pointer to **bool** |  | [optional]

## Methods

### NewAccount

`func NewAccount(accountType string, id string, ) *Account`

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

`func (o *Account) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Account) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Account) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

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

### HasLicense

`func (o *Account) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPlusExpiresAt

`func (o *Account) GetPlusExpiresAt() time.Time`

GetPlusExpiresAt returns the PlusExpiresAt field if non-nil, zero value otherwise.

### GetPlusExpiresAtOk

`func (o *Account) GetPlusExpiresAtOk() (*time.Time, bool)`

GetPlusExpiresAtOk returns a tuple with the PlusExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlusExpiresAt

`func (o *Account) SetPlusExpiresAt(v time.Time)`

SetPlusExpiresAt sets PlusExpiresAt field to given value.

### HasPlusExpiresAt

`func (o *Account) HasPlusExpiresAt() bool`

HasPlusExpiresAt returns a boolean if a field has been set.

### GetManaged

`func (o *Account) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Account) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Account) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Account) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetPremiumData

`func (o *Account) GetPremiumData() int64`

GetPremiumData returns the PremiumData field if non-nil, zero value otherwise.

### GetPremiumDataOk

`func (o *Account) GetPremiumDataOk() (*int64, bool)`

GetPremiumDataOk returns a tuple with the PremiumData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumData

`func (o *Account) SetPremiumData(v int64)`

SetPremiumData sets PremiumData field to given value.

### HasPremiumData

`func (o *Account) HasPremiumData() bool`

HasPremiumData returns a boolean if a field has been set.

### GetQuota

`func (o *Account) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Account) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Account) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Account) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReferralCount

`func (o *Account) GetReferralCount() int32`

GetReferralCount returns the ReferralCount field if non-nil, zero value otherwise.

### GetReferralCountOk

`func (o *Account) GetReferralCountOk() (*int32, bool)`

GetReferralCountOk returns a tuple with the ReferralCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCount

`func (o *Account) SetReferralCount(v int32)`

SetReferralCount sets ReferralCount field to given value.

### HasReferralCount

`func (o *Account) HasReferralCount() bool`

HasReferralCount returns a boolean if a field has been set.

### GetReferralRenewalCountdown

`func (o *Account) GetReferralRenewalCountdown() int32`

GetReferralRenewalCountdown returns the ReferralRenewalCountdown field if non-nil, zero value otherwise.

### GetReferralRenewalCountdownOk

`func (o *Account) GetReferralRenewalCountdownOk() (*int32, bool)`

GetReferralRenewalCountdownOk returns a tuple with the ReferralRenewalCountdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRenewalCountdown

`func (o *Account) SetReferralRenewalCountdown(v int32)`

SetReferralRenewalCountdown sets ReferralRenewalCountdown field to given value.

### HasReferralRenewalCountdown

`func (o *Account) HasReferralRenewalCountdown() bool`

HasReferralRenewalCountdown returns a boolean if a field has been set.

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

### HasRole

`func (o *Account) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTtl

`func (o *Account) GetTtl() time.Time`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Account) GetTtlOk() (*time.Time, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Account) SetTtl(v time.Time)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Account) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUpdated

`func (o *Account) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Account) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Account) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Account) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsage

`func (o *Account) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Account) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Account) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Account) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

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

### HasWarpPlus

`func (o *Account) HasWarpPlus() bool`

HasWarpPlus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


