package naming

import (
	"regexp"
	"strings"
)

const (
	// DNS1123LabelMaxLength is the maximum length for a DNS-1123 label (63 characters)
	DNS1123LabelMaxLength = 63
)

var (
	// dns1123LabelRegex matches valid DNS-1123 labels (lowercase alphanumeric with hyphens)
	dns1123LabelRegex = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
)

// IsDNS1123Label validates if a string is a valid DNS-1123 label
// Must be lowercase alphanumeric characters or '-', start and end with alphanumeric
// Max 63 characters
func IsDNS1123Label(value string) bool {
	if len(value) == 0 || len(value) > DNS1123LabelMaxLength {
		return false
	}
	return dns1123LabelRegex.MatchString(value)
}

// SanitizeDNS1123Label converts a string to a valid DNS-1123 label
// - Converts to lowercase
// - Replaces invalid characters with hyphens
// - Removes leading/trailing hyphens
// - Truncates to 63 characters
func SanitizeDNS1123Label(value string) string {
	// Convert to lowercase
	sanitized := strings.ToLower(value)

	// Replace any non-alphanumeric character (except hyphens) with hyphens
	sanitized = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(sanitized, "-")

	// Remove leading hyphens
	sanitized = strings.TrimLeft(sanitized, "-")

	// Remove trailing hyphens
	sanitized = strings.TrimRight(sanitized, "-")

	// Truncate to maximum length
	if len(sanitized) > DNS1123LabelMaxLength {
		sanitized = sanitized[:DNS1123LabelMaxLength]
		// Ensure we don't end with a hyphen after truncation
		sanitized = strings.TrimRight(sanitized, "-")
	}

	return sanitized
}
