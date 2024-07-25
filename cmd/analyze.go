package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var filePath string
var mode string

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a code file",
	Long:  `This command allows you to analyze a code file for syntax errors or perform an in-depth analysis using a fine-tuned LLM.`,
	Run: func(cmd *cobra.Command, args []string) {
		if filePath == "" {
			fmt.Println("Please provide a file path using the --file flag.")
			return
		}
		if mode != "syntax" && mode != "in-depth" {
			fmt.Println("Invalid mode. Please choose 'syntax' or 'in-depth' using the --mode flag.")
			return
		}

		// Perform analysis based on the mode
		switch mode {
		case "syntax":
			fmt.Println("Performing syntax analysis on file:", filePath)
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
		case "in-depth":
			fmt.Println("Performing in-depth analysis with fine-tuned LLM on file:", filePath)
			// Add in-depth analysis logic using fine-tuned LLM here
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to analyze")
	analyzeCmd.Flags().StringVarP(&mode, "mode", "m", "syntax", "Choose the analysis mode (syntax or in-depth)")
}
