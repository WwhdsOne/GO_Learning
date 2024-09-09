package main

import "log"

func main() {
	defer func() {
		if handler := recover(); handler != nil { // recover()捕获panic异常
			log.Println("Panic:", handler)
		}
		log.Println("This will not be printed")
	}()
	log.Fatal("Fatal")
}
