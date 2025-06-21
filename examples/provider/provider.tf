terraform {
  required_providers {
    arena = {
      source = "unmanned/arena"
    }
  }
}

provider "google" {
  region = "us-central1"
}

data "arena_workspace" "me" {}

resource "arena_agent" "dev" {
  arch = "amd64"
  os   = "linux"
  auth = "google-instance-identity"
}

data "google_compute_default_service_account" "default" {}

resource "google_compute_instance" "dev" {
  zone         = "us-central1-a"
  count        = data.arena_workspace.me.start_count
  name         = "arena-${data.arena_workspace.me.owner}-${data.arena_workspace.me.name}"
  machine_type = "e2-medium"
  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  service_account {
    email  = data.google_compute_default_service_account.default.email
    scopes = ["cloud-platform"]
  }
  metadata_startup_script = arena_agent.dev.init_script
}
