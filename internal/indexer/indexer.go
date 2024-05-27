package indexer

import (
	"encoding/json"
	"github.com/blevesearch/bleve/v2"
	"github.com/olanta/olanta/internal/models"
)

type Indexer struct {
	index bleve.Index
}

func NewIndexer(path string) (*Indexer, error) {
	index, err := bleve.Open(path)
	if err != nil {
		indexMapping := bleve.NewIndexMapping()
		index, err = bleve.New(path, indexMapping)
		if err != nil {
			return nil, err
		}
	}
	return &Indexer{index: index}, nil
}

func (i *Indexer) IndexIssue(issue models.Issue) error {
	issueJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	return i.index.Index(issue.File, issueJSON)
}

func (i *Indexer) SearchIssues(query string) ([]models.Issue, error) {
	queryString := bleve.NewQueryStringQuery(query)
	search := bleve.NewSearchRequest(queryString)
	searchResults, err := i.index.Search(search)
	if err != nil {
		return nil, err
	}

	var issues []models.Issue
	for _, hit := range searchResults.Hits {
		var issue models.Issue
		issueData, ok := hit.Fields[""].([]byte)
		if !ok {
			return nil, err
		}
		err := json.Unmarshal(issueData, &issue)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, nil
}
