package main

import "fmt"

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	
	fmt.Println("greeting is:", greeting)

  brainwash(&greeting)
	
	fmt.Println("greeting is now:", greeting)
}