package main

import (
	"fmt"
)

var balance = 100

func main() {
	amountChan := make(chan int)
	tranChan := make(chan bool)

	for i:=0; i<20; i++ {
		go func(ii int) {
			amountChan <- ii

			if ii == 19 {
				fmt.Println("Should be quittin time")
				tranChan <- true
			}
		}(i)
	}
	breakPoint := false
	for {
		if breakPoint == true {
			break
		}
		select {
			case amount := <-amountChan:
				if balance > 0 && balance-amount >= 0 {
					balance -= amount
					fmt.Println("Transaction success for: ", amount, " balance: ", balance)
				} else {
					fmt.Println("Transaction failed: ", amount, " balance: ", balance)
				}

			case tr := <- tranChan:
				breakPoint = tr
				close(tranChan)
		}
	}
	fmt.Println("Final  balance is: ", balance)
}