package scanner

import (
    "bufio"
    "os"
    "path/filepath"
    "strings"

    "github.com/olanta/olanta/internal/helpers"
    "github.com/olanta/olanta/internal/models"
    "github.com/olanta/olanta/internal/utils"
    "github.com/olanta/olanta/internal/indexer"
)

type JavaScanner struct {
    rules   models.Rules
    indexer *indexer.Indexer
}

func NewJavaScanner(indexer *indexer.Indexer) *JavaScanner {
    rules, err := utils.LoadRules("rules/java_rules.yaml")
    if err != nil {
        helpers.Logger.Errorf("Error loading Java rules: %v", err)
    }
    return &JavaScanner{rules: rules, indexer: indexer}
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
        for _, rule := range s.rules.Rules {
            if strings.Contains(line, rule.Pattern) {
                issue := models.Issue{
                    ID:          generateID(filePath, lineNumber),
                    Description: rule.Description,
                    Severity:    rule.Severity,
                    File:        filePath,
                    Line:        lineNumber,
                    Code:        line,
                }
                issues = append(issues, issue)
                s.indexer.IndexIssue(issue)
            }
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        helpers.Logger.Errorf("Error reading file %s: %v", filePath, err)
    }

    return issues
}

func generateID(filePath string, lineNumber int) string {
    // Implement a function to generate a unique ID for each issue
    return fmt.Sprintf("%s:%d", filePath, lineNumber)
}
