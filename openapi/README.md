# Go API client for swagger

This is a sample proof of concept repo meant to be run on localhost at the moment

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8081/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ServiceApi* | [**ServicesGet**](docs/ServiceApi.md#servicesget) | **Get** /services | List Services
*ServiceApi* | [**ServicesPost**](docs/ServiceApi.md#servicespost) | **Post** /services | Add a new service
*ServiceApi* | [**ServicesPut**](docs/ServiceApi.md#servicesput) | **Put** /services | Update Service
*ServiceApi* | [**ServicesServiceIdDelete**](docs/ServiceApi.md#servicesserviceiddelete) | **Delete** /services/{service_id} | Delete Service
*ServiceApi* | [**ServicesServiceIdGet**](docs/ServiceApi.md#servicesserviceidget) | **Get** /services/{service_id} | Get a service
*ServiceVersionApi* | [**ServicesServiceIdaddVersionPost**](docs/ServiceVersionApi.md#servicesserviceidaddversionpost) | **Post** /services/{service_id}:addVersion | Add a new service version
*ServiceVersionApi* | [**ServicesServiceIdremoveVersionPost**](docs/ServiceVersionApi.md#servicesserviceidremoveversionpost) | **Post** /services/{service_id}:removeVersion | Remove a new service version

## Documentation For Models

 - [AddServiceVersionRequest](docs/AddServiceVersionRequest.md)
 - [CreateServiceRequest](docs/CreateServiceRequest.md)
 - [Empty](docs/Empty.md)
 - [ListServicesResponse](docs/ListServicesResponse.md)
 - [RemoveServiceVersionRequest](docs/RemoveServiceVersionRequest.md)
 - [Service](docs/Service.md)
 - [ServiceVersion](docs/ServiceVersion.md)
 - [UpdateServiceRequest](docs/UpdateServiceRequest.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

diegosegundo2@gmail.com
