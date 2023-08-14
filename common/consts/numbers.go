// Package consts implements some constants
package consts

const (
	// ClientRetries : client retries
	ClientRetries = 3
	// PingDelaySeconds : When access to ping interface, a while back to 200
	PingDelaySeconds = 30
	// ClientRequestTimeout : client request timeout
	ClientRequestTimeout = 30
	// ClientDialTimeout : client dial timeout
	ClientDialTimeout = 30
	// DefaultSelectorTTL : default selector ttl
	DefaultSelectorTTL = 10
	// MaxMsgSize : max msg size
	MaxMsgSize = 10 * 1024 * 1024
)
