package main

import (
	"fmt"
	"os"
	"strings"
)

//echo1
/*func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}*/

//echo2
/*func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}*/

//echo3
func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}