package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	var arrayNumber [20]int

	for i := 0; i < 20; i++ {
		fmt.Printf("Masukan Angka (harus ganjil lho): ")
		fmt.Scan(&arrayNumber[i])

		if arrayNumber[i]%2 == 0 {
			fmt.Println("Uppss.. Program dihentikan, angka yang dimasukan haruslah ganjil. :(")
			os.Exit(0)
		}
	}

	fmt.Println("Output:")
	fmt.Println(arrayNumber)

	fmt.Printf("\n\nSorting...\n\n")

	bubble(arrayNumber)

	selection(arrayNumber)

	insertion(arrayNumber)

	fmt.Println(arrayNumber)
}

// Input : 7 3 9 1 21 13 11 5 15 17 23 25 41 39 99 77 67 55 57 19

func bubble(arrayNumber [20]int) {
	start := time.Now()
	fmt.Println("Bubble Short")
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	length := len(arrayNumber)

	for length > 0 {

		newlength := 0
		for i := 1; i < length; i++ {
			x := arrayNumber[i-1]
			y := arrayNumber[i]

			if x > y {
				arrayNumber[i-1] = y
				arrayNumber[i] = x
				newlength = i
			}
		}

		length = newlength
	}

	fmt.Println(arrayNumber)

	finish := time.Now()
	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	diff := float64(finish.Sub(start)) / 1000000000
	fmt.Println("Waktu yang dihabiskan: ", diff, "detik atau", diff*1000, "milidetik\n")
}

func selection(arrayNumber [20]int) {
	start := time.Now()
	fmt.Println("Selection Short")
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	length := len(arrayNumber)

	for i := 0; i < length; i++ {
		toswap := i
		lowest := arrayNumber[i]

		for j := i; j < length; j++ {

			if arrayNumber[j] < lowest {
				lowest = arrayNumber[j]
				toswap = j
			}
		}

		if toswap != i {
			temp := arrayNumber[i]
			arrayNumber[i] = lowest
			arrayNumber[toswap] = temp
		}
	}

	fmt.Println(arrayNumber)

	finish := time.Now()
	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	diff := float64(finish.Sub(start)) / 1000000000
	fmt.Println("Waktu yang dihabiskan: ", diff, "detik atau", diff*1000, "milidetik\n")
}

func insertion(arrayNumber [20]int) {
	start := time.Now()
	fmt.Println("Insertion Short")
	fmt.Printf("Waktu Mulai: %02d:%02d:%02d \n", start.Hour(), start.Minute(), start.Second())

	length := len(arrayNumber)

	for i := 0; i < length; i++ {
		toswap := i
		lowest := arrayNumber[i]

		for j := i; j >= 0; j-- {

			if arrayNumber[j] > lowest {
				arrayNumber[j+1] = arrayNumber[j]
				toswap = j
			}
		}

		if toswap != i {
			arrayNumber[toswap] = lowest
		}
	}

	fmt.Println(arrayNumber)

	finish := time.Now()
	fmt.Printf("Waktu Selesai: %02d:%02d:%02d \n", finish.Hour(), finish.Minute(), finish.Second())
	diff := float64(finish.Sub(start)) / 1000000000
	fmt.Println("Waktu yang dihabiskan: ", diff, "detik atau", diff*1000, "milidetik\n")
}
