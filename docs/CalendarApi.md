# \CalendarApi

All URIs are relative to *https://esi.tech.ccp.is/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdCalendar**](CalendarApi.md#GetCharactersCharacterIdCalendar) | **Get** /characters/{character_id}/calendar/ | List calendar event summaries
[**GetCharactersCharacterIdCalendarEventId**](CalendarApi.md#GetCharactersCharacterIdCalendarEventId) | **Get** /characters/{character_id}/calendar/{event_id}/ | Get an event
[**PutCharactersCharacterIdCalendarEventId**](CalendarApi.md#PutCharactersCharacterIdCalendarEventId) | **Put** /characters/{character_id}/calendar/{event_id}/ | Respond to an event


# **GetCharactersCharacterIdCalendar**
> []GetCharactersCharacterIdCalendar200OkObject GetCharactersCharacterIdCalendar($characterId, $fromEvent, $datasource)

List calendar event summaries

Get 50 event summaries from the calendar. If no event ID is given, the resource will return the next 50 chronological event summaries from now. If an event ID is specified, it will return the next 50 chronological event summaries from after that event.   ---  Alternate route: `/v1/characters/{character_id}/calendar/`  Alternate route: `/legacy/characters/{character_id}/calendar/`  Alternate route: `/dev/characters/{character_id}/calendar/`   ---  This route is cached for up to 5 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The character to retrieve events from | 
 **fromEvent** | **int32**| The event ID to retrieve events from | [optional] 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**[]GetCharactersCharacterIdCalendar200OkObject**](get_characters_character_id_calendar_200_ok_object.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCalendarEventId**
> GetCharactersCharacterIdCalendarEventIdOk GetCharactersCharacterIdCalendarEventId($characterId, $eventId, $datasource)

Get an event

Get all the information for a specific event  ---  Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`   ---  This route is cached for up to 5 seconds


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The character id requesting the event | 
 **eventId** | **int32**| The id of the event requested | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

[**GetCharactersCharacterIdCalendarEventIdOk**](get_characters_character_id_calendar_event_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdCalendarEventId**
> PutCharactersCharacterIdCalendarEventId($characterId, $eventId, $response, $datasource)

Respond to an event

Set your response status to an event  ---  Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/` 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| The character ID requesting the event | 
 **eventId** | **int32**| The ID of the event requested | 
 **response** | [**PutCharactersCharacterIdCalendarEventIdResponse**](PutCharactersCharacterIdCalendarEventIdResponse.md)| The response value to set, overriding current value. | 
 **datasource** | **string**| The server name you would like data from | [optional] [default to tranquility]

### Return type

void (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

