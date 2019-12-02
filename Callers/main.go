package main

import (
	"runtime"
	"fmt"
)

func main() {
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pc)
		if n <= 0 {
			break
		}
		fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
	}
	// Output:
	// skip = 0, pc = [4304486 4198562 4280114 4289760]
	// skip = 1, pc = [4198562 4280114 4289760]
	// skip = 2, pc = [4280114 4289760]
	// skip = 3, pc = [4289760]
}