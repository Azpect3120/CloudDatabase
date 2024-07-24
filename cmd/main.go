package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Command to execute the shell script
	cmd := exec.Command("./scripts/createUser.sh", "testing", "root")

	// Create buffers to capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Print anything captured from stdout and stderr
	fmt.Println("=== Output from script ===")
	fmt.Println("STDOUT:")
	fmt.Println(stdout.String())
	fmt.Println("STDERR:")
	fmt.Println(stderr.String())

	// Get the exit code
	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			fmt.Printf("Failed to run command: %v\n", err)
			os.Exit(1)
		}
	}

	// Use the exit code in the Go program
	fmt.Printf("Script exited with code: %d\n", exitCode)
}
