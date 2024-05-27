package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/olanta/internal/helpers"
	"github.com/olanta/internal/models"
)

type PythonScanner struct{}

func NewPythonScanner() *PythonScanner {
	return &PythonScanner{}
}

func (s *PythonScanner) Scan(path string) []models.Issue {
	helpers.Logger.Infof("Scanning Python files in %s", path)
	var issues []models.Issue

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".py") {
			fileIssues := s.scanFile(path)
			issues = append(issues, fileIssues...)
		}
		return nil
	})

	if err != nil {
		helpers.Logger.Errorf("Error scanning Python files: %v", err)
	}

	return issues
}

func (s *PythonScanner) scanFile(filePath string) []models.Issue {
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
		if strings.Contains(line, "def ") || strings.Contains(line, "class ") {
			s.checkForUnusedVariables(filePath, line, lineNumber, &issues)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		helpers.Logger.Errorf("Error reading file %s: %v", filePath, err)
	}

	return issues
}

func (s *PythonScanner) checkForUnusedVariables(filePath, line string, lineNumber int, issues *[]models.Issue) {
    // This function must be implemented to check unused variables
	 // For simplicity, we added an example issue
	if strings.Contains(line, "unused_variable") {
		*issues = append(*issues, models.Issue{
			Description: "Unused variable",
			Severity:    "info",
			File:        filePath,
			Line:        lineNumber,
		})
	}
}

func (s *PythonScanner) isImportUsed(filePath, importLine string) bool {
    // This function must be implemented to check if import is used in the file
	// For simplicity, we return false to demonstrate how it works
	return false
}
