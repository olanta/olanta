package models

type Issue struct {
    Description string `json:"description"`
    Severity    string `json:"severity"`
    File        string `json:"file"`
    Line        int    `json:"line"`
    Column      int    `json:"column"`
}

type Rule struct {
    Pattern     string `yaml:"pattern"`
    Description string `yaml:"description"`
    Severity    string `yaml:"severity"`
}

type RulesConfig struct {
    Rules []Rule `yaml:"rules"`
}
