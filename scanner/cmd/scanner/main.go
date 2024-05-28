package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    "github.com/spf13/cobra"
    "github.com/olanta/olanta/scanner/internal/scanner"
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

            var issues []scanner.Issue

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

            submitIssues(coreURL, issues)
        },
    }

    scanCmd.Flags().StringVar(&coreURL, "core-url", "http://localhost:8080/submit", "URL of the Olanta core server")
    rootCmd.AddCommand(scanCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func submitIssues(url string, issues []scanner.Issue) {
    jsonData, err := json.Marshal(issues)
    if err != nil {
        fmt.Printf("Error marshalling issues: %v\n", err)
        os.Exit(1)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Printf("Error submitting issues: %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Failed to submit issues: %s\n", resp.Status)
        os.Exit(1)
    }

    fmt.Println("Issues submitted successfully")
}
