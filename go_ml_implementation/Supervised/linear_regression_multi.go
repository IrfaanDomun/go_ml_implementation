package main

import (
	"fmt"
)
//compute the dot product of two vector
func  dot(x_i interface{} , beta  interface{} )  interface{} {

	switch  x_i.(type) {
	case []uint64:
		res := uint64(0)
		for i := range x_i.([]uint64) {
			res += x_i.([]uint64)[i] * beta.([]uint64)[i]
		}
		return res
	case []uint8:
		res := uint8(0)
		for i := range x_i.([]uint8) {
			res += x_i.([]uint8)[i] * beta.([]uint8)[i]
		}
		return res
	case []uint16:
		res := uint16(0)
		for i := range x_i.([]uint16) {
			res += x_i.([]uint16)[i] * beta.([]uint16)[i]
		}
		return res

	case []uint32:
		res := uint32(0)
		for i := range x_i.([]uint32) {
			res += x_i.([]uint32)[i] * beta.([]uint32)[i]
		}
		return res

	case []int64:
		res := int64(0)
		for i := range x_i.([]int64) {
			res += x_i.([]int64)[i] * beta.([]int64)[i]
		}
		return res

	case []int32:
		res := int32(0)
		for i := range x_i.([]int32) {
			res += x_i.([]int32)[i] * beta.([]int32)[i]
		}
		return res

	case []int16:
		res := int16(0)
		for i := range x_i.([]int16) {
			res += x_i.([]int16)[i] * beta.([]int16)[i]
		}
		return res

	case []int8:
		res := int8(0)
		for i := range x_i.([]int8) {
			res += x_i.([]int8)[i] * beta.([]int8)[i]
		}
		return res

	case []int:
		res := int(0)
		for i := range x_i.([]int) {
			res += x_i.([]int)[i] * beta.([]int)[i]
		}
		return res

	case []float32:
		res := float32(0)
		for i := range x_i.([]float32) {
			res += x_i.([]float32)[i] * beta.([]float32)[i]
		}
		return res

	case []float64:
		res := float64(0)
		for i := range x_i.([]float64) {
			res += x_i.([]float64)[i] * beta.([]float64)[i]
		}
		return res

	default:
		//_ = t
		panic(0)
		//s := x_i.(t)
		//for i := range s {
		//	fmt.Println(i, s[i], t)
		//}
	}

	//t := reflect.ValueOf(beta)
	//fmt.Println(reflect.TypeOf(x_i).Kind(),s.Len())
	//for i := 0; i < s.Len(); i++ {
	//	fmt.Println(s.Index(i), t.Index(i))
	//}
	return x_i
}

//return the  dot product of two vector
func  predict(x_i interface{} , beta  interface{} )  interface{} {
//assumes that the first element of each x_i is 1
	return dot(x_i,beta)
}
//compute the mean error
func  error(x_i,  beta, y_i interface{})interface{}{
	switch x_i.(type){
	case []uint8:
		return y_i.(uint8) - dot(x_i, beta).(uint8)

	case []uint16:
		return y_i.(uint16) - dot(x_i, beta).(uint16)

	case []uint32:
		return y_i.(uint32) - dot(x_i, beta).(uint32)

	case []uint64:
		return y_i.(uint64) - dot(x_i, beta).(uint64)

	case []int8:
		return y_i.(int8) - dot(x_i, beta).(int8)

	case []int16:
		return y_i.(int16) - dot(x_i, beta).(int16)

	case []int32:
		return y_i.(int32) - dot(x_i, beta).(int32)

	case []int64:
		return y_i.(int64) - dot(x_i, beta).(int64)

	case []float32:
		return y_i.(float32) - dot(x_i, beta).(float32)

	case []float64:
		return y_i.(float64) - dot(x_i, beta).(float64)

	case []int:
		return y_i.(int) - dot(x_i, beta).(int)

	default:
		panic(0)
	}
}

//compute the mean square error
func  squared_error(x_i, beta, y_i interface{})interface{}{
	switch x_i.(type){
	case []uint8:
		a:= error(x_i, beta, y_i).(uint8)
		return  a*a
	case []uint16:
		a:= error(x_i, beta, y_i).(uint16)
		return  a*a
	case []uint32:
		a:= error(x_i, beta, y_i).(uint32)
		return  a*a
	case []uint64:
		a:= error(x_i, beta, y_i).(uint64)
		return  a*a
	case []int8:
		a:= error(x_i, beta, y_i).(int8)
		return  a*a
	case []int16:
		a:= error(x_i, beta, y_i).(int16)
		return  a*a
	case []int32:
		a:= error(x_i, beta, y_i).(int32)
		return  a*a
	case []int64:
		a:= error(x_i, beta, y_i).(int64)
		return  a*a
	case []float32:
		a:= error(x_i, beta, y_i).(float32)
		return  a*a
	case []float64:
		a:= error(x_i, beta, y_i).(float64)
		return  a*a
	case []int:
		a:= error(x_i, beta, y_i).(int)
		return  a*a
	default:
		panic(0)
	}
}

//the gradient (with respect to beta)    corresponding to the ith squared error term
//For uintX return int2X, uint64 -> float64
func squared_error_gradient(x_i, beta, y_i interface{}) interface{}{
	switch x_i.(type){
	case []uint8:
		res := []int16 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(uint8)
		for i,j:= range x_i.([]uint8){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int16(v) )
		}
		return res
	case []uint16:
		res := []int32 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(uint16)
		for i,j:= range x_i.([]uint16){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int32(v) )
		}
		return res
	case []uint32:
		res := []int64 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(uint32)
		for i,j:= range x_i.([]uint32){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int64(v) )
		}
		return res
	case []uint64:
		res := []float64 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(uint64)
		for i,j:= range x_i.([]uint64){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * float64(v) )
		}
		return res
	case []int8:
		res := []int8 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(int8)
		for i,j:= range x_i.([]int8){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int8(v) )
		}
		return res
	case []int16:
		res := []int16 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(int16)
		for i,j:= range x_i.([]int16){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int16(v) )
		}
		return res
	case []int32:
		res := []int32 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(int32)
		for i,j:= range x_i.([]int32){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int32(v) )
		}
		return res
	case []int64:
		res := []int64 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(int64)
		for i,j:= range x_i.([]int64){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int64(v) )
		}
		return res
	case []float32:
		res := []float32 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(float32)
		for i,j:= range x_i.([]float32){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * float32(v) )
		}
		return res
	case []float64:
		res := []float64 {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(float64)
		for i,j:= range x_i.([]float64){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * float64(v) )
		}
		return res
	case []int:
		res := []int {}
		//fmt.Println(error(x_i, y_i, beta))
		err := error(x_i, beta , y_i).(int)
		for i,j:= range x_i.([]int){
			fmt.Println(i,j)
			v :=  j * err
			res = append(res, -2.0 * int(v) )
		}
		return res

	default:
		panic(0)
	}

	//return [-2 * x_ij * error(x_i, y_i, beta)            for x_ij in x_i
}


//find the optimal beta using stochastic gradient descent
//func estimate_beta(x, y interface{}){
//	beta_initial = [random.random() for x_i in x[0]]
//	return minimize_stochastic(squared_error, squared_error_gradient, x, y,beta_initial,0.001)}

func main() {
	a := [] float64 {5,6}

	b := [] float64 {6,7}
	c := predict(a,b)
	e := 100.0
	d := squared_error(a,b,e)
	squared_error_gradient(a,b,e)
	fmt.Println(d,c)

	//beta := estimate_beta(x, daily_minutes_good) # [30.63, 0.972, -1.868, 0.911]
}