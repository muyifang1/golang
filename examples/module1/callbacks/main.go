package main

func main() {
	//var a *int
	//*a += 1
	DoOperation(1, increase)
	DoOperation(1, decrease)
}

func increase(a, b int) {
	//return a + b
	println("increase result is:", a+b)
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
