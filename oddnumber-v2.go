package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	loop()

	finish := time.Now()
	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	diff := finish.Sub(start)
	fmt.Println("Perbedaan Waktu: ", diff)
}

func loop() {
	for i := 1; i <= 2000; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
