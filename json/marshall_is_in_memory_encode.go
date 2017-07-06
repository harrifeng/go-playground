package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type User struct {
		Name    string
		Age     int
		Zipcode int
	}

	peter := User{"hello", 12, 23}
	fmt.Println(peter)
	if data, err := json.Marshal(peter); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", data)
		var onePeople User
		json.Unmarshal(data, &onePeople)
		fmt.Println(onePeople)
	}

	os.Exit(0)
}

// <===================OUTPUT===================>
// {hello 12 23}
// {"Name":"hello","Age":12,"Zipcode":23}
// {hello 12 23}
