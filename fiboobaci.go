package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// A function like this together with the non-local variables it references is
// known as a closure. In this case increment and the variable x form the closure.

// One way to use closure is by writing a function which returns another function
// which – when called – can generate a sequence of numbers. For example here's how we might generate all the even numbers:

// makeEvenGenerator returns a function which generates even numbers.
// Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	// f := fib()
	// // Function calls are evaluated left-to-right.
	// for i := 1; i < 90; i++ {
	// 	fmt.Println(i, f())
	// }

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

}
