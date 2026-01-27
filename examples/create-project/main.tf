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

resource "coolify_environment" "production_env" {
  project_id  = "yggg0okkcgcokc80o04o0cc4"
  name        = "Entorno Coolify"
}

resource "coolify_environment" "development_env" {
  project_id  = "yggg0okkcgcokc80o04o0cc4"
  name        = "Entorno Coolify Desarrollo"
}