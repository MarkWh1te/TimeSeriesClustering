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

func readcsv(path string) [][]string {
	// read csv part
	file, err := os.Open(path)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		// fmt.Println("Error:", err)
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

	// sanity check, display to standard output
	// for _, each := range rawCSVdata {
	// fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	// }
	return rawCSVdata
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

func get_centroid() ([][]float64, map[int][]int) {

	datas := make(map[string][]float64)
	datas["a"] = []float64{1.11, 2.22, 3.33, 4.44, 5.55}
	datas["b"] = []float64{2.34, 4.56, 5.12, 6.04, 5.55}
	datas["c"] = []float64{1.55, 2.21, 3.13, 4.24, 5.55}
	datas["d"] = []float64{2.34, 4.56, 5.12, 6.04, 5.55}
	datas["e"] = []float64{11.11, 23.22, 32.33, 41.44, 15.55}
	// fmt.Println(datas)

	raw := readcsv("2016-07-012017-07-01.csv")
	csv_data := dataclean(raw)
	// fmt.Println(csv_data)
	datas = csv_data

	var keys = sorted_keys(datas)
	// fmt.Println(keys)
	var data_list [][]float64
	for _, v := range keys {
		data_list = append(data_list, to_zero(datas[v]))
	}
	// fmt.Println(data_list)
	centroids, assignments := k_means_clust(data_list, 5, 30, 3)
	// fmt.Println("okkkkkkkkkkk")
	// fmt.Println(centroids)
	// fmt.Println(assignments)
	return centroids, assignments
}

func to_zero(arr []float64) []float64 {
	var tmp []float64
	for _, v := range arr {
		tmp = append(tmp, Round(v-arr[0], 2))
	}
	// fmt.Println(tmp)
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
			// fmt.Println(i,j,s1[i]+s2[j],dist)
			values := []float64{DTW[[2]int{i, j - 1}], DTW[[2]int{i - 1, j}], DTW[[2]int{i - 1, j - 1}]}
			// DTW[[2]int{i + 1, j + 1}] = dist + min(values)
			DTW[[2]int{i, j}] = dist + min(values)
			// fmt.Println(i,j,dist,values,min(values))
			// fmt.Println(DTW[[2]int{i, j}])
		}
	}
	return math.Sqrt(DTW[[2]int{len(s1) - 1, len(s2) - 1}])
}

func LB_Keogh(s1 []float64, s2 []float64, r int) float64 {
	// fmt.Println(len(s1),len(s2),r)
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
func k_means_clust(data_list [][]float64, num_clust int, num_iter int, w int) ([][]float64, map[int][]int) {

	// generate init centorids
	rand_keys := (generateRandomNumber(0, len(data_list), num_clust))
	var centroids [][]float64
	for _, v := range rand_keys {
		centroids = append(centroids, data_list[v])
	}
	// fmt.Println("xxxxx", rand_keys,centroids)

	counter := 0
	assignments := make(map[int][]int)
	for i := 0; i < num_iter; i++ {
		counter += 1
		// fmt.Println("counter times", counter, "centroids", centroids)
		// init empty  assignment every iteration
		for k, _ := range centroids {
			assignments[k] = []int{}
		}
		for k, v := range data_list {
			// fmt.Println("pppppppppppppppp")
			min_dist := math.Inf(1)
			var closest_clust int
			// fmt.Println(closest_clust)
			for kk, vv := range centroids {
				// fmt.Println(LB_Keogh(v,vv,w))
				if LB_Keogh(v, vv, w) < min_dist {
					cur_dist := DtwDistance(v, vv)
					// fmt.Println(cur_dist)
					if cur_dist < min_dist {
						min_dist = cur_dist
						closest_clust = kk
					}
				}
			}
			// fmt.Println(closest_clust)
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
