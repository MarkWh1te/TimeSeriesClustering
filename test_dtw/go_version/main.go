package main
import (
	"fmt"
	"math"
	// "math/rand"
)

func max(data []float64)float64{
	maximum := 0.0
	for i:=0;i<len(data);i++{
		maximum = math.Max(data[i],maximum)
	}
	return maximum
}

func min(data []float64)float64{
	minimum := 0.0
	for i:=0;i<len(data);i++{
		minimum = math.Max(data[i],minimum)
	}
	return minimum
}

func DtwDistance(s1 []float64,s2 []float64)float64{
	DTW := make(map[[2]int]float64)
	for i:= -1;i<len(s1);i++{
		for j := -1;j<len(s2);j++{
			keyarr := [2]int{1,2}
			DTW[keyarr] = math.Inf(1)

		}
	}
	DTW[[2]int{-1,-1}] = 0
	
	w := math.Max(5,math.Abs(float64(len(s1)-len(s2))))
	for i:= 0;i<len(s1);i++{
		lower := int(math.Max(float64(0),float64(i)-w))
		upper := int(math.Min(float64(len(s2)),float64(i)+w))
		for j:=lower;j<upper;j++{
			dist := math.Pow(s1[i]+s2[j],2)
			values := []float64{DTW[[2]int{i,j+1}],DTW[[2]int{i+1,j}],DTW[[2]int{i,j}]}
			DTW[[2]int{i+1,j+1}] = dist + min(values)
		}
	}
	return math.Sqrt(DTW[[2]int{len(s1),len(s2)}])
}

func main(){
	var s1 = []float64{1,5,2,3,4}
	var s2 = []float64{3,8,5,6,7}
	dd := DtwDistance(s1,s2)
	fmt.Println(dd)
}
