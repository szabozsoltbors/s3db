resource "aws_s3_bucket" "store" {
  bucket = var.s3_bucket_name

  tags = {
    Name = var.s3_bucket_name
  }
}

resource "aws_s3_bucket_public_access_block" "frontend_access_block" {
  bucket = aws_s3_bucket.store.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_ownership_controls" "frontend_ownership" {
  bucket = aws_s3_bucket.store.id

  rule {
    object_ownership = "BucketOwnerEnforced"
  }
}
