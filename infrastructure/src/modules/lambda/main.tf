# Create the IAM role for Lambda execution
resource "aws_iam_role" "lambda_exec" {
  name = "${var.lambda_name}_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })
}

# Create a custom policy for S3 access and attach it to the Lambda role
resource "aws_iam_policy" "lambda_s3_access" {
  name = "${var.lambda_name}-s3-access-policy"

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "s3:ListBucket"
        ],
        "Resource" : [
          "arn:aws:s3:::${var.s3_bucket_name}"
        ]
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "s3:GetObject",
          "s3:DeleteObject",
          "s3:PutObject"
        ],
        "Resource" : [
          "arn:aws:s3:::${var.s3_bucket_name}/*"
        ]
      },
      {
        "Effect" : "Allow",
        "Action" : "logs:CreateLogGroup",
        "Resource" : "arn:aws:logs:eu-central-1:623110228806:*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" : [
          "arn:aws:logs:eu-central-1:623110228806:log-group:/aws/lambda/${var.lambda_name}:*"
        ]
      }
    ]
  })
}

# Attach the basic execution policy to the Lambda role
resource "aws_iam_role_policy_attachment" "basic_execution" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Attach the custom S3 access policy to the Lambda role
resource "aws_iam_role_policy_attachment" "lambda_s3_attach" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_s3_access.arn
}

# Create the Lambda function
resource "aws_lambda_function" "lambda" {
  function_name = var.lambda_name
  role          = aws_iam_role.lambda_exec.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  architectures = ["x86_64"]

  filename         = var.artifact_path
  source_code_hash = filebase64sha256(var.artifact_path)

  memory_size = 512
  timeout     = 30

  environment {
    variables = {
      S3_BUCKET_NAME = var.s3_bucket_name
    }
  }

  publish = true
}

resource "aws_lambda_alias" "lambda_alias" {
  name             = var.lambda_alias
  function_name    = aws_lambda_function.lambda.function_name
  function_version = aws_lambda_function.lambda.version
}
