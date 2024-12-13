package api

const (
	APIVersion string = "1.0"
	APIStatus  string = "devel"
)

// ServerPut represents the modifiable fields of a server configuration
//
// swagger:model
type ServerPut struct {
	// Server configuration map (refer to doc/server.md)
	// Example: {"core.https_address": ":6443"}
	Config map[string]string `json:"config" yaml:"config"`
}

// ServerUntrusted represents a server configuration for an untrusted client
//
// swagger:model
type ServerUntrusted struct {
	ServerPut `yaml:",inline"`

	// Support status of the current API (one of "devel", "stable" or "deprecated")
	// Read only: true
	// Example: stable
	APIStatus string `json:"api_status" yaml:"api_status"`

	// API version number
	// Read only: true
	// Example: 1.0
	APIVersion string `json:"api_version" yaml:"api_version"`

	// Whether the client is trusted (one of "trusted" or "untrusted")
	// Read only: true
	// Example: untrusted
	Auth string `json:"auth" yaml:"auth"`

	// List of supported authentication methods
	// Read only: true
	// Example: ["tls"]
	//
	// API extension: macaroon_authentication
	AuthMethods []string `json:"auth_methods" yaml:"auth_methods"`
}
