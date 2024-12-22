package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func GetDefaultBranch(username, repo string) (string, error) {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + repo)
	if err != nil {
		return "", err
	}

	data,_ := io.ReadAll(response.Body)
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

func Logic(repoUrl, branch string) {
	gitUrl := strings.Split(repoUrl, "/")

	// gitURL e.g. = https://github.com/1Shubham7/templ8
	username := gitUrl[3]
	reponame := gitUrl[4]

	fmt.Println("init branch:" + branch)
	fmt.Println("init username:" + username)

	if branch == "" {
		defBranch, err := GetDefaultBranch(username, reponame)
		if err != nil {
			panic(err)
		}
		branch = defBranch
	}

	fmt.Println("New branch:", branch)
	fmt.Println("New username:", username)

	// if branch is specified
	doesBranchExists := VerifyBranchName(username, reponame, branch)
	if !doesBranchExists {
		fmt.Printf("branch %v does not exist", branch)
		os.Exit(1)
	}

	fmt.Println("final branch:", branch)
	fmt.Println("final username:", username)
}