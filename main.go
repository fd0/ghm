package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdRoot = cobra.Command{}

// Options collects global settings for the program.
var Options struct {
	Verbose bool
}

func init() {
	cmdRoot.PersistentFlags().BoolVarP(&Options.Verbose, "verbose", "v", false, "be verbose")
}

func main() {
	err := cmdRoot.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
