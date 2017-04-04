package main
import "fmt"
func first(){
	fmt.Println("1st")
}
func second(){
	fmt.Println("2nd")
}
func third(){
	fmt.Println("3th")
}

func main(){
	third()
	defer second()
	first()
}
