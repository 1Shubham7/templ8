/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/1shubham7/templ8/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("BRanch Name: ", cmd.BranchName)
}
