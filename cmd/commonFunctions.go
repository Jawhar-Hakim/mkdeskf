package cmd

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

func askForVariable(variable string) (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ",variable)
	input, _ := reader.ReadString('\n')
	variable = strings.TrimSpace(input)
	return variable
}
