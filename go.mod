module github.com/yourusername/olanta

go 1.22

require (
	github.com/blevesearch/bleve/v2 v2.0.8
	github.com/magefile/mage v1.11.0
	github.com/spf13/cobra v1.4.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/blevesearch/bleve => github.com/blevesearch/bleve/v2 v2.0.8
