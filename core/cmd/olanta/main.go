package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/olanta/olanta/core/internal/server"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "olanta",
        Short: "Olanta core server",
    }

    serveCmd := &cobra.Command{
        Use:   "serve",
        Short: "Start the Olanta server",
        Run: func(cmd *cobra.Command, args []string) {
            server.Start()
        },
    }

    rootCmd.AddCommand(serveCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
