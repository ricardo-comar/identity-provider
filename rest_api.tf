data "archive_file" "lambda_rest_api_zip" {
  type        = "zip"
  source_file = "bin/source_rest_api"
  output_path = "bin/rest_api.zip"
}

// Function
resource "aws_lambda_function" "rest_api" {
  filename          = data.archive_file.lambda_rest_api_zip.output_path
  function_name     = "identity-provider-rest-api"
  description       = "REST API to expose database data"
  role              = aws_iam_role.lambda_role_rest_api.arn
  handler           = "source_rest_api"
  source_code_hash  = filebase64sha256(data.archive_file.lambda_rest_api_zip.output_path)
  runtime           = "go1.x"
  memory_size       = 1024
  timeout           = 30
  depends_on        = [aws_iam_role_policy_attachment.attach_iam_policy_to_iam_role_rest_api]

}


resource "aws_iam_role" "lambda_role_rest_api" {
name   = "lambda_role_rest_api"
assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "lambda.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}

resource "aws_iam_policy" "iam_policy_for_lambda_rest_api" {
 
 name         = "aws_iam_policy_for_terraform_aws_lambda_role_rest_api"
 path         = "/"
 description  = "AWS IAM Policy for managing aws lambda role"
 policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": [
       "logs:CreateLogGroup",
       "logs:CreateLogStream",
       "logs:PutLogEvents"
     ],
     "Resource": "arn:aws:logs:*:*:*",
     "Effect": "Allow"
   },
   {
     "Action": [
        "ec2:CreateNetworkInterface",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DeleteNetworkInterface"
     ],
     "Resource": "*",
     "Effect": "Allow"
   },
   {
     "Action": [
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage",
        "sqs:GetQueueAttributes"
     ],
     "Resource": "*",
     "Effect": "Allow"
   },
   {
     "Action": [
        "dynamodb:Query",
        "dynamodb:Scan"
     ],
     "Resource": "*",
     "Effect": "Allow"
   }
 ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_iam_policy_to_iam_role_rest_api" {
 role        = aws_iam_role.lambda_role_rest_api.name
 policy_arn  = aws_iam_policy.iam_policy_for_lambda_rest_api.arn
}
