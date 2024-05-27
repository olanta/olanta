package main

import (
    "fmt"
    "github.com/magefile/mage/sh"
)

// Build compiles the project.
func Build() error {
    fmt.Println("Building the project...")
    return sh.RunV("go", "build", "-o", "bin/olanta", "./cmd/olanta")
}

// Test runs the tests for the project.
func Test() error {
    fmt.Println("Running tests...")
    return sh.RunV("go", "test", "./...")
}

// Clean removes the binary and other build artifacts.
func Clean() error {
    fmt.Println("Cleaning build artifacts...")
    return sh.Rm("bin/olanta")
}

// GetCurrentDate prints the current date.
func GetCurrentDate() error {
    date, err := sh.Output("date", "+%a %b %d %H:%M:%S %Y")
    if err != nil {
        return err
    }

    fmt.Printf("::set-output name=date::%s\n", date)
    return nil
}
