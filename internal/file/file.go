package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
	return strings.Split(string(dat), "\n")[0]
}
