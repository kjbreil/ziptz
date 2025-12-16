//go:generate go run generate.go

// Package ziptz provides US zip code to IANA timezone lookups.
package ziptz

import (
	"regexp"
)

var zipRegex = regexp.MustCompile(`^\d{5}$`)

// IsValidFormat returns true if the zip code is a valid 5-digit format.
func IsValidFormat(zip string) bool {
	return zipRegex.MatchString(zip)
}

// Lookup returns the IANA timezone for a US zip code (e.g., "America/Chicago").
// Returns an empty string if the zip code is invalid or not found.
func Lookup(zip string) string {
	if !IsValidFormat(zip) {
		return ""
	}
	return zipToTimezone[zip]
}

// LookupWithOk returns the IANA timezone and a boolean indicating if the zip was found.
// Returns ("", false) if the zip code is invalid or not found.
func LookupWithOk(zip string) (string, bool) {
	if !IsValidFormat(zip) {
		return "", false
	}
	tz, ok := zipToTimezone[zip]
	return tz, ok
}

// MustLookup returns the IANA timezone for a US zip code.
// Panics if the zip code is invalid or not found.
func MustLookup(zip string) string {
	tz, ok := LookupWithOk(zip)
	if !ok {
		panic("ziptz: invalid or unknown zip code: " + zip)
	}
	return tz
}

// Count returns the number of zip codes in the database.
func Count() int {
	return len(zipToTimezone)
}
