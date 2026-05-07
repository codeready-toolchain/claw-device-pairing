package client

import (
	"os"
	"testing"
)

func TestNewManager_EnvironmentVariables(t *testing.T) {
	tests := []struct {
		name         string
		namespaceEnv string
		instanceEnv  string
		expectError  bool
		errorMsg     string
	}{
		{
			name:         "missing namespace",
			namespaceEnv: "",
			instanceEnv:  "test-instance",
			expectError:  true,
			errorMsg:     "NAMESPACE environment variable is required",
		},
		{
			name:         "missing instance",
			namespaceEnv: "test-namespace",
			instanceEnv:  "",
			expectError:  true,
			errorMsg:     "CLAW_INSTANCE environment variable is required",
		},
		{
			name:         "both missing",
			namespaceEnv: "",
			instanceEnv:  "",
			expectError:  true,
			errorMsg:     "NAMESPACE environment variable is required",
		},
		{
			name:         "env vars set but no in-cluster config",
			namespaceEnv: "test-namespace",
			instanceEnv:  "test-instance",
			expectError:  true,
			errorMsg:     "failed to load in-cluster configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			if tt.namespaceEnv != "" {
				os.Setenv("NAMESPACE", tt.namespaceEnv)
				defer os.Unsetenv("NAMESPACE")
			} else {
				os.Unsetenv("NAMESPACE")
			}

			if tt.instanceEnv != "" {
				os.Setenv("CLAW_INSTANCE", tt.instanceEnv)
				defer os.Unsetenv("CLAW_INSTANCE")
			} else {
				os.Unsetenv("CLAW_INSTANCE")
			}

			// Create manager
			_, err := NewManager()

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
					return
				}
				if !contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing %q, got %q", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsInner(s, substr)))
}

func containsInner(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
