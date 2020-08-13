terraform {
  required_providers {
    stratos = {
      versions = ["0.2"]
      source = "diehlabs.com/dev/stratos"
    }
  }
}

provider "stratos" {}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
