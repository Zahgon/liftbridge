package server

import (
	"time"

	client "github.com/liftbridge-io/liftbridge-api/v2/go"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"

	proto "github.com/liftbridge-io/liftbridge/server/protocol"
)

const (
	// DefaultNamespace is the default cluster namespace to use if one is not
	// specified.
	DefaultNamespace = "liftbridge-default"

	// DefaultPort is the port to bind to if one is not specified.
	DefaultPort = 9292
)

// Config setting defaults.
const (
	defaultListenAddress                  = "0.0.0.0"
	defaultConnectionAddress              = "localhost"
	defaultReplicaMaxLagTime              = 15 * time.Second
	defaultReplicaMaxLeaderTimeout        = 15 * time.Second
	defaultReplicaMaxIdleWait             = 10 * time.Second
	defaultReplicationMaxBytes            = 1024 * 1024 // 1MB
	defaultRaftSnapshots                  = 2
	defaultRaftCacheSize                  = 512
	defaultMetadataCacheMaxAge            = 2 * time.Minute
	defaultBatchMaxMessages               = 1024
	defaultBatchMaxTime                   = 0 // Disabled by default; set to enable waiting for batches
	defaultReplicaFetchTimeout            = 3 * time.Second
	defaultMinInsyncReplicas              = 1
	defaultRetentionMaxAge                = 7 * 24 * time.Hour
	defaultCleanerInterval                = 5 * time.Minute
	defaultMaxSegmentBytes                = 1024 * 1024 * 256 // 256MB
	defaultMaxSegmentAge                  = defaultRetentionMaxAge
	defaultActivityStreamPublishTimeout   = 5 * time.Second
	defaultActivityStreamPublishAckPolicy = client.AckPolicy_ALL
	defaultCursorsStreamReplicationFactor = maxReplicationFactor
	defaultCursorsStreamAutoPauseTime     = time.Minute
	defaultConcurrencyControl             = false
	defaultEncryption                     = false
	defaultGroupsConsumerTimeout          = 15 * time.Second
	defaultGroupsCoordinatorTimeout       = 15 * time.Second
	defaultTelemetryEnabled               = true
	defaultTelemetryIntervalSeconds       = 86400 // 24 hours
)

