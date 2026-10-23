package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// expandWd replaces "$wd" in a path with the current working directory.
func expandWd(input string) string {
	if !strings.Contains(input, "$wd") {
		return input
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Warning: could not get working directory:", err)
		return input
	}
	return strings.ReplaceAll(input, "$wd", wd)
}

func askForType() string {
	reader := bufio.NewReader(os.Stdin)
	var num int
	var err error
	for {
		fmt.Println("Please choose a number from these:")
		fmt.Println("1: Application")
		fmt.Println("2: Link")
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err == nil && (num == 0 || num == 1 || num == 2) {
			break
		}
	}
	switch num {
	case 1:
		return "Application"
	case 2:
		return "Link"
	case 0:
		os.Exit(0)
	}
	return ""
}

func deskFileLocation() string {
	reader := bufio.NewReader(os.Stdin)
	homeDir, err := os.UserHomeDir()
	location := filepath.Join(homeDir, ".local", "share", "applications") + "/"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Default location:", location)
	fmt.Print("Proceed with this location? [Y/N]: ")

	input, _ := reader.ReadString('\n')
	proceed := strings.ToLower(strings.TrimSpace(input))

	if proceed == "n" || proceed == "no" {
		for {
			fmt.Print("Enter the new location ($wd = current directory): ")
			input, _ = reader.ReadString('\n')
			location = expandWd(strings.TrimSpace(input))
			if location != "" {
				break
			}
			fmt.Println("Location cannot be empty.")
		}
	}

	if !strings.HasSuffix(location, "/") {
		location += "/"
	}
	return location
}

func createFile(appName string, execName string, iconName string,
	version string, comment string, categories string, location string) error {

	var b strings.Builder
	b.WriteString("[Desktop Entry]\n")
	b.WriteString("Type=Application\n")
	if version != "" {
		b.WriteString("Version=" + version + "\n")
	}
	b.WriteString("Name=" + appName + "\n")
	if comment != "" {
		b.WriteString("Comment=" + comment + "\n")
	}
	b.WriteString("Exec=" + execName + "\n")
	if iconName != "" {
		b.WriteString("Icon=" + iconName + "\n")
	}
	b.WriteString("Terminal=false\n")
	if categories != "" {
		b.WriteString("Categories=" + categories + "\n")
	}

	path := location + appName + ".desktop"
	return os.WriteFile(path, []byte(b.String()), 0644)
}
