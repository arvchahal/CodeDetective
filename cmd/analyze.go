package cmd

import (
	"fmt"

	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze [file]",
	Short: "Analyze the syntax of a given file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		analyzeCode(args[0])
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}

func analyzeCode(filePath string) {
	fmt.Printf("Analyzing file %s\n", filePath)
	// Construct the path to the Python script
	scriptPath, err := filepath.Abs("server/analyze_syntax.py")
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
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("analysis result:\n%s", string(output))
}
