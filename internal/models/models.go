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
	Description string
	Severity    string
	File        string
	Line        int
	Column      int
}
