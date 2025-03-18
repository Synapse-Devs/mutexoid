package analyzers

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// registerEnglishCommentsAnalyzer registers the English comments analyzer
func (r *Registry) registerEnglishCommentsAnalyzer() {
	r.register(englishCommentsAnalyzer)
}

var englishCommentsAnalyzer = &analysis.Analyzer{
	Name: "englishcomments",
	Doc:  "Checks that all comments are in English",
	Run:  runEnglishComments,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func runEnglishComments(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, cg := range f.Comments {
			for _, comment := range cg.List {
				text := comment.Text

				// Skip empty comments
				if len(strings.TrimSpace(text)) == 0 {
					continue
				}

				// Skip generate directives
				if isGenerateDirective(text) {
					continue
				}

				// Skip linter directives
				if isLinterDirective(text) {
					continue
				}

				// Skip want directives
				if strings.Contains(text, "want") {
					continue
				}

				// Extract comment text
				text = extractCommentText(text)

				// Skip if empty after extraction
				if text == "" {
					continue
				}

				// Check if comment contains non-English characters
				if containsNonEnglish(text) {
					// Get the position of the next line after the comment
					pos := comment.End()
					nextPos := pass.Fset.File(pos).LineStart(pass.Fset.Position(pos).Line + 1)
					pass.Reportf(nextPos, "comment should be in English")
				}
			}
		}
	}

	return nil, nil
}

func extractCommentText(text string) string {
	// Remove comment markers and trim spaces
	if strings.HasPrefix(text, "/*") {
		text = strings.TrimPrefix(text, "/*")
		text = strings.TrimSuffix(text, "*/")
	} else if strings.HasPrefix(text, "//") {
		text = strings.TrimPrefix(text, "//")
	}

	// Remove want directive if present
	if idx := strings.Index(text, "// want"); idx != -1 {
		text = text[:idx]
	}

	return strings.TrimSpace(text)
}

func isGenerateDirective(text string) bool {
	return regexp.MustCompile(`^//go:generate`).MatchString(text)
}

func isLinterDirective(text string) bool {
	return regexp.MustCompile(`^//\s*(nolint|revive|golint)`).MatchString(text)
}

func containsNonEnglish(text string) bool {
	for _, r := range text {
		// Skip common symbols, numbers, and whitespace
		if r <= 0x7F || unicode.IsNumber(r) || unicode.IsPunct(r) || unicode.IsSpace(r) {
			continue
		}
		// If we find any character outside ASCII range that's not a number or punctuation,
		// we consider it non-English
		return true
	}
	return false
}
