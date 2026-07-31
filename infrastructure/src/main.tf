# Create the bucket for the database document files
module "private_bucket" {
  source         = "./modules/s3"
  s3_bucket_name = var.s3_bucket_name
}

# Create the Lambda function to handle the backend API calls
module "private_lambda" {
  source         = "./modules/lambda"
  lambda_name    = var.lambda_name
  lambda_alias   = var.stage_name
  artifact_path  = abspath("${path.root}/../../worker/build/function.zip")
  s3_bucket_name = var.s3_bucket_name

  depends_on = [module.private_bucket]
}
