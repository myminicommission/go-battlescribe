package battlescribe_test

import (
	"testing"

	"github.com/myminicommission/go-battlescribe"
)

func TestFeed(t *testing.T) {
	data, err := battlescribe.GetFeedData()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(data.Entry) == 0 {
		t.Error("data.Entry had no items...")
		t.FailNow()
	}

	for _, entry := range data.Entry {
		t.Log(entry.ID)
	}
}
