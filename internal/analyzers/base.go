package analyzers

import "golang.org/x/tools/go/analysis"

// Registry holds all available analyzers
type Registry struct {
	analyzers map[string]*analysis.Analyzer
}

// NewRegistry creates a new registry with all available analyzers
func NewRegistry() *Registry {
	r := &Registry{
		analyzers: make(map[string]*analysis.Analyzer),
	}

	// Register all analyzers
	r.registerMutexAnalyzer()
	r.registerEnglishCommentsAnalyzer()
	r.registerTestPackageAnalyzer()
	r.registerConfigValidatorAnalyzer()

	return r
}

// GetAll returns all registered analyzers
func (r *Registry) GetAll() []*analysis.Analyzer {
	var result []*analysis.Analyzer
	for _, a := range r.analyzers {
		result = append(result, a)
	}
	return result
}

// Get returns analyzer by name
func (r *Registry) Get(name string) *analysis.Analyzer {
	return r.analyzers[name]
}

// register adds an analyzer to the registry
func (r *Registry) register(a *analysis.Analyzer) {
	r.analyzers[a.Name] = a
}

// registerTestPackageAnalyzer registers the test package analyzer
func (r *Registry) registerTestPackageAnalyzer() {
	r.register(TestPackageAnalyzer)
}

// registerConfigValidatorAnalyzer registers the config validator analyzer
func (r *Registry) registerConfigValidatorAnalyzer() {
	r.register(ConfigValidatorAnalyzer)
}
