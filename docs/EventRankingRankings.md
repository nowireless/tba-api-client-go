# EventRankingRankings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchesPlayed** | **int32** | Number of matches played by this team. | 
**QualAverage** | **int32** | The average match score during qualifications. Year specific. May be null if not relevant for a given year. | [optional] 
**ExtraStats** | **[]float32** | Additional special data on the team&#39;s performance calculated by TBA. | [optional] 
**SortOrders** | **[]float32** | Additional year-specific information, may be null. See parent &#x60;sort_order_info&#x60; for details. | [optional] 
**Record** | [**WltRecord**](WLT_Record.md) |  | 
**Rank** | **int32** | The team&#39;s rank at the event as provided by FIRST. | 
**Dq** | **int32** | Number of times disqualified. | 
**TeamKey** | **string** | The team with this rank. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


