package main

import "time"
import "log"
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	i := 0
	for _, v := range s {
		sum += v
		for _, v := range s {
			i += v
		}
	}
	c <- sum // send sum to c
}

func sum_n(s []int) int {
	sum := 0
	i := 0
	for _, v := range s {
		sum += v
		for _, v := range s {
			i += v
		}
	}

	return sum
}

func main() {
	s := make([]int, 200000)
	for i, _ := range s {
		s[i] = 1
	}

	start_r := time.Now()

	c := make(chan int)
	go sum(s[:len(s)/4], c)
	go sum(s[len(s)/4:(len(s)/4)*2], c)
	go sum(s[(len(s)/4)*2:(len(s)/4)*3], c)
	go sum(s[(len(s)/4)*3:], c)
	v, w, x, y := <-c, <-c, <-c, <-c // receive from c
	elapsed_r := time.Since(start_r)

	start_n := time.Now()
	n := sum_n(s)
	elapsed_n := time.Since(start_n)

	log.Printf("With go rotune took %s", elapsed_r)
	log.Printf("Normal flow took %s", elapsed_n)

	fmt.Println("Answer from concurrent execution", v+w+x+y)
	fmt.Println("Answer from normal execution", n, "Total ops:", 200000*200000)

}
