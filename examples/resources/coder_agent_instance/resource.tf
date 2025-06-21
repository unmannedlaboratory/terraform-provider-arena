resource "arena_agent" "dev" {
  os   = "linux"
  arch = "amd64"
  auth = "google-instance-identity"
}

resource "google_compute_instance" "dev" {
  zone = "us-central1-a"
}

resource "arena_agent_instance" "dev" {
  agent_id    = arena_agent.dev.id
  instance_id = google_compute_instance.dev.instance_id
}
