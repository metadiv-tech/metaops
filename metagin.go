package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var metaginCmd = &cobra.Command{
	Use:   "metagin",
	Short: "Commands for Metagin backend framework",
	Long:  `Metagin is our backend framework. This command provides operations for managing Metagin projects.`,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Metagin project",
	Long:  `Initialize a new Metagin project by creating docker-compose.yml and .env files.`,
	RunE:  runInit,
}

func init() {
	metaginCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	// Check if files exist and warn user
	if fileExists("docker-compose.yml") || fileExists(".env") {
		fmt.Println("⚠️  Warning: The following files will be overwritten:")
		if fileExists("docker-compose.yml") {
			fmt.Println("  - docker-compose.yml")
		}
		if fileExists(".env") {
			fmt.Println("  - .env")
		}

		fmt.Print("Do you want to continue? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			fmt.Println("Operation cancelled.")
			return nil
		}
	}

	// Ask user for Docker image
	fmt.Print("Please enter the Docker image (default: metadiv/auto-agentic-v0:latest): ")
	reader := bufio.NewReader(os.Stdin)
	imageInput, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading image input: %w", err)
	}

	image := strings.TrimSpace(imageInput)
	if image == "" {
		image = "metadiv/auto-agentic-v0:latest"
	}

	// Create docker-compose.yml
	if err := createDockerCompose(image); err != nil {
		return fmt.Errorf("error creating docker-compose.yml: %w", err)
	}

	// Create .env file
	if err := createEnvFile(); err != nil {
		return fmt.Errorf("error creating .env file: %w", err)
	}

	fmt.Println("✅ Metagin project initialized successfully!")
	fmt.Println("📁 Created files:")
	fmt.Println("  - docker-compose.yml")
	fmt.Println("  - .env")
	fmt.Println("\n🚀 Next steps:")
	fmt.Println("  1. Configure your database settings in .env")
	fmt.Println("  2. Run: docker-compose up -d")

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func createDockerCompose(image string) error {
	content := fmt.Sprintf(`services:
  server:
    image: %s
    ports:
      - "3000:3000"
      - "5000:5000"
    restart: always
    env_file:
      - .env
    volumes:
      - ./app_folder/logs:/app/app_folder/logs
`, image)

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func createEnvFile() error {
	content := `METAGIN_MODE=release
METAGIN_HOST=0.0.0.0
METAGIN_PORT=5000
METAORM_HOST=
METAORM_PORT=
METAORM_USERNAME=
METAORM_PASSWORD=
METAORM_DATABASE=
METAORM_ENCRYPT_KEY=
`

	file, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
