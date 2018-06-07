package main

import (
	"fmt"
)

type Shape interface {
	area() int
}

type twoDShape interface {
	Shape
	perimeter()
}

type Shape1 interface {
	area() int
}

func info(s Shape) {
	fmt.Println(s.area())
}
func info1(s Shape1) {
	fmt.Println(s.area())
}

func infoP(s twoDShape) {
	fmt.Println(s.area())
}

type Circle struct {

}

func (c Circle) area() int {
	fmt.Println("circile area")
	return 0
}


func (c Circle) perimeter() {
	fmt.Println("circile perimeter")
}

type Square struct {

}

func (c Square) area() int {
	fmt.Println("square area")
	return 0
}

func main() {

	s1 := make([]string, 10)
	m1 := make(map[string]int, 20)
	a1 := [10]int{}

	fmt.Println(s1)
	fmt.Println(m1)
	fmt.Println(a1)

	c := Circle{}
	s := Square{}

	info(c)
	info1(c)
	info(s)
	info1(s)

	infoP(c)



	type fullName struct {
		firstName string
		lastName string
	}

	type mix struct {
		field1 int
		string
		fullName
	}

	prashantChoudhari := fullName{"Prashant", "Choudhari"}
	mix1 := mix{12, "hey", prashantChoudhari}
	fmt.Println(mix1.string)
	fmt.Println(prashantChoudhari)
	fmt.Println(prashantChoudhari.firstName)

	type age int

	var myage age = 10

	fmt.Println(myage)


	var slice = make([][]string, 20)
	fmt.Println(cap(slice))
	fmt.Println(cap(slice[0]))

	name := "prashant"
	fmt.Println("Hello World!")
	fmt.Println("My name is:", name)
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My name is %T", name)

	fmt.Scan(&name)
	fmt.Println("This is your new name: ", name)

	sum := 0
	for i := 0; i < 1000; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FizzBuzz")
			sum += i
		case i%3 == 0:
			fmt.Println(i, "Fizz")
			sum += i
		case i%5 == 0:
			fmt.Println(i, "Buzz")
			sum += i

		}
	}
	fmt.Println(sum)

	funVariable := func(name string) {
		fmt.Println("I am annonymous")
	}
	funVariable("")
}
