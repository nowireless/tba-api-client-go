# ApiStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentSeason** | **int32** | Year of the current FRC season. | 
**MaxSeason** | **int32** | Maximum FRC season year for valid queries. | 
**IsDatafeedDown** | **bool** | True if the entire FMS API provided by FIRST is down. | 
**DownEvents** | **[]string** | An array of strings containing event keys of any active events that are no longer updating. | 
**Ios** | [**ApiStatusAppVersion**](API_Status_App_Version.md) |  | 
**Android** | [**ApiStatusAppVersion**](API_Status_App_Version.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


