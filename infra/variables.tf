# GCP Project ID. You have to set this in terraform.tfvars.
variable "project" {}

# Credential file for service accounts.
# First you have to create a service account with sufficient policy,
# then download the key and rename it to the name you specify here.
variable "credentials_file" {
  default = "credentials/key.json"
}

# GCP region designated for Cloud Function. You can use any region you like.
variable "function_region" {
  default = "us-west1"
}

# Region designated for receiver instance. Default is set to apply the GCE free program quota.
variable "receiver_zone" {
  default = "us-west1-b"
}

# Region designated for target instance. You can use any region you like.
variable "target_zone" {}

# Zone of target server.
variable "target_instance_name" {}

variable "discord_token" {}

variable "discord_channel" {}

variable "start_message" {
  default = "サーバーが起動したよ！1分くらい待っててね！"
}

variable "stop_message" {
  default = "サーバーを停止したよ！これで安心して眠れるね！"
}
