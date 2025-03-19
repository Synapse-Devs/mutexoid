package analyzers

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	parallelDiag = "test function %s should call t.Parallel()"
	tableDiag    = "test function %s should use table-driven tests"
)

var ParallelTests = &analysis.Analyzer{
	Name:     "paralleltests",
	Doc:      "Ensures all tests are parallel and use table-driven approach",
	Run:      runParallelTests,
	Requires: []*analysis.Analyzer{},
}

func runParallelTests(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !strings.HasSuffix(filename, "_test.go") {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				if isTestFunction(funcDecl) {
					checkTestFunction(pass, funcDecl)
				}
			}
			return true
		})
	}
	return nil, nil
}

func isTestFile(file *ast.File) bool {
	return file.Name.Name != "" && len(file.Name.Name) > 5 && file.Name.Name[len(file.Name.Name)-5:] == "_test"
}

func isTestFunction(funcDecl *ast.FuncDecl) bool {
	return funcDecl.Name.Name != "TestMain" &&
		len(funcDecl.Name.Name) > 4 &&
		funcDecl.Name.Name[:4] == "Test"
}

func checkTestFunction(pass *analysis.Pass, funcDecl *ast.FuncDecl) {
	hasParallel := false
	hasTableTests := false

	ast.Inspect(funcDecl, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			if isParallelCall(node) {
				hasParallel = true
			}
		case *ast.RangeStmt:
			if isTableTest(node) {
				hasTableTests = true
			}
		}
		return true
	})

	if !hasParallel {
		pass.Reportf(funcDecl.Pos(), parallelDiag, funcDecl.Name.Name)
	}

	if !hasTableTests {
		pass.Reportf(funcDecl.Pos(), tableDiag, funcDecl.Name.Name)
	}
}

func isParallelCall(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if ident, ok := sel.X.(*ast.Ident); ok {
			return ident.Name == "t" && sel.Sel.Name == "Parallel"
		}
	}
	return false
}

func isTableTest(rangeStmt *ast.RangeStmt) bool {
	// Check if we're ranging over a test cases slice/array
	if ident, ok := rangeStmt.X.(*ast.Ident); ok {
		name := ident.Name
		return name == "tests" || name == "testCases" || name == "cases"
	}
	return false
}
