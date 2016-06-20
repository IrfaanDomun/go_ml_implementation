package util_ml

import (
	"time"
	"math/rand"
	"sync"
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
	//now := int64(time.Millisecond)
	rand.Seed(time.Now().Unix())
	arr := arr_in.([]float64)
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//func random_indice(arr [] int, i int)
//return a list of index in random order
func  In_random_index(index int ) [] int {
	list := make([] int, index)
	for i,_ := range list{
		list[i] = i
	}
	w := sync.WaitGroup{}
	//now := int64(time.Millisecond)
	w.Add(index)
	rand.Seed(time.Now().Unix())
	arr := list
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}