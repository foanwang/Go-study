package main
import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println(
	strings.Contains("test", "es"), 
	strings.Count("test","t"),
	strings.HasPrefix("test", "te"), 
	strings.HasSuffix("test", "st"), 
	strings.Index("test", "e"), 
	strings.Join([]string{"a", "b"},"c"), 
	strings.Repeat("a", 5),
	strings.Replace("aaaaa","a","b", 2),
	strings.Split("a-b-c", "-"),
	strings.ToLower("test"),
	)

}
