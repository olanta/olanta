package scanner

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/internal/utils"
)

type JavaScanner struct {
	rules utils.RulesConfig
}

func NewJavaScanner() *JavaScanner {
	rulesData, err := LoadEmbeddedRules("rules/java_rules.yaml")
	if err != nil {
		fmt.Printf("Error loading rules: %v\n", err)
		os.Exit(1)
	}
	rules, err := utils.LoadRules(rulesData)
	if err != nil {
		fmt.Printf("Error parsing rules: %v\n", err)
		os.Exit(1)
	}
	return &JavaScanner{rules: rules}
}

func (s *JavaScanner) Scan(path string) []models.Issue {
	var issues []models.Issue
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return issues
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".java") {
			fileIssues := s.scanFile(path + "/" + file.Name())
			issues = append(issues, fileIssues...)
		}
	}

	return issues
}

func (s *JavaScanner) scanFile(filePath string) []models.Issue {
	var issues []models.Issue
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return issues
	}

	for _, rule := range s.rules.Rules {
		re := regexp.MustCompile(rule.Pattern)
		matches := re.FindAllStringIndex(string(content), -1)
		for _, match := range matches {
			issue := models.Issue{
				Description: rule.Description,
				Severity:    rule.Severity,
				File:        filePath,
				Line:        s.getLine(content, match[0]),
				Column:      match[0],
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

func (s *JavaScanner) getLine(content []byte, offset int) int {
	return strings.Count(string(content[:offset]), "\n") + 1
}
