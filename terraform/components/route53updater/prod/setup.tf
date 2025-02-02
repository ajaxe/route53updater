
terraform {
  backend "s3" {
    bucket  = var.s3_backend_bucket
    key     = var.s3_backend_key
    region  = var.aws_region
    profile = var.aws_profile
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.84.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "3.2.3"
    }
    local = {
      source  = "hashicorp/local"
      version = "2.5.2"
    }
  }
}

provider "aws" {
  profile = var.aws_profile
  region  = var.aws_region
}
