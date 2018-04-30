package main

// https://medium.com/@gautamprajapati/writing-a-simple-e-commerce-website-api-in-go-programming-language-9f671324b4dd
// https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad650
import (
	"errors"
	"fmt"
	"math"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	// basic()
	// structAndPointers()
	// testconcurrency()
	workLoads()

}

func workLoads() {

	number := 50

	jobs := make(chan int, number)
	results := make(chan int, number)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < number; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < number; j++ {
		fmt.Println(j, <-results)
	}

}
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func testconcurrency() {
	// go count("sheep")
	// go count("fish")

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("Sheep")
	// 	wg.Done()
	// }()

	// // wg.Add(1)
	// go func() {
	// 	count("Fish")
	// 	wg.Done()
	// }()

	// wg.Wait()

	// c := make(chan string)
	// go count("sheep", c)

	// msg := <-c
	// fmt.Println(msg)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// c := make(chan string)
	// c <- "hello"

	// msg := <-c
	// fmt.Println(msg)
	//

	// "This above will not work as channel has not pre recievers"

	// c := make(chan string, 2) //buffer channel capacity
	// c <- "hello"
	// c <- "world"

	// msg := <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "1st"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "2nd"
			time.Sleep(time.Second * 2)
		}
	}()

	//here c2 is slow so the c1 is also slow beacuse of that
	//channels are syncronous so until c2 is not recieved it is not back to c1
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	//USING Select

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
func structAndPointers() {
	p := person{name: "Shubhanshu", age: 24}
	fmt.Println(p.name, p.age)

	i := 7
	fmt.Println(i, &i)
	increment(&i)
	fmt.Println(i, &i)

}
func basic() {
	var x int
	var y = 6
	var sum = x + y

	// x := 5 //only for new arguments
	// y := 4
	// sum := x + y

	x = 5
	y = 4
	sum = x + y

	fmt.Println("Hello Go World", sum)
	if sum > 10 {
		fmt.Println("Sum is greater than 10")
	} else if sum == 10 {
		fmt.Println("Sum is 10")
	} else {
		fmt.Println("Sum is less than 10")
	}

	array := []int{1, 2, 3, 4, 5}
	array[0] = 0
	array = append(array, 6)
	fmt.Println(array)

	dict := make(map[string]int)
	dict["1"] = 1
	dict["2"] = 2
	delete(dict, "3")

	fmt.Println(dict)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for index, value := range array {
	// 	fmt.Println("index", index, "value", summ(value, 10))
	// }

	// for key, value := range dict {
	// 	fmt.Println("key", key, "value", value)
	// }

	result, err := sqrt(-65)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Println(findMax(array))

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

}

func summ(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined")
	}
	return math.Sqrt(x), nil
}

func findMax(arr []int) (int, int) {
	max := 0
	indx := 0
	for index, value := range arr {
		if value > max {
			max = value
			indx = index
		}
	}
	return max, indx
}

func increment(x *int) {
	*x++
}
