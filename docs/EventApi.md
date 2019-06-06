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

> []Event GetDistrictEvents(ctx, districtKey, optional)


Gets a list of events in the given district.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDistrictEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetDistrictEventsKeys(ctx, districtKey, optional)


Gets a list of event keys for events in the given district.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictEventsKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDistrictEventsKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []EventSimple GetDistrictEventsSimple(ctx, districtKey, optional)


Gets a short-form list of events in the given district.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictEventsSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDistrictEventsSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> Event GetEvent(ctx, eventKey, optional)


Gets an Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []EliminationAlliance GetEventAlliances(ctx, eventKey, optional)


Gets a list of Elimination Alliances for the given Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventAlliancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventAlliancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Award GetEventAwards(ctx, eventKey, optional)


Gets a list of awards from the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventAwardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventAwardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> EventDistrictPoints GetEventDistrictPoints(ctx, eventKey, optional)


Gets a list of team rankings for the Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventDistrictPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventDistrictPointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> EventInsights GetEventInsights(ctx, eventKey, optional)


Gets a set of Event-specific insights for the given Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetEventMatchTimeseries(ctx, eventKey, optional)


Gets an array of Match Keys for the given event key that have timeseries data. Returns an empty array if no matches have timeseries data. *WARNING:* This is *not* official data, and is subject to a significant possibility of error, or missing data. Do not rely on this data for any purpose. In fact, pretend we made it up. *WARNING:* This endpoint and corresponding data models are under *active development* and may change at any time, including in breaking ways.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventMatchTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventMatchTimeseriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Match GetEventMatches(ctx, eventKey, optional)


Gets a list of matches for the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventMatchesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetEventMatchesKeys(ctx, eventKey, optional)


Gets a list of match keys for the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventMatchesKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventMatchesKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []MatchSimple GetEventMatchesSimple(ctx, eventKey, optional)


Gets a short-form list of matches for the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventMatchesSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventMatchesSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> EventOpRs GetEventOPRs(ctx, eventKey, optional)


Gets a set of Event OPRs (including OPR, DPR, and CCWM) for the given Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventOPRsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventOPRsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> map[string]interface{} GetEventPredictions(ctx, eventKey, optional)


Gets information on TBA-generated predictions for the given Event. Contains year-specific information. *WARNING* This endpoint is currently under development and may change at any time.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventPredictionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventPredictionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> EventRanking GetEventRankings(ctx, eventKey, optional)


Gets a list of team rankings for the Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventRankingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventRankingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> EventSimple GetEventSimple(ctx, eventKey, optional)


Gets a short-form Event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Team GetEventTeams(ctx, eventKey, optional)


Gets a list of `Team` objects that competed in the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventTeamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetEventTeamsKeys(ctx, eventKey, optional)


Gets a list of `Team` keys that competed in the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventTeamsKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventTeamsKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []TeamSimple GetEventTeamsSimple(ctx, eventKey, optional)


Gets a short-form list of `Team` objects that competed in the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventTeamsSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventTeamsSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> map[string]TeamEventStatus GetEventTeamsStatuses(ctx, eventKey, optional)


Gets a key-value list of the event statuses for teams competing at the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetEventTeamsStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventTeamsStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Event GetEventsByYear(ctx, year, optional)


Gets a list of events in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetEventsByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventsByYearOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetEventsByYearKeys(ctx, year, optional)


Gets a list of event keys in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetEventsByYearKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventsByYearKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []EventSimple GetEventsByYearSimple(ctx, year, optional)


Gets a short-form list of events in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetEventsByYearSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventsByYearSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Award GetTeamEventAwards(ctx, teamKey, eventKey, optional)


Gets a list of awards the given team won at the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetTeamEventAwardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventAwardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Match GetTeamEventMatches(ctx, teamKey, eventKey, optional)


Gets a list of matches for the given team and event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetTeamEventMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventMatchesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetTeamEventMatchesKeys(ctx, teamKey, eventKey, optional)


Gets a list of match keys for matches for the given team and event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetTeamEventMatchesKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventMatchesKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Match GetTeamEventMatchesSimple(ctx, teamKey, eventKey, optional)


Gets a short-form list of matches for the given team and event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetTeamEventMatchesSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventMatchesSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> TeamEventStatus GetTeamEventStatus(ctx, teamKey, eventKey, optional)


Gets the competition rank and status of the team at the given event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**eventKey** | **string**| TBA Event Key, eg &#x60;2016nytr&#x60; | 
 **optional** | ***GetTeamEventStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Event GetTeamEvents(ctx, teamKey, optional)


Gets a list of all events this team has competed at.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []Event GetTeamEventsByYear(ctx, teamKey, year, optional)


Gets a list of events this team has competed at in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamEventsByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsByYearOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetTeamEventsByYearKeys(ctx, teamKey, year, optional)


Gets a list of the event keys for events this team has competed at in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamEventsByYearKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsByYearKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []EventSimple GetTeamEventsByYearSimple(ctx, teamKey, year, optional)


Gets a short-form list of events this team has competed at in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamEventsByYearSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsByYearSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []string GetTeamEventsKeys(ctx, teamKey, optional)


Gets a list of the event keys for all events this team has competed at.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamEventsKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> []EventSimple GetTeamEventsSimple(ctx, teamKey, optional)


Gets a short-form list of all events this team has competed at.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamEventsSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsSimpleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

> map[string]TeamEventStatus GetTeamEventsStatusesByYear(ctx, teamKey, year, optional)


Gets a key-value list of the event statuses for events this team has competed at in the given year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
**year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamEventsStatusesByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTeamEventsStatusesByYearOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

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

