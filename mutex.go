package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	money int
)

func init() {
	money = 2500
}

func abhi(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Abhi is giving me %d :%d\n", value, money)
	money += value
	mutex.Unlock()
	wg.Done()
}

func karthik(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Karthik is taking from me %d :%d\n", value, money)
	money -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Simple example of mutex")

	var wg sync.WaitGroup
	wg.Add(2)
	go abhi(500, &wg)
	go karthik(800, &wg)
	wg.Wait()
	fmt.Println("Total money now, I have %d\n", money)
}
