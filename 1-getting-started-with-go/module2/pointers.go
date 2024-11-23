package main

func f(x int) *int {
	y := x + 5
	return &y
}

func main() {
	p1 := f(10)
	p2 := f(100)
	println("p1:", *p1, "p2:", *p2)
}
