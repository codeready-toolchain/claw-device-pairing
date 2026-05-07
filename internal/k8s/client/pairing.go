package client

import (
	"context"
	"fmt"
	"log/slog"

	clawv1alpha1 "github.com/codeready-toolchain/claw-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/xcoulon/claw-device-pairing/internal/k8s/naming"
)

// CreatePairingRequest creates a ClawDevicePairingRequest CR
// Returns error if creation fails
func (m *Manager) CreatePairingRequest(ctx context.Context, requestID string) error {
	if !m.IsEnabled() {
		return fmt.Errorf("Kubernetes client not enabled, cannot create pairing request CR")
	}

	// Sanitize requestID for use as CR name
	crName := naming.SanitizeDNS1123Label(requestID)
	if crName == "" {
		return fmt.Errorf("requestID %q cannot be sanitized to valid DNS-1123 label", requestID)
	}

	// Create the CR object
	cr := &clawv1alpha1.ClawDevicePairingRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: m.namespace,
		},
		Spec: clawv1alpha1.ClawDevicePairingRequestSpec{
			RequestID: requestID,
			Selector:  m.selector,
		},
	}

	// Create the CR using typed client
	err := m.client.Create(ctx, cr)
	if err != nil {
		// Handle AlreadyExists error gracefully
		if errors.IsAlreadyExists(err) {
			slog.Info("ClawDevicePairingRequest already exists, skipping creation", "name", crName, "namespace", m.namespace)
			return nil
		}

		// Log detailed error and return
		slog.Error("failed to create ClawDevicePairingRequest", "error", err, "name", crName, "namespace", m.namespace, "request_id", requestID)
		return fmt.Errorf("failed to create pairing request CR: %w", err)
	}

	slog.Info("ClawDevicePairingRequest created successfully", "name", crName, "namespace", m.namespace, "request_id", requestID)
	return nil
}

// GetPairingRequestStatus fetches a ClawDevicePairingRequest and returns whether
// its "Ready" condition is True.
func (m *Manager) GetPairingRequestStatus(ctx context.Context, requestID string) (ready bool, err error) {
	if !m.IsEnabled() {
		return false, fmt.Errorf("Kubernetes client not enabled, cannot get pairing request status")
	}

	crName := naming.SanitizeDNS1123Label(requestID)
	if crName == "" {
		return false, fmt.Errorf("requestID %q cannot be sanitized to valid DNS-1123 label", requestID)
	}

	cr := &clawv1alpha1.ClawDevicePairingRequest{}
	if err := m.client.Get(ctx, client.ObjectKey{
		Namespace: m.namespace,
		Name:      crName,
	}, cr); err != nil {
		return false, fmt.Errorf("failed to get pairing request CR: %w", err)
	}

	return meta.IsStatusConditionTrue(cr.Status.Conditions, "Ready"), nil
}
