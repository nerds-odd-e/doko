package exercises_qe_jt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(number int)string{
	return "1"
}

func Test1Return1(t *testing.T){
	assert.Equal(t,"1",FizzBuzz(1))
}
