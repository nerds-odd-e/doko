package exercises_qe_jt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(number int)string{
	if number % 15 == 0 {
		return "fizzbuzz"
	}

	if number % 3 == 0 {
		return "fizz"
	}

	if number % 5 == 0 {
		return "buzz"
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
	assert.Equal(t,"buzz",FizzBuzz(5))
}

func Test6ReturnFizz(t *testing.T){
	assert.Equal(t,"fizz",FizzBuzz(6))
}

func Test10ReturnBuzz(t *testing.T){
	assert.Equal(t,"buzz",FizzBuzz(10))
}

func Test15ReturnFizzBuzz(t *testing.T){
	assert.Equal(t,"fizzbuzz",FizzBuzz(15))
}