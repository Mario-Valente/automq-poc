.PHONY: help init plan apply destroy validate fmt clean output refresh import

# Variables
REGION ?= us-east-1
PROFILE ?=
ENV_ID ?= example

# Colors for output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
NC     := \033[0m # No Color

help: ## Show this help message
	@echo '$(GREEN)AutoMQ BYOC Terraform Makefile$(NC)'
	@echo ''
	@echo 'Usage:'
	@echo '  make $(YELLOW)<target>$(NC)'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

init: ## Initialize Terraform
	@echo "$(GREEN)Initializing Terraform...$(NC)"
	terraform init

validate: ## Validate Terraform configuration
	@echo "$(GREEN)Validating Terraform configuration...$(NC)"
	terraform validate

fmt: ## Format Terraform files
	@echo "$(GREEN)Formatting Terraform files...$(NC)"
	terraform fmt -recursive

plan: ## Show Terraform execution plan
	@echo "$(GREEN)Creating execution plan...$(NC)"
	@if [ -n "$(PROFILE)" ]; then \
		terraform plan -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)" -out=tfplan; \
	else \
		terraform plan -var="aws_region=$(REGION)" -out=tfplan; \
	fi

apply: ## Apply Terraform changes
	@echo "$(GREEN)Applying Terraform changes...$(NC)"
	@read -p "Are you sure you want to apply? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		if [ -n "$(PROFILE)" ]; then \
			terraform apply -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)"; \
		else \
			terraform apply -var="aws_region=$(REGION)"; \
		fi \
	else \
		echo "$(RED)Apply cancelled$(NC)"; \
	fi

apply-auto: ## Apply Terraform changes without confirmation
	@echo "$(GREEN)Auto-applying Terraform changes...$(NC)"
	@if [ -n "$(PROFILE)" ]; then \
		terraform apply -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)" -auto-approve; \
	else \
		terraform apply -var="aws_region=$(REGION)" -auto-approve; \
	fi

destroy: ## Destroy Terraform-managed infrastructure
	@echo "$(RED)Destroying infrastructure...$(NC)"
	@read -p "Are you ABSOLUTELY sure you want to destroy? Type 'destroy' to confirm: " confirm; \
	if [ "$$confirm" = "destroy" ]; then \
		if [ -n "$(PROFILE)" ]; then \
			terraform destroy -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)"; \
		else \
			terraform destroy -var="aws_region=$(REGION)"; \
		fi \
	else \
		echo "$(RED)Destroy cancelled$(NC)"; \
	fi

destroy-auto: ## Destroy infrastructure without confirmation (DANGEROUS!)
	@echo "$(RED)Auto-destroying infrastructure...$(NC)"
	@if [ -n "$(PROFILE)" ]; then \
		terraform destroy -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)" -auto-approve; \
	else \
		terraform destroy -var="aws_region=$(REGION)" -auto-approve; \
	fi

output: ## Show Terraform outputs
	@echo "$(GREEN)Terraform outputs:$(NC)"
	@terraform output

output-json: ## Show Terraform outputs in JSON format
	@terraform output -json

refresh: ## Refresh Terraform state
	@echo "$(GREEN)Refreshing Terraform state...$(NC)"
	@if [ -n "$(PROFILE)" ]; then \
		terraform refresh -var="aws_region=$(REGION)" -var="aws_profile=$(PROFILE)"; \
	else \
		terraform refresh -var="aws_region=$(REGION)"; \
	fi

state-list: ## List resources in Terraform state
	@terraform state list

state-show: ## Show detailed state of a resource (use RESOURCE=<name>)
	@terraform state show $(RESOURCE)

clean: ## Clean Terraform cache and plans
	@echo "$(YELLOW)Cleaning Terraform cache...$(NC)"
	rm -rf .terraform
	rm -f tfplan
	rm -f terraform.tfstate.backup
	rm -f .terraform.lock.hcl

graph: ## Generate dependency graph
	@terraform graph | dot -Tpng > graph.png
	@echo "$(GREEN)Graph saved to graph.png$(NC)"

upgrade: ## Upgrade provider versions
	@echo "$(GREEN)Upgrading providers...$(NC)"
	terraform init -upgrade

console: ## Open Terraform console
	@terraform console

import: ## Import existing resource (use RESOURCE=<type.name> ID=<aws-id>)
	@if [ -z "$(RESOURCE)" ] || [ -z "$(ID)" ]; then \
		echo "$(RED)Error: RESOURCE and ID required$(NC)"; \
		echo "Usage: make import RESOURCE=aws_instance.example ID=i-1234567890abcdef0"; \
		exit 1; \
	fi
	@terraform import $(RESOURCE) $(ID)

# Quick deployment workflow
setup: init validate fmt ## Initialize, validate and format
	@echo "$(GREEN)Setup complete!$(NC)"

deploy: setup plan apply ## Complete deployment workflow
	@echo "$(GREEN)Deployment complete!$(NC)"

# Show AutoMQ credentials
show-credentials: ## Display AutoMQ BYOC credentials
	@echo "$(GREEN)AutoMQ BYOC Credentials:$(NC)"
	@echo "Environment ID: $$(terraform output -raw automq_byoc_env_id)"
	@echo "Endpoint:       $$(terraform output -raw automq_byoc_endpoint)"
	@echo "Username:       $$(terraform output -raw automq_byoc_initial_username)"
	@echo "Password:       $$(terraform output -raw automq_byoc_initial_password)"
	@echo "VPC ID:         $$(terraform output -raw automq_byoc_vpc_id)"
	@echo "Instance ID:    $$(terraform output -raw automq_byoc_instance_id)"
