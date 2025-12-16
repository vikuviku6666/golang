package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range nums {
			if nums[i] % 2 == 0 {
				fmt.Printf("%d is even\n", num)
			}else {
				fmt.Printf("%d is Odd\n", num)
			}
	}
}