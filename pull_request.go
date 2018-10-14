package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// PullRequest collects all data needed to operate on a GitHub pull request.
type PullRequest struct {
	Owner, Repo string
	ID          uint
	Response
}

// Response contains the information returned by the GitHub API about a pull request.
type Response struct {
	URL   string `json:"url"`
	State string `json:"state"`
	Title string `json:"title"`
	Head  struct {
		Ref  string `json:"ref"`
		Repo struct {
			SSHURL string `json:"ssh_url"`
			Owner  struct {
				Login string `json:"login"`
			} `json:"owner"`
		} `json:"repo"`
	} `json:"head"`
}

// ParsePullRequestURL extracts a PullRequest from a URL.
func ParsePullRequestURL(inputURL string) (PullRequest, error) {
	url, err := url.Parse(inputURL)
	if err != nil {
		return PullRequest{}, err
	}

	if url.Hostname() != "github.com" {
		return PullRequest{}, fmt.Errorf("invalid hostname %v", url.Hostname())
	}

	components := strings.Split(path.Clean(url.Path), "/")
	if len(components) < 4 {
		return PullRequest{}, errors.New("path has not enough components")
	}

	if components[3] != "pull" {
		return PullRequest{}, errors.New("URL does not point to pull request")
	}

	id, err := strconv.ParseUint(components[4], 10, 32)
	if err != nil {
		return PullRequest{}, fmt.Errorf("unable to parse pull request ID from string %v: %v", components[4], err)
	}

	pr := PullRequest{
		Owner: components[1],
		Repo:  components[2],
		ID:    uint(id),
	}

	return pr, nil
}

// Load calls out to the GitHub API and fills in the data for the pull request.
func (pr *PullRequest) Load() error {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls/%v", pr.Owner, pr.Repo, pr.ID))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %v (%v) received", res.StatusCode, res.Status)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = res.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &pr.Response)
	if err != nil {
		return err
	}

	return nil
}

// Remote returns the name for the remote repository.
func (pr PullRequest) Remote() string {
	return pr.Response.Head.Repo.Owner.Login
}

// RemoteURL returns the URL for the remote repository.
func (pr PullRequest) RemoteURL() string {
	return pr.Response.Head.Repo.SSHURL
}

// Ref returns the name of the reference in the remote repo.
func (pr PullRequest) Ref() string {
	return pr.Response.Head.Ref
}
