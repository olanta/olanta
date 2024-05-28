package scanner

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"

    "github.com/olanta/olanta/scanner/internal/models"
    "github.com/olanta/olanta/scanner/internal/utils"
)

type PythonScanner struct {
    rules utils.RulesConfig
}

func NewPythonScanner() *PythonScanner {
    rules, err := utils.LoadRulesFromYAML(pythonRulesYAML)
    if err != nil {
        fmt.Printf("Error loading rules: %v\n", err)
    }
    return &PythonScanner{rules: rules}
}

func (s *PythonScanner) Scan(path string) []models.Issue {
    var issues []models.Issue
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return issues
    }

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".py") {
            fileIssues := s.scanFile(path + "/" + file.Name())
            issues = append(issues, fileIssues...)
        }
    }

    return issues
}

func (s *PythonScanner) scanFile(filePath string) []models.Issue {
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
        }
    }

    return issues
}

func (s *PythonScanner) getLine(content []byte, offset int) int {
    return strings.Count(string(content[:offset]), "\n") + 1
}
