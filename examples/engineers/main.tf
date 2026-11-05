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

#  data "engineering-bootcamp_engineers" "all" {
#    filter {
#      name   = "name"
#      values = ["Ryan"]
#    }
#  }
