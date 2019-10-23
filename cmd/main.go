package main

import (
	"github.com/TimSatke/selectcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(selectcheck.Analyzer)
}
