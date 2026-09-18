# NotificationAckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FcmMessageId** | **NullableString** |  |
**Success** | **bool** |  |
**Error** | Pointer to **NullableString** |  | [optional]
**NotificationId** | Pointer to **NullableString** |  | [optional]

## Methods

### NewNotificationAckRequest

`func NewNotificationAckRequest(fcmMessageId NullableString, success bool, ) *NotificationAckRequest`

NewNotificationAckRequest instantiates a new NotificationAckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAckRequestWithDefaults

`func NewNotificationAckRequestWithDefaults() *NotificationAckRequest`

NewNotificationAckRequestWithDefaults instantiates a new NotificationAckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFcmMessageId

`func (o *NotificationAckRequest) GetFcmMessageId() string`

GetFcmMessageId returns the FcmMessageId field if non-nil, zero value otherwise.

### GetFcmMessageIdOk

`func (o *NotificationAckRequest) GetFcmMessageIdOk() (*string, bool)`

GetFcmMessageIdOk returns a tuple with the FcmMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmMessageId

`func (o *NotificationAckRequest) SetFcmMessageId(v string)`

SetFcmMessageId sets FcmMessageId field to given value.


### SetFcmMessageIdNil

`func (o *NotificationAckRequest) SetFcmMessageIdNil(b bool)`

 SetFcmMessageIdNil sets the value for FcmMessageId to be an explicit nil

### UnsetFcmMessageId
`func (o *NotificationAckRequest) UnsetFcmMessageId()`

UnsetFcmMessageId ensures that no value is present for FcmMessageId, not even an explicit nil
### GetSuccess

`func (o *NotificationAckRequest) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *NotificationAckRequest) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *NotificationAckRequest) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *NotificationAckRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NotificationAckRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NotificationAckRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NotificationAckRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *NotificationAckRequest) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *NotificationAckRequest) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNotificationId

`func (o *NotificationAckRequest) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotificationAckRequest) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotificationAckRequest) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *NotificationAckRequest) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### SetNotificationIdNil

`func (o *NotificationAckRequest) SetNotificationIdNil(b bool)`

 SetNotificationIdNil sets the value for NotificationId to be an explicit nil

### UnsetNotificationId
`func (o *NotificationAckRequest) UnsetNotificationId()`

UnsetNotificationId ensures that no value is present for NotificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


