package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OlegSadJktu/regexer/internal/engine"
	f "github.com/OlegSadJktu/regexer/internal/file"
)

func main() {

	var regarg string
	var file string
	flag.StringVar(&regarg, "regex", "", "regex string")
	flag.StringVar(&file, "file", "", "regex from file")
	flag.Parse()
	if len(regarg) == 0 && len(file) == 0 {
		fmt.Fprintf(os.Stderr, "Regex or file must not be empty\n")
		os.Exit(1)
	}

	var regexp string
	if len(regarg) != 0 {
		regexp = regarg
	}
	if len(file) != 0 {
    regexp = f.ReadFile(file)
	}
  sum := engine.Count(regexp)
  fmt.Println(sum.String())
}
