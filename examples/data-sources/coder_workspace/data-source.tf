data "arena_workspace" "dev" {
}

resource "kubernetes_pod" "dev" {
  count = data.arena_workspace.dev.transition == "start" ? 1 : 0
}
