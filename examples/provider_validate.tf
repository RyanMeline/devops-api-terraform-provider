 terraform {
   required_providers {
     engineering-bootcamp = {
       source = "liatr.io/terraform/engineering-bootcamp"
     }
   }
 }

 provider "engineering-bootcamp" {
  endpoint = "http://localhost:8080"
 }

data "engineering-bootcamp_engineers" "example" {}

#  data "engineering-bootcamp_engineers" "all" {
#    filter {
#      name   = "name"
#      values = ["Ryan"]
#    }
#  }
