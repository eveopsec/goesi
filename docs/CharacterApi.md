# \CharacterApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterId**](CharacterApi.md#GetCharactersCharacterId) | **Get** /characters/{character_id}/ | Get character&#39;s public information
[**GetCharactersCharacterIdCorporationhistory**](CharacterApi.md#GetCharactersCharacterIdCorporationhistory) | **Get** /characters/{character_id}/corporationhistory/ | Get corporation history
[**GetCharactersCharacterIdPortrait**](CharacterApi.md#GetCharactersCharacterIdPortrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
[**GetCharactersNames**](CharacterApi.md#GetCharactersNames) | **Get** /characters/names/ | Get character names
[**PostCharactersCharacterIdCspa**](CharacterApi.md#PostCharactersCharacterIdCspa) | **Post** /characters/{character_id}/cspa/ | Calculate a CSPA charge cost


# **GetCharactersCharacterId**
> GetCharactersCharacterIdOk GetCharactersCharacterId($characterId, $datasource)

Get character's public information

Public information about a character  ---  Alternate route: `/v3/characters/{character_id}/`  Alternate route: `/legacy/characters/{character_id}/`  Alternate route: `/dev/characters/{character_id}/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCharactersCharacterIdOk**](get_characters_character_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCorporationhistory**
> []GetCharactersCharacterIdCorporationhistory200OkObject GetCharactersCharacterIdCorporationhistory($characterId, $datasource)

Get corporation history

Get a list of all the corporations a character has been a member of  ---  Alternate route: `/v1/characters/{character_id}/corporationhistory/`  Alternate route: `/legacy/characters/{character_id}/corporationhistory/`  Alternate route: `/dev/characters/{character_id}/corporationhistory/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCharactersCharacterIdCorporationhistory200OkObject**](get_characters_character_id_corporationhistory_200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPortrait**
> GetCharactersCharacterIdPortraitOk GetCharactersCharacterIdPortrait($characterId, $datasource)

Get character portraits

Get portrait urls for a character  ---  Alternate route: `/v2/characters/{character_id}/portrait/`  Alternate route: `/dev/characters/{character_id}/portrait/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCharactersCharacterIdPortraitOk**](get_characters_character_id_portrait_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersNames**
> []GetCharactersNames200OkObject GetCharactersNames($characterIds, $datasource)

Get character names

Resolve a set of character IDs to character names  ---  Alternate route: `/v1/characters/names/`  Alternate route: `/legacy/characters/names/`  Alternate route: `/dev/characters/names/`   ---  This route is cached for up to 3600 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterIds** | [**[]int64**](int64.md)| A comma separated list of character IDs | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCharactersNames200OkObject**](get_characters_names_200_ok_object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdCspa**
> PostCharactersCharacterIdCspaCreated PostCharactersCharacterIdCspa($characterId, $characters, $datasource)

Calculate a CSPA charge cost

Takes a source character ID in the url and a set of target character ID's in the body, returns a CSPA charge cost  ---  Alternate route: `/v3/characters/{character_id}/cspa/`  Alternate route: `/legacy/characters/{character_id}/cspa/`  Alternate route: `/dev/characters/{character_id}/cspa/` 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **characters** | [**PostCharactersCharacterIdCspaCharacters**](PostCharactersCharacterIdCspaCharacters.md)| The target characters to calculate the charge for | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**PostCharactersCharacterIdCspaCreated**](post_characters_character_id_cspa_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

