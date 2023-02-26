package main

import (
	"fmt"
	"os"

	"github.com/neuralnorthwest/tpology/cmd"
)

func main() {
	if err := cmd.Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
