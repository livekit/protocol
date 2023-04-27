package pion

import (
	"github.com/pion/logging"
	"go.uber.org/zap/zapcore"

	"github.com/livekit/protocol/logger"
)

var (
	pionLevel           zapcore.Level
	pionIgnoredPrefixes = map[string][]string{
		"ice": {
			"pingAllCandidates called with no candidate pairs",
			"failed to send packet: io: read/write on closed pipe",
			"Ignoring remote candidate with tcpType active",
			"discard message from",
			"Failed to discover mDNS candidate",
			"Failed to read from candidate tcp",
			"remote mDNS candidate added, but mDNS is disabled",
		},
		"pc": {
			"Failed to accept RTCP stream is already closed",
			"Failed to accept RTP stream is already closed",
			"Incoming unhandled RTCP ssrc",
		},
		"tcp_mux": {
			"Error reading first packet from",
			"error closing connection",
		},
		"turn": {
			"error when handling datagram",
		},
	}
)

// implements webrtc.LoggerFactory
type LoggerFactory struct {
	logger logger.Logger
}

func NewLoggerFactory(logger logger.Logger) *LoggerFactory {
	return &LoggerFactory{
		logger: logger,
	}
}

func (f *LoggerFactory) NewLogger(scope string) logging.LeveledLogger {
	return &logAdapter{
		logger:          f.logger.WithName(scope),
		level:           pionLevel,
		ignoredPrefixes: pionIgnoredPrefixes[scope],
	}
}

func SetLogLevel(level string) {
	pionLevel = logger.ParseZapLevel(level)
}
