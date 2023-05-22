package helper

import "fmt"

func init() {
	fmt.Println("Function ini di panggil")
}

func SayHello(name string) string {
	return "Hello " + name
}
