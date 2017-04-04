package main
import "fmt"
func swapstring(x, y string)(string, string){
  return y, x
}
func main(){

  x := "hello"
  y := "world"
  fmt.Println(swapstring(x, y))
}

