# \MatchApi

All URIs are relative to *https://www.thebluealliance.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventMatchTimeseries**](MatchApi.md#GetEventMatchTimeseries) | **Get** /event/{event_key}/matches/timeseries | 
[**GetEventMatches**](MatchApi.md#GetEventMatches) | **Get** /event/{event_key}/matches | 
[**GetEventMatchesKeys**](MatchApi.md#GetEventMatchesKeys) | **Get** /event/{event_key}/matches/keys | 
[**GetEventMatchesSimple**](MatchApi.md#GetEventMatchesSimple) | **Get** /event/{event_key}/matches/simple | 
[**GetMatch**](MatchApi.md#GetMatch) | **Get** /match/{match_key} | 
[**GetMatchSimple**](MatchApi.md#GetMatchSimple) | **Get** /match/{match_key}/simple | 
[**GetMatchTimeseries**](MatchApi.md#GetMatchTimeseries) | **Get** /match/{match_key}/timeseries | 
[**GetTeamEventMatches**](MatchApi.md#GetTeamEventMatches) | **Get** /team/{team_key}/event/{event_key}/matches | 
[**GetTeamEventMatchesKeys**](MatchApi.md#GetTeamEventMatchesKeys) | **Get** /team/{team_key}/event/{event_key}/matches/keys | 
[**GetTeamEventMatchesSimple**](MatchApi.md#GetTeamEventMatchesSimple) | **Get** /team/{team_key}/event/{event_key}/matches/simple | 
[**GetTeamMatchesByYear**](MatchApi.md#GetTeamMatchesByYear) | **Get** /team/{team_key}/matches/{year} | 
[**GetTeamMatchesByYearKeys**](MatchApi.md#GetTeamMatchesByYearKeys) | **Get** /team/{team_key}/matches/{year}/keys | 
[**GetTeamMatchesByYearSimple**](MatchApi.md#GetTeamMatchesByYearSimple) | **Get** /team/{team_key}/matches/{year}/simple | 



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


## GetMatch

> Match GetMatch(ctx, )


Gets a `Match` object for the given match key.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Match**](Match.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchSimple

> MatchSimple GetMatchSimple(ctx, )


Gets a short-form `Match` object for the given match key.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**MatchSimple**](Match_Simple.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchTimeseries

> []map[string]interface{} GetMatchTimeseries(ctx, )


Gets an array of game-specific Match Timeseries objects for the given match key or an empty array if not available. *WARNING:* This is *not* official data, and is subject to a significant possibility of error, or missing data. Do not rely on this data for any purpose. In fact, pretend we made it up. *WARNING:* This endpoint and corresponding data models are under *active development* and may change at any time, including in breaking ways.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

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


## GetTeamMatchesByYear

> []Match GetTeamMatchesByYear(ctx, )


Gets a list of matches for the given team and year.

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


## GetTeamMatchesByYearKeys

> []string GetTeamMatchesByYearKeys(ctx, )


Gets a list of match keys for matches for the given team and year.

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


## GetTeamMatchesByYearSimple

> []MatchSimple GetTeamMatchesByYearSimple(ctx, )


Gets a short-form list of matches for the given team and year.

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

