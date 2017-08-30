package main 
import ( 
	"encoding/json"
	"fmt"
	"net/http"
	// "time"
)
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

// func datetoidx(startTime int,endTime int)(int int){
// 	return staridx
// }
var start0 = readcsv("start0-2011-01-012017-08-01.csv")
var start3 = readcsv("start3-2011-01-012017-08-01.csv")
var start6 = readcsv("start6-2011-01-012017-08-01.csv")
// var data0 = dataclean(start0)
// var data3 = dataclean(start3)
// var data6 = dataclean(start6)


func cluster(w http.ResponseWriter, r *http.Request) {
	// get post args
	r.ParseForm()
	days := r.Form.Get("days") 
	types := r.Form.Get("tpyes") 
	fmt.Println(days,types)

	data0 := ShortData(start3,20,30)
	fmt.Println(data0)


	// fmt.Println(raw)
	// csv_data := dataclean(raw)
	// fmt.Println(csv_data)
	// datas = csv_data

	// get the algorithms answer
	centroids, assignments, keys,data_list := get_centroid(data0,20)
	data:= StockData{Source:data_list,Sort_keys:keys,Cluster:assignments,Centers:centroids}
	// data:= StockData{Source:data0,Sort_keys:keys,Cluster:assignments,Centers:centroids}

	// generate  json data
	jData, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
	}
	// write json data into response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func main() {

	// calculate fucntion runining time 
	// start := time.Now()
	// _, assignments := get_centroid()
	// fmt.Println(assignments)
	// elapsed := time.Since(start)
	// fmt.Println("Binomial took", elapsed)

	// http.HandleFunc("/", handler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/cluster", cluster)
	http.ListenAndServe(":8080", nil)
}