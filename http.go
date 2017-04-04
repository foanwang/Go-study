package main
import("fmt"
	"net/http"
)

type Hello struct{}

func (h Hello)ServrHttp(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, "Hello")
}

func main(){
 	
	http.ListenAndServe("localhost:4000",nil)
}
