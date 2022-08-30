package go_generics

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_AddInt(b *testing.B) {
	// generate random numbers for benchmarking
	numbers := make([]int, 1000000001)
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	for i := 0; i < b.N; i++ {
		numbers[i] = randomizer.Intn(1000) // random numbers between 0-1000
	}
	// regular function
	b.Run("Int_RegularFunction", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInt(numbers[i], numbers[i+1])
		}
	})
	// function with type assertion
	b.Run("Int_TypeAssertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInterface(numbers[i], numbers[i+1])
		}
	})
	// function with generics
	b.Run("Int_Generics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addNumbers(numbers[i], numbers[i+1])
		}
	})
	// function with reflection
	b.Run("Int_Reflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addReflection(numbers[i], numbers[i+1])
		}
	})
}

func Benchmark_AddFloat32(b *testing.B) {
	// generate random numbers for benchmarking
	numbers := make([]float32, 1000000001)
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	for i := 0; i < b.N; i++ {
		numbers[i] = float32(randomizer.Intn(1000)) // random numbers between 0-1000
	}
	// regular function
	b.Run("Float32_RegularFunction", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addFloat32(numbers[i], numbers[i+1])
		}
	})
	// function with type assertion
	b.Run("Float32_TypeAssertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInterface(numbers[i], numbers[i+1])
		}
	})
	// function with generics
	b.Run("Float32_Generics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addNumbers(numbers[i], numbers[i+1])
		}
	})
	// function with reflection
	b.Run("Float32_Reflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addReflection(numbers[i], numbers[i+1])
		}
	})
}

func Benchmark_AddInt64(b *testing.B) {
	// generate random numbers for benchmarking
	numbers := make([]int64, 1000000001)
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	for i := 0; i < b.N; i++ {
		numbers[i] = int64(randomizer.Intn(1000)) // random numbers between 0-1000
	}
	// regular function
	b.Run("Int64_RegularFunction", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInt64(numbers[i], numbers[i+1])
		}
	})
	// function with type assertion
	b.Run("Int64_TypeAssertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInterface(numbers[i], numbers[i+1])
		}
	})
	// function with generics
	b.Run("Int64_Generics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addNumbers(numbers[i], numbers[i+1])
		}
	})
	// function with reflection
	b.Run("Int64_Reflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addReflection(numbers[i], numbers[i+1])
		}
	})
}
