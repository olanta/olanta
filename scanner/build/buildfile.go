//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Build builds the scanner project.
func Build() error {
	fmt.Println("Building the project...")

	output := "bin/scanner"
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", output, "../cmd/scanner")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Clean removes the built binary.
func Clean() error {
	fmt.Println("Cleaning the project...")
	output := "../bin/scanner"
	if runtime.GOOS == "windows" {
		output += ".exe"
	}
	return os.Remove(output)
}

// Run runs the scanner project.
func Run() error {
	fmt.Println("Running the project...")
	output := "../bin/scanner"
	if runtime.GOOS == "windows" {
		output += ".exe"
	}
	return exec.Command(output, "scan", "java", "../path/to/scan").Run()
}
