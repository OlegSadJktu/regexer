package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var reg string
	var file string
	flag.StringVar(&reg, "regex", "", "regex string")
	flag.StringVar(&file, "file", "", "regex from file")
	flag.Parse()
	if len(reg) == 0 && len(file) == 0 {
		fmt.Fprintf(os.Stderr, "Regex or file must not be empty\n")
    os.Exit(1)
	}
	fmt.Println(reg, file)

}
