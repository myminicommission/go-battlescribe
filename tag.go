package battlescribe

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrTagDataIsNil = errors.New("data is nil")
)

// LatestTag takes a repo name and a Feed pointer and extracts the highest value tag for the given repo
func LatestTag(repo string, data *Feed) (string, error) {
	if data == nil {
		return "", ErrTagDataIsNil
	}

	tag := "0"
	replaceStr := fmt.Sprintf("https://github.com/BSData/%s/releases/tag/", repo)

	for _, entry := range data.Entry {
		if strings.Contains(entry.ID, fmt.Sprintf("BSData/%s/", repo)) {
			t := strings.Replace(entry.ID, replaceStr, "", 1)
			if t > tag {
				tag = t
			}
		}
	}

	return tag, nil
}
