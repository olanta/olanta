package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	Pattern     string `yaml:"pattern"`
	Description string `yaml:"description"`
	Severity    string `yaml:"severity"`
}

type RulesConfig struct {
	Rules []Rule `yaml:"rules"`
}

func LoadRulesFromFile(filepath string) (RulesConfig, error) {
	var rules RulesConfig
	data, err := os.ReadFile(filepath)
	if err != nil {
		return rules, err
	}
	err = yaml.Unmarshal(data, &rules)
	return rules, err
}

func LoadRules(data []byte) (RulesConfig, error) {
	var rules RulesConfig
	err := yaml.Unmarshal(data, &rules)
	return rules, err
}
