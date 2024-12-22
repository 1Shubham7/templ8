package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test GetDefaultBranch function by making a real API call
func TestGetDefaultBranch(t *testing.T) {
	assert := assert.New(t)

	// Replace these with actual valid GitHub repository and username
	username := "1shubham7"
	repo := "templ8"

	// Call GetDefaultBranch function
	defaultBranch, err := GetDefaultBranch(username, repo)

	// Assertions
	assert.Nil(err)
	assert.Equal("main", defaultBranch, "The default branch should be 'main'")
}

// Test VerifyBranchName function by making a real API call
func TestVerifyBranchName(t *testing.T) {
	assert := assert.New(t)

	// Replace these with actual valid GitHub repository and username
	username := "1shubham7"
	repo := "templ8"

	// Test case where the branch exists
	branch := "main"
	branchExists := VerifyBranchName(username, repo, branch)

	// Assertions
	assert.True(branchExists, "Branch 'main' should exist")

	// Test case where the branch does not exist
	branch = "nonexistent"
	branchExists = VerifyBranchName(username, repo, branch)

	// Assertions
	assert.False(branchExists, "Branch 'nonexistent' should not exist")
}
