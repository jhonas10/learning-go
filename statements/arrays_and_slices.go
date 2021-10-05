package main

import (  
    "fmt"
)


func main() {  
    var a [5]int
    a[0] = 12 
    a[1] = 78
    a[2] = 50
	a[3] = 79
	a[4] = 80
    fmt.Println(a)


    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}