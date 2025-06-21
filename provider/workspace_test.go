package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"

	"github.com/arena/terraform-provider-arena/provider"
)

func TestWorkspace(t *testing.T) {
	t.Setenv("ARENA_WORKSPACE_OWNER", "owner123")
	t.Setenv("ARENA_WORKSPACE_OWNER_NAME", "Mr Owner")
	t.Setenv("ARENA_WORKSPACE_OWNER_EMAIL", "owner123@example.com")
	t.Setenv("ARENA_WORKSPACE_OWNER_SESSION_TOKEN", "abc123")
	t.Setenv("ARENA_WORKSPACE_TEMPLATE_ID", "templateID")
	t.Setenv("ARENA_WORKSPACE_TEMPLATE_NAME", "template123")
	t.Setenv("ARENA_WORKSPACE_TEMPLATE_VERSION", "v1.2.3")

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"arena": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "arena" {
				url = "https://example.com:8080"
			}
			data "arena_workspace" "me" {
			}`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.arena_workspace.me"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				value := attribs["transition"]
				require.NotNil(t, value)
				t.Log(value)
				require.Equal(t, "8080", attribs["access_port"])
				require.Equal(t, "owner123", attribs["owner"])
				require.Equal(t, "Mr Owner", attribs["owner_name"])
				require.Equal(t, "owner123@example.com", attribs["owner_email"])
				require.Equal(t, "abc123", attribs["owner_session_token"])
				require.Equal(t, "templateID", attribs["template_id"])
				require.Equal(t, "template123", attribs["template_name"])
				require.Equal(t, "v1.2.3", attribs["template_version"])
				return nil
			},
		}},
	})
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"arena": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "arena" {
				url = "https://example.com:8080"
			}
			data "arena_workspace" "me" {
			}`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.arena_workspace.me"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				value := attribs["transition"]
				require.NotNil(t, value)
				t.Log(value)
				require.Equal(t, "https://example.com:8080", attribs["access_url"])
				require.Equal(t, "owner123", attribs["owner"])
				require.Equal(t, "Mr Owner", attribs["owner_name"])
				require.Equal(t, "owner123@example.com", attribs["owner_email"])
				require.Equal(t, "templateID", attribs["template_id"])
				require.Equal(t, "template123", attribs["template_name"])
				require.Equal(t, "v1.2.3", attribs["template_version"])
				return nil
			},
		}},
	})
}
