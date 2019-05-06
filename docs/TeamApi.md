# \TeamApi

All URIs are relative to *https://www.thebluealliance.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDistrictRankings**](TeamApi.md#GetDistrictRankings) | **Get** /district/{district_key}/rankings | 
[**GetDistrictTeams**](TeamApi.md#GetDistrictTeams) | **Get** /district/{district_key}/teams | 
[**GetDistrictTeamsKeys**](TeamApi.md#GetDistrictTeamsKeys) | **Get** /district/{district_key}/teams/keys | 
[**GetDistrictTeamsSimple**](TeamApi.md#GetDistrictTeamsSimple) | **Get** /district/{district_key}/teams/simple | 
[**GetEventTeams**](TeamApi.md#GetEventTeams) | **Get** /event/{event_key}/teams | 
[**GetEventTeamsKeys**](TeamApi.md#GetEventTeamsKeys) | **Get** /event/{event_key}/teams/keys | 
[**GetEventTeamsSimple**](TeamApi.md#GetEventTeamsSimple) | **Get** /event/{event_key}/teams/simple | 
[**GetEventTeamsStatuses**](TeamApi.md#GetEventTeamsStatuses) | **Get** /event/{event_key}/teams/statuses | 
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /team/{team_key} | 
[**GetTeamAwards**](TeamApi.md#GetTeamAwards) | **Get** /team/{team_key}/awards | 
[**GetTeamAwardsByYear**](TeamApi.md#GetTeamAwardsByYear) | **Get** /team/{team_key}/awards/{year} | 
[**GetTeamDistricts**](TeamApi.md#GetTeamDistricts) | **Get** /team/{team_key}/districts | 
[**GetTeamEventAwards**](TeamApi.md#GetTeamEventAwards) | **Get** /team/{team_key}/event/{event_key}/awards | 
[**GetTeamEventMatches**](TeamApi.md#GetTeamEventMatches) | **Get** /team/{team_key}/event/{event_key}/matches | 
[**GetTeamEventMatchesKeys**](TeamApi.md#GetTeamEventMatchesKeys) | **Get** /team/{team_key}/event/{event_key}/matches/keys | 
[**GetTeamEventMatchesSimple**](TeamApi.md#GetTeamEventMatchesSimple) | **Get** /team/{team_key}/event/{event_key}/matches/simple | 
[**GetTeamEventStatus**](TeamApi.md#GetTeamEventStatus) | **Get** /team/{team_key}/event/{event_key}/status | 
[**GetTeamEvents**](TeamApi.md#GetTeamEvents) | **Get** /team/{team_key}/events | 
[**GetTeamEventsByYear**](TeamApi.md#GetTeamEventsByYear) | **Get** /team/{team_key}/events/{year} | 
[**GetTeamEventsByYearKeys**](TeamApi.md#GetTeamEventsByYearKeys) | **Get** /team/{team_key}/events/{year}/keys | 
[**GetTeamEventsByYearSimple**](TeamApi.md#GetTeamEventsByYearSimple) | **Get** /team/{team_key}/events/{year}/simple | 
[**GetTeamEventsKeys**](TeamApi.md#GetTeamEventsKeys) | **Get** /team/{team_key}/events/keys | 
[**GetTeamEventsSimple**](TeamApi.md#GetTeamEventsSimple) | **Get** /team/{team_key}/events/simple | 
[**GetTeamEventsStatusesByYear**](TeamApi.md#GetTeamEventsStatusesByYear) | **Get** /team/{team_key}/events/{year}/statuses | 
[**GetTeamMatchesByYear**](TeamApi.md#GetTeamMatchesByYear) | **Get** /team/{team_key}/matches/{year} | 
[**GetTeamMatchesByYearKeys**](TeamApi.md#GetTeamMatchesByYearKeys) | **Get** /team/{team_key}/matches/{year}/keys | 
[**GetTeamMatchesByYearSimple**](TeamApi.md#GetTeamMatchesByYearSimple) | **Get** /team/{team_key}/matches/{year}/simple | 
[**GetTeamMediaByTag**](TeamApi.md#GetTeamMediaByTag) | **Get** /team/{team_key}/media/tag/{media_tag} | 
[**GetTeamMediaByTagYear**](TeamApi.md#GetTeamMediaByTagYear) | **Get** /team/{team_key}/media/tag/{media_tag}/{year} | 
[**GetTeamMediaByYear**](TeamApi.md#GetTeamMediaByYear) | **Get** /team/{team_key}/media/{year} | 
[**GetTeamRobots**](TeamApi.md#GetTeamRobots) | **Get** /team/{team_key}/robots | 
[**GetTeamSimple**](TeamApi.md#GetTeamSimple) | **Get** /team/{team_key}/simple | 
[**GetTeamSocialMedia**](TeamApi.md#GetTeamSocialMedia) | **Get** /team/{team_key}/social_media | 
[**GetTeamYearsParticipated**](TeamApi.md#GetTeamYearsParticipated) | **Get** /team/{team_key}/years_participated | 
[**GetTeams**](TeamApi.md#GetTeams) | **Get** /teams/{page_num} | 
[**GetTeamsByYear**](TeamApi.md#GetTeamsByYear) | **Get** /teams/{year}/{page_num} | 
[**GetTeamsByYearKeys**](TeamApi.md#GetTeamsByYearKeys) | **Get** /teams/{year}/{page_num}/keys | 
[**GetTeamsByYearSimple**](TeamApi.md#GetTeamsByYearSimple) | **Get** /teams/{year}/{page_num}/simple | 
[**GetTeamsKeys**](TeamApi.md#GetTeamsKeys) | **Get** /teams/{page_num}/keys | 
[**GetTeamsSimple**](TeamApi.md#GetTeamsSimple) | **Get** /teams/{page_num}/simple | 


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

# **GetTeam**
> Team GetTeam(ctx, teamKey, optional)


Gets a `Team` object for the team referenced by the given key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**Team**](Team.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamAwards**
> []Award GetTeamAwards(ctx, teamKey, optional)


Gets a list of awards the given team has won.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamAwardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamAwardsOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamAwardsByYear**
> []Award GetTeamAwardsByYear(ctx, teamKey, year, optional)


Gets a list of awards the given team has won in a given year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamAwardsByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamAwardsByYearOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamDistricts**
> []DistrictList GetTeamDistricts(ctx, teamKey, optional)


Gets an array of districts representing each year the team was in a district. Will return an empty array if the team was never in a district.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamDistrictsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamDistrictsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]DistrictList**](District_List.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventAwards**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventMatches**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventMatchesKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventMatchesSimple**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventStatus**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEvents**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsByYear**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsByYearKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsByYearSimple**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsKeys**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamEventsSimple**
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

# **GetTeamMatchesByYear**
> []Match GetTeamMatchesByYear(ctx, teamKey, year, optional)


Gets a list of matches for the given team and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamMatchesByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMatchesByYearOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMatchesByYearKeys**
> []string GetTeamMatchesByYearKeys(ctx, teamKey, year, optional)


Gets a list of match keys for matches for the given team and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamMatchesByYearKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMatchesByYearKeysOpts struct

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

# **GetTeamMatchesByYearSimple**
> []MatchSimple GetTeamMatchesByYearSimple(ctx, teamKey, year, optional)


Gets a short-form list of matches for the given team and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamMatchesByYearSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMatchesByYearSimpleOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMediaByTag**
> []Media GetTeamMediaByTag(ctx, teamKey, mediaTag, optional)


Gets a list of Media (videos / pictures) for the given team and tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **mediaTag** | **string**| Media Tag which describes the Media. | 
 **optional** | ***GetTeamMediaByTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMediaByTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]Media**](Media.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMediaByTagYear**
> []Media GetTeamMediaByTagYear(ctx, teamKey, mediaTag, year, optional)


Gets a list of Media (videos / pictures) for the given team, tag and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **mediaTag** | **string**| Media Tag which describes the Media. | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamMediaByTagYearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMediaByTagYearOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]Media**](Media.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMediaByYear**
> []Media GetTeamMediaByYear(ctx, teamKey, year, optional)


Gets a list of Media (videos / pictures) for the given team and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
  **year** | **int32**| Competition Year (or Season). Must be 4 digits. | 
 **optional** | ***GetTeamMediaByYearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamMediaByYearOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]Media**](Media.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamRobots**
> []TeamRobot GetTeamRobots(ctx, teamKey, optional)


Gets a list of year and robot name pairs for each year that a robot name was provided. Will return an empty array if the team has never named a robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamRobotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamRobotsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]TeamRobot**](Team_Robot.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamSimple**
> TeamSimple GetTeamSimple(ctx, teamKey, optional)


Gets a `Team_Simple` object for the team referenced by the given key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamSimpleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamSimpleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**TeamSimple**](Team_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamSocialMedia**
> []Media GetTeamSocialMedia(ctx, teamKey, optional)


Gets a list of Media (social media) for the given team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamSocialMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamSocialMediaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

[**[]Media**](Media.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamYearsParticipated**
> []int32 GetTeamYearsParticipated(ctx, teamKey, optional)


Gets a list of years in which the team participated in at least one competition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamKey** | **string**| TBA Team Key, eg &#x60;frc254&#x60; | 
 **optional** | ***GetTeamYearsParticipatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTeamYearsParticipatedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.String**| Value of the &#x60;Last-Modified&#x60; header in the most recently cached response by the client. | 

### Return type

**[]int32**

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

