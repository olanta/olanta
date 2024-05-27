package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/olanta/internal/helpers"
	"github.com/olanta/internal/models"
)

type JavaScanner struct{}

func NewJavaScanner() *JavaScanner {
	return &JavaScanner{}
}

func (s *JavaScanner) Scan(path string) []models.Issue {
	helpers.Logger.Infof("Scanning Java files in %s", path)
	var issues []models.Issue

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".java") {
			fileIssues := s.scanFile(path)
			issues = append(issues, fileIssues...)
		}
		return nil
	})

	if err != nil {
		helpers.Logger.Errorf("Error scanning Java files: %v", err)
	}

	return issues
}

func (s *JavaScanner) scanFile(filePath string) []models.Issue {
	var issues []models.Issue
	file, err := os.Open(filePath)
	if err != nil {
		helpers.Logger.Errorf("Error opening file %s: %v", filePath, err)
		return issues
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "import") && !s.isImportUsed(filePath, line) {
			issues = append(issues, models.Issue{
				Description: "Unused import statement",
				Severity:    "warning",
				File:        filePath,
				Line:        lineNumber,
			})
		}
		if strings.Contains(line, "null") {
			issues = append(issues, models.Issue{
				Description: "Possible null pointer dereference",
				Severity:    "critical",
				File:        filePath,
				Line:        lineNumber,
			})
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		helpers.Logger.Errorf("Error reading file %s: %v", filePath, err)
	}

	return issues
}

func (s *JavaScanner) isImportUsed(filePath, importLine string) bool {
    // This function must be implemented to check if import is used in the file
	 // For simplicity, we return false to demonstrate how it works
	return false
}
