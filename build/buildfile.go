package main

import (
    "fmt"
    "github.com/magefile/mage/sh"
)

// Build compiles the Plant project
func Build() error {
    fmt.Println("Building the project...")
    return sh.RunV("go", "build", "./cmd/olanta")
}

// Test runs the Plant project tests
func Test() error {
    fmt.Println("Running tests...")
    return sh.RunV("go", "test", "./...")
}

// Lint checks the quality of the Planta project code
func Lint() error {
    fmt.Println("Linting the code...")
    return sh.RunV("golangci-lint", "run")
}
