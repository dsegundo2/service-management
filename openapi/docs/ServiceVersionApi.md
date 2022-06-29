# {{classname}}

All URIs are relative to *http://localhost:8081/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesServiceIdaddVersionPost**](ServiceVersionApi.md#ServicesServiceIdaddVersionPost) | **Post** /services/{service_id}:addVersion | Add a new service version
[**ServicesServiceIdremoveVersionPost**](ServiceVersionApi.md#ServicesServiceIdremoveVersionPost) | **Post** /services/{service_id}:removeVersion | Remove a new service version

# **ServicesServiceIdaddVersionPost**
> []ServiceVersion ServicesServiceIdaddVersionPost(ctx, body, serviceId)
Add a new service version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddServiceVersionRequest**](AddServiceVersionRequest.md)| Service info needed to create the service version | 
  **serviceId** | **string**| ID of the service | 

### Return type

[**[]ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesServiceIdremoveVersionPost**
> []Empty ServicesServiceIdremoveVersionPost(ctx, body, serviceId)
Remove a new service version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveServiceVersionRequest**](RemoveServiceVersionRequest.md)| Service info needed to remove the service version | 
  **serviceId** | **string**| ID of the service | 

### Return type

[**[]Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

