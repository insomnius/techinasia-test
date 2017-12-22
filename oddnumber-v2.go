package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	diffparam := make(chan time.Duration)
	finishparam := make(chan time.Time)

	go loop(start, diffparam, finishparam)

	diff, finish := <-diffparam, <-finishparam

	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	fmt.Println("Perbedaan Waktu: ", diff)
}

func loop(start time.Time, diffparam chan time.Duration, finishparam chan time.Time) {
	for i := 1; i <= 2000; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	finishloop := time.Now()
	diffloop := finishloop.Sub(start)

	diffparam <- diffloop
	finishparam <- finishloop
}
