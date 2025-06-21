package provider_test

import (
	"runtime"
	"testing"

	"github.com/arena/terraform-provider-arena/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestProvisioner(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"arena": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "arena" {
			}
			data "arena_provisioner" "me" {
			}`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.arena_provisioner.me"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				require.Equal(t, runtime.GOOS, attribs["os"])
				require.Equal(t, runtime.GOARCH, attribs["arch"])
				return nil
			},
		}},
	})
}
