# {{classname}}

All URIs are relative to *http://localhost:8081/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesGet**](ServiceApi.md#ServicesGet) | **Get** /services | List Services
[**ServicesPost**](ServiceApi.md#ServicesPost) | **Post** /services | Add a new service
[**ServicesPut**](ServiceApi.md#ServicesPut) | **Put** /services | Update Service
[**ServicesServiceIdDelete**](ServiceApi.md#ServicesServiceIdDelete) | **Delete** /services/{service_id} | Delete Service
[**ServicesServiceIdGet**](ServiceApi.md#ServicesServiceIdGet) | **Get** /services/{service_id} | Get a service

# **ServicesGet**
> []ListServicesResponse ServicesGet(ctx, optional)
List Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiServicesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServicesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filters the services by looking for this substring in the service Title | 
 **sort** | **optional.String**| Based on the google api spec. The string value should follow SQL syntax: comma separated list of fields. For example: \&quot;foo,bar\&quot;. The default sorting order is ascending. To specify descending order for a field, a suffix \&quot; desc\&quot; should be appended to the field name. Ex: \&quot;foo desc,bar\&quot; | 
 **limit** | **optional.Int32**| Maximum number of responses to return | [default to 10]
 **offset** | **optional.Int32**| Number of services to skip over before returning the list | [default to 0]

### Return type

[**[]ListServicesResponse**](ListServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPost**
> []Service ServicesPost(ctx, body)
Add a new service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateServiceRequest**](CreateServiceRequest.md)| Service info needed to create the service | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPut**
> []Service ServicesPut(ctx, body)
Update Service

Used to update values of the service. Only Updates the Title and the Description fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateServiceRequest**](UpdateServiceRequest.md)| Service info needed to update the service. Field masks specifies which values to update | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesServiceIdDelete**
> []Empty ServicesServiceIdDelete(ctx, serviceId)
Delete Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| ID of the service to delete | 

### Return type

[**[]Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesServiceIdGet**
> []Service ServicesServiceIdGet(ctx, serviceId, optional)
Get a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| ID of the service | 
 **optional** | ***ServiceApiServicesServiceIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServicesServiceIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadVersions** | **optional.Bool**| Determines whether or not to load the information about a services versions | [default to false]

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

