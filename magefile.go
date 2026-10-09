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

//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/livekit/mageutil"
	"github.com/livekit/protocol/psrpc"
)

var Default = Proto

// install the protoc plugins system-wide, for developers who want them on their PATH
//
// This is purely a developer convenience and is not required to generate anything: no
// build, release or CI step invokes it, and none should. `mage proto` resolves each
// plugin through the tool directives in go.mod (see goToolPath) and passes protoc an
// explicit path, so it works on a clean checkout and ignores whatever these installs put
// in GOPATH/bin.
//
// That separation is why @latest is safe here. These binaries are for ad-hoc use outside
// generation — inspecting a descriptor, compiling a scratch .proto — and are not inputs
// to any committed file. Two things keep them out of the committed output: the tool
// directives above, and the Generate workflow, which regenerates and commits the
// generated tree on every push to a non-main branch, so what lands on main always comes
// from the pinned versions regardless of what a contributor had installed locally.
func Bootstrap() error {
	return mageutil.Run(context.Background(),
		"go install github.com/twitchtv/twirp/protoc-gen-twirp@latest",
		"go install google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"go install github.com/livekit/psrpc/protoc-gen-psrpc@latest",
	)
}

// regenerate protobuf
func Proto() error {
	twirpProtoFiles := []string{
		"cloud_replay.proto",
		"livekit_agent_dispatch.proto",
		"livekit_agentdb.proto",
		"livekit_egress.proto",
		"livekit_ingress.proto",
		"livekit_room.proto",
		"livekit_sip.proto",
		"livekit_cloud_agent.proto",
		"livekit_phone_number.proto",
		"livekit_connector.proto",
		"livekit_connector_whatsapp.proto",
		"livekit_connector_twilio.proto",
		"livekit_agent_simulation.proto",
		"livekit_agent_worker.proto",
	}

	agentProtoFiles := []string{
		"agent/livekit_agent_session.proto",
		"agent/livekit_agent_dev.proto",
		"agent/livekit_agent_inference.proto",
	}

	protoFiles := []string{
		"livekit_agent.proto",
		"livekit_agent_proxy.proto",
		"livekit_analytics.proto",
		"livekit_internal.proto",
		"livekit_models.proto",
		"livekit_rtc.proto",
		"livekit_webhook.proto",
		"livekit_metrics.proto",
		"livekit_token_source.proto",
		"logger/options.proto",
	}
	grpcProtoFiles := []string{
		"infra/link.proto",
		"rpc/analytics.proto",
	}
	psrpcProtoFiles := []string{
		"rpc/agent.proto",
		"rpc/agent_dispatch.proto",
		"rpc/egress.proto",
		"rpc/ingress.proto",
		"rpc/io.proto",
		"rpc/keepalive.proto",
		"rpc/participant.proto",
		"rpc/room.proto",
		"rpc/roommanager.proto",
		"rpc/signal.proto",
		"rpc/whip_signal.proto",
		"rpc/sip.proto",
	}

	// mapped proto directory:
	//    ./protobufs/roomrpc/<name>rpc
	// and generated Go package:
	//    ./livekit/roomrpc/<name>rpc
	roomrpcTypeNames := []string{
		"sip",
	}

	fmt.Println("generating protobuf")
	const target = "./livekit"
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}
	protocGoPath, err := goToolPath("protoc-gen-go")
	if err != nil {
		return err
	}
	twirpPath, err := goToolPath("protoc-gen-twirp")
	if err != nil {
		return err
	}
	protocGrpcGoPath, err := goToolPath("protoc-gen-go-grpc")
	if err != nil {
		return err
	}

	fmt.Println("generating twirp protobuf")
	args := append([]string{
		"--go_out", target,
		"--twirp_out", target,
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=protoc-gen-go=" + protocGoPath,
		"--plugin=protoc-gen-twirp=" + twirpPath,
		"-I=./protobufs",
	}, twirpProtoFiles...)
	cmd := exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("generating protobuf")
	args = append([]string{
		"--go_out", target,
		"--go_opt=paths=source_relative",
		"--plugin=protoc-gen-go=" + protocGoPath,
		"-I=./protobufs",
	}, protoFiles...)
	cmd = exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("generating protobuf (livekit/agent)")
	{
		args := []string{
			"--go_out", target,
			"--go_opt=paths=source_relative",
			"--plugin=protoc-gen-go=" + protocGoPath,
			"-I=./protobufs",
		}
		args = append(args, agentProtoFiles...)
		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	fmt.Println("generating protobuf (livekit/roomrpc)")
	for _, protoName := range roomrpcTypeNames {
		pkgName := "roomrpc/" + protoName + "rpc"
		protoFiles, err := os.ReadDir(filepath.Join("./protobufs", pkgName))
		if err != nil {
			return err
		}
		args := []string{
			"--go_out", target,
			"--go_opt=paths=source_relative",
			"--plugin=protoc-gen-go=" + protocGoPath,
			"-I=./protobufs",
		}
		for _, protoFile := range protoFiles {
			args = append(args, filepath.Join(pkgName, protoFile.Name()))
		}
		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	fmt.Println("generating grpc protobuf")
	args = append([]string{
		"--go_out", ".",
		"--go-grpc_out", ".",
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--plugin=protoc-gen-go=" + protocGoPath,
		"--plugin=protoc-gen-go-grpc=" + protocGrpcGoPath,
		"-I=./protobufs",
	}, grpcProtoFiles...)
	cmd = exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("generating psrpc protobuf")

	psrpcDir, err := mageutil.GetPkgDir("github.com/livekit/psrpc")
	if err != nil {
		return err
	}
	psrpcPath, err := goToolPath("protoc-gen-psrpc")
	if err != nil {
		return err
	}
	if err := psrpc.CheckCompilerVersion(psrpcPath); err != nil {
		return err
	}

	args = append([]string{
		"--go_out", ".",
		"--psrpc_out", ".",
		"--go_opt=paths=source_relative",
		"--psrpc_opt=paths=source_relative",
		"--plugin=protoc-gen-go=" + protocGoPath,
		"--plugin=protoc-gen-psrpc=" + psrpcPath,
		"-I" + psrpcDir + "/protoc-gen-psrpc/options",
		"-I=./protobufs",
	}, psrpcProtoFiles...)
	cmd = exec.Command(protoc, args...)
	mageutil.ConnectStd(cmd)
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}

// run tests
func Test() error {
	cmd := exec.Command("go", "test", "-race", "./...")
	connectStd(cmd)
	return cmd.Run()
}

// helpers

// goToolPath builds a protoc plugin from the tool directives in go.mod and returns its
// path, so generation uses the pinned version rather than whatever happens to be on PATH.
func goToolPath(name string) (string, error) {
	out, err := exec.Command("go", "tool", "-n", name).Output()
	if err != nil {
		return "", fmt.Errorf("resolving tool %s: %w", name, err)
	}
	path := strings.TrimSpace(string(out))
	if path == "" {
		return "", fmt.Errorf("resolving tool %s: no path returned", name)
	}
	return path, nil
}

// getToolPath locates a binary that is not a Go tool, i.e. protoc itself, which CI
// installs with arduino/setup-protoc.
func getToolPath(name string) (string, error) {
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	// check under gopath
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	p := filepath.Join(gopath, "bin", name)
	if _, err := os.Stat(p); err != nil {
		return "", err
	}
	return p, nil
}

func connectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
