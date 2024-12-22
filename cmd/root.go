/*
Copyright © 2024 Shubham Singh <shubhammahar1306@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const logo = `
▀▀█▀▀ ▒█▀▀▀ ▒█▀▄▀█ ▒█▀▀█ ▒█░░░ ▒█▀▀▀█
░▒█░░ ▒█▀▀▀ ▒█▒█▒█ ▒█▄▄█ ▒█░░░ █░▄▄▄█░
░▒█░░ ▒█▄▄▄ ▒█░░▒█ ▒█░░░ ▒█▄▄█ █▄▄▄▄█░
`

var (
	branchName string
	dir        string
)

var rootCmd = &cobra.Command{
	Use:   "templ8",
	Short: "A brief description of your application",
	Long: `🚀 Transform any GitHub repository into your project template

templ8 simplifies project bootstrapping by creating templates from existing GitHub repos.
Just provide a repository URL and branch name - templ8 handles the rest.

Usage:
  templ8 <github-repo-url> -b=<branch-name> -d=<distination>

Happy coding! ✨
`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logo)
		fmt.Println(cmd.Long)
		if len(args) == 0 {
			fmt.Println("GitHub repo url is required")
			return
		}
		input := takeInput()
		if input == "err" {
			fmt.Println("sorry, I couldn't understand, try again with (yes/no)")
			os.Exit(1)
		}
		CreateTemplate(args[0], branchName, input, dir)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func takeInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to init git? (yes/no): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "yes" || input == "y" || input == "ye" || input == "yess" {
		return "yes"
	} else if input == "no" || input == "n" || input == "noo" {
		return "no"
	} else {
		return "err"
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Branch Flag
	rootCmd.Flags().StringVarP(&branchName, "branch", "b", "", "The name of GitHub repo branch you want to download")

	// Directory Flag
	rootCmd.Flags().StringVarP(&dir, "distination", "d", ".", "Where do you want to create your new project")
}
