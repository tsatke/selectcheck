package selectcheck_test

import (
	"path/filepath"
	"testing"

	"github.com/TimSatke/selectcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalysis(t *testing.T) {
	dir, err := filepath.Abs("./testdata")
	if err != nil {
		panic(err)
	}
	analysistest.Run(t, dir, selectcheck.Analyzer, "./...")
}
