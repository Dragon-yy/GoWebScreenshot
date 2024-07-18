package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gowebscreenshot",
	Short: "GoWebScreenshot is a tool for taking website screenshots from a list of domains",
	Long:  `GoWebScreenshot is a tool that reads a list of website domains from a file and takes screenshots of those websites.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(screenshotCmd)
}
