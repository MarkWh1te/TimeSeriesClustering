package main 
import (
	"encoding/json"
	"fmt"
	"net/http"
	// "time"
)

type User struct {
	Id      string
	Balance uint64
}
type Response1 struct {
    Page   int
	Fruits []string
	Sort_keys []string
	tt []int
	// cluster   [][]float64
}
type StockData struct {
	Source    map[string][]float64
	Cluster   map[int][]int
	Centers   [][]float64
	Sort_keys []string
	Id      string
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

func cluster(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	days := r.Form.Get("days") 
	types := r.Form.Get("tyes") 
	fmt.Println(days,types)
	datas := make(map[string][]float64)
	datas["a"] = []float64{1.11, 2.22, 3.33, 4.44, 5.55}
	datas["b"] = []float64{2.34, 4.56, 5.12, 6.04, 5.55}
	datas["c"] = []float64{1.55, 2.21, 3.13, 4.24, 5.55}
	datas["d"] = []float64{2.34, 4.56, 5.12, 6.04, 5.55}
	datas["e"] = []float64{11.11, 23.22, 32.33, 41.44, 15.55}
	// fmt.Println(datas)

	// raw := readcsv("2016-07-012017-07-01.csv")
	// csv_data := dataclean(raw)
	// fmt.Println(csv_data)
	// datas = csv_data

	centroids, assignments, keys := get_centroid()
	// fmt.Println(centroids)
	// fmt.Println(assignments)
	// fmt.Println(keys)
 	// data:= StockData{Source:datas,Sort_keys:keys,Cluster:centroids,Centers:assignments}
 	data:= StockData{Source:datas,Sort_keys:keys,Cluster:assignments,Centers:centroids}
	jData, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
	}
	// fmt.Println("dddd",jData,data)
	// fmt.Println(string(jData))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func main() {
	// cal run time
	// start := time.Now()
	// _, assignments := get_centroid()
	// fmt.Println(assignments)
	// elapsed := time.Since(start)
	// fmt.Println("Binomial took", elapsed)

	// http.HandleFunc("/", handler)
	fs := http.FileServer(http.Dir("static"))
	fmt.Println(fs)
	http.Handle("/", fs)
	http.HandleFunc("/post", testpost)
	http.HandleFunc("/cluster", cluster)
	http.ListenAndServe(":8080", nil)
}
