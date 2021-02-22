# Register200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**GetSourceDevice200ResponseAccount**](GetSourceDevice_200_Response_account.md) |  | 
**Config** | [**GetSourceDevice200ResponseConfig**](GetSourceDevice_200_Response_config.md) |  | 
**Created** | **string** |  | 
**Enabled** | **bool** |  | 
**FcmToken** | **string** |  | 
**Id** | **string** |  | 
**InstallId** | **string** |  | 
**Key** | **string** |  | 
**Locale** | **string** |  | 
**Model** | **string** |  | 
**Name** | **string** |  | 
**Place** | **float32** |  | 
**Token** | **string** |  | 
**Tos** | **string** |  | 
**Type** | **string** |  | 
**Updated** | **string** |  | 
**WaitlistEnabled** | **bool** |  | 
**WarpEnabled** | **bool** |  | 

## Methods

### NewRegister200Response

`func NewRegister200Response(account GetSourceDevice200ResponseAccount, config GetSourceDevice200ResponseConfig, created string, enabled bool, fcmToken string, id string, installId string, key string, locale string, model string, name string, place float32, token string, tos string, type_ string, updated string, waitlistEnabled bool, warpEnabled bool, ) *Register200Response`

NewRegister200Response instantiates a new Register200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegister200ResponseWithDefaults

`func NewRegister200ResponseWithDefaults() *Register200Response`

NewRegister200ResponseWithDefaults instantiates a new Register200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Register200Response) GetAccount() GetSourceDevice200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Register200Response) GetAccountOk() (*GetSourceDevice200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Register200Response) SetAccount(v GetSourceDevice200ResponseAccount)`

SetAccount sets Account field to given value.


### GetConfig

`func (o *Register200Response) GetConfig() GetSourceDevice200ResponseConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Register200Response) GetConfigOk() (*GetSourceDevice200ResponseConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Register200Response) SetConfig(v GetSourceDevice200ResponseConfig)`

SetConfig sets Config field to given value.


### GetCreated

`func (o *Register200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Register200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Register200Response) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetEnabled

`func (o *Register200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Register200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Register200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFcmToken

`func (o *Register200Response) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *Register200Response) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *Register200Response) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.


### GetId

`func (o *Register200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Register200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Register200Response) SetId(v string)`

SetId sets Id field to given value.


### GetInstallId

`func (o *Register200Response) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *Register200Response) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *Register200Response) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.


### GetKey

`func (o *Register200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Register200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Register200Response) SetKey(v string)`

SetKey sets Key field to given value.


### GetLocale

`func (o *Register200Response) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Register200Response) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Register200Response) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetModel

`func (o *Register200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Register200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Register200Response) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *Register200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Register200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Register200Response) SetName(v string)`

SetName sets Name field to given value.


### GetPlace

`func (o *Register200Response) GetPlace() float32`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *Register200Response) GetPlaceOk() (*float32, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *Register200Response) SetPlace(v float32)`

SetPlace sets Place field to given value.


### GetToken

`func (o *Register200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Register200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Register200Response) SetToken(v string)`

SetToken sets Token field to given value.


### GetTos

`func (o *Register200Response) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *Register200Response) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *Register200Response) SetTos(v string)`

SetTos sets Tos field to given value.


### GetType

`func (o *Register200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Register200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Register200Response) SetType(v string)`

SetType sets Type field to given value.


### GetUpdated

`func (o *Register200Response) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Register200Response) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Register200Response) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetWaitlistEnabled

`func (o *Register200Response) GetWaitlistEnabled() bool`

GetWaitlistEnabled returns the WaitlistEnabled field if non-nil, zero value otherwise.

### GetWaitlistEnabledOk

`func (o *Register200Response) GetWaitlistEnabledOk() (*bool, bool)`

GetWaitlistEnabledOk returns a tuple with the WaitlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistEnabled

`func (o *Register200Response) SetWaitlistEnabled(v bool)`

SetWaitlistEnabled sets WaitlistEnabled field to given value.


### GetWarpEnabled

`func (o *Register200Response) GetWarpEnabled() bool`

GetWarpEnabled returns the WarpEnabled field if non-nil, zero value otherwise.

### GetWarpEnabledOk

`func (o *Register200Response) GetWarpEnabledOk() (*bool, bool)`

GetWarpEnabledOk returns a tuple with the WarpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarpEnabled

`func (o *Register200Response) SetWarpEnabled(v bool)`

SetWarpEnabled sets WarpEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


