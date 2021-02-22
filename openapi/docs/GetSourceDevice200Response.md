# GetSourceDevice200Response

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
**Tos** | **string** |  | 
**Type** | **string** |  | 
**Updated** | **string** |  | 
**WaitlistEnabled** | **bool** |  | 
**WarpEnabled** | **bool** |  | 

## Methods

### NewGetSourceDevice200Response

`func NewGetSourceDevice200Response(account GetSourceDevice200ResponseAccount, config GetSourceDevice200ResponseConfig, created string, enabled bool, fcmToken string, id string, installId string, key string, locale string, model string, name string, place float32, tos string, type_ string, updated string, waitlistEnabled bool, warpEnabled bool, ) *GetSourceDevice200Response`

NewGetSourceDevice200Response instantiates a new GetSourceDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceDevice200ResponseWithDefaults

`func NewGetSourceDevice200ResponseWithDefaults() *GetSourceDevice200Response`

NewGetSourceDevice200ResponseWithDefaults instantiates a new GetSourceDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetSourceDevice200Response) GetAccount() GetSourceDevice200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetSourceDevice200Response) GetAccountOk() (*GetSourceDevice200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetSourceDevice200Response) SetAccount(v GetSourceDevice200ResponseAccount)`

SetAccount sets Account field to given value.


### GetConfig

`func (o *GetSourceDevice200Response) GetConfig() GetSourceDevice200ResponseConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSourceDevice200Response) GetConfigOk() (*GetSourceDevice200ResponseConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSourceDevice200Response) SetConfig(v GetSourceDevice200ResponseConfig)`

SetConfig sets Config field to given value.


### GetCreated

`func (o *GetSourceDevice200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetSourceDevice200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetSourceDevice200Response) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetEnabled

`func (o *GetSourceDevice200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSourceDevice200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSourceDevice200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFcmToken

`func (o *GetSourceDevice200Response) GetFcmToken() string`

GetFcmToken returns the FcmToken field if non-nil, zero value otherwise.

### GetFcmTokenOk

`func (o *GetSourceDevice200Response) GetFcmTokenOk() (*string, bool)`

GetFcmTokenOk returns a tuple with the FcmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmToken

`func (o *GetSourceDevice200Response) SetFcmToken(v string)`

SetFcmToken sets FcmToken field to given value.


### GetId

`func (o *GetSourceDevice200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSourceDevice200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSourceDevice200Response) SetId(v string)`

SetId sets Id field to given value.


### GetInstallId

`func (o *GetSourceDevice200Response) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *GetSourceDevice200Response) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *GetSourceDevice200Response) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.


### GetKey

`func (o *GetSourceDevice200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetSourceDevice200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetSourceDevice200Response) SetKey(v string)`

SetKey sets Key field to given value.


### GetLocale

`func (o *GetSourceDevice200Response) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *GetSourceDevice200Response) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *GetSourceDevice200Response) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetModel

`func (o *GetSourceDevice200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetSourceDevice200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetSourceDevice200Response) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *GetSourceDevice200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSourceDevice200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSourceDevice200Response) SetName(v string)`

SetName sets Name field to given value.


### GetPlace

`func (o *GetSourceDevice200Response) GetPlace() float32`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *GetSourceDevice200Response) GetPlaceOk() (*float32, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *GetSourceDevice200Response) SetPlace(v float32)`

SetPlace sets Place field to given value.


### GetTos

`func (o *GetSourceDevice200Response) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *GetSourceDevice200Response) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *GetSourceDevice200Response) SetTos(v string)`

SetTos sets Tos field to given value.


### GetType

`func (o *GetSourceDevice200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSourceDevice200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSourceDevice200Response) SetType(v string)`

SetType sets Type field to given value.


### GetUpdated

`func (o *GetSourceDevice200Response) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetSourceDevice200Response) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetSourceDevice200Response) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetWaitlistEnabled

`func (o *GetSourceDevice200Response) GetWaitlistEnabled() bool`

GetWaitlistEnabled returns the WaitlistEnabled field if non-nil, zero value otherwise.

### GetWaitlistEnabledOk

`func (o *GetSourceDevice200Response) GetWaitlistEnabledOk() (*bool, bool)`

GetWaitlistEnabledOk returns a tuple with the WaitlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistEnabled

`func (o *GetSourceDevice200Response) SetWaitlistEnabled(v bool)`

SetWaitlistEnabled sets WaitlistEnabled field to given value.


### GetWarpEnabled

`func (o *GetSourceDevice200Response) GetWarpEnabled() bool`

GetWarpEnabled returns the WarpEnabled field if non-nil, zero value otherwise.

### GetWarpEnabledOk

`func (o *GetSourceDevice200Response) GetWarpEnabledOk() (*bool, bool)`

GetWarpEnabledOk returns a tuple with the WarpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarpEnabled

`func (o *GetSourceDevice200Response) SetWarpEnabled(v bool)`

SetWarpEnabled sets WarpEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


