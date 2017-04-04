package main
import "fmt"
type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["test"] =Vertex{
		123.234, 12312131,
	}
	fmt.Println(m["test"])
}
