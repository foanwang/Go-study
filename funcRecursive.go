package main
import "fmt"
func recursie(x int)int{
	if x == 0{
		return 1
	}
	return x*recursie(x - 1)
}

func main(){
	fmt.Println(recursie(3))
}
