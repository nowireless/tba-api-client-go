# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | TBA event key with the format yyyy[EVENT_CODE], where yyyy is the year, and EVENT_CODE is the event code of the event. | [default to null]
**Name** | **string** | Official name of event on record either provided by FIRST or organizers of offseason event. | [default to null]
**EventCode** | **string** | Event short code, as provided by FIRST. | [default to null]
**EventType** | **int32** | Event Type, as defined here: https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/event_type.py#L2 | [default to null]
**District** | [***DistrictList**](District_List.md) | The district this event is in, may be null. | [optional] [default to null]
**City** | **string** | City, town, village, etc. the event is located in. | [optional] [default to null]
**StateProv** | **string** | State or Province the event is located in. | [optional] [default to null]
**Country** | **string** | Country the event is located in. | [optional] [default to null]
**StartDate** | **string** | Event start date in &#x60;yyyy-mm-dd&#x60; format. | [default to null]
**EndDate** | **string** | Event end date in &#x60;yyyy-mm-dd&#x60; format. | [default to null]
**Year** | **int32** | Year the event data is for. | [default to null]
**ShortName** | **string** | Same as &#x60;name&#x60; but doesn&#39;t include event specifiers, such as &#39;Regional&#39; or &#39;District&#39;. May be null. | [optional] [default to null]
**EventTypeString** | **string** | Event Type, eg Regional, District, or Offseason. | [default to null]
**Week** | **int32** | Week of the event relative to the first official season event, zero-indexed. Only valid for Regionals, Districts, and District Championships. Null otherwise. (Eg. A season with a week 0 &#39;preseason&#39; event does not count, and week 1 events will show 0 here. Seasons with a week 0.5 regional event will show week 0 for those event(s) and week 1 for week 1 events and so on.) | [optional] [default to null]
**Address** | **string** | Address of the event&#39;s venue, if available. | [optional] [default to null]
**PostalCode** | **string** | Postal code from the event address. | [optional] [default to null]
**GmapsPlaceId** | **string** | Google Maps Place ID for the event address. | [optional] [default to null]
**GmapsUrl** | **string** | Link to address location on Google Maps. | [optional] [default to null]
**Lat** | **float64** | Latitude for the event address. | [optional] [default to null]
**Lng** | **float64** | Longitude for the event address. | [optional] [default to null]
**LocationName** | **string** | Name of the location at the address for the event, eg. Blue Alliance High School. | [optional] [default to null]
**Timezone** | **string** | Timezone name. | [optional] [default to null]
**Website** | **string** | The event&#39;s website, if any. | [optional] [default to null]
**FirstEventId** | **string** | The FIRST internal Event ID, used to link to the event on the FRC webpage. | [optional] [default to null]
**FirstEventCode** | **string** | Public facing event code used by FIRST (on frc-events.firstinspires.org, for example) | [optional] [default to null]
**Webcasts** | [**[]Webcast**](Webcast.md) |  | [optional] [default to null]
**DivisionKeys** | **[]string** | An array of event keys for the divisions at this event. | [optional] [default to null]
**ParentEventKey** | **string** | The TBA Event key that represents the event&#39;s parent. Used to link back to the event from a division event. It is also the inverse relation of &#x60;divison_keys&#x60;. | [optional] [default to null]
**PlayoffType** | **int32** | Playoff Type, as defined here: https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/playoff_type.py#L4, or null. | [optional] [default to null]
**PlayoffTypeString** | **string** | String representation of the &#x60;playoff_type&#x60;, or null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


