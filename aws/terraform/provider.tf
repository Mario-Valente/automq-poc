terraform {
    required_version = ">= 1.4.0"

    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = "~> 6.0"
        }
    }
}

variable "aws_region" {
    type        = string
    default     = "us-east-1"
    description = "AWS region to operate in"
}

variable "aws_profile" {
    type        = string
    default     = ""
    description = "AWS CLI profile name (optional)"
}

provider "aws" {
    region                  = var.aws_region
    profile                 = var.aws_profile != "" ? var.aws_profile : null
}