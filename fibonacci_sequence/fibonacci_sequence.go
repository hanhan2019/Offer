package main

func main() {

}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-2) + fib(N-1)
}

func better(N int) int {
	list := make(map[int]int)
	return fib_better(N, list)
}
func fib_better(N int, list map[int]int) int {
	if N < 2 {
		return N
	}
	if N == 2 {
		return 1
	}
	if list[N] != 0 {
		return list[N]
	}
	list[N] = fib_better(N-1, list) + fib_better(N-2, list)
	return list[N]
}

func fib_dp(N int) int {
	if N < 2 {
		return N
	}
	pre, cur := 0, 1
	for i := 2; i <= N; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
