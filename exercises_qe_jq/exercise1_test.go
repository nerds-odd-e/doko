package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(number int) string{
	if (number == 5){
		return "buzz"
	}
	if (number != 3) {
		return fmt.Sprintf("%d", number);
	}
	return "fizz"
}

func Test3EqualsFizz(t *testing.T){
	assert.Equal(t, FizzBuzz(3),"fizz")
}

func Test1Equals1(t *testing.T){
	assert.Equal(t, FizzBuzz(1), "1")
}
func Test2Equals2(t *testing.T){
	assert.Equal(t, FizzBuzz(2), "2")
}
func Test5EqualsBuzz(t *testing.T){
	assert.Equal(t, FizzBuzz(5), "buzz")
}

