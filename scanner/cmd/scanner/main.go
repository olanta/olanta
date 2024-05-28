package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/internal/scanner"
	"github.com/olanta/olanta/scanner/proto"
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
				printIssues(issues)
			} else {
				fmt.Println("Issues submitted successfully")
			}
		},
	}

	scanCmd.Flags().StringVar(&coreURL, "core-url", "localhost:8080", "URL of the Olanta core server")
	rootCmd.AddCommand(scanCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func submitIssues(coreURL string, issues []models.Issue) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, coreURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewScannerServiceClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var protoIssues []*proto.Issue
	for _, issue := range issues {
		protoIssues = append(protoIssues, &proto.Issue{
			Description: issue.Description,
			Severity:    issue.Severity,
			File:        issue.File,
			Line:        int32(issue.Line),
			Column:      int32(issue.Column),
		})
	}

	_, err = client.Scan(ctx, &proto.ScanRequest{
		Language: "java",
		Path:     "path",
		Issues:   protoIssues,
	})

	if err != nil {
		return fmt.Errorf("could not submit issues: %v", err)
	}

	return nil
}

func printIssues(issues []models.Issue) {
	fmt.Println("Printing issues:")
	for _, issue := range issues {
		fmt.Printf("Description: %s\n", issue.Description)
		fmt.Printf("Severity: %s\n", issue.Severity)
		fmt.Printf("File: %s\n", issue.File)
		fmt.Printf("Line: %d, Column: %d\n", issue.Line, issue.Column)
		fmt.Println("----")
	}
}
