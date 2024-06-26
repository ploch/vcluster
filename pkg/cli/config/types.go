package config

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CLI struct {
	Manager           Manager  `json:"manager,omitempty"`
	PreviousContext   string   `json:"previousContext,omitempty"`
	path              string   `json:"-"`
	Platform          Platform `json:"platform,omitempty"`
	TelemetryDisabled bool     `json:"telemetryDisabled,omitempty"`
}

type Manager struct {
	// Type is the current manager type that is used, either helm or platform
	Type ManagerType `json:"type,omitempty"`
}

type ManagerType string

type Platform struct {
	metav1.TypeMeta `json:",inline"`

	// VirtualClusterAccessPointCertificates is a map of cached certificates for "access point" mode virtual clusters
	VirtualClusterAccessPointCertificates map[string]VirtualClusterCertificatesEntry `json:"virtualClusterAccessPointCertificates,omitempty"`
	// Host is the https endpoint of how to access loft
	Host string `json:"host,omitempty"`
	// LastInstallContext is the last install context
	LastInstallContext string `json:"lastInstallContext,omitempty"`
	// AccessKey is the access key for the given loft host
	AccessKey string `json:"accesskey,omitempty"`
	// VirtualClusterAccessKey is the access key for the given loft host to create virtual clusters
	VirtualClusterAccessKey string `json:"virtualClusterAccessKey,omitempty"`
	// Insecure specifies if the loft instance is insecure
	Insecure bool `json:"insecure,omitempty"`
}

type VirtualClusterCertificatesEntry struct {
	LastRequested   metav1.Time `json:"lastRequested,omitempty"`
	ExpirationTime  time.Time   `json:"expirationTime,omitempty"`
	CertificateData string      `json:"certificateData,omitempty"`
	KeyData         string      `json:"keyData,omitempty"`
}
