package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/olanta/olanta/internal/models"
)

type Indexer struct {
	index bleve.Index
}

type Config struct {
	IndexPath string
}

func NewIndexer(cfg Config) (*Indexer, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(cfg.IndexPath, mapping)
	if err != nil {
		return nil, err
	}

	return &Indexer{index: index}, nil
}

func (i *Indexer) IndexIssue(issue models.Issue) error {
	return i.index.Index(issue.ID, issue)
}
