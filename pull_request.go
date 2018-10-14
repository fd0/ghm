package main

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// PullRequest collects all data needed to operate on a GitHub pull request.
type PullRequest struct {
	Owner, Repo string
	ID          uint
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
