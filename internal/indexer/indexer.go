package indexer

import (
    "github.com/blevesearch/bleve/v2"
    "github.com/olanta/olanta/internal/helpers"
    "github.com/olanta/olanta/internal/models"
    "path/filepath"
)

type Indexer struct {
    index bleve.Index
}

func NewIndexer(indexPath string) (*Indexer, error) {
    index, err := bleve.Open(indexPath)
    if err == bleve.ErrorIndexPathDoesNotExist {
        mapping := bleve.NewIndexMapping()
        index, err = bleve.New(indexPath, mapping)
        if err != nil {
            return nil, err
        }
    } else if err != nil {
        return nil, err
    }

    return &Indexer{index: index}, nil
}

func (i *Indexer) IndexIssue(issue models.Issue) error {
    return i.index.Index(issue.ID, issue)
}

func (i *Indexer) Search(queryStr string) ([]models.Issue, error) {
    query := bleve.NewQueryStringQuery(queryStr)
    searchRequest := bleve.NewSearchRequest(query)
    searchResult, err := i.index.Search(searchRequest)
    if err != nil {
        return nil, err
    }

    var issues []models.Issue
    for _, hit := range searchResult.Hits {
        var issue models.Issue
        err := i.index.Document(hit.ID, &issue)
        if err != nil {
            return nil, err
        }
        issues = append(issues, issue)
    }

    return issues, nil
}
