package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers/configvalidator"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(configvalidator.Analyzer)
}
