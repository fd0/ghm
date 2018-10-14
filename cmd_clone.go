package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdClone = &cobra.Command{
	Use: "clone URL",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Printf("clone\n")
		return nil
	},
}

func init() {
	cmdRoot.AddCommand(cmdClone)
}
