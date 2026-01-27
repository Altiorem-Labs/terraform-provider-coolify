# Define Terraform settings and required providers
terraform {
  required_providers {
    coolify = {
      source = "registry.terraform.io/Altiorem-Labs/coolify"
    }
  }
}

# Configure the Coolify provider with endpoint and token variables
provider "coolify" {
  endpoint = var.coolify_endpoint
  token    = var.coolify_token
}

# Create a project for a web application
resource "coolify_project" "webapp" {
  name        = "Web Application Project"
  description = "Project for managing web application environments"
}

# Create staging environment
resource "coolify_environments" "staging" {
  project_id = coolify_project.webapp.id
  name       = "Staging"

  depends_on = [coolify_project.webapp]
}

# Create development environment
resource "coolify_environments" "development" {
  project_id = coolify_project.webapp.id
  name       = "Development"

  depends_on = [coolify_project.webapp]
}

# Create testing environment
resource "coolify_environments" "testing" {
  project_id = coolify_project.webapp.id
  name       = "Testing"

  depends_on = [coolify_project.webapp]
}