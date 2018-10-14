package main

import (
	"os"

	"github.com/spf13/cobra"
)

var cmdRoot = cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
}

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
		os.Exit(1)
	}
}
