package main 
import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)
func k_means_clust_new(data_list map[int][]float64, num_clust int, num_iter int, w int) ([][]float64, map[int][]int, map[int]float64) {

	// generate init centorids
	var keys []int
	for k, _ := range data_list {
		keys = append(keys, k)
	}
	rand_keys := (generateRandomNumber(0, len(keys), num_clust))
	var centroids [][]float64
	for _, v := range rand_keys {
		centroids = append(centroids, data_list[keys[v]])
	}

	counter := 0
	assignments := make(map[int][]int)
	sumdistance := make(map[int]float64)
	for i := 0; i < num_iter; i++ {
		counter += 1
		// fmt.Println("counter times", counter, "centroids", centroids)
		// init empty  assignment every iteration
		for k, _ := range centroids {
			assignments[k] = []int{}
			sumdistance[k] = 0.0
		}
		//cal distance to centroids
		for k, v := range data_list {
			min_dist := math.Inf(1)
			var closest_clust int
			for kk, vv := range centroids {
				if LB_Keogh(v, vv, w) < min_dist {
					cur_dist := DtwDistance(v, vv)
					if cur_dist < min_dist {
						min_dist = cur_dist
						closest_clust = kk
					}
				}
			}
			assignments[closest_clust] = append(assignments[closest_clust], k)
			sumdistance[closest_clust] += min_dist
		}

		for k, v := range assignments {
			sumdistance[k] = sumdistance[k] / float64(len(v))
			var clust_sum []float64
			for _, vv := range v {
				for kkk, vvv := range data_list[vv] {
					if len(clust_sum) < kkk+1 {
						clust_sum = append(clust_sum, 0)
					}
					clust_sum[kkk] += vvv
				}
			}
			for kk, vv := range clust_sum {
				centroids[k][kk] = vv / float64(len(v))
			}

		}
	}
	return centroids, assignments, sumdistance
}

func bisecting_k_means_clust(data_list map[int][]float64, num_clust int, num_iter int, w int) ([][]float64, [][]int) {
	//var clust_sum []float64
	// generate init centorids
	var allcentroids [][]float64
	var allassign [][]int
	var alldistance []float64
	for {
		var lastd = math.Inf(1)
		lastassign := make(map[int][]int)
		lastdistance := make(map[int]float64)
		var lastcentroids [][]float64
		//get min value
		for i := 0; i < 20; i++ {
			centroids, assignments, sumdistance := k_means_clust_new(data_list, 2, num_iter, w)
			nowdistance := sum([]float64{sumdistance[0], sumdistance[1]})
			if nowdistance < lastd {
				lastd = nowdistance
				lastcentroids = centroids
				lastassign = assignments
				lastdistance = sumdistance
			}
		}
		// 汇总
		for k, v := range lastcentroids {
			allcentroids = append(allcentroids, v)
			allassign = append(allassign, lastassign[k])
			alldistance = append(alldistance, lastdistance[k])
		}
		//fmt.Println(allcentroids)
		if len(allcentroids) == num_clust {
			break
		}
		// cal max
		var maxdistance = math.Inf(-1)
		var max_index = 0

		for k, v := range alldistance {
			if v > maxdistance {
				maxdistance = v
				max_index = k
			}
		}
		max_index_list := allassign[max_index]
		tmp := make(map[int][]float64)
		for _, v := range max_index_list {
			tmp[v] = data_list[v]
		}
		data_list = tmp
		allcentroids = DeleteSlice(allcentroids, max_index)
		allassign = DeleteSlice3(allassign, max_index)
		alldistance = DeleteSlice2(alldistance, max_index)

		//DeleteSlice(allassign,max_index)
		//DeleteSlice(,max_index)
	}

	fmt.Println(allassign)
	return allcentroids, allassign
}

//删除切片
func DeleteSlice(sss [][]float64, index int) ([][]float64) {
	//sliceValue := reflect.ValueOf(slice)
	length := len(sss)
	if sss == nil || length == 0 || (length-1) < index {
		return nil
	}
	if length-1 == index {
		return sss[0:index]
	} else if (length - 1) >= index {
		//return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
		tmp := sss[0: index]
		for _, v := range sss[index+1:length] {
			tmp = append(tmp, v)
		}
		return tmp
	}
	return nil
}
func DeleteSlice2(sss []float64, index int) ([]float64) {
	//sliceValue := reflect.ValueOf(slice)
	length := len(sss)
	if sss == nil || length == 0 || (length-1) < index {
		return nil
	}
	if length-1 == index {
		return sss[0:index]
	} else if (length - 1) >= index {
		//return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
		tmp := sss[0: index]
		for _, v := range sss[index+1:length] {
			tmp = append(tmp, v)
		}
		return tmp
	}
	return nil
}

