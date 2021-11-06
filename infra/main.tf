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

resource "google_compute_instance" "instance_receiver" {
  name = "server-kun-receiver"
  machine_type = "e2-micro"

  zone = var.gce_zone

  network_interface {
    network = "default"

    access_config {
      
    }
  }

  tags = [ "server-kun" ]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }
}
