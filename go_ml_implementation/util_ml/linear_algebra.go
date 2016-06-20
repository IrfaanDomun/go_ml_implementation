package util_ml

import (
	"github.com/gonum/floats"
	"math"
)

////compute the dot product of two vector
//func  Dot(x_i float64 , beta  float64 )  float64{
//	res := float64(0)
//	for i := range x_i {
//		res += x_i[i] * beta[i]
//	}
//	return res
//}

//def vector_subtract(v, w):
//"""subtracts two vectors componentwise"""
func Vector_subtract(v,w [] float64) []float64{
	res := make([] float64,len(v))
	for i,_ := range v{
		res[i] = v[i] - w[i]
	}
	return res
	//return [v_i - w_i for v_i, w_i in zip(v,w)]
}

//def squared_distance(v, w):
func Squared_distance(v,w [] float64) float64{
	//return sum_of_squares(vector_subtract(v, w))
	return  Sum_of_squares(Vector_subtract(v,w))
}
//def distance(v, w):
func Distance(v, w [] float64 ) float64{
	return math.Sqrt( Squared_distance(v,w))
	//return math.sqrt(squared_distance(v, w))
}
//def sum_of_squares(v):
func Sum_of_squares(x [] float64) float64{
	return floats.Dot(x,x)
}

func Negate_vector(v [] float64) [] float64{
	res := make([] float64, len(v))
	for i,j := range v{
		res[i] = -j
	}
	return res
}
//"""v_1 * v_1 + ... + v_n * v_n"""
//return dot(v, v)
