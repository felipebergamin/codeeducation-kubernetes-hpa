package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func sum() float64 {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}

func handler(response http.ResponseWriter, request *http.Request) {
	sum()
	fmt.Fprint(response, "CodeEducation Rocks!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
