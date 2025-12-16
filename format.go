package ziptz

import (
	"fmt"
	"time"
)

// IANA returns the IANA timezone name for a zip code (e.g., "America/Chicago").
// This is an alias for Lookup.
func IANA(zip string) string {
	return Lookup(zip)
}

// Location returns the *time.Location for a zip code.
// Returns an error if the zip code is invalid, not found, or the timezone cannot be loaded.
func Location(zip string) (*time.Location, error) {
	tz := Lookup(zip)
	if tz == "" {
		return nil, fmt.Errorf("ziptz: invalid or unknown zip code: %s", zip)
	}
	return time.LoadLocation(tz)
}

// Abbreviation returns the timezone abbreviation for a zip code at the current time
// (e.g., "CST", "CDT", "PST", "PDT").
// Returns an empty string if the zip code is invalid or not found.
func Abbreviation(zip string) string {
	return AbbreviationAt(zip, time.Now())
}

// AbbreviationAt returns the timezone abbreviation for a zip code at a specific time.
// Returns an empty string if the zip code is invalid or not found.
func AbbreviationAt(zip string, t time.Time) string {
	loc, err := Location(zip)
	if err != nil {
		return ""
	}
	name, _ := t.In(loc).Zone()
	return name
}

// Offset returns the UTC offset as a string for a zip code at the current time
// (e.g., "-06:00", "-05:00", "+00:00").
// Returns an empty string if the zip code is invalid or not found.
func Offset(zip string) string {
	return OffsetAt(zip, time.Now())
}

// OffsetAt returns the UTC offset as a string for a zip code at a specific time.
// Returns an empty string if the zip code is invalid or not found.
func OffsetAt(zip string, t time.Time) string {
	loc, err := Location(zip)
	if err != nil {
		return ""
	}
	_, offset := t.In(loc).Zone()
	hours := offset / 3600
	minutes := (offset % 3600) / 60
	if minutes < 0 {
		minutes = -minutes
	}
	return fmt.Sprintf("%+03d:%02d", hours, minutes)
}

// OffsetSeconds returns the UTC offset in seconds for a zip code at the current time.
// Returns 0 if the zip code is invalid or not found.
func OffsetSeconds(zip string) int {
	return OffsetSecondsAt(zip, time.Now())
}

// OffsetSecondsAt returns the UTC offset in seconds for a zip code at a specific time.
// Returns 0 if the zip code is invalid or not found.
func OffsetSecondsAt(zip string, t time.Time) int {
	loc, err := Location(zip)
	if err != nil {
		return 0
	}
	_, offset := t.In(loc).Zone()
	return offset
}

// IsDST returns true if the zip code's timezone is currently observing daylight saving time.
// Returns false if the zip code is invalid or not found.
func IsDST(zip string) bool {
	return IsDSTAt(zip, time.Now())
}

// IsDSTAt returns true if the zip code's timezone is observing daylight saving time
// at the specified time.
// Returns false if the zip code is invalid or not found.
func IsDSTAt(zip string, t time.Time) bool {
	loc, err := Location(zip)
	if err != nil {
		return false
	}
	// Get the offset at this time and compare to standard time offset
	// Standard time is typically in January
	jan := time.Date(t.Year(), time.January, 1, 12, 0, 0, 0, loc)
	_, janOffset := jan.Zone()
	_, currentOffset := t.In(loc).Zone()
	return currentOffset != janOffset
}
