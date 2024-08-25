// Viết ham nhập 2 cạnh của hình chữ nhật, in ra diện tích, chu vi

package main

import (
	"fmt"
)

func calculateArea(w, h *float64) float64 {
	return *w * *h
}

func calculatePerimeter(w, h *float64) float64 {
	return 2 * (*w + *h)
}

func main() {
	var w, h float64
	fmt.Print("Enter width: ")
	fmt.Scanln(&w)
	fmt.Print("Enter height: ")
	fmt.Scanln(&h)
	fmt.Println("Area: ", calculateArea(&w, &h))
	fmt.Println("Perimeter: ", calculatePerimeter(&w, &h))
}
