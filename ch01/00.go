package main

import "fmt"

func main() {
	w := "stressed"
	r := ""
	for i := len(w) - 1; 0 <= i; i-- {
		r += w[i : i+1]
	}
	fmt.Printf("%v > %+v\n", w, r)
}
