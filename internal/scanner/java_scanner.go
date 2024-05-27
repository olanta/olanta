package scanner

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/olanta/olanta/internal/indexer"
	"github.com/olanta/olanta/internal/models"
	"github.com/olanta/olanta/internal/utils"
)

type JavaScanner struct {
	indexer *indexer.Indexer
	rules   utils.RulesConfig
}

func NewJavaScanner(indexer *indexer.Indexer) *JavaScanner {
	rules, err := utils.LoadRulesFromFile("rules/java_rules.yaml")
	if err != nil {
		fmt.Printf("Error loading rules: %v\n", err)
	}
	return &JavaScanner{indexer: indexer, rules: rules}
}

func (s *JavaScanner) Scan(path string) []models.Issue {
	var issues []models.Issue
	files, err := ioutil.ReadDir(path)
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
	content, err := ioutil.ReadFile(filePath)
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
			s.indexer.IndexIssue(issue)
		}
	}

	return issues
}

func (s *JavaScanner) getLine(content []byte, offset int) int {
	return strings.Count(string(content[:offset]), "\n") + 1
}
