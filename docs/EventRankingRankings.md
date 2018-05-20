# EventRankingRankings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dq** | **int32** | Number of times disqualified. | [default to null]
**MatchesPlayed** | **int32** | Number of matches played by this team. | [default to null]
**QualAverage** | **int32** | The average match score during qualifications. Year specific. May be null if not relevant for a given year. | [optional] [default to null]
**Rank** | **int32** | The team&#39;s rank at the event as provided by FIRST. | [default to null]
**Record** | [***WltRecord**](WLT_Record.md) |  | [default to null]
**ExtraStats** | **[]float32** | Additional special data on the team&#39;s performance calculated by TBA. | [optional] [default to null]
**SortOrders** | **[]float32** | Additional year-specific information, may be null. See parent &#x60;sort_order_info&#x60; for details. | [optional] [default to null]
**TeamKey** | **string** | The team with this rank. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


