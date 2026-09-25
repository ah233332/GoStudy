package main

func add(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans = ans + i
	}
	return ans
}
