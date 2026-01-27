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

# Create a generic resource group for production services
resource "coolify_project" "production_infrastructure" {
  name        = "Production Infrastructure"
  description = "Core infrastructure resources managed via Terraform"
}