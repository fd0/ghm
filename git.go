package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Git runs "git" with the given arguments.
func Git(args ...string) ([]byte, error) {
	if Options.Verbose {
		fmt.Printf("run %v %v\n", "git", args)
	}

	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

// GitQuiet runs "git" with the given arguments and surpresses output sent to stderr.
func GitQuiet(args ...string) ([]byte, error) {
	if Options.Verbose {
		fmt.Printf("run quiet %v %v\n", "git", args)
	}

	cmd := exec.Command("git", args...)
	return cmd.Output()
}

// AddRemote adds a new remote.
func AddRemote(name, url string) error {
	buf, err := GitQuiet("remote", "get-url", name)
	if err == nil {
		haveURL := strings.TrimSpace(string(buf))
		if haveURL != url {
			return fmt.Errorf("remote %v already exists, but has wrong URL %v", name, haveURL)
		}

		// remote already exists and has the correct URL, nothing to do
		return nil
	}

	_, err = Git("remote", "add", name, url)
	if err != nil {
		return err
	}

	return nil
}
