package configvalidator

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "configvalidator",
	Doc:  "Ensures that struct fields with koanf/json tags also have validation tags",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			structType, ok := n.(*ast.StructType)
			if !ok {
				return true
			}

			for _, field := range structType.Fields.List {
				if field.Tag == nil {
					continue
				}

				tagValue := field.Tag.Value
				// Remove backticks
				tagValue = strings.Trim(tagValue, "`")

				// Check if field has koanf or json tag
				hasConfigTag := strings.Contains(tagValue, "koanf:") || strings.Contains(tagValue, "json:")

				if hasConfigTag {
					// Check if field has validate tag
					if !strings.Contains(tagValue, "validate:") {
						pass.Reportf(field.Pos(), "config field %s should have a validate tag", field.Names[0].Name)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
