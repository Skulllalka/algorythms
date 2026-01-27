package main

import "fmt"

func main() {
	s := make([]int, 0, 2) // len = 0, cap = 2
	s = append(s, 1, 2)    // len =2, cap=2  s[1, 2]

	a := s[:1]  // based on s , len =1, cap = 2, a[1]2
	b := s[1:2] // based on s, len=1, cap=1, b[2]

	a = append(a, 10) // based on s, len=2, cap = 2, a[1,10], s[1,10], b[10]
	b = append(b, 20) // new array , len=2, cap=2, b[10, 20]

	s = append(s, 30) // new array , len =3, cap =4 , s[1,10,30]0

	fmt.Println("s:", s) //[1,10,30]0
	fmt.Println("a:", a) //[1,10]
	fmt.Println("b:", b) //[10, 20]
}
