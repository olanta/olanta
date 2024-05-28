package scanner

import (
    "github.com/olanta/olanta/scanner/internal/models"
)

type Scanner interface {
    Scan(path string) []models.Issue
}
