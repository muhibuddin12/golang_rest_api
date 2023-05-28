package main

import (
	"fmt"

	"github.com/muhibuddin12/golang_rest_api/helper"
)

func Pow(number *int) {
	*number *= *number
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(helper.SayHello("Muhibuddin"))

	number := 12

	Pow(&number)
	fmt.Println(number)
}
