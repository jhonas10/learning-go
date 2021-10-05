package main
import "fmt"
type student struct {
	firstName string
	lastName  string
}
func main() {
	Peter := student { firstName: "Peter", lastName: "Parker"}
	Peter.print()
}
func (s student) print() {
	fmt.Println("First Name: ", s.firstName, " Last Name: ", s.lastName)
}