package main

func main() {
	ch := make(chan int)
	x := 233
	ch <- x

}
