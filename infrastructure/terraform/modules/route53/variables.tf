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
variable "public_ip" {
  description = "public ip address"
  type        = string
}

variable "base" {
  description = "url for base domain"
  type        = string
}

variable "www" {
  description = "url for www domain"
  type        = string
}