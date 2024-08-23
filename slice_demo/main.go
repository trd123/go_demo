package main

import "fmt"

/**
 * 实验: 切片越界的情况
 */
func main() {
	arr := []int{1, 2, 3, 4, 5}
	s1 := arr[1:3]
	s2 := s1[3:4]
	fmt.Printf("arr: %v, arr地址: %p\n", arr, &arr) // arr: [1 2 3 4 5]
	fmt.Printf("s1: %v, s1地址: %p\n", s1, s1)      // s1: [2 3]
	fmt.Printf("s2: %v, s2地址: %p\n", s2, s2)      // 预计-越界,实际-s2: [5]
}
