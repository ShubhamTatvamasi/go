package reversestring

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func ReverseString(mystring string) {

	p(mystring)
	// p(s.ToLower(mystring))
	// p(s.Repeat(mystring, 3))
	// const backward = s.Split(mystring)
	p(s.Split(mystring, ""))

	myletters := s.Split(mystring, "")

	// backwards := []
	var backwards []string

	for _, myletter := range myletters {
		// backwards :=
		backwards = append(myletters)
		fmt.Print(myletter)
	}

	fmt.Print(backwards)

	// p(myarray)

	// mystring.Split()

	// const array = mystring

	// for i, v := range mystring {

	// 	p(i, v)

	// }

}
