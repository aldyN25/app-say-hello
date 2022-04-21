package main

import (
	"fmt"

	p "github.com/aldyN25/mod_Say_Hello/v2"
)

func main() {
	name := "aldy"
	p.SayHello(name)
	fmt.Println(p.SayHello(name))
}
