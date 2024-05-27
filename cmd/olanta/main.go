package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/olanta/olanta/internal/indexer"
    "github.com/olanta/olanta/internal/scanner"
    "github.com/olanta/olanta/internal/helpers"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "olanta",
        Short: "Olanta is a static code analysis tool to find code smells and bugs.",
    }

    scanCmd := &cobra.Command{
        Use:   "scan [language] [path]",
        Short: "Scan the specified path for code smells and bugs.",
        Args:  cobra.ExactArgs(2),
        Run: func(cmd *cobra.Command, args []string) {
            language := args[0]
            path := args[1]

            index, err := indexer.NewIndexer(indexer.IndexPath)
            if err != nil {
                helpers.Logger.Errorf("Error creating indexer: %v", err)
                os.Exit(1)
            }

            switch language {
            case "java":
                scanner := scanner.NewJavaScanner(index)
                issues := scanner.Scan(path)
                for _, issue := range issues {
                    fmt.Printf("%+v\n", issue)
                }
            case "python":
                scanner := scanner.NewPythonScanner(index)
                issues := scanner.Scan(path)
                for _, issue := range issues {
                    fmt.Printf("%+v\n", issue)
                }
            default:
                fmt.Printf("Unsupported language: %s\n", language)
                os.Exit(1)
            }
        },
    }

    rootCmd.AddCommand(scanCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
