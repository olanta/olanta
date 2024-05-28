package analyzer

import (
    "github.com/olanta/olanta/config"
    "github.com/olanta/olanta/internal/helpers"
    "github.com/olanta/olanta/internal/scanner"
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
