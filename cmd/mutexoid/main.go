package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	registry := analyzers.NewRegistry()
	singlechecker.Main(registry.Get("mutexoid"))
}
