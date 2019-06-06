# \ListApi

All URIs are relative to *https://www.thebluealliance.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDistrictEvents**](ListApi.md#GetDistrictEvents) | **Get** /district/{district_key}/events | 
[**GetDistrictEventsKeys**](ListApi.md#GetDistrictEventsKeys) | **Get** /district/{district_key}/events/keys | 
[**GetDistrictEventsSimple**](ListApi.md#GetDistrictEventsSimple) | **Get** /district/{district_key}/events/simple | 
[**GetDistrictRankings**](ListApi.md#GetDistrictRankings) | **Get** /district/{district_key}/rankings | 
[**GetDistrictTeams**](ListApi.md#GetDistrictTeams) | **Get** /district/{district_key}/teams | 
[**GetDistrictTeamsKeys**](ListApi.md#GetDistrictTeamsKeys) | **Get** /district/{district_key}/teams/keys | 
[**GetDistrictTeamsSimple**](ListApi.md#GetDistrictTeamsSimple) | **Get** /district/{district_key}/teams/simple | 
[**GetEventTeams**](ListApi.md#GetEventTeams) | **Get** /event/{event_key}/teams | 
[**GetEventTeamsKeys**](ListApi.md#GetEventTeamsKeys) | **Get** /event/{event_key}/teams/keys | 
[**GetEventTeamsSimple**](ListApi.md#GetEventTeamsSimple) | **Get** /event/{event_key}/teams/simple | 
[**GetEventTeamsStatuses**](ListApi.md#GetEventTeamsStatuses) | **Get** /event/{event_key}/teams/statuses | 
[**GetEventsByYear**](ListApi.md#GetEventsByYear) | **Get** /events/{year} | 
[**GetEventsByYearKeys**](ListApi.md#GetEventsByYearKeys) | **Get** /events/{year}/keys | 
[**GetEventsByYearSimple**](ListApi.md#GetEventsByYearSimple) | **Get** /events/{year}/simple | 
[**GetTeamEventsStatusesByYear**](ListApi.md#GetTeamEventsStatusesByYear) | **Get** /team/{team_key}/events/{year}/statuses | 
[**GetTeams**](ListApi.md#GetTeams) | **Get** /teams/{page_num} | 
[**GetTeamsByYear**](ListApi.md#GetTeamsByYear) | **Get** /teams/{year}/{page_num} | 
[**GetTeamsByYearKeys**](ListApi.md#GetTeamsByYearKeys) | **Get** /teams/{year}/{page_num}/keys | 
[**GetTeamsByYearSimple**](ListApi.md#GetTeamsByYearSimple) | **Get** /teams/{year}/{page_num}/simple | 
[**GetTeamsKeys**](ListApi.md#GetTeamsKeys) | **Get** /teams/{page_num}/keys | 
[**GetTeamsSimple**](ListApi.md#GetTeamsSimple) | **Get** /teams/{page_num}/simple | 



## GetDistrictEvents

> []Event GetDistrictEvents(ctx, )


Gets a list of events in the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictEventsKeys

> []string GetDistrictEventsKeys(ctx, )


Gets a list of event keys for events in the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictEventsSimple

> []EventSimple GetDistrictEventsSimple(ctx, districtKey)


Gets a short-form list of events in the given district.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 

### Return type

[**[]EventSimple**](Event_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictRankings

> []DistrictRanking GetDistrictRankings(ctx, )


Gets a list of team district rankings for the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DistrictRanking**](District_Ranking.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictTeams

> []Team GetDistrictTeams(ctx, )


Gets a list of `Team` objects that competed in events in the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictTeamsKeys

> []string GetDistrictTeamsKeys(ctx, )


Gets a list of `Team` objects that competed in events in the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistrictTeamsSimple

> []TeamSimple GetDistrictTeamsSimple(ctx, )


Gets a short-form list of `Team` objects that competed in events in the given district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TeamSimple**](Team_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTeams

> []Team GetEventTeams(ctx, )


Gets a list of `Team` objects that competed in the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTeamsKeys

> []string GetEventTeamsKeys(ctx, )


Gets a list of `Team` keys that competed in the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTeamsSimple

> []TeamSimple GetEventTeamsSimple(ctx, )


Gets a short-form list of `Team` objects that competed in the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TeamSimple**](Team_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTeamsStatuses

> map[string]TeamEventStatus GetEventTeamsStatuses(ctx, )


Gets a key-value list of the event statuses for teams competing at the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]TeamEventStatus**](Team_Event_Status.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByYear

> []Event GetEventsByYear(ctx, )


Gets a list of events in the given year.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByYearKeys

> []string GetEventsByYearKeys(ctx, )


Gets a list of event keys in the given year.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByYearSimple

> []EventSimple GetEventsByYearSimple(ctx, )


Gets a short-form list of events in the given year.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EventSimple**](Event_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEventsStatusesByYear

> map[string]TeamEventStatus GetTeamEventsStatusesByYear(ctx, )


Gets a key-value list of the event statuses for events this team has competed at in the given year.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]TeamEventStatus**](Team_Event_Status.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> []Team GetTeams(ctx, )


Gets a list of `Team` objects, paginated in groups of 500.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByYear

> []Team GetTeamsByYear(ctx, )


Gets a list of `Team` objects that competed in the given year, paginated in groups of 500.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByYearKeys

> []string GetTeamsByYearKeys(ctx, )


Gets a list Team Keys that competed in the given year, paginated in groups of 500.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsByYearSimple

> []TeamSimple GetTeamsByYearSimple(ctx, )


Gets a list of short form `Team_Simple` objects that competed in the given year, paginated in groups of 500.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TeamSimple**](Team_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsKeys

> []string GetTeamsKeys(ctx, )


Gets a list of Team keys, paginated in groups of 500. (Note, each page will not have 500 teams, but will include the teams within that range of 500.)

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsSimple

> []TeamSimple GetTeamsSimple(ctx, )


Gets a list of short form `Team_Simple` objects, paginated in groups of 500.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TeamSimple**](Team_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

