# s3db - infrastructure

#### Prerequisites

Before creating the infrastructure of the application the following steps needs to be done:

```sh
# Set your AWS PROFILE and the AWS ACCESS KEYS as environment variables
export AWS_PROFILE=aws_profile
export AWS_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"
export AWS_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
export AWS_SESSION_TOKEN="IQoJb3JpZ2luX2VjE..."   # Only for temporary credentials

# Create the S3 bucket (if missing)
aws s3api create-bucket \
  --bucket szs-terraform-s3db-state \
  --region eu-central-1 \
  --create-bucket-configuration LocationConstraint=eu-central-1

#### Create or update the infrastructure

Use the following commands to plan and apply terraform changes.

```sh
ENV=dev REGION=eu-central-1 make init
ENV=dev REGION=eu-central-1 make plan
ENV=dev REGION=eu-central-1 make apply
ENV=dev REGION=eu-central-1 make destroy
```
