package fizzbuzz

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGetOneReturnOne(t *testing.T) {
	assert.Equal(t, "", fizzbuzz(1))
}