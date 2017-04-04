package main
import "fmt"
func printSlice(s string, x []int){
	fmt.Println(s, " len=", len(x), "cap=", cap[x], x)
}

func main(){
	a := make([]int ,5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c" ,c)
	d := c[2:5]
	printSlice("d", d)
}
