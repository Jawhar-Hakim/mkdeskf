package cmd

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
	"strconv"
)

func askForType() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please choose a number from these:")
	fmt.Println("1: Application Type")
	fmt.Println("2: Link")
	fmt.Println("3: Directory")
	fmt.Print("Your choice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err!=nil {
		fmt.Println("Please enter a number")
	}
	for{
		fmt.Println("Please choose a number from these:")
		fmt.Println("1: Application Type")
		fmt.Println("2: Link")
		fmt.Println("3: Directory")
		fmt.Println("0: Quit the app")
		fmt.Print("Your choice: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err = strconv.Atoi(input)
		if err != nil  {
			fmt.Println("Please enter one of these numbers")
		}
	}
	switch num {
	case 1:
		return "Application"
	case 2:
		return "Link"
	case 3:
		return "Directory"
	case 0:
		os.Exit(0)
		return 0
	default:
		fmt.Println("Choose one of these choices")
		return ""
	}
}

func askForInputs() (string,string,string,string,string,string,bool,error) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are your files in this directory? [Y/N]:")
		input, _ := reader.ReadString('\n')
		thisDirectory := strings.TrimSpace(input)
		thisDirectory = strings.ToLower(thisDirectory)
		var dir string = ""
		var err error
		if thisDirectory == "y" || thisDirectory == "yes" {
			dir, err = os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				return "","","","","","",false,err
			}
			fmt.Println("Current directory:", dir)
		} else{
			//get directory
			fmt.Print("Directory: ")
			input,_ = reader.ReadString('\n')
			dir = strings.TrimSpace(input)

		}

		
		// application name
		fmt.Print("application name: ")
		input, _ = reader.ReadString('\n')
		appName := strings.TrimSpace(input)

		// get file name
		fmt.Print("Executable file name: ")
		input, _ = reader.ReadString('\n')
		execName := strings.TrimSpace(input)
		if dir!="" {
			execName = dir+"/"+execName
		}

		// get icon name
		fmt.Print("Icon file name: ")
		input,_ = reader.ReadString('\n')
		iconName := strings.TrimSpace(input)
		if dir!=""{
			iconName = dir+"/"+iconName
		}

		//get version
		fmt.Print("Version: ")
		input,_ =reader.ReadString('\n')
		version := strings.TrimSpace(input)

		//get comment
		fmt.Print("Comment: ")
		input,_ = reader.ReadString('\n')
		comment := strings.TrimSpace(input)

		//get categories
		fmt.Print("Categories(seperated with a semicolon): ")
		input,_ = reader.ReadString('\n')
		categories := strings.TrimSpace(input)



		fmt.Println("Are these correct?")
		fmt.Println("Application name:", appName)
		fmt.Println("Executable:", execName)
		fmt.Println("Icon:", iconName)
		fmt.Println("Version:",version)
		fmt.Println("Comment:",comment)
		fmt.Println("Categories:",categories)
		fmt.Print("[Y/N]: ")
		input, _ = reader.ReadString('\n')
		correctRes := strings.TrimSpace(input)
		correctRes = strings.ToLower(correctRes)
		correct := correctRes == "y" || correctRes == "yes"
		return appName, execName, iconName, version, comment, categories, correct, nil
}

func deskFileLocation() string {
	reader := bufio.NewReader(os.Stdin)
	homeDir, err := os.UserHomeDir()
	location := filepath.Join(homeDir, ".local", "share", "applications") + "/"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("This is the default location:", location)
	fmt.Print("Proceed with this location? [Y/N]: ")
	
	input, _ := reader.ReadString('\n')
	proceed := strings.TrimSpace(input)
	proceed = strings.ToLower(proceed)
	
	if proceed == "n" || proceed == "no" {
		fmt.Print("Enter the new location: ")
		input, _ = reader.ReadString('\n')
		location = strings.TrimSpace(input)
	}
	
	
	for location == "" {
		fmt.Println("Empty location is not allowed. Please try again.")
		fmt.Print("Enter the new location: ")
		input, _ := reader.ReadString('\n')
		location = strings.TrimSpace(input)
	}

	if location[len(location)-1:] != "/" {
		location += "/"
	}
	
	return location
}

func createFile(appName string, execName string, iconName string,
	version string, comment string, categories string, location string)error{
	content := []byte("[Desktop Entry]\nType=Application\nVersion="+version+
	"\nName="+appName+"\nComment="+comment+"\nExec="+execName+"\nIcon="+iconName+
	"\nTerminal=false\nCategories:"+categories)
	file, err := os.Create(location+appName+".desktop")

  if err != nil {
  	return err
  }

  defer file.Close()
	err = os.WriteFile(location+appName+".desktop",content,0644)
	if err != nil {
		return err
	}
	return nil
}





















