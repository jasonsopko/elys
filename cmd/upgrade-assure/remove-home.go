package main

import (
	"log"
	"os/exec"
)

func removeHome(homePath string) {
	// Command and arguments
	args := []string{"-rf", homePath}

	// Execute the command
	if err := exec.Command("rm", args...).Run(); err != nil {
		log.Fatalf(ColorRed+"Command execution failed: %v", err)
	}

	// If execution reaches here, the command was successful
	log.Printf(ColorYellow+"removed home path %s successfully", homePath)
}
