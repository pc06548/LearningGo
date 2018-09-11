package main

import (
	"fmt"
)

type report struct {
	Name string
}
/*
func getReport(filename string) (rep report, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			// invalidate rep
			rep = nil
			// return the modified err and rep
		}
	}()
	panic("Report format not recognized.")
	// rest of the getReport function, which can try to out-of-bound-access a slice
	///...
}

func main() {
	_, err := getReport("hello")
	fmt.Println("err", err)
}

func df (err error) func(){
	return func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			// invalidate rep
			// rep = nil
			// return the modified err and rep
		}
	}
}*/

func getReport(filename string) (rep *report) {
	dfc := df(rep)
	defer dfc()
	panic("Report format not recognized.")
	// rest of the getReport function, which can try to out-of-bound-access a slice
	///...
}

func main() {
	err := getReport("hello")
	fmt.Println("err", err)
}

func df (err *report) func(){
	return func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			err = &report{Name:"tt"}
		}
	}
}