package utils

import "time"

func GetDateFromString(dateString time.Time) string {
	if dateString.IsZero() {
		dateString = time.Now()
	}
	theTime := dateString.Format("2006/01/02")
	return theTime
}

func GetTimeFromString(dateString time.Time) string {
	if dateString.IsZero() {
		dateString = time.Now()
	}
	theTime := dateString.Format("15:04:05")
	return theTime
}
