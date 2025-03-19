package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	registry := analyzers.NewRegistry()
	multichecker.Main(registry.GetAll()...)
}
