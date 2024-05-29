package submitter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/proto"
	"google.golang.org/grpc"
)

func SubmitOrPrintIssues(coreURL string, issues []models.Issue) {
	// Conectar ao servidor gRPC
	conn, err := grpc.Dial(coreURL, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Error connecting to core server: %v", err)
		printIssues(issues)
		return
	}
	defer conn.Close()

	client := proto.NewScannerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	protoIssues := convertToProtoIssues(issues)
	_, err = client.Scan(ctx, &proto.ScanRequest{
		Language: "java",
		Path:     "path",
		Issues:   protoIssues,
	})
	if err != nil {
		log.Printf("Error submitting issues: %v", err)
		printIssues(issues)
		return
	}

	fmt.Println("Issues submitted successfully")
}

func convertToProtoIssues(issues []models.Issue) []*proto.Issue {
	var protoIssues []*proto.Issue
	for _, issue := range issues {
		protoIssue := &proto.Issue{
			Description: issue.Description,
			Severity:    issue.Severity,
			File:        issue.File,
			Line:        int32(issue.Line),
			Column:      int32(issue.Column),
		}
		protoIssues = append(protoIssues, protoIssue)
	}
	return protoIssues
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
