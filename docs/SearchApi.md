# \SearchApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdSearch**](SearchApi.md#GetCharactersCharacterIdSearch) | **Get** /characters/{character_id}/search/ | Search on a string
[**GetSearch**](SearchApi.md#GetSearch) | **Get** /search/ | Search on a string


# **GetCharactersCharacterIdSearch**
> GetCharactersCharacterIdSearchOk GetCharactersCharacterIdSearch($characterId, $search, $categories, $language, $strict, $datasource)

Search on a string

Search for entities that match a given sub-string.  ---  Alternate route: `/v2/characters/{character_id}/search/`  Alternate route: `/dev/characters/{character_id}/search/` 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **search** | **string**| The string to search on | 
 **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **language** | **string**| Search locale | [optional] [default to en-us]
 **strict** | **bool**| Whether the search should be a strict match | [optional] [default to false]
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCharactersCharacterIdSearchOk**](get_characters_character_id_search_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearch**
> GetSearchOk GetSearch($search, $categories, $language, $strict, $datasource)

Search on a string

Search for entities that match a given sub-string.  ---  Alternate route: `/v1/search/`  Alternate route: `/legacy/search/`  Alternate route: `/dev/search/` 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string**| The string to search on | 
 **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **language** | **string**| Search locale | [optional] [default to en-us]
 **strict** | **bool**| Whether the search should be a strict match | [optional] [default to false]
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetSearchOk**](get_search_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

