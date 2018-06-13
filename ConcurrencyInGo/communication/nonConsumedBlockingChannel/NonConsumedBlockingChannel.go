package main

/*
	This code results in dead lock*/


	func doNothingF(doNothing chan string) {
		//doNothing := make(chan string)
		doNothing <- "Hi"
		//fmt.Println(<- doNothing)

	}

func main() {
	doNothing := make(chan string)
	go doNothingF(doNothing)

}
