package main
import "fmt"
func main(){
	p:=[]int{2, 3, 4, 5, 7, 9, 13}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++{
		fmt.Println("p[", i,"] ==",  p[i])
	}
}
