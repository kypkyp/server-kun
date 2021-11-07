terraform {
  required_version = "~> 1.0.10"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file(var.credentials_file)

  project = var.project
}

# receiver
resource "google_compute_instance" "instance_receiver" {
  name         = "server-kun-receiver"
  zone         = var.receiver_zone
  machine_type = "e2-micro"
  tags         = ["http-server", "https-server"]

  network_interface {
    network = "default"

    access_config {
    }
  }

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }
}

# Storage bucket cloud functions sources are stored.
resource "google_storage_bucket" "bucket_source" {
  name     = "server-kun-source"
  location = var.function_region
}

# Start function
module "cf_start" {
  source               = "./cloudfunctions"
  source_bucket_name   = google_storage_bucket.bucket_source.name
  function_region      = var.function_region
  function_name        = "start"
  description          = "A simple server starter by server-kun"
  entry_point          = "Start"
  target_instance_name = var.target_instance_name
  target_zone          = var.target_zone
  project              = var.project
}

# Stop function
module "cf_stop" {
  source               = "./cloudfunctions"
  source_bucket_name   = google_storage_bucket.bucket_source.name
  function_region      = var.function_region
  function_name        = "stop"
  description          = "A simple server stopper by server-kun"
  entry_point          = "Stop"
  target_instance_name = var.target_instance_name
  target_zone          = var.target_zone
  project              = var.project
}
