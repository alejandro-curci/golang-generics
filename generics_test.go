package go_generics

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type genericSuite struct {
	suite.Suite
}

func TestGenerics(t *testing.T) {
	suite.Run(t, new(genericSuite))
}

func (t *genericSuite) Test_ReverseFunc() {
	// random slices
	nums := []int{1, 2, 3, 4}
	words := []string{"hello", "my", "friend"}
	anything := []interface{}{"string", 45, true}

	// regular functions
	_ = reverseInt(nums)
	_ = reverseString(words)
	_ = reverseInterface(anything)

	// explicit instantiation
	_ = reverse[int](nums)
	_ = reverse[string](words)

	// type inference
	_ = reverse(nums)
	_ = reverse(words)
}

// interface{}/any --> type parameter vs function argument
func (t *genericSuite) Test_TypeParameters_VS_RegularArguments() {
	printInterface(12, "abc", struct{}{}) // different types
	// printAny(12, "abc", struct{}{})          // compile error

	printAny(12, 4, 8)              // same type --> T is int
	printAny("abc", "bbbb", "cccc") // same type --> T is string

}

func (t *genericSuite) Test_AddReflection() {
	t.Equal(10, addReflection(4, 6))
	t.Equal(int64(15), addReflection(int64(3), int64(12)))
	t.Equal(float32(20), addReflection(float32(7), float32(13)))
	t.Equal(4.6, addReflection(3.9, 0.7))
	t.Nil(addReflection(int64(5), int32(8)))
}

func (t *genericSuite) Test_TypeApproximation() {
	/*
		x, y := Point(4), Point(9)
		_ = Min(x, y) // without approximation, compile error
	*/
	a, b := Floint(4), Floint(9)
	_ = Min(a, b) // with approximation, OK
}
