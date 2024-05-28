package submitter

import (
	"context"
	"fmt"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/proto"
	"google.golang.org/grpc"
)

func SubmitOrPrintIssues(coreURL string, issues []models.Issue) {
	err := submitIssuesGRPC(coreURL, issues)
	if err != nil {
		fmt.Printf("Error submitting issues: %v\n", err)
		printIssues(issues)
	}
}

func submitIssuesGRPC(url string, issues []models.Issue) error {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewScannerServiceClient(conn)

	protoIssues := convertToProtoIssues(issues)
	_, err = client.Scan(context.Background(), &proto.ScanRequest{
		Language: "java",
		Path:     "path",
		Issues:   protoIssues,
	})
	return err
}

func printIssues(issues []models.Issue) {
	fmt.Println("Printing issues:")
	for _, issue := range issues {
		fmt.Printf("\nDescription: %s\nSeverity: %s\nFile: %s\nLine: %d, Column: %d\n----\n",
			issue.Description, issue.Severity, issue.File, issue.Line, issue.Column)
	}
}

func convertToProtoIssues(issues []models.Issue) []*proto.Issue {
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
	return protoIssues
}
