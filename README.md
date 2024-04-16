# goTimeZoneApi

The endpoint is be exposed as /api/time
<BR>
The output is encoded in JSON <BR>
Example:
<BR>
{<BR>
    "current_time": "2021-08-09 11:18:06 +0000 UTC"<BR>
}<BR>
<BR>


A user can request the current time in another timezone. /api/time?tz=America/New_York
A list of TZ database names is available at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

 A user can request time in multiple timezones i.e.<BR> /api/time?tz=America/New_York,Asia/Kolkata

API Response:

{
    "Asia/Kolkata": "2021-08-09 01:23:42 +0530 IST",<BR>
    "America/New_York": "2021-08-09 01:23:42 -0400 EDT"
}
