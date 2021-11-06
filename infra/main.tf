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

# start
data "archive_file" "archive_start" {
  type        = "zip"
  source_dir  = "../functions/start"
  output_path = "../functions/start.zip"
}

resource "google_storage_bucket" "bucket_source" {
  name     = "server-kun-source"
  location = var.cf_region
}

resource "google_storage_bucket_object" "object_source_start" {
  name   = "start.${data.archive_file.archive_start.output_md5}.zip"
  bucket = google_storage_bucket.bucket_source.name
  source = data.archive_file.archive_start.output_path
}

resource "google_cloudfunctions_function" "function_start" {
  name                  = "server-kun-start"
  description           = "Server-kun start"
  region                = var.cf_region
  runtime               = "go113"
  available_memory_mb   = 128
  timeout               = 30
  source_archive_bucket = google_storage_bucket.bucket_source.name
  source_archive_object = google_storage_bucket_object.object_source_start.name
  trigger_http          = true
  entry_point           = "Start"

  environment_variables = {
    SERVER_PROJECT  = var.project
    SERVER_ZONE     = var.target_zone
    SERVER_INSTANCE = var.target_instance_name
  }
}
