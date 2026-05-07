package version

// Version information set at build time via ldflags
var (
	// CommitHash is the git commit hash
	CommitHash = "unknown"

	// BuildTime is the build timestamp
	BuildTime = "unknown"
)
