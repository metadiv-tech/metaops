# MetaOps CLI Tool

MetaOps is a CLI tool that helps developers manage development operations for Metadiv Technology's backend framework and services.

## Installation

### Install via go install (recommended)

```bash
go install github.com/metadiv-tech/metaops@latest
```

### Build from source

```bash
git clone <repository-url>
cd metaops
go install .
```

## Usage

### Initialize a Metagin Project

The `metaops metagin init` command initializes a new Metagin project by creating the necessary configuration files.

```bash
metaops metagin init
```

This command will:

1. **Prompt for Docker image**: You'll be asked to provide the Docker image for your project. If no input is provided, it defaults to `metadiv/auto-agentic-v0:latest`.

2. **Create docker-compose.yml**: A Docker Compose configuration file with the following structure:
   ```yaml
   services:
     server:
       image: <your-specified-image>
       ports:
         - "3000:3000"
         - "5000:5000"
       restart: always
       env_file:
         - .env
       volumes:
         - ./app_folder/logs:/app/app_folder/logs
   ```

3. **Create .env file**: An environment configuration file with the following variables:
   ```env
   METAGIN_MODE=release
   METAGIN_HOST=0.0.0.0
   METAGIN_PORT=5000
   METAORM_HOST=
   METAORM_PORT=
   METAORM_USERNAME=
   METAORM_PASSWORD=
   METAORM_DATABASE=
   METAORM_ENCRYPT_KEY=
   ```

### File Overwrite Protection

If `docker-compose.yml` or `.env` files already exist in the current directory, the command will:
- Display a warning listing which files will be overwritten
- Prompt for confirmation before proceeding
- Allow you to cancel the operation if needed

### Example Usage

```bash
$ metaops metagin init
Please enter the Docker image (default: metadiv/auto-agentic-v0:latest): mycompany/custom-image:v1.0.0
✅ Metagin project initialized successfully!
📁 Created files:
  - docker-compose.yml
  - .env

🚀 Next steps:
  1. Configure your database settings in .env
  2. Run: docker-compose up -d
```

### Warning Example

```bash
$ metaops metagin init
⚠️  Warning: The following files will be overwritten:
  - docker-compose.yml
  - .env
Do you want to continue? [y/N]: y
Please enter the Docker image (default: metadiv/auto-agentic-v0:latest): 
✅ Metagin project initialized successfully!
...
```

## Commands

### `metaops metagin init`

Initialize a new Metagin project with Docker Compose and environment configuration.

**Usage:**
```bash
metaops metagin init
```

**Interactive prompts:**
- Docker image selection
- Confirmation for file overwriting (if files exist)

## Development

### Requirements

- Go 1.21 or later
- Dependencies managed via Go modules

### Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI framework

### Building

```bash
go install .
```

### Testing

```bash
go test ./...
```

## License

Copyright (c) 2025 Metadiv Technology Limited. All rights reserved. 