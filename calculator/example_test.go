package calculator

import "fmt"

func ExampleBorderlen() {
	fmt.Printf("%.2f", Borderlen(30))
	// Output: 19.42
}

func ExampleDiameter() {
	fmt.Printf("%.2f", Diameter(10))
	// Output: 3.57
}

func ExampleDidgits() {
	Didgits(456)
	// Output: В этом числе:\n сотен: 4,\n десятков: 5,\n единиц: 6. \n
}
