# Team

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | TBA team key with the format &#x60;frcXXXX&#x60; with &#x60;XXXX&#x60; representing the team number. | [default to null]
**TeamNumber** | **int32** | Official team number issued by FIRST. | [default to null]
**Nickname** | **string** | Team nickname provided by FIRST. | [optional] [default to null]
**Name** | **string** | Official long name registered with FIRST. | [default to null]
**City** | **string** | City of team derived from parsing the address registered with FIRST. | [optional] [default to null]
**StateProv** | **string** | State of team derived from parsing the address registered with FIRST. | [optional] [default to null]
**Country** | **string** | Country of team derived from parsing the address registered with FIRST. | [optional] [default to null]
**Address** | **string** | Will be NULL, for future development. | [optional] [default to null]
**PostalCode** | **string** | Postal code from the team address. | [optional] [default to null]
**GmapsPlaceId** | **string** | Will be NULL, for future development. | [optional] [default to null]
**GmapsUrl** | **string** | Will be NULL, for future development. | [optional] [default to null]
**Lat** | **float64** | Will be NULL, for future development. | [optional] [default to null]
**Lng** | **float64** | Will be NULL, for future development. | [optional] [default to null]
**LocationName** | **string** | Will be NULL, for future development. | [optional] [default to null]
**Website** | **string** | Official website associated with the team. | [optional] [default to null]
**RookieYear** | **int32** | First year the team officially competed. | [default to null]
**Motto** | **string** | Team&#39;s motto as provided by FIRST. | [optional] [default to null]
**HomeChampionship** | [***interface{}**](interface{}.md) | Location of the team&#39;s home championship each year as a key-value pair. The year (as a string) is the key, and the city is the value. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


