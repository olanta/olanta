package utils

import "gopkg.in/yaml.v2"

type Rule struct {
    Pattern     string `yaml:"pattern"`
    Description string `yaml:"description"`
    Severity    string `yaml:"severity"`
}

type RulesConfig struct {
    Rules []Rule `yaml:"rules"`
}

func LoadRulesFromYAML(data []byte) (RulesConfig, error) {
    var rules RulesConfig
    err := yaml.Unmarshal(data, &rules)
    return rules, err
}