// Config setting key names.
const (
	configListen              = "listen"
	configHost                = "host"
	configPort                = "port"
	configDataDir             = "data.dir"
	configMetadataCacheMaxAge = "metadata.cache.max.age"

	configLoggingLevel    = "logging.level"
	configLoggingRecovery = "logging.recovery"
	configLoggingRaft     = "logging.raft"
	configLoggingNATS     = "logging.nats"

	configBatchMaxMessages = "batch.max.messages"
	configBatchMaxTime     = "batch.max.time"

	configTLSKey                = "tls.key"
	configTLSCert               = "tls.cert"
	configTLSClientAuthEnabled  = "tls.client.auth.enabled"
	configTLSClientAuthCA       = "tls.client.auth.ca"
	configTLSClientAuthzEnabled = "tls.client.authz.enabled"
	configTLSClientAuthzModel   = "tls.client.authz.model"
	configTLSClientAuthzPolicy  = "tls.client.authz.policy"

	configNATSServers        = "nats.servers"
	configNATSUser           = "nats.user"
	configNATSPassword       = "nats.password"
	configNATSCert           = "nats.tls.cert"
	configNATSKey            = "nats.tls.key"
	configNATSCA             = "nats.tls.ca"
	configNATSEmbedded       = "nats.embedded"
	configNATSEmbeddedConfig = "nats.embedded.config"

	configStreamsRetentionMaxBytes             = "streams.retention.max.bytes"
	configStreamsRetentionMaxMessages          = "streams.retention.max.messages"
	configStreamsRetentionMaxAge               = "streams.retention.max.age"
	configStreamsCleanerInterval               = "streams.cleaner.interval"
	configStreamsSegmentMaxBytes               = "streams.segment.max.bytes"
	configStreamsSegmentMaxAge                 = "streams.segment.max.age"
	configStreamsCompactEnabled                = "streams.compact.enabled"
	configStreamsCompactMaxGoroutines          = "streams.compact.max.goroutines"
	configStreamsAutoPauseTime                 = "streams.auto.pause.time"
	configStreamsAutoPauseDisableIfSubscribers = "streams.auto.pause.disable.if.subscribers"
	configStreamsConcurrencyControl            = "streams.concurrency.control"
	configStreamsEncryption                    = "streams.encryption"

	configClusteringServerID                = "clustering.server.id"
	configClusteringNamespace               = "clustering.namespace"
	configClusteringRaftSnapshotRetain      = "clustering.raft.snapshot.retain"
	configClusteringRaftSnapshotThreshold   = "clustering.raft.snapshot.threshold"
	configClusteringRaftCacheSize           = "clustering.raft.cache.size"
	configClusteringRaftBootstrapSeed       = "clustering.raft.bootstrap.seed"
	configClusteringRaftBootstrapPeers      = "clustering.raft.bootstrap.peers"
	configClusteringRaftMaxQuorumSize       = "clustering.raft.max.quorum.size"
	configClusteringReplicaMaxLagTime       = "clustering.replica.max.lag.time"
	configClusteringReplicaMaxLeaderTimeout = "clustering.replica.max.leader.timeout"
	configClusteringReplicaMaxIdleWait      = "clustering.replica.max.idle.wait"
	configClusteringReplicaFetchTimeout     = "clustering.replica.fetch.timeout"
	configClusteringMinInsyncReplicas       = "clustering.min.insync.replicas"
	configClusteringReplicationMaxBytes     = "clustering.replication.max.bytes"

	configActivityStreamEnabled          = "activity.stream.enabled"
	configActivityStreamPublishTimeout   = "activity.stream.publish.timeout"
	configActivityStreamPublishAckPolicy = "activity.stream.publish.ack.policy"

	configCursorsStreamPartitions        = "cursors.stream.partitions"
	configCursorsStreamReplicationFactor = "cursors.stream.replication.factor"
	configCursorsStreamAutoPauseTime     = "cursors.stream.auto.pause.time"

	configGroupsConsumerTimeout    = "groups.consumer.timeout"
	configGroupsCoordinatorTimeout = "groups.coordinator.timeout"

	configTelemetryEnabled         = "telemetry.enabled"
	configTelemetryIntervalSeconds = "telemetry.interval.seconds"
)

var configKeys = map[string]struct{}{
	configListen:                               {},
	configHost:                                 {},
	configPort:                                 {},
	configDataDir:                              {},
	configMetadataCacheMaxAge:                  {},
	configLoggingLevel:                         {},
	configLoggingRecovery:                      {},
	configLoggingRaft:                          {},
	configLoggingNATS:                          {},
	configBatchMaxMessages:                     {},
	configBatchMaxTime:                         {},
	configTLSKey:                               {},
	configTLSCert:                              {},
	configTLSClientAuthEnabled:                 {},
	configTLSClientAuthCA:                      {},
	configTLSClientAuthzEnabled:                {},
	configTLSClientAuthzModel:                  {},
	configTLSClientAuthzPolicy:                 {},
	configNATSServers:                          {},
	configNATSUser:                             {},
	configNATSPassword:                         {},
	configNATSCert:                             {},
	configNATSKey:                              {},
	configNATSCA:                               {},
	configNATSEmbedded:                         {},
	configNATSEmbeddedConfig:                   {},
	configStreamsRetentionMaxBytes:             {},
	configStreamsRetentionMaxMessages:          {},
	configStreamsRetentionMaxAge:               {},
	configStreamsCleanerInterval:               {},
	configStreamsSegmentMaxBytes:               {},
	configStreamsSegmentMaxAge:                 {},
	configStreamsCompactEnabled:                {},
	configStreamsConcurrencyControl:            {},
	configStreamsEncryption:                    {},
	configStreamsCompactMaxGoroutines:          {},
	configStreamsAutoPauseTime:                 {},
	configStreamsAutoPauseDisableIfSubscribers: {},
	configClusteringServerID:                   {},
	configClusteringNamespace:                  {},
	configClusteringRaftSnapshotRetain:         {},
	configClusteringRaftSnapshotThreshold:      {},
	configClusteringRaftCacheSize:              {},
	configClusteringRaftBootstrapSeed:          {},
	configClusteringRaftBootstrapPeers:         {},
	configClusteringRaftMaxQuorumSize:          {},
	configClusteringReplicaMaxLagTime:          {},
	configClusteringReplicaMaxLeaderTimeout:    {},
	configClusteringReplicaMaxIdleWait:         {},
	configClusteringReplicaFetchTimeout:        {},
	configClusteringMinInsyncReplicas:          {},
	configClusteringReplicationMaxBytes:        {},
	configActivityStreamEnabled:                {},
	configActivityStreamPublishTimeout:         {},
	configActivityStreamPublishAckPolicy:       {},
	configCursorsStreamPartitions:              {},
	configCursorsStreamReplicationFactor:       {},
	configCursorsStreamAutoPauseTime:           {},
	configGroupsConsumerTimeout:                {},
	configGroupsCoordinatorTimeout:             {},
	configTelemetryEnabled:                     {},
	configTelemetryIntervalSeconds:             {},
}

