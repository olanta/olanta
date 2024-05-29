package main

import (
	"fmt"
	"os"

	"github.com/olanta/olanta/scanner/internal/executor"
	"github.com/olanta/olanta/scanner/internal/submitter"
	"github.com/spf13/cobra"
)

func main() {
	var coreURL string

	rootCmd := &cobra.Command{
		Use:   "scanner",
		Short: "Olanta scanner",
	}

	scanCmd := &cobra.Command{
		Use:   "scan [language] [path]",
		Short: "Scan the specified path for code smells and bugs.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			language := args[0]
			path := args[1]

			issues := executor.ExecuteScan(language, path)
			submitter.SubmitOrPrintIssues(coreURL, issues)
		},
	}

	scanCmd.Flags().StringVar(&coreURL, "core-url", "localhost:8080", "URL of the Olanta core server")
	rootCmd.AddCommand(scanCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
