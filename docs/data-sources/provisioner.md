---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "arena_provisioner Data Source - terraform-provider-arena"
subcategory: ""
description: |-
  Use this data source to get information about the Arena provisioner.
---

# arena_provisioner (Data Source)

Use this data source to get information about the Arena provisioner.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `arch` (String) The architecture of the host. This exposes `runtime.GOARCH` (see https://pkg.go.dev/runtime#pkg-constants).
- `id` (String) The ID of this resource.
- `os` (String) The operating system of the host. This exposes `runtime.GOOS` (see https://pkg.go.dev/runtime#pkg-constants).
