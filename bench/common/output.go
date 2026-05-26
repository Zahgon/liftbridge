package common

// BenchmarkResult holds the formatted benchmark results.
type BenchmarkResult struct {
	Duration          string  `json:"duration"`
	TotalMessages     int64   `json:"total_messages"`
	TotalBytes        int64   `json:"total_bytes"`
	MessagesPerSecond float64 `json:"messages_per_second"`
	BytesPerSecond    float64 `json:"bytes_per_second"`
	MBPerSecond       float64 `json:"mb_per_second"`
	LatencyMin        string  `json:"latency_min,omitempty"`
	LatencyMean       string  `json:"latency_mean,omitempty"`
	LatencyP50        string  `json:"latency_p50,omitempty"`
	LatencyP95        string  `json:"latency_p95,omitempty"`
	LatencyP99        string  `json:"latency_p99,omitempty"`
	LatencyP999       string  `json:"latency_p999,omitempty"`
	LatencyMax        string  `json:"latency_max,omitempty"`
	Errors            int64   `json:"errors"`
}

// PrintProducerResults outputs the producer benchmark results.
func PrintProducerResults(stats *Stats, format string) { _ = "STUB: not implemented"; return }

// Include latency stats if we have samples

// PrintConsumerResults outputs the consumer benchmark results.
func PrintConsumerResults(stats *Stats, format string) { _ = "STUB: not implemented"; return }

func printJSON(result BenchmarkResult) { _ = "STUB: not implemented"; return }

func printTextProducer(r BenchmarkResult) { _ = "STUB: not implemented"; return }

func printTextConsumer(r BenchmarkResult) { _ = "STUB: not implemented"; return }
