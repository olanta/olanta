package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Build builds the scanner project
func Build() error {
	fmt.Println("Building the project...")

	outputDir := "../bin"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	output := "../bin/olanta"
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", output, "../cmd/scanner")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build project: %w", err)
	}

	fmt.Println("Build completed successfully.")
	return nil
}

// Clean removes the build artifacts
func Clean() {
	fmt.Println("Cleaning the project...")
	sh.Rm("../bin")
}

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build
