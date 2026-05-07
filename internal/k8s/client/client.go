package client

import (
	"fmt"
	"log/slog"
	"os"

	clawv1alpha1 "github.com/codeready-toolchain/claw-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// InstanceLabelKey is the label key for the Claw instance
	InstanceLabelKey = "claw.sandbox.redhat.com/instance"
	// AppLabelKey is the label key for the app
	AppLabelKey = "app"
)

// Manager handles Kubernetes client operations and configuration
type Manager struct {
	client    client.Client
	namespace string
	selector  metav1.LabelSelector
}

// NewManager creates a new Kubernetes client manager
// Requires NAMESPACE and CLAW_INSTANCE environment variables to be set
func NewManager() (*Manager, error) {
	// Read namespace from environment (required)
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		return nil, fmt.Errorf("NAMESPACE environment variable is required")
	}

	// Read instance label from environment (required)
	instanceValue := os.Getenv("CLAW_INSTANCE")
	if instanceValue == "" {
		return nil, fmt.Errorf("CLAW_INSTANCE environment variable is required")
	}

	// Construct label selector
	selector := metav1.LabelSelector{
		MatchLabels: map[string]string{
			AppLabelKey:      "claw",
			InstanceLabelKey: instanceValue,
		},
	}

	// Load in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load in-cluster configuration: %w (ensure service account credentials are mounted at /var/run/secrets/kubernetes.io/serviceaccount/)", err)
	}

	// Create scheme with Claw CRD types
	scheme := runtime.NewScheme()
	if err := clawv1alpha1.AddToScheme(scheme); err != nil {
		return nil, fmt.Errorf("failed to add Claw types to scheme: %w", err)
	}

	// Create controller-runtime client
	k8sClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	slog.Info("Kubernetes client initialized", "namespace", namespace, "instance_label", instanceValue)

	return &Manager{
		client:    k8sClient,
		namespace: namespace,
		selector:  selector,
	}, nil
}

// IsEnabled returns true if Kubernetes client is available
func (m *Manager) IsEnabled() bool {
	return m.client != nil
}

// GetNamespace returns the configured namespace
func (m *Manager) GetNamespace() string {
	return m.namespace
}

// GetSelector returns the configured label selector
func (m *Manager) GetSelector() metav1.LabelSelector {
	return m.selector
}
