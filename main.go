package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ {
		f := float64(i)
		fmt.Printf("%d%%, %d%% { background-image: radial-gradient(transparent %2.1f%%, hsl(0, 0%%, 100%%) %2.1f%%); }\n", i, (100 - i), (35 + (f / 2)), (35 + (f / 2)))
	}
}
