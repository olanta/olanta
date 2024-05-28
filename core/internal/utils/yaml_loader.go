package utils

import (
	"github.com/olanta/olanta/internal/helpers"
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
		helpers.Logger.Errorf(helpers.ErrReadingFile, err)
		return rules, err
	}
	err = yaml.Unmarshal(data, &rules)
	if err != nil {
		helpers.Logger.Errorf(helpers.ErrLoadingRules, err)
		return rules, err
	}
	return rules, nil
}
