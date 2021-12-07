//go:build mage
// +build mage

package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/target"
)

var Default = Proto

// regenerate protobuf
func Proto() error {
	updated, err := target.Path("proto/livekit_models.pb.go",
		"livekit_internal.proto",
		"livekit_models.proto",
		"livekit_recording.proto",
		"livekit_room.proto",
		"livekit_rtc.proto",
		"livekit_webhook.proto",
		"livekit_analytics.proto",
	)
	if err != nil {
		return err
	}
	if !updated {
		return nil
	}

	fmt.Println("generating protobuf")
	target := "proto"
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
	// generate twirp-related protos
	cmd := exec.Command(protoc,
		"--go_out", target,
		"--twirp_out", target,
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=go="+protocGoPath,
		"--plugin=twirp="+twirpPath,
		"-I=.",
		"livekit_recording.proto",
		"livekit_room.proto",
	)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	// generate basic protobuf
	cmd = exec.Command(protoc,
		"--go_out", target,
		"--go-grpc_out", target,
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--plugin=go="+protocGoPath,
		"--plugin=go-grpc="+protocGrpcGoPath,
		"-I=.",
		"livekit_internal.proto",
		"livekit_models.proto",
		"livekit_rtc.proto",
		"livekit_webhook.proto",
		"livekit_analytics.proto",
	)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
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
