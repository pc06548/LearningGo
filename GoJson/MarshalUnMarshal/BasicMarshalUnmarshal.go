package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName, LastName string
	Age int
}

func main() {
	me := Person{ FirstName:"Prashant", LastName:"Choudhari", Age:27}
	meInJson, _ := json.Marshal(me)
	fmt.Println(string(meInJson))

	var myClone Person
	json.Unmarshal(meInJson, &myClone)
	println(myClone.FirstName)
}
