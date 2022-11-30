
resource "aws_api_gateway_rest_api" "idp_api" {
  name = "idp_api"
}

resource "aws_api_gateway_resource" "idp_resource" {
  path_part   = "employees"
  parent_id   = aws_api_gateway_rest_api.idp_api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.idp_api.id
}

resource "aws_api_gateway_method" "idp_method" {
  rest_api_id   = aws_api_gateway_rest_api.idp_api.id
  resource_id   = aws_api_gateway_resource.idp_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "idp_integration" {
  rest_api_id             = aws_api_gateway_rest_api.idp_api.id
  resource_id             = aws_api_gateway_resource.idp_resource.id
  http_method             = aws_api_gateway_method.idp_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.rest_api.invoke_arn
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rest_api.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.idp_api.execution_arn}/*/*/*"
}

resource "aws_api_gateway_deployment" "rest_api_deploy" {
  depends_on = [aws_api_gateway_integration.idp_integration]

  rest_api_id = aws_api_gateway_rest_api.idp_api.id
  stage_name  = "v1"

}

resource "aws_api_gateway_rest_api_policy" "rest_api_policy" {
  rest_api_id = aws_api_gateway_rest_api.idp_api.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": "*",
            "Action": "execute-api:Invoke",
            "Resource": [
                "${aws_api_gateway_rest_api.idp_api.execution_arn}/*"
            ]
        },
        {
            "Effect": "Deny",
            "Principal": "*",
            "Action": "execute-api:Invoke",
            "Resource": [
                "${aws_api_gateway_rest_api.idp_api.execution_arn}/*"
            ],
            "Condition": {
                "NotIpAddress": {
                    "aws:SourceIp": [
                      "3.135.189.93/32"
                    ]
                }
            }
        }
    ]
}
EOF

  #3.135.189.93 - Squid Sectools Ohio - Hom

}

output "url" {
  value = "${aws_api_gateway_deployment.rest_api_deploy.invoke_url}${aws_api_gateway_resource.idp_resource.path}"
}
