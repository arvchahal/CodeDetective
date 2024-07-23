package main

import (
	"fmt"
	"os"
	"os/exec"
)

func analyzeCode(filepath string) {
	fmt.Printf("Analyzing file %s\n", filepath)
	cmd := exec.Command("python3", "./server/init.py", filepath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	fmt.Printf("analysis result:\n%s", string(output))

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path to analyze (supported types: .py, .jsx,.ts,.tsx, .cpp, .h)")
		os.Exit(1)
	}
	filePath := os.Args[1]
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not Found error")
		os.Exit(1)
	}
	analyzeCode(filePath)
}
