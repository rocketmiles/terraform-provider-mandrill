terraform {
  required_providers {
    mandrill = {
      source = "rocketmiles/mandrill"
    }
  }
  required_version = ">= 1.0.0"
}

provider "mandrill" {
  api_key = "API_KEY"
}

resource "mandrill_sending_domain" "test" {
  domain_name = "test.com"
}
