package main

import "fmt"

func add(x int,y int) int {
   return x + y
}

func main(){
   fmt.Println("X = 2, Y = 3, X+Y =", add(2, 3))
}

