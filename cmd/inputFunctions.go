package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func appInput() (appName, execName, iconName, version, comment, categories string, correct bool, err error) {
	reader := bufio.NewReader(os.Stdin)
	var dir string

	fmt.Print("Are your files in this directory? [Y/N]: ")
	input, _ := reader.ReadString('\n')
	inThisDir := strings.ToLower(strings.TrimSpace(input))

	if inThisDir == "y" || inThisDir == "yes" {
		dir, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println("Current directory:", dir)
	} else {
		fmt.Print("Directory ($wd = current directory): ")
		input, _ = reader.ReadString('\n')
		dir = expandWd(strings.TrimSpace(input))
	}

	execName = expandWd(askForVariable("Executable file name"))
	if dir != "" {
		execName = dir + "/" + execName
	}

	appName = askForVariable("Application Name")

	rawIcon := expandWd(askForVariable("Icon ($wd = current directory, optional)"))
	if rawIcon != "" && dir != "" && !filepath.IsAbs(rawIcon) {
		iconName = dir + "/" + rawIcon
	} else {
		iconName = rawIcon
	}

	version = askForVariable("Version (optional)")
	comment = askForVariable("Comment (optional)")
	categories = askForVariable("Categories (optional, e.g. Utility;)")

	fmt.Println("\nAre these correct?")
	fmt.Println("Application Name:", appName)
	fmt.Println("Executable:", execName)
	if iconName != "" {
		fmt.Println("Icon:", iconName)
	}
	if version != "" {
		fmt.Println("Version:", version)
	}
	if comment != "" {
		fmt.Println("Comment:", comment)
	}
	if categories != "" {
		fmt.Println("Categories:", categories)
	}

	fmt.Print("[Y/N]: ")
	input, _ = reader.ReadString('\n')
	res := strings.ToLower(strings.TrimSpace(input))
	correct = res == "y" || res == "yes"
	return
}
