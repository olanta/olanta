package utils

import (
    "github.com/olanta/olanta/internal/helpers"
    "github.com/olanta/olanta/internal/models"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

// LoadRulesFromFile lê e deserializa as regras de um arquivo YAML
func LoadRulesFromFile(filepath string) (models.Rules, error) {
    var rules models.Rules
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
