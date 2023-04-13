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

	"github.com/livekit/mageutil"
)

var Default = Proto

func Bootstrap() error {
	return mageutil.Run(context.Background(),
		"go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.3",
		"go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1",
		"go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0",
		"go install github.com/livekit/psrpc/protoc-gen-psrpc@v0.2.10",
	)
}

// regenerate protobuf
func Proto() error {
	twirpProtoFiles := []string{
		"livekit_egress.proto",
		"livekit_ingress.proto",
		"livekit_room.proto",
	}
	grpcProtoFiles := []string{
		"livekit_analytics.proto",
		"livekit_internal.proto",
		"livekit_models.proto",
		"livekit_rpc_internal.proto",
		"livekit_rtc.proto",
		"livekit_webhook.proto",
	}
	psrpcProtoFiles := []string{
		"rpc/egress.proto",
		"rpc/ingress.proto",
		"rpc/io.proto",
		"rpc/signal.proto",
	}

	fmt.Println("generating protobuf")
	target := "livekit"
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}
	protocGoPath, err := getToolPath("protoc-gen-go")
	if err != nil {
		return err
	}
	twirpPath, err := getToolPath("protoc-gen-twirp")
	if err != nil {
		return err
	}
	protocGrpcGoPath, err := getToolPath("protoc-gen-go-grpc")
	if err != nil {
		return err
	}
	fmt.Println("generating twirp protobuf")
	args := append([]string{
		"--go_out", target,
		"--twirp_out", target,
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=twirp=" + twirpPath,
		"-I=.",
	}, twirpProtoFiles...)
	cmd := exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("generating grpc protobuf")
	args = append([]string{
		"--go_out", target,
		"--go-grpc_out", target,
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=go-grpc=" + protocGrpcGoPath,
		"-I=.",
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
	psrpcPath, err := mageutil.GetToolPath("protoc-gen-psrpc")
	if err != nil {
		return err
	}

	args = append([]string{
		"--go_out", ".",
		"--psrpc_out", ".",
		"--go_opt=paths=source_relative",
		"--psrpc_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=psrpc=" + psrpcPath,
		"-I" + psrpcDir + "/protoc-gen-psrpc/options",
		"-I=.",
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
	cmd := exec.Command("go", "test", "./...")
	connectStd(cmd)
	return cmd.Run()
}

// helpers

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
