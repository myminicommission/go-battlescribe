package battlescribe_test

import (
	"testing"

	"github.com/myminicommission/go-battlescribe"
)

func TestGetRepos(t *testing.T) {
	repos, err := battlescribe.GetRepos()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(repos) == 0 {
		t.Error("repos slice is empty")
		t.FailNow()
	}
}
