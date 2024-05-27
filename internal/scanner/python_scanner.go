package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/olanta/internal/helpers"
	"github.com/olanta/internal/models"
	"github.com/olanta/internal/utils"
)

type PythonScanner struct {
	rules models.Rules
}

func NewPythonScanner() *PythonScanner {
	rules, err := utils.LoadRules("rules/python_rules.yaml")
	if err != nil {
		helpers.Logger.Errorf("Error loading Python rules: %v", err)
	}
	return &PythonScanner{rules: rules}
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
		for _, rule := range s.rules.Rules {
			if strings.Contains(line, rule.Pattern) {
				issues = append(issues, models.Issue{
					Description: rule.Description,
					Severity:    rule.Severity,
					File:        filePath,
					Line:        lineNumber,
				})
			}
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		helpers.Logger.Errorf("Error reading file %s: %v", filePath, err)
	}

	return issues
}
