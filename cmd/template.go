package cmd

import (
	"fmt"
	"os"
	"strings"
)

func createTemplate(repoUrl, branch, initGit, destination string) {
	gitUrl := strings.Split(repoUrl, "/")

	// gitURL e.g. = https://github.com/1Shubham7/templ8
	username := gitUrl[3]
	reponame := gitUrl[4]

	fmt.Println("init branch:" + branch)
	fmt.Println("init username:" + username)

	// if branch not specified
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

	var destPath string

	// if destination not specified
	if destination == "." {
		destPath = reponame
	} else {
		destPath = destination
	}

	fmt.Println("Destpath: ", destPath)

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
	zipPath := tempDir + "/" + "file.zip"
	fmt.Println("Zippath is: ", zipPath)
	fmt.Println("Branch here is: ", branch)

	err = downloadFile(zipPath, fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Everything is fine")

	// unzip the file
	unzipPath := tempDir + "/" + "unzipped"
	_, err = unzip(zipPath, unzipPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Works fine till unzip")

	files, err := os.ReadDir(unzipPath)
	if err != nil {
		panic(err)
	}
	repoDir := files[0]

	// move the folder
	repoDirPath := unzipPath + "/" + repoDir.Name()
	// moves file in param 1 file to param 2
	err = os.Rename(repoDirPath, "./"+destPath)
	if err != nil {
		panic(err)
	}

	// remove the temp folder
	defer os.RemoveAll(tempDir)

	// Init git
	isGitInstalled := CheckIsGitInstalled()
	if isGitInstalled == false {
		fmt.Println("git not found. Repository will not be intialized automatically.")
	}

	if isGitInstalled && initGit == "yes" {
		InitGit(destPath)
	}
}
