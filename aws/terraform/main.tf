module "automq-byoc" {
  source = "AutoMQ/automq-byoc-environment/aws"

  # Set the identifier for the environment to be installed. This ID will be used for naming internal resources. The environment ID supports only uppercase and lowercase English letters, numbers, and hyphens (-). It must start with a letter and is limited to a length of 32 characters.
  automq_byoc_env_id                       = "example"

  # Set the target regionId of aws
  cloud_provider_region                    = "us-east-1"

  # Optional: Add additional tags to all resources
  additional_tags = {
    Environment = "Production"
    Project     = "MyProject"
    Owner       = "TeamA"
    CostCenter  = "Engineering"
  }
}

# Necessary outputs
output "automq_byoc_env_id" {
  value = module.automq-byoc.automq_byoc_env_id
}

output "automq_byoc_endpoint" {
  value = module.automq-byoc.automq_byoc_endpoint
}

output "automq_byoc_initial_username" {
  value = module.automq-byoc.automq_byoc_initial_username
}

output "automq_byoc_initial_password" {
  value = module.automq-byoc.automq_byoc_initial_password
}

output "automq_byoc_vpc_id" {
  value = module.automq-byoc.automq_byoc_vpc_id
}

output "automq_byoc_instance_id" {
  value = module.automq-byoc.automq_byoc_instance_id
}
