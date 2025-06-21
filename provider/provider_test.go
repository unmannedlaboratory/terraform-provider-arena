package provider_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/unmannedlaboratory/terraform-provider-arena/provider"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	tfProvider := provider.New()
	err := tfProvider.InternalValidate()
	require.NoError(t, err)
}
