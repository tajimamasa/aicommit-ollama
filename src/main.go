package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	message := flag.String("ollamaHost", "http://localhost:11434/", "Ollama host")
	model := flag.String("ollamaModel", "llama3.2:latest", "Ollama model")
	flag.Parse()
	fmt.Println("Ollama Host: " + *message)
	fmt.Println("Ollama Model: " + *model)

	// Execute the command to get the current branch name
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, branchErr := branchCmd.CombinedOutput()
	if branchErr != nil {
		fmt.Println("Error executing git branch command:", branchErr)
		return
	}

	// Print the current branch name
	fmt.Println("Current Branch: " + string(branchOutput))

	// Execute the git diff command
	cmd := exec.Command("git", "diff", "--name-status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing git diff:", err)
		return
	}

	// Print the output of the git diff command
	fmt.Println(string(output))
}
