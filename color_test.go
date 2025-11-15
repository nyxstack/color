package color

import (
	"testing"
)

func TestBasicColors(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{"Reset", Reset, "\x1b[0m"},
		{"Red", Red, "\x1b[31m"},
		{"Green", Green, "\x1b[32m"},
		{"Blue", Blue, "\x1b[34m"},
		{"BrightRed", BrightRed, "\x1b[91m"},
		{"BrightGreen", BrightGreen, "\x1b[92m"},
		{"BrightBlue", BrightBlue, "\x1b[94m"},
		{"Bold", Bold, "\x1b[1m"},
		{"Underline", Underline, "\x1b[4m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.color != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, tt.color)
			}
		})
	}
}

func TestBackgroundColors(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{"BgRed", BgRed, "\x1b[41m"},
		{"BgGreen", BgGreen, "\x1b[42m"},
		{"BgBlue", BgBlue, "\x1b[44m"},
		{"BgBrightRed", BgBrightRed, "\x1b[101m"},
		{"BgBrightGreen", BgBrightGreen, "\x1b[102m"},
		{"BgBrightBlue", BgBrightBlue, "\x1b[104m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.color != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, tt.color)
			}
		})
	}
}

func TestColor256(t *testing.T) {
	tests := []struct {
		input    uint8
		expected string
	}{
		{0, "\x1b[38;5;0m"},
		{1, "\x1b[38;5;1m"},
		{15, "\x1b[38;5;15m"},
		{255, "\x1b[38;5;255m"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := Color256(tt.input)
			if result != tt.expected {
				t.Errorf("Color256(%d) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRGB(t *testing.T) {
	tests := []struct {
		r, g, b  uint8
		expected string
	}{
		{255, 0, 0, "\x1b[38;2;255;0;0m"},         // Pure red
		{0, 255, 0, "\x1b[38;2;0;255;0m"},         // Pure green
		{0, 0, 255, "\x1b[38;2;0;0;255m"},         // Pure blue
		{128, 128, 128, "\x1b[38;2;128;128;128m"}, // Gray
		{255, 255, 255, "\x1b[38;2;255;255;255m"}, // White
		{0, 0, 0, "\x1b[38;2;0;0;0m"},             // Black
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := RGB(tt.r, tt.g, tt.b)
			if result != tt.expected {
				t.Errorf("RGB(%d, %d, %d) = %q, expected %q",
					tt.r, tt.g, tt.b, result, tt.expected)
			}
		})
	}
}

func TestHex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"FF0000", "\x1b[38;2;255;0;0m"},     // Red
		{"#FF0000", "\x1b[38;2;255;0;0m"},    // Red with hash
		{"00FF00", "\x1b[38;2;0;255;0m"},     // Green
		{"0000FF", "\x1b[38;2;0;0;255m"},     // Blue
		{"FFFFFF", "\x1b[38;2;255;255;255m"}, // White
		{"000000", "\x1b[38;2;0;0;0m"},       // Black
		{"808080", "\x1b[38;2;128;128;128m"}, // Gray
		{"", ""},                             // Empty string
		{"FFF", ""},                          // Invalid length
		{"GGGGGG", "\x1b[38;2;0;0;0m"},       // Invalid hex (defaults to 0,0,0)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Hex(tt.input)
			if result != tt.expected {
				t.Errorf("Hex(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHexEdgeCases(t *testing.T) {
	// Test edge cases
	if result := Hex(""); result != "" {
		t.Errorf("Hex(\"\") should return empty string, got %q", result)
	}

	if result := Hex("ABC"); result != "" {
		t.Errorf("Hex(\"ABC\") should return empty string for invalid length, got %q", result)
	}

	if result := Hex("ABCDEFG"); result != "" {
		t.Errorf("Hex(\"ABCDEFG\") should return empty string for invalid length, got %q", result)
	}
}

func TestColor256Table(t *testing.T) {
	// Test that the color256Table has the correct length
	if len(color256Table) != 256 {
		t.Errorf("color256Table should have 256 entries, got %d", len(color256Table))
	}

	// Test a few specific entries
	expectedEntries := map[int]string{
		0:   "\x1b[38;5;0m",
		1:   "\x1b[38;5;1m",
		255: "\x1b[38;5;255m",
	}

	for index, expected := range expectedEntries {
		if color256Table[index] != expected {
			t.Errorf("color256Table[%d] = %q, expected %q",
				index, color256Table[index], expected)
		}
	}
}

// Benchmark tests
func BenchmarkRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGB(255, 128, 64)
	}
}

func BenchmarkHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hex("FF8040")
	}
}

func BenchmarkColor256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color256(128)
	}
}

// Test color combinations
func TestColorCombinations(t *testing.T) {
	// Test that we can combine colors and styles
	redBold := Red + Bold + "Hello" + Reset
	expectedLength := len(Red) + len(Bold) + len("Hello") + len(Reset)

	if len(redBold) != expectedLength {
		t.Errorf("Color combination length mismatch: got %d, expected %d",
			len(redBold), expectedLength)
	}

	// Test that the combination contains all parts
	if !containsSubstring(redBold, Red) {
		t.Error("Red color code not found in combination")
	}
	if !containsSubstring(redBold, Bold) {
		t.Error("Bold style code not found in combination")
	}
	if !containsSubstring(redBold, "Hello") {
		t.Error("Text not found in combination")
	}
	if !containsSubstring(redBold, Reset) {
		t.Error("Reset code not found in combination")
	}
}

// Helper function to check if a string contains a substring
func containsSubstring(str, substr string) bool {
	return len(str) >= len(substr) &&
		(str == substr ||
			str[:len(substr)] == substr ||
			str[len(str)-len(substr):] == substr ||
			hasSubstring(str, substr))
}

func hasSubstring(str, substr string) bool {
	if len(substr) > len(str) {
		return false
	}
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
