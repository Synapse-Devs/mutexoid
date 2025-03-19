package analyzers

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "uselesscomments",
	Doc:  "Detects redundant comments that don't add value to code understanding",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.GenDecl:
				if node.Doc != nil {
					checkComment(pass, node.Doc, node)
				}
				// Check specs inside GenDecl
				for _, spec := range node.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok {
						if ts.Doc != nil {
							checkComment(pass, ts.Doc, ts)
						}
						// Check interface methods
						if iface, ok := ts.Type.(*ast.InterfaceType); ok {
							if iface.Methods != nil {
								for _, method := range iface.Methods.List {
									if method.Doc != nil {
										checkComment(pass, method.Doc, method)
									}
								}
							}
						}
					}
				}
			case *ast.FuncDecl:
				if node.Doc != nil {
					checkComment(pass, node.Doc, node)
				}
			}
			return true
		})
	}
	return nil, nil
}

func checkComment(pass *analysis.Pass, doc *ast.CommentGroup, node ast.Node) {
	for _, comment := range doc.List {
		text := comment.Text

		// Skip if comment is a directive like //go:generate
		if len(text) > 2 && text[2] == ':' {
			continue
		}

		switch n := node.(type) {
		case *ast.GenDecl:
			if isUselessComment(text, n) {
				pass.Reportf(comment.Pos(), "redundant comment that restates what is evident from the code")
			}
		case *ast.FuncDecl:
			if isUselessFunctionComment(text, n) {
				pass.Reportf(comment.Pos(), "redundant comment that restates function signature or obvious behavior")
			}
		case *ast.TypeSpec:
			if isUselessTypeComment(text, n) {
				pass.Reportf(comment.Pos(), "redundant comment that restates type definition")
			}
		case *ast.Field:
			if isUselessMethodComment(text, n) {
				pass.Reportf(comment.Pos(), "redundant comment that restates method signature")
			}
		}
	}
}

func cleanComment(text string) string {
	// Remove comment markers and trim spaces
	text = strings.TrimPrefix(text, "//")
	text = strings.TrimPrefix(text, "/*")
	text = strings.TrimSuffix(text, "*/")
	return strings.TrimSpace(text)
}

func isUselessComment(text string, decl *ast.GenDecl) bool {
	text = cleanComment(text)

	// Check for type declarations
	if decl.Tok == token.TYPE {
		for _, spec := range decl.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				name := ts.Name.Name
				// Check if comment just repeats type name
				if strings.Contains(strings.ToLower(text), strings.ToLower(name)) &&
					(strings.HasSuffix(text, name) || strings.HasPrefix(text, name)) {
					return true
				}
				// Check for common useless patterns
				if strings.HasSuffix(text, "interface") || strings.HasSuffix(text, "struct") {
					return true
				}
			}
		}
	}
	return false
}

func isUselessFunctionComment(text string, fn *ast.FuncDecl) bool {
	text = cleanComment(text)
	name := fn.Name.Name

	// Skip comments that contain specific details about implementation
	if strings.Contains(strings.ToLower(text), "according to") ||
		strings.Contains(strings.ToLower(text), "performs") ||
		strings.Contains(strings.ToLower(text), "validates") ||
		strings.Contains(strings.ToLower(text), "calculates") {
		return false
	}

	lowerText := strings.ToLower(text)
	lowerName := strings.ToLower(name)

	if strings.Contains(lowerText, lowerName) {
		// Check for patterns like "GetUser gets the user"
		prefix := strings.TrimSpace(strings.Split(lowerText, lowerName)[0])
		if prefix == "" || strings.HasSuffix(prefix, "method") || strings.HasSuffix(prefix, "function") {
			return true
		}

		// Check if the rest of the comment just describes obvious behavior
		parts := strings.Split(lowerText, lowerName)
		if len(parts) > 1 {
			rest := strings.TrimSpace(parts[1])
			if rest == "" || strings.HasPrefix(rest, "gets") || strings.HasPrefix(rest, "sets") {
				return true
			}
		}
	}

	// Check for common useless patterns
	commonPatterns := []string{
		"implements the",
		"extends the",
		"is a function that",
		"method that",
		"function that",
		"this function",
		"this method",
	}

	for _, pattern := range commonPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}

	return false
}

func isUselessTypeComment(text string, typ *ast.TypeSpec) bool {
	text = cleanComment(text)
	name := typ.Name.Name

	// Check if comment just repeats type name
	if strings.Contains(strings.ToLower(text), strings.ToLower(name)) {
		words := strings.Fields(text)
		if len(words) <= 3 { // Short comments that just mention the type name
			return true
		}
	}

	// Check for common useless patterns
	commonPatterns := []string{
		"represents a",
		"defines a",
		"is a type that",
		"is an interface that",
		"is a struct that",
	}

	for _, pattern := range commonPatterns {
		if strings.Contains(strings.ToLower(text), pattern) {
			return true
		}
	}

	return false
}

func isUselessMethodComment(text string, field *ast.Field) bool {
	text = cleanComment(text)

	if len(field.Names) == 0 {
		return false
	}

	name := field.Names[0].Name
	lowerText := strings.ToLower(text)
	lowerName := strings.ToLower(name)

	// Check if comment just repeats method name with common verbs
	if strings.Contains(lowerText, lowerName) {
		// Check for patterns like "GetData gets the data"
		prefix := strings.TrimSpace(strings.Split(lowerText, lowerName)[0])
		if prefix == "" || strings.HasSuffix(prefix, "method") {
			return true
		}

		// Check if the rest of the comment just describes obvious behavior
		parts := strings.Split(lowerText, lowerName)
		if len(parts) > 1 {
			rest := strings.TrimSpace(parts[1])
			if rest == "" || strings.HasPrefix(rest, "gets") || strings.HasPrefix(rest, "sets") {
				return true
			}
		}
	}

	// Check for common useless patterns
	commonPatterns := []string{
		"implements the",
		"extends the",
		"method that",
		"returns the",
		"gets the",
		"sets the",
		"this method",
	}

	for _, pattern := range commonPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}

	return false
}
