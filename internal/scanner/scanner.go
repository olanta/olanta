package scanner

import (
	"github.com/olanta/olanta/internal/models"
)

type Scanner interface {
	Scan(path string) []models.Issue
}