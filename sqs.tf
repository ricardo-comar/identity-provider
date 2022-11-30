
resource "aws_sqs_queue" "sqs_queue" {
  name = "identity-provider-sqs-employees"
  # delay_seconds             = 5
  max_message_size          = 20480
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

}
