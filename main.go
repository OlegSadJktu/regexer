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
		var err error
		regexp, err = f.ReadFile(file)
		if err != nil {
			fmt.Printf("invalid file: %v\n", err)
			return
		}
	}
	sum, err := engine.Count(regexp)
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
	} else {
		fmt.Println(sum.String())
	}
}
