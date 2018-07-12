package main

import (
	"time"
	"fmt"
)

func main() {

	t1 := time.Now()
	t2 := time.Now()
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("is t1 after t2", t1.After(t2))
	fmt.Println("is t2 after t1", t2.After(t1))
	ticker := time.NewTicker(1 * time.Second)

	loop:
	for {
		select {
			case v := <- ticker.C :
				fmt.Println("ticker ticked ", v) // this keeps on ticking
				break loop
		}
	}

	timer := time.NewTimer(1*time.Second)

	loop1:
	for {
		select {
			case m := <- timer.C :
				fmt.Println("timer ticked ", m)
				timer.Reset(2 * time.Second) // must do this to get the next tick. otherwise timer is not continuous
				break loop1
		}
	}

	timerAfterFunc := time.AfterFunc(3 * time.Second, func(){ fmt.Println("I am executed after 3 secs")})
	select {
		case m := <- timerAfterFunc.C :
			fmt.Println("I am never ticked ", m)
		case <- time.After(5*time.Second): //It is equivalent to NewTimer(d).C.
	}

}