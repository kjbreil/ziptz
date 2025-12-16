package ziptz

import (
	"testing"
)

func TestIsValidFormat(t *testing.T) {
	tests := []struct {
		zip  string
		want bool
	}{
		{"60601", true},
		{"00501", true},
		{"99950", true},
		{"1234", false},
		{"123456", false},
		{"abcde", false},
		{"1234a", false},
		{"", false},
		{"60601-1234", false},
	}

	for _, tt := range tests {
		t.Run(tt.zip, func(t *testing.T) {
			if got := IsValidFormat(tt.zip); got != tt.want {
				t.Errorf("IsValidFormat(%q) = %v, want %v", tt.zip, got, tt.want)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	tests := []struct {
		zip  string
		want string
	}{
		{"60601", "America/Chicago"},
		{"10001", "America/New_York"},
		{"90210", "America/Los_Angeles"},
		{"85001", "America/Phoenix"},
		{"00501", "America/New_York"},
		{"96701", "Pacific/Honolulu"},
		{"99501", "America/Anchorage"},
		// Invalid formats
		{"1234", ""},
		{"abcde", ""},
		{"", ""},
		// Unknown zip
		{"00000", ""},
	}

	for _, tt := range tests {
		t.Run(tt.zip, func(t *testing.T) {
			if got := Lookup(tt.zip); got != tt.want {
				t.Errorf("Lookup(%q) = %q, want %q", tt.zip, got, tt.want)
			}
		})
	}
}

func TestLookupWithOk(t *testing.T) {
	// Valid zip
	tz, ok := LookupWithOk("60601")
	if !ok || tz != "America/Chicago" {
		t.Errorf("LookupWithOk(60601) = (%q, %v), want (America/Chicago, true)", tz, ok)
	}

	// Invalid format
	tz, ok = LookupWithOk("1234")
	if ok || tz != "" {
		t.Errorf("LookupWithOk(1234) = (%q, %v), want (\"\", false)", tz, ok)
	}

	// Unknown zip
	tz, ok = LookupWithOk("00000")
	if ok || tz != "" {
		t.Errorf("LookupWithOk(00000) = (%q, %v), want (\"\", false)", tz, ok)
	}
}

func TestMustLookup(t *testing.T) {
	// Valid zip should not panic
	tz := MustLookup("60601")
	if tz != "America/Chicago" {
		t.Errorf("MustLookup(60601) = %q, want America/Chicago", tz)
	}

	// Invalid format should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustLookup(1234) should have panicked")
		}
	}()
	MustLookup("1234")
}

func TestMustLookupUnknown(t *testing.T) {
	// Unknown zip should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustLookup(00000) should have panicked")
		}
	}()
	MustLookup("00000")
}

func TestCount(t *testing.T) {
	count := Count()
	if count < 40000 || count > 50000 {
		t.Errorf("Count() = %d, expected between 40000 and 50000", count)
	}
}

func BenchmarkLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lookup("60601")
	}
}

func BenchmarkIsValidFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidFormat("60601")
	}
}
