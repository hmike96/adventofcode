
package 

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get the file path from command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}
	filePath := os.Args[1]

	// Read the file contents
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Process the file contents
	// TODO: Add your code here

	// Print the file contents
	fmt.Println(string(fileContents))
}
