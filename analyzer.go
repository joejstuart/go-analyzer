package analyzer

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "ecanalyzer",
		Doc:   "calculates complexity",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			_, ok := n.(*ast.ReturnStmt)
			if !ok {
				return true
			}

			intf, ok := n.(*ast.InterfaceType)
			if !ok {
				return true
			}

			pass.Reportf(intf.Pos(), "Interface returned from function %v", intf)
			return true

		})
	}
	return nil, nil
}
