package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultBranch(t *testing.T) {
	assert := assert.New(t)

	username := "1shubham7"
	repo := "templ8"

	defaultBranch, err := GetDefaultBranch(username, repo)

	assert.Nil(err)
	assert.Equal("main", defaultBranch)
}

func TestVerifyBranchName(t *testing.T) {
	assert := assert.New(t)

	username := "1shubham7"
	repo := "templ8"

	branch := "main"
	branchExists := VerifyBranchName(username, repo, branch)

	assert.True(branchExists)

	branch = "nonexistent"
	branchExists = VerifyBranchName(username, repo, branch)

	assert.False(branchExists)
}
