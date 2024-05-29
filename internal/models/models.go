package models

import "github.com/google/uuid"

type Issue struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
}

func NewIssue(description, severity, file string, line, column int) Issue {
	return Issue{
		ID:          uuid.New().String(),
		Description: description,
		Severity:    severity,
		File:        file,
		Line:        line,
		Column:      column,
	}
}
