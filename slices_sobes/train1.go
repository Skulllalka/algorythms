// package main

// import "fmt"

// func main() {
// 	s := make([]int, 0, 5) // len= 0, cap = 5
// 	s = append(s, 1, 2, 3) // len = 3, cap = 5 , [1,2,3]0,0

// 	a := s[1:3] // len = 2, cap = 4, [2,3]0, 0, based on s

// 	a[0] = 10        // a[10,3] 0, 0, s[1,10,3]0,0
// 	a = append(a, 4) // a[10,3,4]0; s[1,10,3],4,0

// 	s = append(s, 5, 6) // len =5 , cap=5 ,s[1,10,3,5,6], a=[10,3,5],6

// 	a[1] = 20 // a[10,20,5]6, s[1,10,20,5,6]

// 	fmt.Println("s:", s)
// 	fmt.Println("a:", a)
// }
