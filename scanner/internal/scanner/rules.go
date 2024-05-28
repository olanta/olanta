package scanner

import (
	"embed"
	"fmt"
)

//go:embed rules/*
var embeddedRules embed.FS

func LoadEmbeddedRules(fileName string) ([]byte, error) {
	data, err := embeddedRules.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading embedded rule file %s: %w", fileName, err)
	}
	return data, nil
}
