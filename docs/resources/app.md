---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "arena_app Resource - terraform-provider-arena"
subcategory: ""
description: |-
  Use this resource to define shortcuts to access applications in a workspace.
---

# arena_app (Resource)

Use this resource to define shortcuts to access applications in a workspace.

## Example Usage

```terraform
data "arena_workspace" "me" {}

resource "arena_agent" "dev" {
  os             = "linux"
  arch           = "amd64"
  dir            = "/workspace"
  startup_script = <<EOF
curl -fsSL https://code-server.dev/install.sh | sh
code-server --auth none --port 13337
EOF
}

resource "arena_app" "code-server" {
  agent_id     = arena_agent.dev.id
  slug         = "code-server"
  display_name = "VS Code"
  icon         = "${data.arena_workspace.me.access_url}/icon/code.svg"
  url          = "http://localhost:13337"
  share        = "owner"
  subdomain    = false
  healthcheck {
    url       = "http://localhost:13337/healthz"
    interval  = 5
    threshold = 6
  }
}

resource "arena_app" "vim" {
  agent_id     = arena_agent.dev.id
  slug         = "vim"
  display_name = "Vim"
  icon         = "${data.arena_workspace.me.access_url}/icon/vim.svg"
  command      = "vim"
}

resource "arena_app" "intellij" {
  agent_id     = arena_agent.dev.id
  icon         = "${data.arena_workspace.me.access_url}/icon/intellij.svg"
  slug         = "intellij"
  display_name = "JetBrains IntelliJ"
  command      = "projector run"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `agent_id` (String) The "id" property of a "arena_agent" resource to associate with.
- `slug` (String) A hostname-friendly name for the app. This is used in URLs to access the app. May contain alphanumerics and hyphens. Cannot start/end with a hyphen or contain two consecutive hyphens.

### Optional

- `command` (String) A command to run in a terminal opening this app. In the web, this will open in a new tab. In the CLI, this will SSH and execute the command. Either "command" or "url" may be specified, but not both.
- `display_name` (String) A display name to identify the app. Defaults to the slug.
- `external` (Boolean) Specifies whether "url" is opened on the client machine instead of proxied through the workspace.
- `healthcheck` (Block Set, Max: 1) HTTP health checking to determine the application readiness. (see [below for nested schema](#nestedblock--healthcheck))
- `icon` (String) A URL to an icon that will display in the dashboard. View built-in icons here: https://github.com/unmannedlaboratory/arena/tree/main/site/static/icon. Use a built-in icon with `data.arena_workspace.me.access_url + "/icon/<path>"`.
- `name` (String, Deprecated) A display name to identify the app.
- `order` (Number) The order determines the position of app in the UI presentation. The lowest order is shown first and apps with equal order are sorted by name (ascending order).
- `relative_path` (Boolean, Deprecated) Specifies whether the URL will be accessed via a relative path or wildcard. Use if wildcard routing is unavailable. Defaults to true.
- `share` (String) Determines the "level" which the application is shared at. Valid levels are "owner" (default), "authenticated" and "public". Level "owner" disables sharing on the app, so only the workspace owner can access it. Level "authenticated" shares the app with all authenticated users. Level "public" shares it with any user, including unauthenticated users. Permitted application sharing levels can be configured site-wide via a flag on `arena server` (Enterprise only).
- `subdomain` (Boolean) Determines whether the app will be accessed via it's own subdomain or whether it will be accessed via a path on Arena. If wildcards have not been setup by the administrator then apps with "subdomain" set to true will not be accessible. Defaults to false.
- `url` (String) An external url if "external=true" or a URL to be proxied to from inside the workspace. This should be of the form "http://localhost:PORT[/SUBPATH]". Either "command" or "url" may be specified, but not both.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--healthcheck"></a>
### Nested Schema for `healthcheck`

Required:

- `interval` (Number) Duration in seconds to wait between healthcheck requests.
- `threshold` (Number) Number of consecutive heathcheck failures before returning an unhealthy status.
- `url` (String) HTTP address used determine the application readiness. A successful health check is a HTTP response code less than 500 returned before healthcheck.interval seconds.
