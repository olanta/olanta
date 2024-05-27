package helpers

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
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
    data, err := ioutil.ReadFile(filepath)
    if err != nil {
        Logger.Errorf("Error reading rules file: %v", err)
        return rules, err
    }
    err = yaml.Unmarshal(data, &rules)
    if err != nil {
        Logger.Errorf("Error unmarshalling rules: %v", err)
        return rules, err
    }
    return rules, nil
}
