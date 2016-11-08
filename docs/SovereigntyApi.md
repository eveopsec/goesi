# \SovereigntyApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSovereigntyCampaigns**](SovereigntyApi.md#GetSovereigntyCampaigns) | **Get** /sovereignty/campaigns/ | List sovereignty campaigns
[**GetSovereigntyStructures**](SovereigntyApi.md#GetSovereigntyStructures) | **Get** /sovereignty/structures/ | List sovereignty structures


# **GetSovereigntyCampaigns**
> []GetSovereigntyCampaigns200OkObject GetSovereigntyCampaigns($datasource)

List sovereignty campaigns

Shows sovereignty data for campaigns.  ---  Alternate route: `/v1/sovereignty/campaigns/`  Alternate route: `/legacy/sovereignty/campaigns/`  Alternate route: `/dev/sovereignty/campaigns/`   ---  This route is cached for up to 5 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetSovereigntyCampaigns200OkObject**](get_sovereignty_campaigns_200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSovereigntyStructures**
> []GetSovereigntyStructures200OkObject GetSovereigntyStructures($datasource)

List sovereignty structures

Shows sovereignty data for structures.  ---  Alternate route: `/v1/sovereignty/structures/`  Alternate route: `/legacy/sovereignty/structures/`  Alternate route: `/dev/sovereignty/structures/`   ---  This route is cached for up to 120 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetSovereigntyStructures200OkObject**](get_sovereignty_structures_200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

