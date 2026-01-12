output "instance_profile_name" {
  value       = aws_iam_instance_profile.ec2_ses_profile.name
  description = "Instance profile name to attach to EC2"
}