func DeleteSlice3(sss [][]int, index int) ([][]int) {
	//sliceValue := reflect.ValueOf(slice)
	length := len(sss)
	if sss == nil || length == 0 || (length-1) < index {
		return nil
	}
	if length-1 == index {
		return sss[0:index]
	} else if (length - 1) >= index {
		//return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
		tmp := sss[0: index]
		for _, v := range sss[index+1:length] {
			tmp = append(tmp, v)
		}
		return tmp
	}
	return nil
}

func ShortData(data map[string][]float64,start int,end int)(map[string][]float64){
	shortData := make(map[string][]float64)
	for i := range data{
		shortData[i] = data[i][start:end]
	}
	return shortData
}

func readcsv(path string) map[string][]float64{
	// read csv part
	file, err := os.Open(path)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("error:", err)
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// rawCSVdata = rawCSVdata[:30]
	// rawCSVdata = rawCSVdata

	// sanity check, display to standard output
	// for _, each := range rawCSVdata {
	// fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	// }
	newdata := dataclean(rawCSVdata)
	return newdata
}

func dataclean(raw [][]string) map[string][]float64 {
	csv_data := make(map[string][]float64)
	for _, line := range raw {
		for _, nums := range line {
			if n, err := strconv.ParseFloat(nums, 64); err == nil {
				csv_data[line[0]] = append([]float64{n}, csv_data[line[0]]...)
			}
		}
	}
	return csv_data
}

func get_centroid(datas map[string][]float64,n int) ([][]float64, map[int][]int,[]string,map[string][]float64) {
	var keys = sorted_keys(datas)
	// fmt.Println(keys)
	var data_list [][]float64
	// var data_map map[string][]float64
	data_map := make(map[string][]float64)
	for _, v := range keys {
		a := to_zero(datas[v])
		data_list = append(data_list, a)
		data_map[v] = a
	}
	centroids, assignments := k_means_clust(data_list, n, 100, 3)
	return centroids, assignments,keys,data_map
}
func get_centroid_rate(datas map[string][]float64,n int) ([][]float64, map[int][]int,[]string,map[string][]float64) {
	var keys = sorted_keys(datas)
	var data_list [][]float64
	data_map := make(map[string][]float64)
	for _, v := range keys {
		a := to_rate(datas[v])
		data_list = append(data_list, a)
		data_map[v] = a
	}
	centroids, assignments := k_means_clust(data_list, n, 100, 3)
	return centroids, assignments,keys,data_map
}
func get_centroid_new(datas map[string][]float64,n int) ([][]float64, map[int][]int,[]string,map[string][]float64) {
	var keys = sorted_keys(datas)
	data_list := make(map[int][]float64)
	data_map := make(map[string][]float64)
	for k, v := range keys {
		a := to_zero(datas[v])
		data_list[k] = a
		data_map[v] = a
	}
	centroids, assignments := bisecting_k_means_clust(data_list, n, 20, 3)
	newassignments := make(map[int][]int)
	for k,v := range assignments{
		newassignments[k] = v
	}
	return centroids, newassignments,keys,data_map
}


func to_zero(arr []float64) []float64 {
	var tmp []float64
	for _, v := range arr {
		tmp = append(tmp, Round(v-arr[0], 2))
	}
	return tmp
}

func to_rate(arr []float64) []float64 {
	var tmp []float64
	for k, v := range arr {
		if k!=0{
		tmp = append(tmp,Round((v-arr[k-1])/arr[k-1],2))
		}
	}
	return tmp
}

func sorted_keys(m map[string][]float64) []string {
	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	return sorted_keys
}

func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func max(data []float64) float64 {
	maximum := math.Inf(-1)
	for i := 0; i < len(data); i++ {
		maximum = math.Max(data[i], maximum)
	}
	return maximum
}

func min(data []float64) float64 {
	minimum := math.Inf(1)
	for i := 0; i < len(data); i++ {
		minimum = math.Min(data[i], minimum)
	}
	return minimum
}

func sum(data []float64) float64 {
	all := 0.0
	for i := 0; i < len(data); i++ {
		all += data[i]
	}
	return all
}

