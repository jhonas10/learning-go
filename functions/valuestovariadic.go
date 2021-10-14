

package main

import "fmt"

func main() {
	
	variadicExample()
	variadicExample("hi", "how")
	variadicExample("hi", "how", "are")
	variadicExample("hi", "how", "are", "you?")
}

func variadicExample(s ...string) {
	fmt.Println(s)
}

