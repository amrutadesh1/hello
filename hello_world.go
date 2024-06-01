package main

import (
	"fmt"
	jc "hello/json_code"
)

func main() {
	fmt.Println("Hello World!!")
	jc.ReadFile("data/users.json")
}
