package aocutil

import (
	"fmt"
	"time"

	"github.com/loov/hrtime"
)

var startTime time.Duration = hrtime.Now()

func AOCFinish[T any](parts ...T) {
	elapsed := hrtime.Since(startTime)
	fmt.Println("Solutions:")
	for i, p := range parts {
		fmt.Printf(" Part %v: '%v'\n", i, p)
	}
	fmt.Printf("found in %v seconds\n", elapsed.Seconds())
}
