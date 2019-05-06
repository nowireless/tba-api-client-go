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


# **GetDistrictEvents**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictEventsKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictEventsSimple**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictRankings**
> []DistrictRanking GetDistrictRankings(ctx, districtKey, optional)


Gets a list of team district rankings for the given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictRankingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDistrictRankingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]DistrictRanking**](District_Ranking.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictTeams**
> []Team GetDistrictTeams(ctx, districtKey, optional)


Gets a list of `Team` objects that competed in events in the given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDistrictTeamsOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictTeamsKeys**
> []string GetDistrictTeamsKeys(ctx, districtKey, optional)


Gets a list of `Team` objects that competed in events in the given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictTeamsKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDistrictTeamsKeysOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictTeamsSimple**
> []TeamSimple GetDistrictTeamsSimple(ctx, districtKey, optional)


Gets a short-form list of `Team` objects that competed in events in the given district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **districtKey** | **string**| TBA District Key, eg &#x60;2016fim&#x60; | 
 **optional** | ***GetDistrictTeamsSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDistrictTeamsSimpleOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventTeams**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventTeamsKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventTeamsSimple**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventTeamsStatuses**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsByYear**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsByYearKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsByYearSimple**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsStatusesByYear**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeams**
> []Team GetTeams(ctx, pageNum, optional)


Gets a list of `Team` objects, paginated in groups of 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsByYear**
> []Team GetTeamsByYear(ctx, year, pageNum, optional)


Gets a list of `Team` objects that competed in the given year, paginated in groups of 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsByYearOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsByYearKeys**
> []string GetTeamsByYearKeys(ctx, year, pageNum, optional)


Gets a list Team Keys that competed in the given year, paginated in groups of 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsByYearKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsByYearKeysOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsByYearSimple**
> []TeamSimple GetTeamsByYearSimple(ctx, year, pageNum, optional)


Gets a list of short form `Team_Simple` objects that competed in the given year, paginated in groups of 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsByYearSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsByYearSimpleOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsKeys**
> []string GetTeamsKeys(ctx, pageNum, optional)


Gets a list of Team keys, paginated in groups of 500. (Note, each page will not have 500 teams, but will include the teams within that range of 500.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsKeysOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsSimple**
> []TeamSimple GetTeamsSimple(ctx, pageNum, optional)


Gets a list of short form `Team_Simple` objects, paginated in groups of 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageNum** | **int32**| Page number of results to return, zero-indexed | 
 **optional** | ***GetTeamsSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamsSimpleOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

