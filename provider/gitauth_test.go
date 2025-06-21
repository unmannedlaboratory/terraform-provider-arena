package provider_test

import (
	"testing"

	"github.com/arena/terraform-provider-arena/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/require"
)

func TestGitAuth(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"arena": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "arena" {
			}
			data "arena_git_auth" "github" {
				id = "github"
			}
			`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.arena_git_auth.github"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				require.Equal(t, "github", attribs["id"])

				return nil
			},
		}},
	})
}
