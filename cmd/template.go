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

	// if branch not specified
	if branch == "" {
		defBranch, err := GetDefaultBranch(username, reponame)
		if err != nil {
			panic(err)
		}
		branch = defBranch
	}

	// if branch is specified
	doesBranchExists := VerifyBranchName(username, reponame, branch)
	if !doesBranchExists {
		fmt.Printf("branch %v does not exist", branch)
		os.Exit(1)
	}

	fmt.Println("Branch specified/Default branch:", branch)
	fmt.Println("Username:", username)

	var destPath string

	// if destination not specified
	if destination == "." {
		destPath = reponame
	} else {
		destPath = destination
	}

	fmt.Println("Destination folder specified/ default destination:", destPath)

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
	err = downloadFile(zipPath, fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Downloading the compressed file...")

	// unzip the file
	unzipPath := tempDir + "/" + "unzipped"
	_, err = unzip(zipPath, unzipPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decompressing the file...")

	files, err := os.ReadDir(unzipPath)
	if err != nil {
		panic(err)
	}
	repoDir := files[0]

	fmt.Println("Moving template to the destination...")

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
