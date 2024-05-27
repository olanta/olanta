package main

import (
    "github.com/olanta/config"
    "github.com/olanta/internal/analyzer"
    "github.com/olanta/internal/helpers"
)

func main() {
    config := config.Load()
    analyzer := analyzer.NewAnalyzer(config)
    helpers.Logger.Info("Starting Olanta analysis")
    analyzer.Run()
}
