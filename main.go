package main

func main() {

}

func worker(jobs <-chan int, results chan<- int) {

}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
