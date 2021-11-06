# GCP Project ID. You have to set this in secret.tfvars.
variable "project" {}

# Credential file for service accounts.
# First you have to create a service account with sufficient policy,
# then download the key and rename it to the name you specify here.
variable "credentials_file" {
  default = "key.json"
}

# GCP region designated for GCE. Default is set to apply the GCE free program quota.
variable "gce_zone" {
  default = "us-west1-b"
}

# GCP region designated for Cloud Function. You can use any region you like.
variable "cf_region" {
  default = "us-west1"
}