func DtwDistance(s1 []float64, s2 []float64) float64 {
	DTW := make(map[[2]int]float64)
	for i := -1; i < len(s1); i++ {
		for j := -1; j < len(s2); j++ {
			keyarr := [2]int{i, j}
			DTW[keyarr] = math.Inf(1)
		}
	}
	DTW[[2]int{-1, -1}] = 0

	w := math.Max(5, math.Abs(float64(len(s1)-len(s2))))
	for i := 0; i < len(s1); i++ {
		lower := int(math.Max(float64(0), float64(i)-w))
		upper := int(math.Min(float64(len(s2)), float64(i)+w))
		for j := lower; j < upper; j++ {
			dist := math.Pow(s1[i]-s2[j], 2)
			values := []float64{DTW[[2]int{i, j - 1}], DTW[[2]int{i - 1, j}], DTW[[2]int{i - 1, j - 1}]}
			DTW[[2]int{i, j}] = dist + min(values)
		}
	}
	return math.Sqrt(DTW[[2]int{len(s1) - 1, len(s2) - 1}])
}

func LB_Keogh(s1 []float64, s2 []float64, r int) float64 {
	LB_sum := 0.0
	for ind, i := range s1 {
		start := 0
		if ind-r >= 0 {
			start = ind - r
		}
		// fmt.Println(start,ind+r)
		end := ind + r
		if end >= len(s1) {
			end = len(s1)
		}
		lower_bound := (min(s2[start:end]))
		upper_bound := (max(s2[start:end]))

		if i > upper_bound {
			LB_sum = LB_sum + math.Pow((i-upper_bound), 2)
		} else if i < lower_bound {
			LB_sum = LB_sum + math.Pow((i-lower_bound), 2)
		}
	}
	return LB_sum
}
func rand_centroids(data_list [][]float64, num_clust int)[][]float64{

	// generate init centorids
	rand_keys := (generateRandomNumber(0, len(data_list), num_clust))
	var centroids [][]float64
	for _, v := range rand_keys {
		centroids = append(centroids, data_list[v])
	}
	// fmt.Println(centroids)
	return centroids
}

func get_maxline(data_list [][]float64,centroids [][]float64,ignore_keys []int)([]float64,int){
	cur_dis:=math.Inf(-1)
	max_index:=0
	for k,v :=range data_list{
		pass :=false
		for _,ii :=range ignore_keys{
			if k == ii{
				pass = true
				break
			}
		}
		if pass{
			continue
		}
		var dis float64
		for _,vv :=range centroids{
			dis += DtwDistance(vv,v)
		}
		if dis>cur_dis{
			cur_dis=dis
			max_index=k
		}
	}
	return data_list[max_index],max_index
}

func kpp_centroids(data_list [][]float64, num_clust int)[][]float64{

	// generate init centorids
	rand_keys := (generateRandomNumber(0, len(data_list), 1))
	rand_key := rand_keys[0]

	var centroids [][]float64
	var ignore_keys [] int
	centroids = append(centroids,data_list[rand_key])
	ignore_keys = append(ignore_keys,rand_key)
	for i:=0;i<num_clust-1;i++{
		max_line,max_index := get_maxline(data_list,centroids,ignore_keys)
		centroids = append(centroids,max_line)
		ignore_keys = append(ignore_keys,max_index)
	}
	return centroids
}

func k_means_clust(data_list [][]float64, num_clust int, num_iter int, w int) ([][]float64, map[int][]int) {

	centroids:=kpp_centroids(data_list,num_clust)
	// generate init centorids
	// rand_keys := (generateRandomNumber(0, len(data_list), num_clust))
	// var centroids [][]float64
	// for _, v := range rand_keys {
	// 	centroids = append(centroids, data_list[v])
	// }
	// fmt.Println("xxxxx", rand_keys,centroids)

	counter := 0
	assignments := make(map[int][]int)
	for i := 0; i < num_iter; i++ {
		counter += 1
		// init empty  assignment every iteration
		for k, _ := range centroids {
			assignments[k] = []int{}
		}
		for k, v := range data_list {
			min_dist := math.Inf(1)
			var closest_clust int
			for kk, vv := range centroids {
				if LB_Keogh(v, vv, w) < min_dist {
					cur_dist := DtwDistance(v, vv)
					if cur_dist < min_dist {
						min_dist = cur_dist
						closest_clust = kk
					}
				}
			}
			assignments[closest_clust] = append(assignments[closest_clust], k)
			// fmt.Println("iter times",i,"assignment",assignments)
		}

		for k, v := range assignments {
			var clust_sum []float64
			for _, vv := range v {
				for kkk, vvv := range data_list[vv] {
					if len(clust_sum) < kkk+1 {
						clust_sum = append(clust_sum, 0)
					}
					clust_sum[kkk] += vvv
				}
			}
			for kk, vv := range clust_sum {
				centroids[k][kk] = vv / float64(len(v))
			}

		}
	}
	return centroids, assignments
}

func get_stock_map(stocklist []int)(map[int]int){
	stockmap := make(map[int]int)
	for k,v := range stocklist{
		stockmap[v] = k
	}
	return stockmap
}