package cmd

import (
	"fmt"
	"os"
	"strings"
)

// ensureScheme prepends "https://" to a URL if it has no scheme.
// If the site only serves HTTP, the browser will follow the redirect automatically.
func ensureScheme(url string) string {
	if url == "" {
		return url
	}
	lower := strings.ToLower(url)
	if strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://") {
		return url
	}
	return "https://" + url
}

func linkInput() (name, icon, version, comment, url string, correct bool) {
	name = askForVariable("Link Name")
	icon = askForVariable("Icon (optional)")
	version = askForVariable("Version (optional)")
	comment = askForVariable("Comment (optional)")
	url = ensureScheme(askForVariable("URL"))

	fmt.Println("\nAre these correct?")
	fmt.Println("Link Name:", name)
	if icon != "" {
		fmt.Println("Icon:", icon)
	}
	if version != "" {
		fmt.Println("Version:", version)
	}
	if comment != "" {
		fmt.Println("Comment:", comment)
	}
	fmt.Println("URL:", url)

	res := strings.ToLower(askForVariable("[Y/N]"))
	correct = res == "y" || res == "yes"
	return
}

func createFileLink(name, icon, version, comment, url, location string) error {
	var b strings.Builder
	b.WriteString("[Desktop Entry]\n")
	b.WriteString("Type=Application\n")
	if version != "" {
		b.WriteString("Version=" + version + "\n")
	}
	b.WriteString("Name=" + name + "\n")
	if comment != "" {
		b.WriteString("Comment=" + comment + "\n")
	}
	b.WriteString("Exec=xdg-open " + url + "\n")
	if icon != "" {
		b.WriteString("Icon=" + icon + "\n")
	}
	b.WriteString("Terminal=false\n")

	path := location + name + ".desktop"
	return os.WriteFile(path, []byte(b.String()), 0644)
}
