package exercises_qe_jt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(number int)string{
	if number == 3 {
		return "fizz"
	}

	return fmt.Sprintf("%v",number)
}

func Test1Return1(t *testing.T){
	assert.Equal(t,"1",FizzBuzz(1))
}

func Test2Return2(t *testing.T){
	assert.Equal(t,"2",FizzBuzz(2))
}
func Test3ReturnFizz(t *testing.T){
	assert.Equal(t,"fizz",FizzBuzz(3))
}

func Test5ReturnBuzz(t *testing.T){
	assert.Equal(t,"5",FizzBuzz(5))
}