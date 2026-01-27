package main

import "fmt"

func main() {
	s := make([]int, 0, 4) // len = 0, cap = 4
	s = append(s, 1, 2, 3) // len= 3, cap = 4 [1,2,3] 0

	a := s[:2]  // based on s, len = 2, cap = 4, a[1, 2]
	b := s[1:3] // based on s, len 2, cap = 3, b[2, 3]

	a = append(a, 10) // based on s, len = 3 , cap = 4, a[1,2,10], s[1,2,10], b[2,10]
	b = append(b, 20) // based on s, len = 3 , cap = 3, b[2, 10, 20], s[1,2,10],20, a[1,2,10],20

	s = append(s, 30) // len = 4 , cap = 4 , s[1,2,10,30], b[2,10,30]

	fmt.Println("s:", s) //[1,2,10,30]
	fmt.Println("a:", a) //[1,2,10]
	fmt.Println("b:", b) //[2,10,30]
}
