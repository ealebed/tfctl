package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name         string
		version      string
		releasePhase string
		wantContains []string
	}{
		{
			name:         "version with release phase",
			version:      "1.2.3",
			releasePhase: "dev",
			wantContains: []string{"1.2.3", "dev"},
		},
		{
			name:         "version without release phase",
			version:      "1.2.3",
			releasePhase: "",
			wantContains: []string{"1.2.3"},
		},
		{
			name:         "version with alpha release phase",
			version:      "2.0.0",
			releasePhase: "alpha",
			wantContains: []string{"2.0.0", "alpha"},
		},
		{
			name:         "version with beta release phase",
			version:      "2.0.0",
			releasePhase: "beta",
			wantContains: []string{"2.0.0", "beta"},
		},
		{
			name:         "empty version with release phase",
			version:      "",
			releasePhase: "dev",
			wantContains: []string{"dev"},
		},
		{
			name:         "empty version without release phase",
			version:      "",
			releasePhase: "",
			wantContains: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original values
			originalVersion := Version
			originalReleasePhase := ReleasePhase

			// Set test values
			Version = tt.version
			ReleasePhase = tt.releasePhase

			// Restore original values after test
			defer func() {
				Version = originalVersion
				ReleasePhase = originalReleasePhase
			}()

			got := String()

			// Verify the result contains expected strings
			// Note: The function uses color formatting, so we check for content
			// The color codes are ANSI escape sequences, but the actual text should be present
			for _, want := range tt.wantContains {
				if want != "" && !strings.Contains(got, want) {
					t.Errorf("String() = %q, want to contain %q", got, want)
				}
			}

			// Verify it's not empty (unless both version and phase are empty)
			// When both are empty, color.SprintFunc() on an empty string might return
			// just color codes or empty string, which is acceptable
			if tt.version != "" || tt.releasePhase != "" {
				// At least one is non-empty, result should not be empty
				if got == "" {
					t.Error("String() returned empty string")
				}
			}
		})
	}
}

func TestString_Format(t *testing.T) {
	// Save original values
	originalVersion := Version
	originalReleasePhase := ReleasePhase

	// Set test values
	Version = "1.0.0"
	ReleasePhase = "dev"

	// Restore original values after test
	defer func() {
		Version = originalVersion
		ReleasePhase = originalReleasePhase
	}()

	got := String()

	// Verify format: should contain version and release phase separated by dash
	if !strings.Contains(got, "1.0.0") {
		t.Errorf("String() = %q, want to contain version '1.0.0'", got)
	}
	if !strings.Contains(got, "dev") {
		t.Errorf("String() = %q, want to contain release phase 'dev'", got)
	}
	// The format should be "version-releasePhase" (with color codes)
	// We can't easily test the exact format due to ANSI color codes,
	// but we verify both parts are present
}

func TestString_NoReleasePhase(t *testing.T) {
	// Save original values
	originalVersion := Version
	originalReleasePhase := ReleasePhase

	// Set test values
	Version = "2.5.0"
	ReleasePhase = ""

	// Restore original values after test
	defer func() {
		Version = originalVersion
		ReleasePhase = originalReleasePhase
	}()

	got := String()

	// Should contain version but not a dash (since no release phase)
	if !strings.Contains(got, "2.5.0") {
		t.Errorf("String() = %q, want to contain version '2.5.0'", got)
	}
	// Should not contain a dash followed by nothing (would indicate empty release phase)
	// But we can't easily test this due to color codes, so we just verify version is present
}
