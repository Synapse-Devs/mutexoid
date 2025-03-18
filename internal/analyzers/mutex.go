package analyzers

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// registerMutexAnalyzer registers the mutex analyzer
func (r *Registry) registerMutexAnalyzer() {
	r.register(mutexAnalyzer)
}

var mutexAnalyzer = &analysis.Analyzer{
	Name: "mutexoid",
	Doc:  "Checks for incorrect mutex usage patterns",
	Run:  runMutexAnalysis,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func runMutexAnalysis(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		structType, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		for _, field := range structType.Fields.List {
			if isMutexField(field) && !isPointerType(field.Type) {
				pass.Reportf(field.Pos(), "mutex should be a pointer type")
			}
		}
	})

	return nil, nil
}

func isMutexField(field *ast.Field) bool {
	if sel, ok := field.Type.(*ast.SelectorExpr); ok {
		if ident, ok := sel.X.(*ast.Ident); ok {
			return ident.Name == "sync" && (sel.Sel.Name == "Mutex" || sel.Sel.Name == "RWMutex")
		}
	}
	return false
}

func isPointerType(expr ast.Expr) bool {
	_, ok := expr.(*ast.StarExpr)
	return ok
}
