package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	go loop()

	finish := time.Now()
	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	diff := finish.Sub(start)
	fmt.Println("Perbedaan Waktu: ", diff)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func loop() {
	for i := 1; i <= 2000; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
