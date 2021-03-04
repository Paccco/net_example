package main

func hello(i int) {
	println(i)
}

func main() {
	i := 5
	defer hello(i)
	i = 10 + i
}