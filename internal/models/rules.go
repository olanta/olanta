package models

type Rule struct {
	Description string `yaml:"description"`
	Pattern     string `yaml:"pattern"`
	Severity    string `yaml:"severity"`
}

type Rules struct {
	Rules []Rule `yaml:"rules"`
}
