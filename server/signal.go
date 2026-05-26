package server

// handleSignals sets up a handler for SIGINT to do a graceful shutdown.
func (s *Server) handleSignals() { _ = "STUB: not implemented"; return }

// Use a naked goroutine instead of startGoroutine because this stops the
// server which would cause a deadlock.

// Reload authz permissions from storage

// LoadPolicy, if fails, will likely throw panic
// Casbin raise a panic if it fails to load the policy, i.e: policy file is corrupted,
// Refer to issue: https://github.com/casbin/casbin/issues/640
