package analyzer

import (
    "github.com/olanta/config"
    "github.com/olanta/internal/helpers"
    "github.com/olanta/internal/scanner"
)

type Analyzer struct {
    config *config.Config
}

func NewAnalyzer(cfg *config.Config) *Analyzer {
    return &Analyzer{
        config: cfg,
    }
}

func (a *Analyzer) Run() {
    helpers.Logger.Info(helpers.MsgStartAnalysis)
    s := scanner.NewScanner(a.config)
    s.Scan()
    helpers.Logger.Info(helpers.MsgEndAnalysis)
}
