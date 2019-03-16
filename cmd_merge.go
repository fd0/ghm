package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdMerge = &cobra.Command{
	Use: "merge URL",
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

		msg := fmt.Sprintf("Merge pull request #%v from %v/%v\n\n%v",
			pr.ID, pr.Remote(), pr.Ref(), pr.Title)

		fmt.Printf("merge %v/%v\n", pr.Remote(), pr.Ref())
		buf, err := Git("merge", "--no-ff", "--message", msg, pr.Remote()+"/"+pr.Ref())

		if err != nil {
			fmt.Fprintf(os.Stderr, "merge failed, please fix:\n%s\n", buf)
			return err
		}

		return nil
	},
}

func init() {
	cmdRoot.AddCommand(cmdMerge)
}
