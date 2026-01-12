# 1. Create the IAM Role that the EC2 instance will assume
resource "aws_iam_role" "ec2_ses_role" {
  name = "ec2-ses-sender-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

# 2. Create the Policy specifically for Sending Email
# Note: ses:SendEmail covers both v1 and v2 SDKs
resource "aws_iam_policy" "ses_send_policy" {
  name        = "ses-send-email-policy"
  description = "Allows sending of emails via SES"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid      = "AllowSendingEmails"
        Effect   = "Allow"
        Action   = [
          "ses:SendEmail",
          "ses:SendRawEmail"
        ]
        # Resource = "*" 
        Resource = "arn:aws:ses:us-east-1:432883629663:identity/partiburo.com"
      }
    ]
  })
}

# 3. Attach the Policy to the Role
resource "aws_iam_role_policy_attachment" "attach_ses_policy" {
  role       = aws_iam_role.ec2_ses_role.name
  policy_arn = aws_iam_policy.ses_send_policy.arn
}

# 4. Create the Instance Profile (This is what you attach to the EC2)
resource "aws_iam_instance_profile" "ec2_ses_profile" {
  name = "ec2-ses-instance-profile"
  role = aws_iam_role.ec2_ses_role.name
}