# GetClientConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**ClientConfig**](ClientConfig.md) |  |
**Success** | **bool** |  |
**Errors** | **[]map[string]interface{}** |  |
**Messages** | **[]map[string]interface{}** |  |

## Methods

### NewGetClientConfig200Response

`func NewGetClientConfig200Response(result ClientConfig, success bool, errors []map[string]interface{}, messages []map[string]interface{}, ) *GetClientConfig200Response`

NewGetClientConfig200Response instantiates a new GetClientConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientConfig200ResponseWithDefaults

`func NewGetClientConfig200ResponseWithDefaults() *GetClientConfig200Response`

NewGetClientConfig200ResponseWithDefaults instantiates a new GetClientConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetClientConfig200Response) GetResult() ClientConfig`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetClientConfig200Response) GetResultOk() (*ClientConfig, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetClientConfig200Response) SetResult(v ClientConfig)`

SetResult sets Result field to given value.


### GetSuccess

`func (o *GetClientConfig200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetClientConfig200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetClientConfig200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrors

`func (o *GetClientConfig200Response) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetClientConfig200Response) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetClientConfig200Response) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.


### GetMessages

`func (o *GetClientConfig200Response) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetClientConfig200Response) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetClientConfig200Response) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


