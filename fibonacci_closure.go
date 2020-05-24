package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	result := []int{0, 1}
	return func(num int) int {
		if num > -1 {
			switch num {
			case 0:
				return result[num]
			case 1:
				return result[num]
			default:
				result = append(result, result[num-1]+result[num-2])
				return result[num]
			}
		} else {
			fmt.Println("args need to be bigger than -1")
			return 0
		}
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
