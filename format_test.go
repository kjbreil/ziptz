package ziptz

import (
	"testing"
	"time"
)

func TestIANA(t *testing.T) {
	// IANA is just an alias for Lookup
	if got := IANA("60601"); got != "America/Chicago" {
		t.Errorf("IANA(60601) = %q, want America/Chicago", got)
	}
	if got := IANA("invalid"); got != "" {
		t.Errorf("IANA(invalid) = %q, want empty string", got)
	}
}

func TestLocation(t *testing.T) {
	loc, err := Location("60601")
	if err != nil {
		t.Fatalf("Location(60601) error: %v", err)
	}
	if loc.String() != "America/Chicago" {
		t.Errorf("Location(60601) = %q, want America/Chicago", loc.String())
	}

	// Invalid zip
	_, err = Location("invalid")
	if err == nil {
		t.Error("Location(invalid) should return error")
	}

	// Unknown zip
	_, err = Location("00000")
	if err == nil {
		t.Error("Location(00000) should return error")
	}
}

func TestAbbreviationAt(t *testing.T) {
	// Winter time (standard time) - January
	winter := time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC)

	// Chicago in winter should be CST
	abbr := AbbreviationAt("60601", winter)
	if abbr != "CST" {
		t.Errorf("AbbreviationAt(60601, winter) = %q, want CST", abbr)
	}

	// Summer time (daylight saving time) - July
	summer := time.Date(2024, time.July, 15, 12, 0, 0, 0, time.UTC)

	// Chicago in summer should be CDT
	abbr = AbbreviationAt("60601", summer)
	if abbr != "CDT" {
		t.Errorf("AbbreviationAt(60601, summer) = %q, want CDT", abbr)
	}

	// Arizona doesn't observe DST (except Navajo Nation)
	abbr = AbbreviationAt("85001", summer)
	if abbr != "MST" {
		t.Errorf("AbbreviationAt(85001, summer) = %q, want MST", abbr)
	}

	// Invalid zip
	if got := AbbreviationAt("invalid", winter); got != "" {
		t.Errorf("AbbreviationAt(invalid, winter) = %q, want empty", got)
	}
}

func TestOffsetAt(t *testing.T) {
	// Winter time - January
	winter := time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC)

	// Chicago in winter (CST = UTC-6)
	offset := OffsetAt("60601", winter)
	if offset != "-06:00" {
		t.Errorf("OffsetAt(60601, winter) = %q, want -06:00", offset)
	}

	// Summer time - July
	summer := time.Date(2024, time.July, 15, 12, 0, 0, 0, time.UTC)

	// Chicago in summer (CDT = UTC-5)
	offset = OffsetAt("60601", summer)
	if offset != "-05:00" {
		t.Errorf("OffsetAt(60601, summer) = %q, want -05:00", offset)
	}

	// New York in winter (EST = UTC-5)
	offset = OffsetAt("10001", winter)
	if offset != "-05:00" {
		t.Errorf("OffsetAt(10001, winter) = %q, want -05:00", offset)
	}

	// Los Angeles in winter (PST = UTC-8)
	offset = OffsetAt("90210", winter)
	if offset != "-08:00" {
		t.Errorf("OffsetAt(90210, winter) = %q, want -08:00", offset)
	}

	// Invalid zip
	if got := OffsetAt("invalid", winter); got != "" {
		t.Errorf("OffsetAt(invalid, winter) = %q, want empty", got)
	}
}

func TestOffsetSecondsAt(t *testing.T) {
	winter := time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC)

	// Chicago in winter (CST = UTC-6 = -21600 seconds)
	offset := OffsetSecondsAt("60601", winter)
	if offset != -21600 {
		t.Errorf("OffsetSecondsAt(60601, winter) = %d, want -21600", offset)
	}

	// Invalid zip
	if got := OffsetSecondsAt("invalid", winter); got != 0 {
		t.Errorf("OffsetSecondsAt(invalid, winter) = %d, want 0", got)
	}
}

func TestIsDSTAt(t *testing.T) {
	winter := time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC)
	summer := time.Date(2024, time.July, 15, 12, 0, 0, 0, time.UTC)

	// Chicago observes DST
	if IsDSTAt("60601", winter) {
		t.Error("IsDSTAt(60601, winter) = true, want false")
	}
	if !IsDSTAt("60601", summer) {
		t.Error("IsDSTAt(60601, summer) = false, want true")
	}

	// Arizona (Phoenix) does not observe DST
	if IsDSTAt("85001", winter) {
		t.Error("IsDSTAt(85001, winter) = true, want false")
	}
	if IsDSTAt("85001", summer) {
		t.Error("IsDSTAt(85001, summer) = true, want false")
	}

	// Invalid zip
	if IsDSTAt("invalid", summer) {
		t.Error("IsDSTAt(invalid, summer) = true, want false")
	}
}

func BenchmarkLocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Location("60601")
	}
}

func BenchmarkAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abbreviation("60601")
	}
}
