package main

import "fmt"

// func climbStairs(n int) int {
// 	// recursion functions always need a base case
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// if n == 3 {
// 	return 3
// }

// 	// recursive case
// 	waysToTop := climbStairs(n-1) + climbStairs(n-2)
// 	return waysToTop
// }

func climbStairs(n int, memo map[int]int) int {
	// check if the result is already computed
	if val, exists := memo[n]; exists {
		return val
	}

	// base cases
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}

	// recursive case with memoization
	memo[n] = climbStairs(n-1, memo) + climbStairs(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	fmt.Println("You are climbing a staircase. It takes n steps to reach the top.")
	fmt.Println("Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?")

	// test cases
	for i := 1; i <= 10; i++ {
		fmt.Println("n =", i, "ways to top =", climbStairs(i, memo))
	}
}
