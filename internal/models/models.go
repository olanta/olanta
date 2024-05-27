package models

type Rule struct {
    Pattern     string `yaml:"pattern"`
    Description string `yaml:"description"`
    Severity    string `yaml:"severity"`
}

type Rules struct {
    Rules []Rule `yaml:"rules"`
}

type Issue struct {
    ID          string `json:"id"`
    Description string `json:"description"`
    Severity    string `json:"severity"`
    File        string `json:"file"`
    Line        int    `json:"line"`
    Code        string `json:"code"`
}
