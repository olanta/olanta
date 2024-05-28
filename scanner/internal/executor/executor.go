package executor

import (
	"fmt"
	"os"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/internal/scanner"
)

func ExecuteScan(language, path string) []models.Issue {
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

	return issues
}
