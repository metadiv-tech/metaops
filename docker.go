package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Commands for Docker operations",
	Long:  `Docker commands provide operations for managing Docker installation and operations.`,
}

var dockerInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Install Docker on Linux",
	Long:  `Install Docker on Linux by executing the official Docker installation script for Ubuntu/Debian systems.`,
	RunE:  runDockerInit,
}

func init() {
	dockerCmd.AddCommand(dockerInitCmd)
}

func runDockerInit(cmd *cobra.Command, args []string) error {
	fmt.Println("🐳 Starting Docker installation on Linux...")
	fmt.Println("This will install Docker CE, Docker CLI, and Docker Compose")
	fmt.Println()

	// List of commands to execute
	commands := [][]string{
		{"sudo", "apt-get", "update"},
		{"sudo", "apt-get", "-y", "install", "ca-certificates", "curl", "gnupg", "lsb-release"},
		{"sudo", "mkdir", "-p", "/etc/apt/keyrings"},
		{"bash", "-c", "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg"},
		{"bash", "-c", `echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`},
		{"sudo", "apt-get", "update"},
		{"sudo", "apt-get", "-y", "install", "docker-ce", "docker-ce-cli", "containerd.io", "docker-compose-plugin"},
		{"sudo", "apt", "-y", "install", "docker-compose"},
	}

	// Execute each command
	for i, cmdArgs := range commands {
		fmt.Printf("📋 Step %d/%d: %s\n", i+1, len(commands), formatCommand(cmdArgs))

		var execCmd *exec.Cmd
		if cmdArgs[0] == "bash" && cmdArgs[1] == "-c" {
			// For bash commands with pipes and redirects
			execCmd = exec.Command("bash", "-c", cmdArgs[2])
		} else {
			execCmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
		}

		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("failed to execute command '%s': %w", formatCommand(cmdArgs), err)
		}
		fmt.Println("✅ Completed")
		fmt.Println()
	}

	fmt.Println("🎉 Docker installation completed successfully!")
	fmt.Println("\n🚀 Next steps:")
	fmt.Println("  1. Add your user to the docker group: sudo usermod -aG docker $USER")
	fmt.Println("  2. Log out and log back in (or run: newgrp docker)")
	fmt.Println("  3. Test the installation: docker --version")
	fmt.Println("  4. Test Docker Compose: docker-compose --version")

	return nil
}

func formatCommand(cmdArgs []string) string {
	if len(cmdArgs) == 0 {
		return ""
	}
	if cmdArgs[0] == "bash" && len(cmdArgs) >= 3 && cmdArgs[1] == "-c" {
		return cmdArgs[2]
	}
	result := cmdArgs[0]
	for _, arg := range cmdArgs[1:] {
		result += " " + arg
	}
	return result
}
