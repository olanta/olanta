package scanner

import (
    "github.com/olanta/olanta/config"
    "github.com/olanta/olanta/internal/helpers"
    "github.com/olanta/olanta/internal/models"
    "path/filepath"
)

type Scanner struct {
    config       *config.Config
    javaScanner  *JavaScanner
    pythonScanner *PythonScanner
}

func NewScanner(cfg *config.Config) *Scanner {
    return &Scanner{
        config:       cfg,
        javaScanner:  NewJavaScanner(),
        pythonScanner: NewPythonScanner(),
    }
}

func (s *Scanner) Scan() {
    issues := []models.Issue{}

    // Scan for Java files
    javaIssues := s.javaScanner.Scan(s.config.ProjectPath)
    issues = append(issues, javaIssues...)

    // Scan for Python files
    pythonIssues := s.pythonScanner.Scan(s.config.ProjectPath)
    issues = append(issues, pythonIssues...)

    for _, issue := range issues {
        helpers.Logger.Infof("Found issue: %s with severity: %s", issue.Description, issue.Severity)
    }
}
