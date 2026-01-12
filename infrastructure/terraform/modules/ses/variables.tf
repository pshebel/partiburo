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
variable "name_prefix" {
  description = "prefix for resources created for a project"
  type        = string
  default     = "partiburo"
}