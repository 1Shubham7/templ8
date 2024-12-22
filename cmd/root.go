/*
Copyright © 2024 Shubham Singh <shubhammahar1306@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const logo = `
▀▀█▀▀ ▒█▀▀▀ ▒█▀▄▀█ ▒█▀▀█ ▒█░░░ ▒█▀▀▀█
░▒█░░ ▒█▀▀▀ ▒█▒█▒█ ▒█▄▄█ ▒█░░░ █░▄▄▄█░
░▒█░░ ▒█▄▄▄ ▒█░░▒█ ▒█░░░ ▒█▄▄█ █▄▄▄▄█░
`

var (
	branchName string
	dir string
)

// rootCmd represents the base command when called without any subcommands
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
		fmt.Println(logo)
		fmt.Println(cmd.Long)
		if len(args) == 0 {
			fmt.Println("GitHub repo is required")
			return
		}
		Logic(args[0], branchName)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	// Branch Flag
	rootCmd.Flags().StringVarP(&branchName, "branch", "b", "", "The name of GitHub repo branch you want to download")

	// Directory Flag
	rootCmd.Flags().StringVarP(&dir, "distination", "d", ".", "Where do you want to create your new project")
}
