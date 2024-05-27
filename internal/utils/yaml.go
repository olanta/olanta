package utils

import (
	"github.com/olanta/internal/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadRules(filePath string) (models.Rules, error) {
	var rules models.Rules
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return rules, err
	}
	err = yaml.Unmarshal(data, &rules)
	return rules, err
}