// StreamsConfig contains settings for controlling the message log for streams.
type StreamsConfig struct {
	RetentionMaxBytes             int64
	RetentionMaxMessages          int64
	RetentionMaxAge               time.Duration
	CleanerInterval               time.Duration
	SegmentMaxBytes               int64
	SegmentMaxAge                 time.Duration
	Compact                       bool
	CompactMaxGoroutines          int
	AutoPauseTime                 time.Duration
	AutoPauseDisableIfSubscribers bool
	MinISR                        int
	ConcurrencyControl            bool
	Encryption                    bool
}

// RetentionString returns a human-readable string representation of the
// retention policy.
func (l StreamsConfig) RetentionString() string { _ = "STUB: not implemented"; return "" }

// AutoPauseString returns a human-readable string representation of the auto
// pause setting.
func (l StreamsConfig) AutoPauseString() string { _ = "STUB: not implemented"; return "" }

// ApplyOverrides applies the values from the StreamConfig protobuf to the
// StreamsConfig struct. If the value is present in the request's config
// section, it will be set in StreamsConfig.
func (l *StreamsConfig) ApplyOverrides(c *proto.StreamConfig) { _ = "STUB: not implemented"; return }

// By default, duration configuration is considered as milliseconds.

// ClusteringConfig contains settings for controlling cluster behavior.
type ClusteringConfig struct {
	ServerID                string
	Namespace               string
	RaftSnapshots           int
	RaftSnapshotThreshold   uint64
	RaftCacheSize           int
	RaftBootstrapSeed       bool
	RaftBootstrapPeers      []string
	RaftMaxQuorumSize       uint
	ReplicaMaxLagTime       time.Duration
	ReplicaMaxLeaderTimeout time.Duration
	ReplicaFetchTimeout     time.Duration
	ReplicaMaxIdleWait      time.Duration
	MinISR                  int
	ReplicationMaxBytes     int64
}

// ActivityStreamConfig contains settings for controlling activity stream
// behavior.
type ActivityStreamConfig struct {
	Enabled          bool
	PublishTimeout   time.Duration
	PublishAckPolicy client.AckPolicy
}

// CursorsStreamConfig contains settings for controlling cursors stream
// behavior.
type CursorsStreamConfig struct {
	Partitions        int32
	ReplicationFactor int32
	AutoPauseTime     time.Duration
}

// GroupsConfig contains settings for controlling consumer group behavior.
type GroupsConfig struct {
	ConsumerTimeout    time.Duration
	CoordinatorTimeout time.Duration
}

// TelemetryConfig contains settings for controlling telemetry behavior.
type TelemetryConfig struct {
	Enabled         bool
	IntervalSeconds int
}

