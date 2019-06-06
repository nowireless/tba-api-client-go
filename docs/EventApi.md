# \EventApi

All URIs are relative to *https://www.thebluealliance.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDistrictEvents**](EventApi.md#GetDistrictEvents) | **Get** /district/{district_key}/events | 
[**GetDistrictEventsKeys**](EventApi.md#GetDistrictEventsKeys) | **Get** /district/{district_key}/events/keys | 
[**GetDistrictEventsSimple**](EventApi.md#GetDistrictEventsSimple) | **Get** /district/{district_key}/events/simple | 
[**GetEvent**](EventApi.md#GetEvent) | **Get** /event/{event_key} | 
[**GetEventAlliances**](EventApi.md#GetEventAlliances) | **Get** /event/{event_key}/alliances | 
[**GetEventAwards**](EventApi.md#GetEventAwards) | **Get** /event/{event_key}/awards | 
[**GetEventDistrictPoints**](EventApi.md#GetEventDistrictPoints) | **Get** /event/{event_key}/district_points | 
[**GetEventInsights**](EventApi.md#GetEventInsights) | **Get** /event/{event_key}/insights | 
[**GetEventMatchTimeseries**](EventApi.md#GetEventMatchTimeseries) | **Get** /event/{event_key}/matches/timeseries | 
[**GetEventMatches**](EventApi.md#GetEventMatches) | **Get** /event/{event_key}/matches | 
[**GetEventMatchesKeys**](EventApi.md#GetEventMatchesKeys) | **Get** /event/{event_key}/matches/keys | 
[**GetEventMatchesSimple**](EventApi.md#GetEventMatchesSimple) | **Get** /event/{event_key}/matches/simple | 
[**GetEventOPRs**](EventApi.md#GetEventOPRs) | **Get** /event/{event_key}/oprs | 
[**GetEventPredictions**](EventApi.md#GetEventPredictions) | **Get** /event/{event_key}/predictions | 
[**GetEventRankings**](EventApi.md#GetEventRankings) | **Get** /event/{event_key}/rankings | 
[**GetEventSimple**](EventApi.md#GetEventSimple) | **Get** /event/{event_key}/simple | 
[**GetEventTeams**](EventApi.md#GetEventTeams) | **Get** /event/{event_key}/teams | 
[**GetEventTeamsKeys**](EventApi.md#GetEventTeamsKeys) | **Get** /event/{event_key}/teams/keys | 
[**GetEventTeamsSimple**](EventApi.md#GetEventTeamsSimple) | **Get** /event/{event_key}/teams/simple | 
[**GetEventTeamsStatuses**](EventApi.md#GetEventTeamsStatuses) | **Get** /event/{event_key}/teams/statuses | 
[**GetEventsByYear**](EventApi.md#GetEventsByYear) | **Get** /events/{year} | 
[**GetEventsByYearKeys**](EventApi.md#GetEventsByYearKeys) | **Get** /events/{year}/keys | 
[**GetEventsByYearSimple**](EventApi.md#GetEventsByYearSimple) | **Get** /events/{year}/simple | 
[**GetTeamEventAwards**](EventApi.md#GetTeamEventAwards) | **Get** /team/{team_key}/event/{event_key}/awards | 
[**GetTeamEventMatches**](EventApi.md#GetTeamEventMatches) | **Get** /team/{team_key}/event/{event_key}/matches | 
[**GetTeamEventMatchesKeys**](EventApi.md#GetTeamEventMatchesKeys) | **Get** /team/{team_key}/event/{event_key}/matches/keys | 
[**GetTeamEventMatchesSimple**](EventApi.md#GetTeamEventMatchesSimple) | **Get** /team/{team_key}/event/{event_key}/matches/simple | 
[**GetTeamEventStatus**](EventApi.md#GetTeamEventStatus) | **Get** /team/{team_key}/event/{event_key}/status | 
[**GetTeamEvents**](EventApi.md#GetTeamEvents) | **Get** /team/{team_key}/events | 
[**GetTeamEventsByYear**](EventApi.md#GetTeamEventsByYear) | **Get** /team/{team_key}/events/{year} | 
[**GetTeamEventsByYearKeys**](EventApi.md#GetTeamEventsByYearKeys) | **Get** /team/{team_key}/events/{year}/keys | 
[**GetTeamEventsByYearSimple**](EventApi.md#GetTeamEventsByYearSimple) | **Get** /team/{team_key}/events/{year}/simple | 
[**GetTeamEventsKeys**](EventApi.md#GetTeamEventsKeys) | **Get** /team/{team_key}/events/keys | 
[**GetTeamEventsSimple**](EventApi.md#GetTeamEventsSimple) | **Get** /team/{team_key}/events/simple | 
[**GetTeamEventsStatusesByYear**](EventApi.md#GetTeamEventsStatusesByYear) | **Get** /team/{team_key}/events/{year}/statuses | 



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


## GetEvent

> Event GetEvent(ctx, )


Gets an Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventAlliances

> []EliminationAlliance GetEventAlliances(ctx, )


Gets a list of Elimination Alliances for the given Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EliminationAlliance**](Elimination_Alliance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventAwards

> []Award GetEventAwards(ctx, )


Gets a list of awards from the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Award**](Award.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventDistrictPoints

> EventDistrictPoints GetEventDistrictPoints(ctx, )


Gets a list of team rankings for the Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventDistrictPoints**](Event_District_Points.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventInsights

> EventInsights GetEventInsights(ctx, )


Gets a set of Event-specific insights for the given Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventInsights**](Event_Insights.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventMatchTimeseries

> []string GetEventMatchTimeseries(ctx, )


Gets an array of Match Keys for the given event key that have timeseries data. Returns an empty array if no matches have timeseries data. *WARNING:* This is *not* official data, and is subject to a significant possibility of error, or missing data. Do not rely on this data for any purpose. In fact, pretend we made it up. *WARNING:* This endpoint and corresponding data models are under *active development* and may change at any time, including in breaking ways.

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


## GetEventMatches

> []Match GetEventMatches(ctx, )


Gets a list of matches for the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Match**](Match.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventMatchesKeys

> []string GetEventMatchesKeys(ctx, )


Gets a list of match keys for the given event.

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


## GetEventMatchesSimple

> []MatchSimple GetEventMatchesSimple(ctx, )


Gets a short-form list of matches for the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MatchSimple**](Match_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventOPRs

> EventOpRs GetEventOPRs(ctx, )


Gets a set of Event OPRs (including OPR, DPR, and CCWM) for the given Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventOpRs**](Event_OPRs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventPredictions

> map[string]interface{} GetEventPredictions(ctx, )


Gets information on TBA-generated predictions for the given Event. Contains year-specific information. *WARNING* This endpoint is currently under development and may change at any time.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventRankings

> EventRanking GetEventRankings(ctx, )


Gets a list of team rankings for the Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventRanking**](Event_Ranking.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSimple

> EventSimple GetEventSimple(ctx, )


Gets a short-form Event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventSimple**](Event_Simple.md)

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


## GetTeamEventAwards

> []Award GetTeamEventAwards(ctx, )


Gets a list of awards the given team won at the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Award**](Award.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEventMatches

> []Match GetTeamEventMatches(ctx, )


Gets a list of matches for the given team and event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Match**](Match.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEventMatchesKeys

> []string GetTeamEventMatchesKeys(ctx, )


Gets a list of match keys for matches for the given team and event.

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


## GetTeamEventMatchesSimple

> []Match GetTeamEventMatchesSimple(ctx, )


Gets a short-form list of matches for the given team and event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Match**](Match.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEventStatus

> TeamEventStatus GetTeamEventStatus(ctx, )


Gets the competition rank and status of the team at the given event.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TeamEventStatus**](Team_Event_Status.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEvents

> []Event GetTeamEvents(ctx, )


Gets a list of all events this team has competed at.

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


## GetTeamEventsByYear

> []Event GetTeamEventsByYear(ctx, )


Gets a list of events this team has competed at in the given year.

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


## GetTeamEventsByYearKeys

> []string GetTeamEventsByYearKeys(ctx, )


Gets a list of the event keys for events this team has competed at in the given year.

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


## GetTeamEventsByYearSimple

> []EventSimple GetTeamEventsByYearSimple(ctx, )


Gets a short-form list of events this team has competed at in the given year.

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


## GetTeamEventsKeys

> []string GetTeamEventsKeys(ctx, )


Gets a list of the event keys for all events this team has competed at.

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


## GetTeamEventsSimple

> []EventSimple GetTeamEventsSimple(ctx, )


Gets a short-form list of all events this team has competed at.

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

