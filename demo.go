package main

import (
	"fmt"
	"time"
)

func main() {

	//go的多线程-协程
	go printgo()

	for i := 0; i < 100; i++ {
		fmt.Println("main======>", i)
	}

	time.Sleep(time.Second * 3)

}

func printgo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("printgo======>", i)
	}
}
