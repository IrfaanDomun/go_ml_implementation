package Linear_regression

import (
	"math"
	"Machine-learning-Go/go_ml_implementation/util_ml"
//	"fmt"
)

func Predict(alpha, beta,x_i float64) float64 {
	return  beta * x_i + alpha
	//return beta * x_i + alpha
	//
}
//def error(alpha, beta, x_i, y_i):
func Error( alpha, beta, x_i, y_i float64) float64 {
	return  y_i - Predict(alpha, beta, x_i)
	//return y_i - predict(alpha, beta, x_i)
	//
}
//def sum_of_squared_errors(alpha, beta, x, y):
func Sum_of_squared_errors( alpha, beta float64, x,y [] float64) float64{
	res := 0.0
	for i,_ := range x{
		res += math.Pow(Error(alpha,beta,x[i],y[i]),2)
	}
	return res
	//return sum(error(alpha, beta, x_i, y_i) ** 2
	//for x_i, y_i in zip(x, y))
}

//"""given training values for x and y,
//find the least-squares values of alpha and beta"""
func Least_squares_fit(x,y [] float64) (float64,float64) {

	beta  := util_ml.Correlation(x, y) * util_ml.Standard_deviation(y) / util_ml.Standard_deviation(x)
	alpha := util_ml.Mean(y) - beta * util_ml.Mean(x)
	return alpha, beta
}


//def total_sum_of_squares(y):
func Total_sum_of_squares( y []float64) float64{
	//"""the total squared variation of y_i's from their mean"""
	res := 0.0
	for _,j := range util_ml.De_mean(y) {
		res += j*j
	}
	return res
	//return sum(v ** 2 for v in de_mean(y))
}

