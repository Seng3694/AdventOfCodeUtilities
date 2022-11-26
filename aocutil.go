package aocutil

import (
	"fmt"
	"time"
)

var startTime time.Time = time.Now()

func AOCFinish[T any](parts ...T) {
	elapsed := time.Since(startTime)
	fmt.Println("Solutions:")
	for i, p := range parts {
		fmt.Printf(" Part %v: '%v'\n", i, p)
	}
	fmt.Printf("found in %v seconds\n", elapsed.Seconds())
}
