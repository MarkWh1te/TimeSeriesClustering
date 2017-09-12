package main 
import ( 
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "time"
)
type StockData struct {
	Source    map[string][]float64
	Origin   map[string][]float64
	Cluster   map[int][]int
	Centers   [][]float64
	Sort_keys []string
	Id      string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func concatemaps(x map[string][]float64,y map[string][]float64)map[string][]float64{
	new := make(map[string][]float64)
	for k,v := range x{
		new[k] = v
	}
	for k,v := range y{
		new[k] = v
	}
	return new
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
			//see k v
            return true
        }
    }
    return false
}


func packData(sw string,data map[string][]float64)map[string][]float64{
	new := make(map[string][]float64)
	stocklist :=  indumap[sw]
	for k,v := range(data){
		if stringInSlice(k[:len(k)-3],stocklist){
			new[k[:len(k)-3]] = v
		}
	}
	fmt.Println("nums",len(stocklist),len(data),len(new))
	return new
}

var start0 = readcsv("start0-2011-01-012017-08-01.csv")
var start3 = readcsv("start3-2011-01-012017-08-01.csv")
var start6 = readcsv("start6-2011-01-012017-08-01.csv")
var start03 = concatemaps(start0,start3)
var start06 = concatemaps(start0,start6)
var start36 = concatemaps(start3,start6)
var start036 = concatemaps(start03,start6)

var codemap = map[string]map[string][]float64{
	"0":start0,
	"3":start3,
	"6":start6,
	"6,3":start36,
	"0,3":start03,
	"0,6":start06,
	"0,6,3":start036,
}
// var data0 = dataclean(start0)
// var data3 = dataclean(start3)
// var data6 = dataclean(start6)

func timeToIndex(starttime float64,endtime float64)(int,int){
	var min_idx,max_idx int
	// starttime = float64(starttime)
	// endtime = float64(endtime)
	for k,v := range stocklist{
		min_idx = k
		if starttime == v{
			break
		}
		if v > starttime{
			min_idx = k-1
			break
		}
	}
	for k,v := range stocklist{
		max_idx = k
		if endtime == v{
			max_idx = k
			break
		}
		if endtime < v{
			max_idx = k-1
			break
		}
	}
	fmt.Println(min_idx,max_idx)
	return min_idx,max_idx
}


func cluster(w http.ResponseWriter, r *http.Request) {
	// get post args
	fmt.Println("test")
	r.ParseForm()
	start_date := r.Form.Get("start_date") 
	sw := r.Form.Get("sw") 
	end_date := r.Form.Get("end_date") 
	types := r.Form.Get("types") 
	stock := r.Form.Get("stock")
	methods := r.Form.Get("method")
	// centroids, assignments, keys,data_list := get_centroid(rawdata,int(typesint))
	fmt.Println(start_date,end_date,types,stock,methods,sw)


	// init data
	stock = "0,6,3"
	datas := codemap[stock]
	start_float,_ := strconv.ParseFloat(start_date,64)
	end_float,_ := strconv.ParseFloat(end_date,64)
	min_idx,max_idx := timeToIndex(start_float,end_float)
	new_datas := packData(sw,datas)
	rawdata := ShortData(new_datas,min_idx,max_idx)
	// rawdata := ShortData(datas,20,30)

	// fmt.Println(stock)
	// fmt.Println(start36)
	// fmt.Println(codemap["6,3"])
	// fmt.Println(raw)
	// csv_data := dataclean(raw)
	// fmt.Println(csv_data)
	// datas = csv_data

	// get the algorithms answer
	typesint,_ := strconv.ParseInt(types,10,64)

	// centroids, assignments, keys,data_list := get_centroid_new(rawdata,int(typesint))
	if methods =="0"{
		centroids, assignments, keys,data_list := get_centroid(rawdata,int(typesint))
		data:= StockData{Origin:rawdata,Source:data_list,Sort_keys:keys,Cluster:assignments,Centers:centroids}
	// generate  json data
	jData, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
	}
	// fmt.Println(keys)

	// write json data into response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

	}else{
		centroids, assignments, keys,data_list := get_centroid_rate(rawdata,int(typesint))
		data:= StockData{Origin:rawdata,Source:data_list,Sort_keys:keys,Cluster:assignments,Centers:centroids}
	// generate  json data
	jData, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
	}
	// fmt.Println(keys)

	// write json data into response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

	}
	// data:= StockData{Source:data0,Sort_keys:keys,Cluster:assignments,Centers:centroids}
}


func main() {

	// calculate fucntion runining time kk
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