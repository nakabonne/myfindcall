// The findcall command runs the findcall analyzer.
package main

import (
	"github.com/nakabonne/myfindcall"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(myfindcall.Analyzer) }
