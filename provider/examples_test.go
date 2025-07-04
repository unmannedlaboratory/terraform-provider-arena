package provider_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/unmannedlaboratory/terraform-provider-arena/provider"
)

func TestExamples(t *testing.T) {
	t.Parallel()

	t.Run("arena_parameter", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			Providers: map[string]*schema.Provider{
				"arena": provider.New(),
			},
			IsUnitTest: true,
			Steps: []resource.TestStep{{
				Config: mustReadFile(t, "../examples/resources/arena_parameter/resource.tf"),
			}},
		})
	})
}

func mustReadFile(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(content)
}
