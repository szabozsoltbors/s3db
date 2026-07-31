variable "profile" {
  description = "The AWS profile to use for authentication"
  type        = string
}

variable "region" {
  description = "The AWS region where the resources will be created"
  type        = string
}

variable "s3_bucket_name" {
  description = "The name of the S3 bucket to be created for storing assets"
  type        = string
}

variable "lambda_name" {
  description = "The name of the Lambda function to be created"
  type        = string
}

variable "stage_name" {
  description = "The name of the API Gateway stage to be created"
  type        = string
}
