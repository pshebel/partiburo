terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket = "partiburo-tfstate"
    key    = "dev/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = var.region
}
# data "aws_ami" "amazon_linux" {
#   most_recent = true
#   owners      = ["amazon"]

#   filter {
#     name   = "name"
#     values = ["amzn2-ami-hvm-*-x86_64-gp2"]
#   }

#   filter {
#     name   = "virtualization-type"
#     values = ["hvm"]
#   }
# }

locals {
  vpc_cidr = "10.1.0.0/16"
}

# network
module "vpc" {
  source = "../../modules/vpc"

  vpc_cidr = local.vpc_cidr

  name_prefix = "partiburo-${var.environment}-vpc"
  region      = var.region
  environment = var.environment
}

module "ses" {
  source = "../../modules/ses"

  name_prefix = "partiburo"
  region      = var.region
  environment = var.environment
}

# compute
module "monolith" {
  source = "../../modules/monolith"

  size          = "t3.small"
  name_prefix   = "partiburo-${var.environment}-monolith"
  # ami           = data.aws_ami.amazon_linux.id
  ami            = "ami-0fd3ac4abb734302a" # rhel 10 ami
  region        = var.region
  environment   = var.environment
  subnet_id     = module.vpc.public_subnets[0]
  vpc_id        = module.vpc.vpc_id
  profile_name = module.ses.instance_profile_name
}


# DNS
module "route53" {
  source = "../../modules/route53"
  base = "partiburo.com"
  www = "www.partiburo.com"
  public_ip = module.monolith.public_ip
}