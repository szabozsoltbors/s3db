variable "lambda_name" {
  description = "The name of the lambda function"
  type        = string
}

variable "lambda_alias" {
  description = "The name of the lambda alias"
  type        = string
}

variable "artifact_path" {
  description = "The path to the lambda deployment package"
  type        = string
}

variable "s3_bucket_name" {
  description = "The name of the S3 bucket to store asset files uploaded by the Lambda function"
  type        = string
}
