module "lambda" {
  source = "github.com/pbs/terraform-aws-lambda-module?ref=1.3.0"

  handler  = "index.handler"
  filename = "./artifacts/handler.zip"
  runtime  = "python3.9"

  architectures = ["arm64"]

  add_vpc_config = true

  environment_vars = {
    "REDIS_PRIMARY_ENDPOINT" = module.redis.primary_endpoint_address
    "REDIS_READER_ENDPOINT"  = module.redis.reader_endpoint_address
  }

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}

module "redis" {
  source = "github.com/pbs/terraform-aws-elasticache-redis-standalone-module?ref=0.0.3"

  private_hosted_zone = var.private_hosted_zone

  parameter_group_version = "7"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}

module "sg_rule" {
  source = "../.."

  security_group_id = module.redis.sg_ids[0]

  description = "Allow Lambda ${module.lambda.name} to access Redis"

  port                     = 6379
  source_security_group_id = module.lambda.sg
}
