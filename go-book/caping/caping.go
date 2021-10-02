package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("Q2 Len: %v | Q2 Cap: %v | Q2: %v\n", len(q2), cap(q2), q2)
	fmt.Printf("Summer Len: %v | Summer Cap: %v | Summer: %v\n", len(summer), cap(summer), summer)

	defer func() {
		recover()

		increaseSliceCap(summer)
	}()
	fmt.Printf("Over Len Q2: %v", q2[4])
}

func increaseSliceCap(summer []string) {
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
