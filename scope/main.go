package main

import "fmt"

var y int = 20

func main()  {
	x := 10
	fmt.Println(x)
	fmt.Println(y)
	//O z esta em outro arquivo, mas o mesmo pacote
	fmt.Println(z)
}

func printX()  {
	fmt.Println(y)
}