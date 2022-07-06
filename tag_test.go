package battlescribe_test

import (
	"testing"

	"github.com/myminicommission/go-battlescribe"
)

func TestLatestTag(t *testing.T) {
	minTagValue := "1.7.0"
	repo := "wh40k"
	data, err := battlescribe.GetFeedData()
	if err != nil {
		t.Error("could not get feed data", err)
		t.FailNow()
	}

	tagName, err := battlescribe.LatestTag(repo, data)
	if err != nil {
		t.Error("could not find tag", err)
		t.FailNow()
	}

	if tagName < minTagValue {
		t.Errorf("tag value was less than expected. Expected: > %s. Actual: %s", minTagValue, tagName)
		t.FailNow()
	}
}
