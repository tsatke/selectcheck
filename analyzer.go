package selectcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer checks whether select statements have cases and reports if they
// don't.
var Analyzer = &analysis.Analyzer{
	Name: "selectcheck",
	Doc:  "check if select statements have cases",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{
		(*ast.SelectStmt)(nil),
	}, func(n ast.Node) {
		sel := n.(*ast.SelectStmt)
		if len(sel.Body.List) == 0 {
			pass.Reportf(sel.Pos(), "select statement has no cases")
		}
	})
	return nil, nil
}

func testFunction() {
	ch := make(chan struct{})
	select {
	default:
	}
	select {
	case <-ch:
	}
	select {}
}
