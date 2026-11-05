 terraform {
   required_providers {
     devops = {
       source = "liatr.io/terraform/engineering-bootcamp"
     }
   }
 }

 provider "devops" {
  endpoint = "http://localhost:8080"
 }

data "devops_engineers" "example" {}

resource "devops_engineer" "ryan" {
  name  = "Ryan"
  email = "ryan@meline.com"
}

#  data "devops_engineers" "all" {
#    filter {
#      name   = "name"
#      values = ["Ryan"]
#    }
#  }
