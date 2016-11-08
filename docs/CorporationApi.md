# \CorporationApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorporationsCorporationId**](CorporationApi.md#GetCorporationsCorporationId) | **Get** /corporations/{corporation_id}/ | Get corporation information
[**GetCorporationsCorporationIdAlliancehistory**](CorporationApi.md#GetCorporationsCorporationIdAlliancehistory) | **Get** /corporations/{corporation_id}/alliancehistory/ | Get alliance history
[**GetCorporationsCorporationIdIcons**](CorporationApi.md#GetCorporationsCorporationIdIcons) | **Get** /corporations/{corporation_id}/icons/ | Get corporation icon
[**GetCorporationsCorporationIdMembers**](CorporationApi.md#GetCorporationsCorporationIdMembers) | **Get** /corporations/{corporation_id}/members/ | Get corporation members
[**GetCorporationsCorporationIdRoles**](CorporationApi.md#GetCorporationsCorporationIdRoles) | **Get** /corporations/{corporation_id}/roles/ | Get corporation members
[**GetCorporationsNames**](CorporationApi.md#GetCorporationsNames) | **Get** /corporations/names/ | Get corporation names


# **GetCorporationsCorporationId**
> GetCorporationsCorporationIdOk GetCorporationsCorporationId($corporationId, $datasource)

Get corporation information

Public information about a corporation  ---  Alternate route: `/v2/corporations/{corporation_id}/`  Alternate route: `/dev/corporations/{corporation_id}/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An Eve corporation ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCorporationsCorporationIdOk**](get_corporations_corporation_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdAlliancehistory**
> []GetCorporationsCorporationIdAlliancehistory200OkObject GetCorporationsCorporationIdAlliancehistory($corporationId, $datasource)

Get alliance history

Get a list of all the alliances a corporation has been a member of  ---  Alternate route: `/v1/corporations/{corporation_id}/alliancehistory/`  Alternate route: `/legacy/corporations/{corporation_id}/alliancehistory/`  Alternate route: `/dev/corporations/{corporation_id}/alliancehistory/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCorporationsCorporationIdAlliancehistory200OkObject**](get_corporations_corporation_id_alliancehistory__200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdIcons**
> GetCorporationsCorporationIdIconsOk GetCorporationsCorporationIdIcons($corporationId, $datasource)

Get corporation icon

Get the icon urls for a corporation  ---  Alternate route: `/v1/corporations/{corporation_id}/icons/`  Alternate route: `/legacy/corporations/{corporation_id}/icons/`  Alternate route: `/dev/corporations/{corporation_id}/icons/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCorporationsCorporationIdIconsOk**](get_corporations_corporation_id_icons_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdMembers**
> []GetCorporationsCorporationIdMembers200OkObject GetCorporationsCorporationIdMembers($corporationId, $datasource)

Get corporation members

Read the current list of members if the calling character is a member.  ---  Alternate route: `/v2/corporations/{corporation_id}/members/`  Alternate route: `/legacy/corporations/{corporation_id}/members/`  Alternate route: `/dev/corporations/{corporation_id}/members/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| A corporation ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCorporationsCorporationIdMembers200OkObject**](get_corporations_corporation_id_members__200_ok_object.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdRoles**
> []GetCorporationsCorporationIdRoles200OkObject GetCorporationsCorporationIdRoles($corporationId, $datasource)

Get corporation members

Return the roles of all members if the character has the personnel manager role or any grantable role.  ---  Alternate route: `/v1/corporations/{corporation_id}/roles/`  Alternate route: `/legacy/corporations/{corporation_id}/roles/`  Alternate route: `/dev/corporations/{corporation_id}/roles/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| A corporation ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCorporationsCorporationIdRoles200OkObject**](get_corporations_corporation_id_roles__200_ok_object.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsNames**
> []GetCorporationsNames200OkObject GetCorporationsNames($corporationIds, $datasource)

Get corporation names

Resolve a set of corporation IDs to corporation names  ---  Alternate route: `/v1/corporations/names/`  Alternate route: `/legacy/corporations/names/`  Alternate route: `/dev/corporations/names/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationIds** | [**[]int64**](int64.md)| A comma separated list of corporation IDs | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCorporationsNames200OkObject**](get_corporations_names_200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

