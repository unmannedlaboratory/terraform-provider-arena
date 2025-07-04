data "arena_workspace" "me" {
}

resource "arena_agent" "dev" {
  os   = "linux"
  arch = "amd64"
  dir  = "/workspace"
  display_apps {
    vscode          = true
    vscode_insiders = false
    web_terminal    = true
    ssh_helper      = false
  }

  metadata {
    display_name = "CPU Usage"
    key          = "cpu_usage"
    script       = "arena stat cpu"
    interval     = 10
    timeout      = 1
    order        = 2
  }
  metadata {
    display_name = "RAM Usage"
    key          = "ram_usage"
    script       = "arena stat mem"
    interval     = 10
    timeout      = 1
    order        = 1
  }

  order = 1
}

resource "kubernetes_pod" "dev" {
  count = data.arena_workspace.me.start_count
  spec {
    container {
      command = ["sh", "-c", arena_agent.dev.init_script]
      env {
        name  = "ARENA_AGENT_TOKEN"
        value = arena_agent.dev.token
      }
    }
  }
}
