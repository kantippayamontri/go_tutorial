package main

import "fmt"
import "rsc.io/quote/v4"

func main(){ //A main function executes by default when you run the main package.
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}


