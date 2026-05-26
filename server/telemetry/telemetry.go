package telemetry

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/liftbridge-io/liftbridge/server/logger"
)

const (
	// DefaultEndpoint is the telemetry collection endpoint (hardcoded, not configurable)
	DefaultEndpoint = "https://telemetry.basekick.net/api/v1/liftbridge/telemetry"

	// DefaultInterval is the telemetry reporting interval
	DefaultInterval = 24 * time.Hour

	// instanceIDFile stores the persistent instance ID
	instanceIDFile = ".instance_id"
)

// Config holds telemetry configuration
type Config struct {
	Enabled  bool          // Enable telemetry (default: true)
	Interval time.Duration // Reporting interval (default: 24h)
	DataDir  string        // Directory for instance ID file
}

// DefaultConfig returns the default telemetry configuration
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// Collector collects and sends telemetry data
type Collector struct {
	config     *Config
	instanceID string
	version    string
	startTime  time.Time

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	client *http.Client
	logger logger.Logger
}

// TelemetryPayload represents the data sent to the telemetry endpoint
type TelemetryPayload struct {
	InstanceID        string  `json:"instance_id"`
	Timestamp         string  `json:"timestamp"`
	LiftbridgeVersion string  `json:"liftbridge_version"`
	OS                OSInfo  `json:"os"`
	CPU               CPUInfo `json:"cpu"`
	Memory            MemInfo `json:"memory"`
}

// OSInfo contains operating system information
type OSInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Platform     string `json:"platform"`
}

// CPUInfo contains CPU information
type CPUInfo struct {
	PhysicalCores *int `json:"physical_cores"`
	LogicalCores  *int `json:"logical_cores"`
	FrequencyMHz  *int `json:"frequency_mhz"`
}

// MemInfo contains memory information
type MemInfo struct {
	TotalGB *float64 `json:"total_gb"`
}

// New creates a new telemetry collector
func New(cfg *Config, version string, log logger.Logger) (*Collector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load or generate instance ID

// Start begins periodic telemetry collection
func (c *Collector) Start() { _ = "STUB: not implemented"; return }

// Stop stops the telemetry collector
func (c *Collector) Stop() { _ = "STUB: not implemented"; return }

// GetInstanceID returns the instance ID
func (c *Collector) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

func (c *Collector) run() {
	_ = "STUB: not implemented"

	// Send initial telemetry beacon
	return
}

// Then send periodically

func (c *Collector) sendTelemetry() { _ = "STUB: not implemented"; return }

func (c *Collector) collectPayload() *TelemetryPayload {
	_ = "STUB: not implemented"
	// Get CPU info
	return nil
}

// Get memory info (total system memory via runtime)

// Build OS platform string

// Not easily available in Go without cgo

func loadOrCreateInstanceID(dataDir string) (string, error) {
	_ = "STUB: not implemented"
	// Ensure data directory exists
	return "", nil
}

// Try to load existing ID

// Generate new ID as UUID format

// Save ID

func generateUUID() (string, error) {
	_ = "STUB: not implemented"
	// Generate UUID v4 format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	return "", nil
}

// Set version (4) and variant bits
// Version 4
// Variant is 10
