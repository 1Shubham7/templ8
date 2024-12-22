package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

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

	// create a temp folder
	tempDir := ".temp"
	err := os.Mkdir(tempDir, 0755)
	// Owner: 7 → 4 + 2 + 1 (read, write, and execute)
	// Group: 5 → 4 + 1 (read and execute)
	// Others: 5 → 4 + 1 (read and execute)
	if err != nil {
		panic(err)
	}

	// download acrhive repo
	fileUrl := "https://github.com/" + username + "/" + reponame + "/archive/refs/heads/" + branch + ".zip"
	zipPath := tempDir + "/" + "file.zip"}
	DownloadFile(zipPath, fileUrl)
}