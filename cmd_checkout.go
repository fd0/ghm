package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var cmdCheckout = &cobra.Command{
	Use: "checkout URL",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("first (and only) argument must be the Pull Request URL")
		}

		pr, err := ParsePullRequestURL(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("query GitHub API for PR #%d\n", pr.ID)
		err = pr.Load()
		if err != nil {
			return err
		}

		if pr.Response.State != "open" {
			return fmt.Errorf("state %q of PR #%d is invalid", pr.Response.State, pr.ID)
		}

		fmt.Printf("add remote %v\n", pr.Remote())
		err = AddRemote(pr.Remote(), pr.RemoteURL())
		if err != nil {
			return err
		}

		_, err = Git("fetch", "--quiet", pr.Remote())
		if err != nil {
			return err
		}

		fmt.Printf("checkout tracking branch %v/%v\n", pr.Remote(), pr.Ref())
		_, err = Git("checkout", "-b", pr.Ref(), pr.Remote()+"/"+pr.Ref())
		return err
	},
}

func init() {
	cmdRoot.AddCommand(cmdCheckout)
}
