provider "aws" {
  region = var.region
}

# Security Group for UI Host
resource "aws_security_group" "monolith" {
  vpc_id      = var.vpc_id

  ingress {
    description = "http"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "https"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "ssh"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.name_prefix}-sg"
    Environment = var.environment
  }
}

resource "aws_instance" "monolith" {
  ami                    = var.ami
  instance_type          = var.size
  key_name               = var.key_pair_name
  vpc_security_group_ids = [aws_security_group.monolith.id]
  subnet_id              = var.subnet_id

  tags = {
    Name        = var.name_prefix
    Environment = var.environment
  }
}

resource "aws_ebs_volume" "monolith" {
  availability_zone = "us-east-1a"
  size              = 40

  tags = {
    Name = "${var.name_prefix}-ebs-volume"
  }
}

resource "aws_volume_attachment" "ebs_att" {
  device_name = "/dev/sdh"
  volume_id   = aws_ebs_volume.monolith.id
  instance_id = aws_instance.monolith.id
}

resource "aws_eip" "lb" {
  instance = aws_instance.monolith.id
  domain   = "vpc"
}