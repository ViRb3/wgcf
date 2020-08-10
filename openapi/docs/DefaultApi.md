# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](DefaultApi.md#GetAccount) | **Get** /v0a977/reg/{sourceDeviceId}/account | GetAccount
[**GetBoundDevices**](DefaultApi.md#GetBoundDevices) | **Get** /v0a977/reg/{sourceDeviceId}/account/devices | GetBoundDevices
[**GetClientConfig**](DefaultApi.md#GetClientConfig) | **Get** /v0a977/client_config | GetClientConfig
[**GetSourceDevice**](DefaultApi.md#GetSourceDevice) | **Get** /v0a977/reg/{sourceDeviceId} | GetSourceDevice
[**Register**](DefaultApi.md#Register) | **Post** /v0a977/reg | Register
[**ResetAccountLicense**](DefaultApi.md#ResetAccountLicense) | **Post** /v0a977/reg/{sourceDeviceId}/account/license | ResetAccountLicense
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Put** /v0a977/reg/{sourceDeviceId}/account | UpdateAccount
[**UpdateBoundDevice**](DefaultApi.md#UpdateBoundDevice) | **Patch** /v0a977/reg/{sourceDeviceId}/account/reg/{boundDeviceId} | UpdateBoundDevice
[**UpdateSourceDevice**](DefaultApi.md#UpdateSourceDevice) | **Patch** /v0a977/reg/{sourceDeviceId} | UpdateSourceDevice



## GetAccount

> GetAccount200Response GetAccount(ctx, sourceDeviceId)

GetAccount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 

### Return type

[**GetAccount200Response**](GetAccount_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoundDevices

> []GetBoundDevices200Response GetBoundDevices(ctx, sourceDeviceId)

GetBoundDevices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 

### Return type

[**[]GetBoundDevices200Response**](GetBoundDevices_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientConfig

> GetClientConfig200Response GetClientConfig(ctx, )

GetClientConfig

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetClientConfig200Response**](GetClientConfig_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceDevice

> GetSourceDevice200Response GetSourceDevice(ctx, sourceDeviceId)

GetSourceDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 

### Return type

[**GetSourceDevice200Response**](GetSourceDevice_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> Register200Response Register(ctx, optional)

Register

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**optional.Interface of RegisterRequest**](RegisterRequest.md)|  | 

### Return type

[**Register200Response**](Register_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAccountLicense

> ResetAccountLicense200Response ResetAccountLicense(ctx, sourceDeviceId)

ResetAccountLicense

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 

### Return type

[**ResetAccountLicense200Response**](ResetAccountLicense_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount200Response UpdateAccount(ctx, sourceDeviceId, optional)

UpdateAccount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 
 **optional** | ***UpdateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountRequest** | [**optional.Interface of UpdateAccountRequest**](UpdateAccountRequest.md)|  | 

### Return type

[**UpdateAccount200Response**](UpdateAccount_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBoundDevice

> []UpdateBoundDevice200Response UpdateBoundDevice(ctx, sourceDeviceId, boundDeviceId, optional)

UpdateBoundDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 
**boundDeviceId** | **string**|  | 
 **optional** | ***UpdateBoundDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBoundDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBoundDeviceRequest** | [**optional.Interface of UpdateBoundDeviceRequest**](UpdateBoundDeviceRequest.md)|  | 

### Return type

[**[]UpdateBoundDevice200Response**](UpdateBoundDevice_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceDevice

> UpdateSourceDevice200Response UpdateSourceDevice(ctx, sourceDeviceId, optional)

UpdateSourceDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string**|  | 
 **optional** | ***UpdateSourceDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSourceDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSourceDeviceRequest** | [**optional.Interface of UpdateSourceDeviceRequest**](UpdateSourceDeviceRequest.md)|  | 

### Return type

[**UpdateSourceDevice200Response**](UpdateSourceDevice_200_Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

