package main

func FibRecursion(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return FibRecursion(n-1) + FibRecursion(n-2)
}

func FibIteration(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first := 0
	second := 1
	for i := 0; i < n; i++ {
		first, second = second, first+second
	}
	return first
}
