package testcases

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "testcases",
	Doc:  "Checks that table-driven tests contain at least one positive and one negative test case",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if !isTestFile(file) {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok || !isTestFunction(funcDecl) {
				return true
			}

			checkTableDrivenTests(pass, funcDecl)
			return true
		})
	}

	return nil, nil
}

func isTestFile(f *ast.File) bool {
	return f != nil && f.Name != nil
}

func isTestFunction(fn *ast.FuncDecl) bool {
	return fn.Name != nil && len(fn.Name.Name) > 4 && fn.Name.Name[:4] == "Test"
}

func checkTableDrivenTests(pass *analysis.Pass, fn *ast.FuncDecl) {
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		if compLit, ok := n.(*ast.CompositeLit); ok {
			if isTableTest(compLit) {
				hasPositive, hasNegative := checkTestCases(compLit)
				if !hasPositive || !hasNegative {
					pass.Reportf(fn.Pos(), "table-driven test %q should contain at least one positive and one negative test case", fn.Name.Name)
				}
			}
		}
		return true
	})
}

func isTableTest(compLit *ast.CompositeLit) bool {
	if arrayType, ok := compLit.Type.(*ast.ArrayType); ok {
		if _, ok := arrayType.Elt.(*ast.StructType); ok {
			return true
		}
	}
	return false
}

func checkTestCases(compLit *ast.CompositeLit) (hasPositive, hasNegative bool) {
	for _, elt := range compLit.Elts {
		if lit, ok := elt.(*ast.CompositeLit); ok {
			for _, field := range lit.Elts {
				if kv, ok := field.(*ast.KeyValueExpr); ok {
					if key, ok := kv.Key.(*ast.Ident); ok {
						if key.Name == "wantErr" || key.Name == "shouldFail" || key.Name == "isError" || key.Name == "expectError" {
							if val, ok := kv.Value.(*ast.Ident); ok {
								if val.Name == "true" {
									hasNegative = true
								} else if val.Name == "false" {
									hasPositive = true
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
