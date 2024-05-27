package main

import (
    "fmt"
    "os"

    "github.com/magefile/mage/mg"
    "github.com/magefile/mage/sh"
)

func Build() error {
    fmt.Println("Building the project...")
    return sh.RunV("go", "build", "-o", "olanta", "./cmd/olanta")
}

func Test() error {
    fmt.Println("Running tests...")
    return sh.RunV("go", "test", "./...")
}

func Clean() error {
    fmt.Println("Cleaning build artifacts...")
    return os.RemoveAll("olanta")
}

var Default = Build
