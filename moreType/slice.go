package moreType

import "fmt"

func Test2() {
	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Aguest", 9: "September", 10: "October", 11: "November", 12: "December"}

	fmt.Println(months[1:4])

}
