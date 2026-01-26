# Define required providers
terraform {
  required_providers {
    coolify = {
        source = "registry.terraform.io/Altiorem-Labs/coolify"
    }
  }
}

# Configure the Coolify provider
provider "coolify" {
  # The API endpoint of your Coolify instance
  endpoint = "https://demo.coolify.io"
  # The API token for authentication
  token    = "tk_test_123456"
}