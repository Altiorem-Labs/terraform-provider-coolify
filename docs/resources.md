# Available Resources

This document provides an overview of the currently available resources in the Coolify Terraform provider.

## Resources Summary

| Resource Name | Description | Status |
|---------------|-------------|--------|
| [`coolify_project`](resources/project.md) | Manage Coolify projects | Available |
| [`coolify_environment`](resources/environment.md) | Manage Coolify environments within projects | Available |

## Resource Details

### coolify_project

Manages a Coolify project resource.

📖 **[View complete documentation](resources/project.md)**

**Arguments:**
- `name` (Required) - The name of the project
- `description` (Optional) - Description of the project

**Attributes:**
- `id` - Unique identifier (UUID) for the project

**Example Usage:**
```hcl
resource "coolify_project" "example" {
  name        = "my-awesome-project"
  description = "My awesome Coolify project"
}
```

### coolify_environment

Manages a Coolify environment within a project.

📖 **[View complete documentation](resources/environment.md)**

> **Note**: Due to Coolify API limitations, environments cannot be updated after creation.

**Arguments:**
- `project_id` (Required) - The ID of the project to create the environment in
- `name` (Required) - The name of the environment

**Attributes:**
- `id` - Unique identifier (UUID) for the environment

**Example Usage:**
```hcl
resource "coolify_environment" "production" {
  project_id = coolify_project.example.id
  name       = "production"
}
```

## Data Sources

> Data sources are not yet available in this provider but are planned for future releases.