package utils

import "time"

// Sun, 06 Nov 1994 08:49:37 GMT    ; RFC 822, updated by RFC 1123
func StringToDateRFC1123(str string) time.Time {
	t, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", str)
	return t
}
