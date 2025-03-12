// Copyright 2025 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package medialogutils

import (
	"fmt"
	"strings"

	"github.com/livekit/protocol/logger"
)

// CmdLogger logs cmd outputs
type CmdLogger struct {
	name string
}

func NewCmdLogger(name string) *CmdLogger {
	return &CmdLogger{
		name: name,
	}
}

func (l *CmdLogger) Write(p []byte) (int, error) {
	logger.Infow(fmt.Sprintf("%s: %s", l.name, string(p)))
	return len(p), nil
}

// HandlerLogger catches stray outputs from egress handlers
type HandlerLogger struct {
	logger logger.Logger
}

func NewHandlerLogger(keyAndValues ...any) *HandlerLogger {
	return &HandlerLogger{
		logger: logger.GetLogger().WithValues(keyAndValues...),
	}
}

var downgrade = map[string]bool{
	"turnc ": true,
	"ice ER": true,
	"SDK 20": true,
	"(egres": true,
}

func (l *HandlerLogger) Write(p []byte) (n int, err error) {
	s := strings.Split(strings.TrimSuffix(string(p), "\n"), "\n")
	for _, line := range s {
		switch {
		case strings.HasSuffix(line, "}"):
			// (probably) normal log
			fmt.Println(line)
		case strings.HasPrefix(line, "0:00:"):
			// ignore cuda and template not mapped gstreamer warnings
			continue
		case len(line) > 6 && downgrade[line[:6]]:
			// downgrade turn, ice, and sdk errors
			l.logger.Infow(line)
		default:
			// panics and unexpected errors
			l.logger.Errorw(line, nil)
		}
	}

	return len(p), nil
}
