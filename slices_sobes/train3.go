package main

import "fmt"

func mutate(x []int) []int {
	x = append(x, 100) // len = 2, cap = 3 , x[1,100]0, s[1,100]0
	x[0] = 50          // x[50,100]0, s[50, 100],0
	return x
}

func main() {
	s := make([]int, 0, 3) // len = 0 , cap = 3
	s = append(s, 1, 2)    // len = 2, cap =3 , s[1,2]0

	a := s[:1] // based on s, len = 1, cap = 3, a[1]2,0

	a = mutate(a)    // new arr len = 2, cap = 3, a[50,100],0;
	s = append(s, 3) //s[50, 100, 3]

	fmt.Println("s:", s)
	fmt.Println("a:", a)

	fmt.Println("cap a", a[0:3])
}
