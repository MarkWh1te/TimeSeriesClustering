package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)
type User struct{
    Id      string
    Balance uint64
}

func main() {
    u := User{Id: "US123", Balance: 8}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(u)
    // res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
	res, _ := http.Post("http://127.0.0.1:8080/post", "application/json; charset=utf-8", b)
	fmt.Println(res)
}