# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](DefaultApi.md#GetAccount) | **Get** /{apiVersion}/reg/{sourceDeviceId}/account | GetAccount
[**GetBoundDevices**](DefaultApi.md#GetBoundDevices) | **Get** /{apiVersion}/reg/{sourceDeviceId}/account/devices | GetBoundDevices
[**GetClientConfig**](DefaultApi.md#GetClientConfig) | **Get** /{apiVersion}/client_config | GetClientConfig
[**GetSourceDevice**](DefaultApi.md#GetSourceDevice) | **Get** /{apiVersion}/reg/{sourceDeviceId} | GetSourceDevice
[**Register**](DefaultApi.md#Register) | **Post** /{apiVersion}/reg | Register
[**ResetAccountLicense**](DefaultApi.md#ResetAccountLicense) | **Post** /{apiVersion}/reg/{sourceDeviceId}/account/license | ResetAccountLicense
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Put** /{apiVersion}/reg/{sourceDeviceId}/account | UpdateAccount
[**UpdateBoundDevice**](DefaultApi.md#UpdateBoundDevice) | **Patch** /{apiVersion}/reg/{sourceDeviceId}/account/reg/{boundDeviceId} | UpdateBoundDevice
[**UpdateSourceDevice**](DefaultApi.md#UpdateSourceDevice) | **Patch** /{apiVersion}/reg/{sourceDeviceId} | UpdateSourceDevice



## GetAccount

> GetAccount200Response GetAccount(ctx, sourceDeviceId, apiVersion).Execute()

GetAccount

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceDeviceId := "sourceDeviceId_example" // string | 
    apiVersion := "apiVersion_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAccount(context.Background(), sourceDeviceId, apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: GetAccount200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string** |  | 
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> []GetBoundDevices200Response GetBoundDevices(ctx, sourceDeviceId, apiVersion).Execute()

GetBoundDevices

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceDeviceId := "sourceDeviceId_example" // string | 
    apiVersion := "apiVersion_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBoundDevices(context.Background(), sourceDeviceId, apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBoundDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoundDevices`: []GetBoundDevices200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBoundDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string** |  | 
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoundDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> GetClientConfig200Response GetClientConfig(ctx, apiVersion).Execute()

GetClientConfig

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiVersion := "apiVersion_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetClientConfig(context.Background(), apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClientConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientConfig`: GetClientConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClientConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> GetSourceDevice200Response GetSourceDevice(ctx, apiVersion, sourceDeviceId).Execute()

GetSourceDevice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiVersion := "apiVersion_example" // string | 
    sourceDeviceId := "sourceDeviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSourceDevice(context.Background(), apiVersion, sourceDeviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceDevice`: GetSourceDevice200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSourceDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** |  | 
**sourceDeviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> Register200Response Register(ctx, apiVersion).RegisterRequest(registerRequest).Execute()

Register

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiVersion := "apiVersion_example" // string | 
    registerRequest := *openapiclient.NewRegisterRequest("FcmToken_example", "InstallId_example", "Key_example", "Locale_example", "Model_example", "Tos_example", "Type_example") // RegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Register(context.Background(), apiVersion).RegisterRequest(registerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: Register200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Register`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  | 

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

> ResetAccountLicense200Response ResetAccountLicense(ctx, sourceDeviceId, apiVersion).Execute()

ResetAccountLicense

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceDeviceId := "sourceDeviceId_example" // string | 
    apiVersion := "apiVersion_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetAccountLicense(context.Background(), sourceDeviceId, apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetAccountLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetAccountLicense`: ResetAccountLicense200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetAccountLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string** |  | 
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAccountLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> UpdateAccount200Response UpdateAccount(ctx, sourceDeviceId, apiVersion).UpdateAccountRequest(updateAccountRequest).Execute()

UpdateAccount

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceDeviceId := "sourceDeviceId_example" // string | 
    apiVersion := "apiVersion_example" // string | 
    updateAccountRequest := *openapiclient.NewUpdateAccountRequest("License_example") // UpdateAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccount(context.Background(), sourceDeviceId, apiVersion).UpdateAccountRequest(updateAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: UpdateAccount200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string** |  | 
**apiVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

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

> []UpdateBoundDevice200Response UpdateBoundDevice(ctx, sourceDeviceId, apiVersion, boundDeviceId).UpdateBoundDeviceRequest(updateBoundDeviceRequest).Execute()

UpdateBoundDevice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceDeviceId := "sourceDeviceId_example" // string | 
    apiVersion := "apiVersion_example" // string | 
    boundDeviceId := "boundDeviceId_example" // string | 
    updateBoundDeviceRequest := *openapiclient.NewUpdateBoundDeviceRequest() // UpdateBoundDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBoundDevice(context.Background(), sourceDeviceId, apiVersion, boundDeviceId).UpdateBoundDeviceRequest(updateBoundDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBoundDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBoundDevice`: []UpdateBoundDevice200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBoundDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceDeviceId** | **string** |  | 
**apiVersion** | **string** |  | 
**boundDeviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBoundDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateBoundDeviceRequest** | [**UpdateBoundDeviceRequest**](UpdateBoundDeviceRequest.md) |  | 

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

> UpdateSourceDevice200Response UpdateSourceDevice(ctx, apiVersion, sourceDeviceId).UpdateSourceDeviceRequest(updateSourceDeviceRequest).Execute()

UpdateSourceDevice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiVersion := "apiVersion_example" // string | 
    sourceDeviceId := "sourceDeviceId_example" // string | 
    updateSourceDeviceRequest := *openapiclient.NewUpdateSourceDeviceRequest("Key_example") // UpdateSourceDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSourceDevice(context.Background(), apiVersion, sourceDeviceId).UpdateSourceDeviceRequest(updateSourceDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSourceDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourceDevice`: UpdateSourceDevice200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSourceDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** |  | 
**sourceDeviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSourceDeviceRequest** | [**UpdateSourceDeviceRequest**](UpdateSourceDeviceRequest.md) |  | 

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

