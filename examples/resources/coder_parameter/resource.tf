provider "arena" {}

data "arena_parameter" "example" {
  name        = "Region"
  description = "Specify a region to place your workspace."
  mutable     = false
  type        = "string"
  default     = "asia-central1-a"
  option {
    value = "us-central1-a"
    name  = "US Central"
    icon  = "/icon/usa.svg"
  }
  option {
    value = "asia-central1-a"
    name  = "Asia"
    icon  = "/icon/asia.svg"
  }
}

data "arena_parameter" "ami" {
  name        = "Machine Image"
  description = <<-EOT
    # Provide the machine image
    See the [registry](https://container.registry.blah/namespace) for options.
    EOT
  option {
    value = "ami-xxxxxxxx"
    name  = "Ubuntu"
    icon  = "/icon/ubuntu.svg"
  }
}

data "arena_parameter" "is_public_instance" {
  name    = "Is public instance?"
  type    = "bool"
  icon    = "/icon/docker.svg"
  default = false
}

data "arena_parameter" "cores" {
  name    = "CPU Cores"
  type    = "number"
  icon    = "/icon/cpu.svg"
  default = 3
  order   = 10
}

data "arena_parameter" "disk_size" {
  name    = "Disk Size"
  type    = "number"
  default = "5"
  order   = 8
  validation {
    # This can apply to number.
    min       = 0
    max       = 10
    monotonic = "increasing"
  }
}

data "arena_parameter" "cat_lives" {
  name    = "Cat Lives"
  type    = "number"
  default = "9"
  validation {
    # This can apply to number.
    min       = 0
    max       = 10
    monotonic = "decreasing"
  }
}

data "arena_parameter" "fairy_tale" {
  name      = "Fairy Tale"
  type      = "string"
  mutable   = true
  default   = "Hansel and Gretel"
  ephemeral = true
}

data "arena_parameter" "users" {
  name         = "system_users"
  display_name = "System users"
  type         = "list(string)"
  default      = jsonencode(["root", "user1", "user2"])
}