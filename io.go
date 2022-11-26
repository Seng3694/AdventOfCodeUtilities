package aocutil

import (
	"bufio"
	"os"
)

func FileReadAll[T string | []byte](path string) T {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return T(data)
}

func FileReadAllLines(path string) <-chan string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	ch := make(chan string)

	go func() {
		defer file.Close()
		defer close(ch)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}
