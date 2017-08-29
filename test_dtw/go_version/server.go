package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Id      string
	Balance uint64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func testpost(w http.ResponseWriter, r *http.Request) {
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
	// s1 := []float64{1,2,3,4}
	// s2 := []float64{4,3,2,1}
	// v1 := LB_Keogh(s1,s2,5)
	// v2 := DtwDistance(s1,s2)
	// fmt.Println(v1,v2)
	start := time.Now()
	_, assignments := get_centroid()
	// centroids, assignments := get_centroid()
	// fmt.Println("okkkkkkkkkkk")
	// fmt.Println(centroids)
	fmt.Println(assignments)
	elapsed := time.Since(start)
	fmt.Println("Binomial took", elapsed)
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/post", testpost)
	// http.ListenAndServe(":8080", nil)
}
