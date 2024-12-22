package cmd

import (
	"encoding/json"
	"io"
	"net/http"
)

type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func GetDefaultBranch(username, repo string) (string, error) {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + repo)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	data, _ := io.ReadAll(response.Body)
	var repoDetails GitHubResponse
	err = json.Unmarshal(data, &repoDetails)
	if err != nil {
		return "", err
	}

	return repoDetails.DefaultBranch, nil
}

func VerifyBranchName(username, repo, branch string) bool {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + repo + "/branches/" + branch)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == http.StatusNotFound {
		return false
	}

	return true
}