// Config contains all settings for a Liftbridge Server.
type Config struct {
	Listen               HostPort
	Host                 string
	Port                 int
	LogLevel             uint32
	LogRecovery          bool
	LogRaft              bool
	LogNATS              bool
	LogSilent            bool
	DataDir              string
	BatchMaxMessages     int
	BatchMaxTime         time.Duration
	MetadataCacheMaxAge  time.Duration
	TLSKey               string
	TLSCert              string
	TLSClientAuth        bool
	TLSClientAuthCA      string
	TLSClientAuthz       bool
	TLSClientAuthzModel  string
	TLSClientAuthzPolicy string
	NATS                 nats.Options
	EmbeddedNATS         bool
	EmbeddedNATSConfig   string
	Streams              StreamsConfig
	Clustering           ClusteringConfig
	ActivityStream       ActivityStreamConfig
	CursorsStream        CursorsStreamConfig
	Groups               GroupsConfig
	Telemetry            TelemetryConfig
}

// NewDefaultConfig creates a new Config with default settings.
func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// BatchMaxTime defaults to 0 (no wait)

// NATSServersString returns a human-readable string representation of the
// list of NATS servers.
func (c Config) NATSServersString() string { _ = "STUB: not implemented"; return "" }

// GetListenAddress returns the address and port to listen to.
func (c Config) GetListenAddress() HostPort { _ = "STUB: not implemented"; return *new(HostPort) }

// GetConnectionAddress returns the host if specified and listen otherwise.
func (c Config) GetConnectionAddress() HostPort { _ = "STUB: not implemented"; return *new(HostPort) }

// GetLogLevel converts the level string to its corresponding int value. It
// returns an error if the level is invalid.
func GetLogLevel(level string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// NewConfig creates a new Config with default settings and applies any
// settings from the given configuration file.
func NewConfig(configFile string) (*Config, error) {
	_ = "STUB: not implemented" // nolint: gocyclo
	return nil, nil
}

// Return default config if config file is not given.

// Expect a yaml config file.

// Allow overriding config with environment variables

// Parse the config file.

// Validate config settings.

// Reset SegmentMaxAge since this will get overwritten later.

// Process parsed config file here with v.

// If SegmentMaxAge is not set, default it to the retention time.

// parseNATSConfig parses the `nats` section of a config file and populates the
// given nats.Options.
func parseNATSConfig(config *Config, v *viper.Viper) error { _ = "STUB: not implemented"; return nil }

// NATS TLS config
// Both Cert and Key must be presented

// Load cert and key file

// Load CACert if available

// Load CA cert

// parseStreamConfig parses the `streams` section of a config file and
// populates the given Config.
func parseStreamsConfig(config *Config, v *viper.Viper) error {
	_ = "STUB: not implemented"
	return nil
}

// parseClusteringConfig parses the `clustering` section of a config file and
// populates the given Config.
func parseClusteringConfig(config *Config, v *viper.Viper) error {
	_ = "STUB: not implemented" // nolint: gocyclo
	return nil
}

// parseActivityStreamConfig parses the `activitystream` section of a config
// file and populates the given Config.
func parseActivityStreamConfig(config *Config, v *viper.Viper) error {
	_ = "STUB: not implemented" // nolint: gocyclo
	return nil
}

// parseCursorsStreamConfig parses the `cursors` section of a config file and
// populates the given Config.
func parseCursorsStreamConfig(config *Config, v *viper.Viper) error {
	_ = "STUB: not implemented" // nolint: gocyclo
	return nil
}

// parseGroupsConfig parses the `groups` section of a config file and populates
// the given Config.
func parseGroupsConfig(config *Config, v *viper.Viper) error {
	_ = "STUB: not implemented" // nolint: gocyclo
	return nil
}

// parseTelemetryConfig parses the `telemetry` section of a config file and
// populates the given Config.
func parseTelemetryConfig(config *Config, v *viper.Viper) { _ = "STUB: not implemented"; return }

// HostPort is simple struct to hold parsed listen/addr strings.
type HostPort struct {
	Host string
	Port int
}

// parseListen will parse the `listen` option containing the host and port.
func parseListen(v *viper.Viper) (*HostPort, error) { _ = "STUB: not implemented"; return nil, nil }

// Only a port

// parseAckPolicy will parse the activity stream's `ack.policy` option
// containing the ack policy to use when publishing activity events.
func parseAckPolicy(v *viper.Viper) (client.AckPolicy, error) {
	_ = "STUB: not implemented"
	return *new(client.AckPolicy), nil
}
