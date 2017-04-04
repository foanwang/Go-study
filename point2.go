package main
import "fmt"
type Vertex struct{
	x int
	y int
}
func main(){
	v := new(Vertex)
	fmt.Println(v)	
	v.x, v.y =1,23
	fmt.Println(v)
}
