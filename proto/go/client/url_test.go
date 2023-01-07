package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractOffset(t *testing.T) {

	cases := []struct {
		url            string
		expectedOffset int
		err            error
	}{
		{
			url:            "https://pokeapi.co/api/v2/ability/?offset=100&limit=50",
			expectedOffset: 100,
			err:            nil,
		},
		{
			url:            "https://pokeapi.co/api/v2/ability/?limit=50",
			expectedOffset: 0,
			err:            noOffsetErr,
		},
	}

	for _, tt := range cases {
		actual, err := extractOffset(tt.url)
		if tt.err == nil {
			require.NoError(t, err)
			assert.Equal(t, tt.expectedOffset, actual)
		} else {
			assert.Equal(t, tt.err, err)
			assert.Equal(t, actual, tt.expectedOffset)
		}
	}
}
