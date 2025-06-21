provider "arena" {
}

data "arena_git_auth" "github" {
  # Matches the ID of the git auth provider in Arena.
  id = "github"
}

resource "arena_agent" "dev" {
  os   = "linux"
  arch = "amd64"
  dir  = "~/arena"
  env = {
    GITHUB_TOKEN : data.arena_git_auth.github.access_token
  }
  startup_script = <<EOF
if [ ! -d ~/arena ]; then
    git clone https://github.com/unmanned/arena
fi
EOF
}