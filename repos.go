package battlescribe

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ReposResponse struct {
	Name                string       `json:"name"`
	Description         string       `json:"description"`
	BattleScribeVersion string       `json:"battleScribeVersion"`
	RepositorySourceURL string       `json:"repositorySourceUrl"`
	WebsiteURL          string       `json:"websiteUrl"`
	GithubURL           string       `json:"githubUrl"`
	DiscordURL          string       `json:"discordUrl"`
	FeedURL             string       `json:"feedUrl"`
	TwitterURL          string       `json:"twitterUrl"`
	FacebookURL         string       `json:"facebookUrl"`
	Repositories        []Repository `json:"repositories"`
}

type Repository struct {
	Name                  string        `json:"name"`
	Description           string        `json:"description"`
	BattleScribeVersion   string        `json:"battleScribeVersion"`
	Version               string        `json:"version"`
	LastUpdated           string        `json:"lastUpdated"`
	LastUpdateDescription string        `json:"lastUpdateDescription"`
	IndexURL              string        `json:"indexUrl"`
	RepositoryURL         string        `json:"repositoryUrl"`
	GithubURL             string        `json:"githubUrl"`
	FeedURL               string        `json:"feedUrl"`
	BugTrackerURL         string        `json:"bugTrackerUrl"`
	ReportBugURL          string        `json:"reportBugUrl"`
	Archived              bool          `json:"archived"`
	RepositoryFiles       []interface{} `json:"repositoryFiles"`
}

const (
	reposURL = "http://battlescribedata.appspot.com/repos"
)

var (
	ErrReposCouldNotFetchRepos     = errors.New("could not fetch repos")
	ErrReposCouldNotDecodeResponse = errors.New("could not decode repos response")
)

// GetRepos calls the /repos api and returns the repositories in the response
func GetRepos() ([]Repository, error) {
	var reposResponse ReposResponse
	resp, err := http.Get(reposURL)
	if err != nil {
		return nil, ErrReposCouldNotFetchRepos
	}

	jsonDecoder := json.NewDecoder(resp.Body)
	err = jsonDecoder.Decode(&reposResponse)
	if err != nil {
		return nil, ErrReposCouldNotDecodeResponse
	}

	return reposResponse.Repositories, nil
}
