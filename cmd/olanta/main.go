package main

import (
    "github.com/olanta/internal/indexer"
    "github.com/olanta/internal/scanner"
    "github.com/olanta/internal/utils"
    "log"
)

func main() {
    // Initialize the indexer
    idx, err := indexer.NewIndexer(indexer.IndexPath)
    if err != nil {
        log.Fatalf("Erro ao inicializar o indexador: %v", err)
    }

    // Initialize scanners for Java and Python
    javaScanner := scanner.NewJavaScanner(idx)
    pythonScanner := scanner.NewPythonScanner(idx)

    // Path of the project to be analyzed
    projectPath := "path/to/your/project"

    // Run the analysis
    javaIssues := javaScanner.Scan(projectPath)
    pythonIssues := pythonScanner.Scan(projectPath)

    // Log the results
    utils.Logger.Infof("Java Issues: %v", javaIssues)
    utils.Logger.Infof("Python Issues: %v", pythonIssues)
}
