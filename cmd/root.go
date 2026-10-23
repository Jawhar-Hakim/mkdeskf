package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkdeskf",
	Short: "Create .desktop files for Linux applications and links",
	Run: func(cmd *cobra.Command, args []string) {
		appType := askForType()
		location := deskFileLocation()

		switch appType {
		case "Application":
			var appName, execName, iconName, version, comment, categories string
			var correct bool
			var err error
			for {
				appName, execName, iconName, version, comment, categories, correct, err = appInput()
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
				if correct {
					break
				}
			}
			if err := createFile(appName, execName, iconName, version, comment, categories, location); err != nil {
				fmt.Println("Error writing file:", err)
			} else {
				fmt.Println("Done!")
			}

		case "Link":
			var name, icon, version, comment, url string
			var correct bool
			for {
				name, icon, version, comment, url, correct = linkInput()
				if correct {
					break
				}
			}
			if err := createFileLink(name, icon, version, comment, url, location); err != nil {
				fmt.Println("Error writing file:", err)
			} else {
				fmt.Println("Done!")
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
