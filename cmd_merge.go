package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdMerge = &cobra.Command{
	Use: "merge URL",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Printf("merge\n")
		return nil
	},
}

func init() {
	cmdRoot.AddCommand(cmdMerge)
}
