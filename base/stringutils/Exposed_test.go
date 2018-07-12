package stringutils

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 7, Sum(2,5))
	require.Equal(t, 7, Sum(2,5))

}
