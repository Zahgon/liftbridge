package common

// PreparedMessage holds a pre-generated message ready for publishing.
type PreparedMessage struct {
	Key   []byte
	Value []byte
}

// PreGenerateMessages creates all message batches upfront.
// This allows benchmarks to measure pure ingestion performance
// without including data generation time.
func PreGenerateMessages(numMessages, messageSize, batchSize int) [][]PreparedMessage {
	_ = "STUB: not implemented"
	return nil
}

// Allocate batches

// Pre-generate the payload template once

// Handle remainder

// PreGenerateMessagesFlat creates all messages as a flat slice.
// Useful when batch structure isn't needed.
func PreGenerateMessagesFlat(numMessages, messageSize int) []PreparedMessage {
	_ = "STUB: not implemented"
	return nil
}

// Pre-generate the payload template once

// generatePayload creates a new payload by copying and slightly modifying the template.
// This is more efficient than calling crypto/rand for every message.
func generatePayload(size int, template []byte) []byte { _ = "STUB: not implemented"; return nil }

// Add some variation by modifying a few bytes

// TotalMessageCount returns the total number of messages across all batches.
func TotalMessageCount(batches [][]PreparedMessage) int { _ = "STUB: not implemented"; return 0 }

// TotalByteSize returns the total bytes across all messages.
func TotalByteSize(batches [][]PreparedMessage) int64 { _ = "STUB: not implemented"; return 0 }
