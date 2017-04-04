package main
import "fmt"
func add(args...int)int{
	total := 0
	for _, v:= range args{
		total += v
	}
	return total
}

func main(){
	xs :=[]int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(add(xs...))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4, 5))
}
