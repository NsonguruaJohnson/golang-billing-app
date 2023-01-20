package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("tom cruise")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("Nsongurua johnson")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Nsongurua Akpan Jimmy Johnson")
	fmt.Println(fn3, sn3)

	fn4, sn4 := getInitials("Nsongurua")
	fmt.Println(fn4, sn4)

}
