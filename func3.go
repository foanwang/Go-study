package main
import "fmt"
func temp()(int, int){
	return 8, 7
}
func main(){
	x, y := temp()
	fmt.Println(x, y)
}
