package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsed, _ := time.Parse(layout, date);
	return parsed;
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// "December 9, 2112 11:59:59"
	layout := "January 2, 2006 15:04:05";
	parsed, _ := time.Parse(layout, date);
	return time.Now().After(parsed);
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Thursday, May 13, 2010 20:32:00
	layout := "Monday, January 2, 2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed.Hour() >= 12 && parsed.Hour() < 18;
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return "You have an appointment on " + Schedule(date).Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year();
	anniversary := fmt.Sprintf("%d-09-15", year);
	parsed, _ := time.Parse("2006-01-02", anniversary);
	return parsed;
}
