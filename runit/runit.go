package main

import (
	"fmt"
	"github.com/dinnerTime/goSandbox"
)

func main() {
	foo := goSandbox.MyFoo{Bar: "some string"}
	fmt.Println("Value is ", foo.Bar)
}
