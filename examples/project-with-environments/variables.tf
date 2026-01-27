# Coolify Instance Configuration
variable "coolify_endpoint" {
  description = "The URL of your Coolify instance (e.g., https://app.coolify.io)"
  type        = string
}

# Authentication Configuration
variable "coolify_token" {
  description = "Your Coolify API Token for authentication"
  type        = string
  sensitive   = true
}