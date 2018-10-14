package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParsePullRequest(t *testing.T) {
	var tests = []struct {
		urls []string
		want PullRequest
	}{
		{
			[]string{
				"https://github.com/restic/restic/pull/1234",
				"https://github.com/restic/restic/pull/1234/files",
				"https://github.com/restic/restic/pull/1234#issue-209345125",
			},
			PullRequest{
				Owner: "restic",
				Repo:  "restic",
				ID:    1234,
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			for _, url := range test.urls {
				t.Run("", func(t *testing.T) {
					pr, err := ParsePullRequestURL(url)
					if err != nil {
						t.Fatal(err)
					}

					if !cmp.Equal(test.want, pr) {
						t.Error(cmp.Diff(test.want, pr))
					}
				})
			}
		})
	}
}
