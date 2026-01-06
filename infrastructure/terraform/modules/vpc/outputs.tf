output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.main.id
}

output "public_subnets" {
    description = "public subnet"
    value = aws_subnet.public[*].id
}

output "private_subnets" {
    description = "private subnet"
    value = aws_subnet.private[*].id
}

output "private_rt" {
    description = "private route table"
    value = aws_route_table.private.id
}