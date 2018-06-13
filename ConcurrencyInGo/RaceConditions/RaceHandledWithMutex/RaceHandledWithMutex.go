package main

import (
	"math/rand"
	"fmt"
	"sync"
)

var balance = 100
var mutex = sync.Mutex{}

func main() {
	wt := sync.WaitGroup{}
	wt.Add(20)
	for i:=0; i<20; i++ {
		withDrawAmount := rand.Intn(25)
		go withDraw(withDrawAmount, &wt)
	}
	wt.Wait()
	fmt.Println("Final  balance is: ", balance)
}

func withDraw(amount int, wt *sync.WaitGroup) {
	mutex.Lock()
	if balance > 0 && balance - amount > 0 {
		balance -= amount
		fmt.Println("Transaction success for: ", amount, " balance: ",balance)
	} else {
		fmt.Println("Transaction failed: ", amount, " balance: ",balance)
	}
	mutex.Unlock()
	wt.Done()

}