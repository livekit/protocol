// Copyright 2023 LiveKit, Inc.
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

package psrpc

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/livekit/psrpc/version"
)

func CheckCompilerVersion(path string) error {
	b, err := exec.Command(path, "--version").Output()
	if err != nil {
		return err
	}

	b = bytes.Trim(b, "\r\n")
	if string(b) != version.Version {
		return fmt.Errorf("found psrpc compiler version %s need %s. please run:\ngo install github.com/livekit/psrpc/protoc-gen-psrpc", string(b), version.Version)
	}
	return nil
}
