// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	// mage:import
	"github.com/magefile/mage/sh"
)

// Build compiles the code

func Build() error {
	fmt.Println("Building the project...")

	binaryName := "../bin/olanta"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", binaryName, "../cmd/olanta")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Test runs the tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}
