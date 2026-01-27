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
  value       = coolify_environments.staging.id
}

output "development_environment_id" {
  description = "The ID of the development environment"
  value       = coolify_environments.development.id
}

output "testing_environment_id" {
  description = "The ID of the testing environment"
  value       = coolify_environments.testing.id
}

# Summary of all environments
output "environments" {
  description = "List of all created environments with their details"
  value = {
    staging = {
      id   = coolify_environments.staging.id
      name = coolify_environments.staging.name
    }
    development = {
      id   = coolify_environments.development.id
      name = coolify_environments.development.name
    }
    testing = {
      id   = coolify_environments.testing.id
      name = coolify_environments.testing.name
    }
  }
}