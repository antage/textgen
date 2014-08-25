package main

func For(start int, end int) <-chan int {
	step := 1
	if end < start {
		step = -1
	}

	c := make(chan int)

	go func(start int, end int, c chan int) {
		i := start
		for {
			c <- i
			i += step
			if end < start {
				if i < end {
					close(c)
					return
				}
			} else {
				if i > end {
					close(c)
					return
				}
			}
		}
	}(start, end, c)

	return c
}

func List(items ...interface{}) []interface{} {
	return items
}
