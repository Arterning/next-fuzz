package calc

import "fmt"

// private var belong to package
var home = "home"

var Age = 90

// Add 首字母大写表示public
func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func PrintWelcome() {
	println("welcome to " + home)
}

// package init method would be executed..
func init() {
	fmt.Println("calc pacakge init...")
}
