package go_generics

import (
	"fmt"
	"reflect"
)

// 1) Predefined Constraints --> Signed, Unsigned, Integer, Float, Ordered, etc
type (
	Float interface {
		~float32 | ~float64
	}
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Ordered interface {
		Integer | Float | ~string
	}
	Integer interface {
		Signed | Unsigned
	}
)

// 2) Reverse func example --> one function per type VS one generic function
// reverseInt reverses a slice of int
func reverseInt(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for i, n := range nums {
		result[length-i-1] = n
	}
	return result
}

// reverseString reverses a slice of string
func reverseString(words []string) []string {
	length := len(words)
	result := make([]string, length)
	for i, w := range words {
		result[length-i-1] = w
	}
	return result
}

// reverseInterface reverses a slice of interface{}
func reverseInterface(anything []any) []any {
	length := len(anything)
	result := make([]any, length)
	for i, a := range anything {
		result[length-i-1] = a
	}
	return result
}

// reverse returns the reversed slice of T
func reverse[T any](list []T) []T {
	length := len(list)
	result := make([]T, length)
	for i, el := range list {
		result[length-i-1] = el
	}
	return result
}

// 3) Difference between "type parameters" and "regular function arguments"
// printInterface prints arguments of type interface{}
// "a", "b" and "c" can have different types during implementation
func printInterface(a, b, c any) {
	fmt.Println(a, b, c)
}

// printAny prints arguments of type T
// after instantiation, T becomes a concrete type
// "a", "b" and "c" have the same type T during implementation
func printAny[T any](a, b, c T) {
	fmt.Println(a, b, c)
}

// 4) Type Approximation --> Floint implements Number but Point doesn't
type (
	Point  int64   // int64 is the underlying type
	Floint float64 // float64 is the underlying type
	Number interface {
		int | int64 | float32 | ~float64
	}
)

func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// 5) Comparison among different solutions for "adding two numbers"
// 5.A - One function per concrete type
func addInt(m, n int) int {
	return m + n
}

func addInt64(m, n int64) int64 {
	return m + n
}

func addFloat32(m, n float32) float32 {
	return m + n
}

// 5.B - Solution with type assertions
// addInterface returns the sum of two numbers using type assertion.
// Type assertion converts an interface to the given type if possible
// remember the syntax --> "interface.(type)"
func addInterface(m, n any) any {
	switch m.(type) {
	case int:
		return m.(int) + n.(int)
	case int64:
		return m.(int64) + n.(int64)
	case float32:
		return m.(float32) + n.(float32)
	case float64:
		return m.(float64) + n.(float64)
	default:
		return nil
	}
}

// 5.C - Solution with reflection
// addReflection returns the sum of two numbers using reflection.
// it inspects each variable's type at runtime and sums them if possible
func addReflection(m, n any) any {
	mKind := reflect.TypeOf(m).Kind()
	nKind := reflect.TypeOf(n).Kind()
	if mKind != nKind {
		return nil
	}
	switch mKind {
	case reflect.Int, reflect.Int64:
		mNum := reflect.ValueOf(m).Int()
		nNum := reflect.ValueOf(n).Int()
		if mKind == reflect.Int {
			return int(mNum + nNum)
		}
		return mNum + nNum
	case reflect.Float32, reflect.Float64:
		mNum := reflect.ValueOf(m).Float()
		nNum := reflect.ValueOf(n).Float()
		if mKind == reflect.Float32 {
			return float32(mNum + nNum)
		}
		return mNum + nNum
	default:
		return nil
	}
}

// 5.D - Solution with generics
// addNumbers returns the sum of two numbers using generics.
func addNumbers[T Number](m, n T) T {
	return m + n
}

// 6) Exercise 1 -> "Write a generic function which determines whether a list of elements contains certain element"
func contains[T comparable](list []T, target T) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}
	return false
}

// Exercise 2 -> "Write a generic function which extracts a map's keys into a slice"
func keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return result
}

// NOTE -> [comparable] is the constraint for types that support equality operators == and !=
// it cannot be used with slices, maps and functions
