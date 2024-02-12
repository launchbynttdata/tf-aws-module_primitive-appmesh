package common

import (
	"testing"

	"github.com/nexient-llc/lcaf-component-terratest-common/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestAlwaysSucceeds", func(t *testing.T) {
		assert.Equal(t, "foo", "foo", "Should always be the same!")
		assert.NotEqual(t, "foo", "bar", "Should never be the same!")
	})

	// Implement your own tests below this line. See the README.md file
	// at https://github.com/nexient-llc/lcaf-component-terratest-common
	// for more details on writing tests.
}

func TestNonComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestAlwaysSucceeds", func(t *testing.T) {
		assert.Equal(t, "foo", "foo", "Should always be the same!")
		assert.NotEqual(t, "foo", "bar", "Should never be the same!")
	})

	// Implement your own tests below this line. See the README.md file
	// at https://github.com/nexient-llc/lcaf-component-terratest-common
	// for more details on writing tests.
}
