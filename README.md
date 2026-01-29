# Terraform Provider for Coolify

An open source Terraform provider for managing [Coolify](https://coolify.io) resources, developed by [Altiorem Labs](https://github.com/Altiorem-Labs).

> ⚠️ **Warning**: This provider is currently under development and is **not ready for production use**.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Basic Usage](#basic-usage)
- [Available Resources](#available-resources)
- [Examples](#examples)
- [Contributing](#contributing)
- [Development Status](#development-status)
- [License](#license)
- [Support](#support)

## Overview

This Terraform provider enables you to manage your Coolify infrastructure as code. Coolify is a self-hosted alternative to services like Netlify and Vercel, providing an easy way to deploy and manage applications.

## Features

- ✅ Project management
- ✅ Environment management
- 🚧 More resources coming soon...

## Quick Start

### Installation

Add the following to your Terraform configuration:

```hcl
terraform {
  required_providers {
    coolify = {
      source = "altiorem/coolify"
      version = "~> 0.1.0"
    }
  }
}

provider "coolify" {
  # Configuration options
}
```

### Basic Usage

```hcl
# Create a new project
resource "coolify_project" "my_project" {
  name        = "my-awesome-project"
  description = "My awesome Coolify project"
}

# Create an environment within the project
resource "coolify_environment" "production" {
  project_id = coolify_project.my_project.id
  name       = "production"
}
```

## Available Resources

For a complete list of available resources and their configuration options, see [Available Resources](docs/resources.md).

## Examples

Check out the [examples](examples/) directory for more comprehensive usage examples:

- [Basic Setup](examples/basic-setup/) - Simple project and environment setup
- [Create Project](examples/create-project/) - Project creation with variables
- [Project with Environments](examples/project-with-environments/) - Multi-environment project setup

## Contributing

🚧 **Under Construction** - Contribution guidelines coming soon!

This section is currently being developed. We welcome contributions and will provide detailed guidelines soon.

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## Support

- 📚 Documentation: [Available Resources](docs/resources.md)
- 🐛 Issues: [GitHub Issues](https://github.com/Altiorem-Labs/terraform-provider-coolify/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/Altiorem-Labs/terraform-provider-coolify/discussions)
