//go:build mage

package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

func Build() error {
    fmt.Println("Building the project...")

    output := "../bin/olanta"
    if runtime.GOOS == "windows" {
        output = "../bin/olanta.exe"
    }

    cmd := exec.Command("go", "build", "-o", output, "../cmd/olanta")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
