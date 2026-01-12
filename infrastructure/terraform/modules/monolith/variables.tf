variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}
variable "environment" {
  description = "environment"
  type        = string
  default     = "dev"
}
variable "profile_name" {
  description = "instance profile name to attach to ec2"
  type        = string
}
variable "name_prefix" {
  description = "prefix for resources created for a project"
  type        = string
  default     = "partiburo"
}

variable "key_pair_name" {
  description = "project name"
  type        = string
  default     = "partiburo"
}

variable "ami" {
  description = "the id for the aws linux image"
  type        = string
}

variable "subnet_id" {
    description = "private subnet id for backend"
    type = string
}

variable "size" {
  description = "size of the ec2 instance used for the bastion"
  type        = string
  default = "t3.micro"
}

variable "vpc_id" {
  description = "id for the previously created vpc"
  type        = string
}
