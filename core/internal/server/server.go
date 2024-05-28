package server

import (
    "encoding/json"
    "net/http"

    "github.com/olanta/olanta/core/internal/indexer"
    "github.com/olanta/olanta/core/internal/models"
    "github.com/olanta/olanta/core/internal/helpers"
)

func Start() {
    http.HandleFunc("/submit", submitHandler)
    helpers.Logger.Infof("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    var issues []models.Issue
    if err := json.NewDecoder(r.Body).Decode(&issues); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    index, err := indexer.NewIndexer(indexer.IndexPath)
    if err != nil {
        helpers.Logger.Errorf("Error creating indexer: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for _, issue := range issues {
        if err := index.Index(issue); err != nil {
            helpers.Logger.Errorf("Error indexing issue: %v", err)
        }
    }

    w.WriteHeader(http.StatusOK)
}
