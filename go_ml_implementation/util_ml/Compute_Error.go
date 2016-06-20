package util_ml

//import "math"
//
//func Squared_error( x_i,y_i, float64, theta [] float64) float64 {
//	alpha := theta[0]
//	beta := theta[1]
//	//alpha, beta = theta
//	res := Error(alpha,beta,x_i,y_i)
//	return res * res
//	//return error(alpha, beta, x_i, y_i) ** 2
//	//
//}
////def error(alpha, beta, x_i, y_i):
//func Error( alpha, beta, x_i, y_i float64) float64 {
//	return  y_i - Predict(alpha, beta, x_i)
//	//return y_i - predict(alpha, beta, x_i)
//	//
//}
////def sum_of_squared_errors(alpha, beta, x, y):
//func Sum_of_squared_errors( alpha, beta float64, x,y [] float64) {
//	res := 0.0
//	for i,_ := range x{
//		res += math.Pow(Error(alpha,beta,x[i],y[i]),2)
//	}
//	return res
//	//return sum(error(alpha, beta, x_i, y_i) ** 2
//	//for x_i, y_i in zip(x, y))
//}