# \DistrictApi

All URIs are relative to *https://www.thebluealliance.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDistrictEvents**](DistrictApi.md#GetDistrictEvents) | **Get** /district/{district_key}/events | 
[**GetDistrictEventsKeys**](DistrictApi.md#GetDistrictEventsKeys) | **Get** /district/{district_key}/events/keys | 
[**GetDistrictEventsSimple**](DistrictApi.md#GetDistrictEventsSimple) | **Get** /district/{district_key}/events/simple | 
[**GetDistrictRankings**](DistrictApi.md#GetDistrictRankings) | **Get** /district/{district_key}/rankings | 
[**GetDistrictTeams**](DistrictApi.md#GetDistrictTeams) | **Get** /district/{district_key}/teams | 
[**GetDistrictTeamsKeys**](DistrictApi.md#GetDistrictTeamsKeys) | **Get** /district/{district_key}/teams/keys | 
[**GetDistrictTeamsSimple**](DistrictApi.md#GetDistrictTeamsSimple) | **Get** /district/{district_key}/teams/simple | 
[**GetDistrictsByYear**](DistrictApi.md#GetDistrictsByYear) | **Get** /districts/{year} | 
[**GetEventDistrictPoints**](DistrictApi.md#GetEventDistrictPoints) | **Get** /event/{event_key}/district_points | 
[**GetTeamDistricts**](DistrictApi.md#GetTeamDistricts) | **Get** /team/{team_key}/districts | 



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


## GetDistrictsByYear

> []DistrictList GetDistrictsByYear(ctx, )


Gets a list of districts and their corresponding district key, for the given year.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DistrictList**](District_List.md)

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


## GetTeamDistricts

> []DistrictList GetTeamDistricts(ctx, )


Gets an array of districts representing each year the team was in a district. Will return an empty array if the team was never in a district.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DistrictList**](District_List.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

