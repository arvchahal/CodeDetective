package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func analyzeCode(filePath string) {
	fmt.Printf("Analyzing file %s\n", filePath)
	// Construct the path to the Python script
	scriptPath, err := filepath.Abs("../server/analyze_syntax.py")
	if err != nil {
		fmt.Printf("Error finding script path: %s\n", err)
		return
	}
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("Error finding file path: %s\n", err)
		return
	}

	cmd := exec.Command("python3", scriptPath, absFilePath)
	fmt.Printf("cmd: %s\n", cmd)
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
