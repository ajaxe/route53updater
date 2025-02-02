module "main" {
  source = "github.com/ajaxe/terraform/modules/route53updater"
  #source = "../../../../../../../../../github-terraform/modules/route53updater" # relative path to local setup

  environment    = var.environment
  lambda_folder  = abspath("../../../../lambda/")
  pre_shared_key = var.pre_shared_key
  hosted_zone_id = var.hosted_zone_id
}
