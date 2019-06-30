package seven

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// Sieve ...
func Sieve() int {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 10000; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	prime := <-ch
	return prime
}