//def r_squared(alpha, beta, x, y):
//"""the fraction of variation in y captured by the model, which equals
//1 - the fraction of variation in y not captured by the model"""
func R_squared(alpha, beta float64, x,y [] float64) float64{
	return 1.0 - ( Sum_of_squared_errors(alpha,beta,x,y) / Total_sum_of_squares(y))

	//return 1.0 - (sum_of_squared_errors(alpha, beta, x, y) /
	//total_sum_of_squares(y))
}
//def squared_error(x_i, y_i, theta):
func Squared_error( x_i,y_i float64, theta [] float64) float64 {
	alpha := theta[0]
	beta := theta[1]
	//alpha, beta = theta
	res := Error(alpha,beta,x_i,y_i)
	return res * res
	//return error(alpha, beta, x_i, y_i) ** 2
	//
}
//def squared_error_gradient(x_i, y_i, theta):
func Squared_error_gradient(x_i,y_i float64, theta [] float64)([]float64) {
	alpha := theta[0]
	beta := theta[1]
	res := make([] float64,2)

	res[0] = -2.0* Error(alpha,beta,x_i,y_i)
	res[1] = res[0]*x_i

	return res


	//alpha, beta = theta
	//return [-2 * error(alpha, beta, x_i, y_i),       # alpha partial derivative
	//-2 * error(alpha, beta, x_i, y_i) * x_i] # beta partial derivative
	//
}
//if __name__ == "__main__":
//func main() {
//
//	//num_friends_good = [49,41,40,25,21,21,19,19,18,18,16,15,15,15,15,14,14,13,13,13,13,12,12,11,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,8,8,8,8,8,8,8,8,8,8,8,8,8,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
//	num_friends_good := [] float64{49,41,40,25,21,21,19,19,18,18,16,15,15,15,15,14,14,13,13,13,13,12,12,11,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,8,8,8,8,8,8,8,8,8,8,8,8,8,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
//	//daily_minutes_good = [68.77,51.25,52.08,38.36,44.54,57.13,51.4,41.42,31.22,34.76,54.01,38.79,47.59,49.1,27.66,41.03,36.73,48.65,28.12,46.62,35.57,32.98,35,26.07,23.77,39.73,40.57,31.65,31.21,36.32,20.45,21.93,26.02,27.34,23.49,46.94,30.5,33.8,24.23,21.4,27.94,32.24,40.57,25.07,19.42,22.39,18.42,46.96,23.72,26.41,26.97,36.76,40.32,35.02,29.47,30.2,31,38.11,38.18,36.31,21.03,30.86,36.07,28.66,29.08,37.28,15.28,24.17,22.31,30.17,25.53,19.85,35.37,44.6,17.23,13.47,26.33,35.02,32.09,24.81,19.33,28.77,24.26,31.98,25.73,24.86,16.28,34.51,15.23,39.72,40.8,26.06,35.76,34.76,16.13,44.04,18.03,19.65,32.62,35.59,39.43,14.18,35.24,40.13,41.82,35.45,36.07,43.67,24.61,20.9,21.9,18.79,27.61,27.21,26.61,29.77,20.59,27.53,13.82,33.2,25,33.1,36.65,18.63,14.87,22.2,36.81,25.53,24.62,26.25,18.21,28.08,19.42,29.79,32.8,35.99,28.32,27.79,35.88,29.06,36.28,14.1,36.63,37.49,26.9,18.58,38.48,24.48,18.95,33.55,14.24,29.04,32.51,25.63,22.22,19,32.73,15.16,13.9,27.2,32.01,29.27,33,13.74,20.42,27.32,18.23,35.35,28.48,9.08,24.62,20.12,35.26,19.92,31.02,16.49,12.16,30.7,31.22,34.65,13.13,27.51,33.2,31.57,14.1,33.42,17.44,10.12,24.42,9.82,23.39,30.93,15.03,21.67,31.09,33.29,22.61,26.89,23.48,8.38,27.81,32.35,23.84]
//	daily_minutes_good := [] float64{68.77,51.25,52.08,38.36,44.54,57.13,51.4,41.42,31.22,34.76,54.01,38.79,47.59,49.1,27.66,41.03,36.73,48.65,28.12,46.62,35.57,32.98,35,26.07,23.77,39.73,40.57,31.65,31.21,36.32,20.45,21.93,26.02,27.34,23.49,46.94,30.5,33.8,24.23,21.4,27.94,32.24,40.57,25.07,19.42,22.39,18.42,46.96,23.72,26.41,26.97,36.76,40.32,35.02,29.47,30.2,31,38.11,38.18,36.31,21.03,30.86,36.07,28.66,29.08,37.28,15.28,24.17,22.31,30.17,25.53,19.85,35.37,44.6,17.23,13.47,26.33,35.02,32.09,24.81,19.33,28.77,24.26,31.98,25.73,24.86,16.28,34.51,15.23,39.72,40.8,26.06,35.76,34.76,16.13,44.04,18.03,19.65,32.62,35.59,39.43,14.18,35.24,40.13,41.82,35.45,36.07,43.67,24.61,20.9,21.9,18.79,27.61,27.21,26.61,29.77,20.59,27.53,13.82,33.2,25,33.1,36.65,18.63,14.87,22.2,36.81,25.53,24.62,26.25,18.21,28.08,19.42,29.79,32.8,35.99,28.32,27.79,35.88,29.06,36.28,14.1,36.63,37.49,26.9,18.58,38.48,24.48,18.95,33.55,14.24,29.04,32.51,25.63,22.22,19,32.73,15.16,13.9,27.2,32.01,29.27,33,13.74,20.42,27.32,18.23,35.35,28.48,9.08,24.62,20.12,35.26,19.92,31.02,16.49,12.16,30.7,31.22,34.65,13.13,27.51,33.2,31.57,14.1,33.42,17.44,10.12,24.42,9.82,23.39,30.93,15.03,21.67,31.09,33.29,22.61,26.89,23.48,8.38,27.81,32.35,23.84}
//	fmt.Println(num_friends_good)
//	fmt.Println(daily_minutes_good)
//	//
//	alpha, beta := Least_squares_fit(num_friends_good, daily_minutes_good)
//	fmt.Println( "alpha", alpha)
//	fmt.Println( "beta", beta)
//
//	fmt.Println( "r-squared", R_squared(alpha, beta, num_friends_good, daily_minutes_good))
//	//
//	//
//	fmt.Println( "gradient descent:")
//	// choose random value to start
//	rand.Seed(0)
//	//random.seed(0)
//	theta := [] float64{rand.Float64(), rand.Float64()}
//	alpha, beta = Minimize_stochastic(Squared_error,
//	Squared_error_gradient,
//	num_friends_good,
//	daily_minutes_good,
//	theta,
//	0.0001)
//	//fmt.Println( "alpha", alpha
//	//fmt.Println( "beta", beta
//}
