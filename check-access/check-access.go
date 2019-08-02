package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var administrators = [3]string{"Julien"}

var firstName string
var forbiddenMessage string
var welcomeMessage string

func checkAccess(firstName string, administrators [3]string) string {
	for _, adminName := range administrators {
		if strings.Compare(adminName, firstName) == 0 {
			welcomeMessage := "Welcome " + firstName + "!"
			return welcomeMessage
		}
	}

	forbiddenMessage := "Access forbidden for " + firstName + "!"
	return forbiddenMessage
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter firstname: ")
	firstName, _ := reader.ReadString('\n')
	// convert CRLF to LF
	firstName = strings.Replace(firstName, "\n", "", -1)
	message := checkAccess(firstName, administrators)
	fmt.Println(message)
}
