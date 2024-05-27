package utils

import (
    "github.com/olanta/internal/models"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

func LoadRules(path string) (models.Rules, error) {
    var rules models.Rules
    data, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("Error reading YAML file: %s\n", err)
        return rules, err
    }
    err = yaml.Unmarshal(data, &rules)
    if err != nil {
        log.Printf("Error unmarshalling YAML file: %s\n", err)
        return rules, err
    }
    return rules, nil
}