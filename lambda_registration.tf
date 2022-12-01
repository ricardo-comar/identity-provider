data "archive_file" "lambda_registration_zip" {
  type        = "zip"
  source_file = "bin/lambda_registration"
  output_path = "bin/registration.zip"
}

// Function
resource "aws_lambda_function" "registration" {
  filename         = data.archive_file.lambda_registration_zip.output_path
  function_name    = "identity-provider-registration"
  description      = "Employee Registration Lambda"
  role             = aws_iam_role.lambda_role_registration.arn
  handler          = "lambda_registration"
  source_code_hash = filebase64sha256(data.archive_file.lambda_registration_zip.output_path)
  runtime          = "go1.x"
  memory_size      = 1024
  timeout          = 30
  depends_on       = [aws_iam_role_policy_attachment.attach_iam_policy_to_iam_role_registration]

  environment {
    variables = {
      EMPLOYEE_QUEUE = aws_sqs_queue.sqs_queue.url
    }
  }
}



resource "aws_iam_role" "lambda_role_registration" {
  name               = "lambda_role_registration"
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

resource "aws_iam_policy" "iam_policy_for_lambda_registration" {

  name        = "aws_iam_policy_for_terraform_aws_lambda_role_registration"
  path        = "/"
  description = "AWS IAM Policy for managing aws lambda role"
  policy      = <<EOF
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
        "lambda:InvokeFunction"
     ],
     "Resource": "*",
     "Effect": "Allow"
   },
   {
     "Action": [
        "sqs:SendMessage",
        "sqs:GetQueueAttributes"
     ],
     "Resource": "*",
     "Effect": "Allow"
   }
 ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_iam_policy_to_iam_role_registration" {
  role       = aws_iam_role.lambda_role_registration.name
  policy_arn = aws_iam_policy.iam_policy_for_lambda_registration.arn
}

resource "aws_cloudwatch_event_rule" "registration_lambda_event_rule" {
  name                = "identity-provider-lambda-event-rule"
  description         = "retry scheduled every 2 min"
  schedule_expression = "rate(2 minutes)"
}

resource "aws_cloudwatch_event_target" "registration_lambda_target" {
  arn       = aws_lambda_function.registration.arn
  target_id = aws_lambda_function.registration.id
  rule      = aws_cloudwatch_event_rule.registration_lambda_event_rule.name
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_registration_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.registration.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.registration_lambda_event_rule.arn
}
