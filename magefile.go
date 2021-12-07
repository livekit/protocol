//go:build mage
// +build mage

package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
)

var Default = Proto

// regenerate protobuf
func Proto() error {
	twirpProtoFiles := []string{
		"proto/livekit/livekit_recording.proto",
		"proto/livekit/livekit_room.proto",
	}
	allProtoFiles, err := filepath.Glob("proto/livekit/*.proto")
	if err != nil {
		return err
	}
	target := "proto"
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
		"--go_out", ".",
		"--twirp_out", ".",
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=twirp=" + twirpPath,
		"-I=.",
	}, twirpProtoFiles...)
	cmd := exec.Command(protoc, args...)
	fmt.Println("CMD", cmd)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	if true {
		return nil
	}
	fmt.Println("generating basic protobuf")
	args = append([]string{
		"--go_out", ".",
		"--go-grpc_out", ".",
		"--go_opt=paths=import",
		"--go-grpc_opt=paths=import",
		"--plugin=go=" + protocGoPath,
		"--plugin=go-grpc=" + protocGrpcGoPath,
		"-I=.",
	}, allProtoFiles...)
	cmd = exec.Command(protoc, args...)
	fmt.Println("CMD", cmd)
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
