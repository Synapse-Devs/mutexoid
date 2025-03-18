package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis"
)

// AnalyzerPlugin exports all analyzers as a golangci-lint plugin
type AnalyzerPlugin struct{}

// GetAnalyzers returns all available analyzers
func (*AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return analyzers.NewRegistry().GetAll()
}

// This must be defined and named 'New' for golangci-lint to load it
var New AnalyzerPlugin
