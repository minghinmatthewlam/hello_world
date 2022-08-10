package neetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationString(t *testing.T) {
	s1, s2 := "ab", "eidbaooo"
	assert.True(t, checkInclusion(s1, s2))
}
