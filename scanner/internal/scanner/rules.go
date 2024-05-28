package scanner

import (
	"embed"
	"fmt"
)

//go:embed rules/*
var embeddedRules embed.FS

func LoadEmbeddedRules(filename string) ([]byte, error) {
	data, err := embeddedRules.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading embedded rules: %w", err)
	}
	return data, nil
}
