package main
import (
"fmt"
	"net/http"
	"encoding/json"
)
type User struct{
    Id      string
    Balance uint64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func testpost(w http.ResponseWriter,r *http.Request){
	var t User
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Id)

	// fmt.Println(r.Body)
	u := User{Id: "US123", Balance: 8}
	json.NewEncoder(w).Encode(u)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/post", testpost)
    http.ListenAndServe(":8080", nil)
}