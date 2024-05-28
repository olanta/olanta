package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/internal/scanner"
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

			var issues []models.Issue

			switch language {
			case "java":
				s := scanner.NewJavaScanner()
				issues = s.Scan(path)
			case "python":
				s := scanner.NewPythonScanner()
				issues = s.Scan(path)
			default:
				fmt.Printf("Unsupported language: %s\n", language)
				os.Exit(1)
			}

			if err := submitIssues(coreURL, issues); err != nil {
				fmt.Printf("Error submitting issues: %v\n", err)
				fmt.Println("Printing issues:")
				printIssues(issues)
			}
		},
	}

	scanCmd.Flags().StringVar(&coreURL, "core-url", "http://localhost:8080/submit", "URL of the Olanta core server")
	rootCmd.AddCommand(scanCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func submitIssues(url string, issues []models.Issue) error {
	jsonData, err := json.Marshal(issues)
	if err != nil {
		return fmt.Errorf("error marshalling issues: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error submitting issues: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to submit issues: %s", resp.Status)
	}

	fmt.Println("Issues submitted successfully")
	return nil
}

func printIssues(issues []models.Issue) {
	for _, issue := range issues {
		fmt.Printf("Description: %s\n", issue.Description)
		fmt.Printf("Severity: %s\n", issue.Severity)
		fmt.Printf("File: %s\n", issue.File)
		fmt.Printf("Line: %d, Column: %d\n", issue.Line, issue.Column)
		fmt.Println("----")
	}
}
