output "public_ip" {
    description = "monolith ip address"
    value = aws_eip.lb.public_ip
}