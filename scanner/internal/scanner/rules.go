package scanner

import _ "embed"

//go:embed rules/java_rules.yaml
var javaRulesYAML []byte

//go:embed rules/python_rules.yaml
var pythonRulesYAML []byte
