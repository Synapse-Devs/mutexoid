package analyzers

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

var TestPackageAnalyzer = &analysis.Analyzer{
	Name: "testpackage",
	Doc:  "Ensures that *_test.go files use *_test package names",
	Run:  runTestPackage,
}

func runTestPackage(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		filename := pass.Fset.File(file.Pos()).Name()
		if !strings.HasSuffix(filename, "_test.go") {
			continue
		}

		if !strings.HasSuffix(file.Name.Name, "_test") {
			pass.Reportf(file.Name.Pos(), "test file %s should have a package name with '_test' suffix", filename)
		}
	}

	return nil, nil
}
