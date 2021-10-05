package main
import (
	"fmt"
	"time"
	"strings"
)
// go routines
func main() {
	go my_name_low("Jhonatan")
	var  wait string
	fmt.Scanln(&wait)
}

func my_name_low(name string)  {
	letter:= strings.Split(name,"")

	for _,letter:=range (letter){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}

	
}