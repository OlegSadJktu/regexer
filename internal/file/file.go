package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) (string, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return "", err

	}
	fmt.Println(string(dat))
	return strings.Split(string(dat), "\n")[0], nil
}
