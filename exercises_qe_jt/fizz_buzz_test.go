package exercises_qe_jt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(number int)string{
	if number == 1 {
		return "1"
	}
	return "2"
}

func Test1Return1(t *testing.T){
	assert.Equal(t,"1",FizzBuzz(1))
}

func Test2Return2(t *testing.T){
	assert.Equal(t,"2",FizzBuzz(2))
}