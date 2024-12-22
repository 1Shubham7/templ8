package cmd

import (
	"os"
	"os/exec"
)

func CheckIsGitInstalled() bool {

	// Searches for the executable named git in the system's PATH environment.
	binPath, _ := exec.LookPath("git")
	if binPath != "" {
		return true
	} else {
		return false 
	}
}

func InitGit(path string) error {

	// 	Searches for the executable named git in the system's PATH environment.
	//  Returns the full path to the git binary (e.g., /usr/bin/git).
	gitBin, _ := exec.LookPath("git")

	// use the git cli to initialize a Git repo
	cmd := &exec.Cmd{
		Path:   gitBin,
		Args:   []string{gitBin, "init", path},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err := cmd.Run()
	return err
}