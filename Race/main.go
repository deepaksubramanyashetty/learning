// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("goroutines :", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()

			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("goroutines :", runtime.NumGoroutine())
	fmt.Println("counts :", counter)
}
