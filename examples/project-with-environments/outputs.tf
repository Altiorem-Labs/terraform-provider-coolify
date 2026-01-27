# Project Information
output "project_id" {
  description = "The ID of the created project"
  value       = coolify_project.webapp.id
}

output "project_name" {
  description = "The name of the created project"
  value       = coolify_project.webapp.name
}


output "staging_environment_id" {
  description = "The ID of the staging environment"
  value       = coolify_environment.staging.id
}

output "development_environment_id" {
  description = "The ID of the development environment"
  value       = coolify_environment.development.id
}

output "testing_environment_id" {
  description = "The ID of the testing environment"
  value       = coolify_environment.testing.id
}

# Summary of all environments
output "environments" {
  description = "List of all created environments with their details"
  value = {
    staging = {
      id   = coolify_environment.staging.id
      name = coolify_environment.staging.name
    }
    development = {
      id   = coolify_environment.development.id
      name = coolify_environment.development.name
    }
    testing = {
      id   = coolify_environment.testing.id
      name = coolify_environment.testing.name
    }
  }
}