// +build mage

package main

import (
	"fmt"
    "os"

	// mage:import
	"github.com/magefile/mage/sh"
)

// Build compiles the code
func Build() error {
	fmt.Println("Building the project...")
	if err := os.MkdirAll("bin", os.ModePerm); err != nil {
		return err
	}
	return sh.RunV("go", "build", "-o", "bin/olanta", "../cmd/olanta")
}

// Test runs the tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}
