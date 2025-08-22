# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBoundDevice**](DefaultAPI.md#DeleteBoundDevice) | **Delete** /{apiVersion}/reg/{sourceDeviceId}/account/reg/{boundDeviceId} | DeleteBoundDevice
[**GetAccount**](DefaultAPI.md#GetAccount) | **Get** /{apiVersion}/reg/{sourceDeviceId}/account | GetAccount
[**GetBoundDevices**](DefaultAPI.md#GetBoundDevices) | **Get** /{apiVersion}/reg/{sourceDeviceId}/account/devices | GetBoundDevices
[**GetClientConfig**](DefaultAPI.md#GetClientConfig) | **Get** /{apiVersion}/client_config | GetClientConfig
[**GetSourceDevice**](DefaultAPI.md#GetSourceDevice) | **Get** /{apiVersion}/reg/{sourceDeviceId} | GetSourceDevice
[**Register**](DefaultAPI.md#Register) | **Post** /{apiVersion}/reg | Register
[**ResetAccountLicense**](DefaultAPI.md#ResetAccountLicense) | **Post** /{apiVersion}/reg/{sourceDeviceId}/account/license | ResetAccountLicense
[**UpdateAccount**](DefaultAPI.md#UpdateAccount) | **Put** /{apiVersion}/reg/{sourceDeviceId}/account | UpdateAccount
[**UpdateBoundDevice**](DefaultAPI.md#UpdateBoundDevice) | **Patch** /{apiVersion}/reg/{sourceDeviceId}/account/reg/{boundDeviceId} | UpdateBoundDevice
[**UpdateSourceDevice**](DefaultAPI.md#UpdateSourceDevice) | **Patch** /{apiVersion}/reg/{sourceDeviceId} | UpdateSourceDevice



## DeleteBoundDevice

> DeleteBoundDevice(ctx, sourceDeviceId, apiVersion, boundDeviceId).Execute()

DeleteBoundDevice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 
	boundDeviceId := "boundDeviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteBoundDevice(context.Background(), sourceDeviceId, apiVersion, boundDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteBoundDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiDeleteBoundDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, sourceDeviceId, apiVersion).Execute()

GetAccount

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAccount(context.Background(), sourceDeviceId, apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAccount`: %v\n", resp)
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

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoundDevices

> []BoundDevice GetBoundDevices(ctx, sourceDeviceId, apiVersion).Execute()

GetBoundDevices

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBoundDevices(context.Background(), sourceDeviceId, apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBoundDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoundDevices`: []BoundDevice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBoundDevices`: %v\n", resp)
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

[**[]BoundDevice**](BoundDevice.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClientConfig(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClientConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientConfig`: GetClientConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClientConfig`: %v\n", resp)
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

[**GetClientConfig200Response**](GetClientConfig200Response.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | 
	sourceDeviceId := "sourceDeviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSourceDevice(context.Background(), apiVersion, sourceDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSourceDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceDevice`: GetSourceDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSourceDevice`: %v\n", resp)
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

[**GetSourceDevice200Response**](GetSourceDevice200Response.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | 
	registerRequest := *openapiclient.NewRegisterRequest("FcmToken_example", "InstallId_example", "Key_example", "Locale_example", "Model_example", "Tos_example", "Type_example") // RegisterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Register(context.Background(), apiVersion).RegisterRequest(registerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Register``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Register`: Register200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Register`: %v\n", resp)
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

[**Register200Response**](Register200Response.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ResetAccountLicense(context.Background(), sourceDeviceId, apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResetAccountLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAccountLicense`: ResetAccountLicense200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ResetAccountLicense`: %v\n", resp)
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

[**ResetAccountLicense200Response**](ResetAccountLicense200Response.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest("License_example") // UpdateAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateAccount(context.Background(), sourceDeviceId, apiVersion).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: UpdateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAccount`: %v\n", resp)
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

[**UpdateAccount200Response**](UpdateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBoundDevice

> []BoundDevice UpdateBoundDevice(ctx, sourceDeviceId, apiVersion, boundDeviceId).UpdateBoundDeviceRequest(updateBoundDeviceRequest).Execute()

UpdateBoundDevice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceDeviceId := "sourceDeviceId_example" // string | 
	apiVersion := "apiVersion_example" // string | 
	boundDeviceId := "boundDeviceId_example" // string | 
	updateBoundDeviceRequest := *openapiclient.NewUpdateBoundDeviceRequest() // UpdateBoundDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateBoundDevice(context.Background(), sourceDeviceId, apiVersion, boundDeviceId).UpdateBoundDeviceRequest(updateBoundDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateBoundDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBoundDevice`: []BoundDevice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateBoundDevice`: %v\n", resp)
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

[**[]BoundDevice**](BoundDevice.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | 
	sourceDeviceId := "sourceDeviceId_example" // string | 
	updateSourceDeviceRequest := *openapiclient.NewUpdateSourceDeviceRequest("Key_example") // UpdateSourceDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateSourceDevice(context.Background(), apiVersion, sourceDeviceId).UpdateSourceDeviceRequest(updateSourceDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSourceDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceDevice`: UpdateSourceDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSourceDevice`: %v\n", resp)
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

[**UpdateSourceDevice200Response**](UpdateSourceDevice200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

