package naming

import (
	"strings"
	"testing"
)

func TestIsDNS1123Label(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid lowercase alphanumeric",
			input: "valid-label-123", //nolint:goconst
			want:  true,
		},
		{
			name:  "valid single character",
			input: "a",
			want:  true,
		},
		{
			name:  "valid with hyphens",
			input: "test-123-abc",
			want:  true,
		},
		{
			name:  "invalid uppercase",
			input: "Invalid-Label",
			want:  false,
		},
		{
			name:  "invalid starts with hyphen",
			input: "-invalid",
			want:  false,
		},
		{
			name:  "invalid ends with hyphen",
			input: "invalid-",
			want:  false,
		},
		{
			name:  "invalid special characters",
			input: "invalid_label",
			want:  false,
		},
		{
			name:  "invalid empty string",
			input: "",
			want:  false,
		},
		{
			name:  "invalid too long",
			input: strings.Repeat("a", 64),
			want:  false,
		},
		{
			name:  "valid max length",
			input: strings.Repeat("a", 63),
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDNS1123Label(tt.input); got != tt.want {
				t.Errorf("IsDNS1123Label(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSanitizeDNS1123Label(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "already valid lowercase",
			input: "valid-label-123", //nolint:goconst
			want:  "valid-label-123", //nolint:goconst
		},
		{
			name:  "uppercase to lowercase",
			input: "UPPERCASE-LABEL",
			want:  "uppercase-label",
		},
		{
			name:  "mixed case to lowercase",
			input: "MixedCase123",
			want:  "mixedcase123",
		},
		{
			name:  "underscores to hyphens",
			input: "label_with_underscores",
			want:  "label-with-underscores",
		},
		{
			name:  "special characters removed",
			input: "label@with#special$chars",
			want:  "label-with-special-chars",
		},
		{
			name:  "leading hyphens removed",
			input: "---leading-hyphens",
			want:  "leading-hyphens",
		},
		{
			name:  "trailing hyphens removed",
			input: "trailing-hyphens---",
			want:  "trailing-hyphens",
		},
		{
			name:  "multiple consecutive hyphens preserved",
			input: "label--with--double--hyphens",
			want:  "label--with--double--hyphens",
		},
		{
			name:  "truncate to 63 chars",
			input: strings.Repeat("a", 100),
			want:  strings.Repeat("a", 63),
		},
		{
			name:  "truncate and remove trailing hyphen",
			input: strings.Repeat("a", 62) + "-" + strings.Repeat("b", 10),
			want:  strings.Repeat("a", 62),
		},
		{
			name:  "spaces to hyphens",
			input: "label with spaces",
			want:  "label-with-spaces",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "only special characters",
			input: "@#$%",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeDNS1123Label(tt.input); got != tt.want {
				t.Errorf("SanitizeDNS1123Label(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSanitizeDNS1123Label_AlwaysValid(t *testing.T) {
	// Property-based test: sanitized output should always be valid or empty
	inputs := []string{
		"ValidLabel123",
		"UPPERCASE",
		"invalid_chars!@#",
		strings.Repeat("a", 100),
		"---leading",
		"trailing---",
		"  spaces  ",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			sanitized := SanitizeDNS1123Label(input)
			if sanitized != "" && !IsDNS1123Label(sanitized) {
				t.Errorf("SanitizeDNS1123Label(%q) = %q, which is not a valid DNS-1123 label", input, sanitized)
			}
		})
	}
}
