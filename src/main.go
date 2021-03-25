package main

import(
	"fmt"
)

func main() {
	fmt.Print("Insert your name : ")
	var nama string
	fmt.Scanln(&nama)
	
	fmt.Print("Your name is : "+nama)
}