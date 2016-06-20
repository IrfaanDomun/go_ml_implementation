package util_ml

import (
	"time"
	"math/rand"
)

//get a random number between a range
func Random(min, max int) float64 {
	rand.Seed(time.Now().Unix())
	return rand.Float64()*float64(max - min) + float64(min)
}



//""
//"generator that returns the elements of data in random order"
//""
func  In_random_order(arr_in interface{}) interface{} {
	now := int64(time.Millisecond)
	rand.Seed(now)
	arr := arr_in.([]float64)
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}