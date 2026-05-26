package client

import (
	"context"
	"testing"

	clawv1alpha1 "github.com/codeready-toolchain/claw-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCreatePairingRequest_Success(t *testing.T) {
	// Create scheme and fake client
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	// Create manager with fake client
	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
		selector: metav1.LabelSelector{
			MatchLabels: map[string]string{
				AppLabelKey:      "claw",
				InstanceLabelKey: testInstance,
			},
		},
	}

	// Test CR creation
	ctx := context.Background()
	requestID := "test-request-123"

	err := manager.CreatePairingRequest(ctx, requestID)
	if err != nil {
		t.Fatalf("CreatePairingRequest failed: %v", err)
	}

	// Verify CR was created
	cr := &clawv1alpha1.ClawDevicePairingRequest{}
	err = fakeClient.Get(ctx, client.ObjectKey{
		Namespace: testNamespace,
		Name:      requestID,
	}, cr)
	if err != nil {
		t.Fatalf("Failed to get created CR: %v", err)
	}

	// Verify CR fields
	if cr.Spec.RequestID != requestID {
		t.Errorf("Expected requestID %q, got %q", requestID, cr.Spec.RequestID)
	}
	if cr.Spec.Selector.MatchLabels[InstanceLabelKey] != testInstance {
		t.Errorf("Expected instance label %q, got %q", testInstance, cr.Spec.Selector.MatchLabels[InstanceLabelKey])
	}
}

func TestCreatePairingRequest_DisabledClient(t *testing.T) {
	// Create disabled manager (nil client)
	manager := &Manager{
		client: nil,
	}

	ctx := context.Background()
	err := manager.CreatePairingRequest(ctx, "test-request")

	if err == nil {
		t.Error("Expected error for disabled client, got nil")
	}

	expectedError := "Kubernetes client not enabled"
	if err.Error() != expectedError+", cannot create pairing request CR" {
		t.Errorf("Expected error containing %q, got %q", expectedError, err.Error())
	}
}

func TestCreatePairingRequest_InvalidRequestID(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
	}

	ctx := context.Background()
	// Request ID with only special characters that sanitize to empty string
	requestID := "@#$%"

	err := manager.CreatePairingRequest(ctx, requestID)
	if err == nil {
		t.Error("Expected error for invalid requestID, got nil")
	}

	if !contains(err.Error(), "cannot be sanitized") {
		t.Errorf("Expected sanitization error, got: %v", err)
	}
}

func TestCreatePairingRequest_SanitizesName(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
	}

	ctx := context.Background()
	// Request ID that needs sanitization
	requestID := "Test-Request-UPPERCASE"
	expectedName := "test-request-uppercase"

	err := manager.CreatePairingRequest(ctx, requestID)
	if err != nil {
		t.Fatalf("CreatePairingRequest failed: %v", err)
	}

	// Verify CR was created with sanitized name
	cr := &clawv1alpha1.ClawDevicePairingRequest{}
	err = fakeClient.Get(ctx, client.ObjectKey{
		Namespace: testNamespace,
		Name:      expectedName,
	}, cr)
	if err != nil {
		t.Fatalf("Failed to get created CR: %v", err)
	}

	if cr.Name != expectedName {
		t.Errorf("Expected name %q, got %q", expectedName, cr.Name)
	}
}

func TestGetPairingRequestStatus_Ready(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	cr := &clawv1alpha1.ClawDevicePairingRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-request-123",
			Namespace: testNamespace,
		},
		Status: clawv1alpha1.ClawDevicePairingRequestStatus{
			Conditions: []metav1.Condition{
				{
					Type:   "Ready",
					Status: metav1.ConditionTrue,
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(cr).
		WithStatusSubresource(cr).
		Build()

	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
	}

	ready, err := manager.GetPairingRequestStatus(context.Background(), "test-request-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ready {
		t.Error("expected ready=true, got false")
	}
}

func TestGetPairingRequestStatus_NotReady(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	cr := &clawv1alpha1.ClawDevicePairingRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-request-456",
			Namespace: testNamespace,
		},
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(cr).
		Build()

	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
	}

	ready, err := manager.GetPairingRequestStatus(context.Background(), "test-request-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ready {
		t.Error("expected ready=false, got true")
	}
}

func TestGetPairingRequestStatus_NotFound(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = clawv1alpha1.AddToScheme(scheme)

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	manager := &Manager{
		client:    fakeClient,
		namespace: testNamespace,
	}

	_, err := manager.GetPairingRequestStatus(context.Background(), "nonexistent")
	if err == nil {
		t.Error("expected error for non-existent CR, got nil")
	}
}
