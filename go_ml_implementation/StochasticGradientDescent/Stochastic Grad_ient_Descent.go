package StochasticGradientDescent

import (

	"math"
)



//from __future__ import division
//from collections import Counter
//from linear_algebra import distance, vector_subtract, scalar_multiply
//import math, random
//
//def sum_of_squares(v):
//"""computes the sum of squared elements in v"""
//func Sum_of_squares(v [] float64){
//	res := make([] float64,len(v))
//	for i,j := range v{
//		res[i] = j*j
//	}
//	return util_ml.Sum(res)
//	//return sum(v_i ** 2 for v_i in v)
//	//
//}

//def difference_quotient(f, x, h):
//return (f(x + h) - f(x)) / h
//
//def plot_estimated_derivative():
//
//def square(x):
//return x * x
//
//def derivative(x):
//return 2 * x
//
//derivative_estimate = lambda x: difference_quotient(square, x, h=0.00001)
//
//# plot to show they're basically the same
//import matplotlib.pyplot as plt
//x = range(-10,10)
//plt.plot(x, map(derivative, x), 'rx')           # red  x
//plt.plot(x, map(derivative_estimate, x), 'b+')  # blue +
//plt.show()                                      # purple *, hopefully
//
//def partial_difference_quotient(f, v, i, h):
//
//# add h to just the i-th element of v
//w = [v_j + (h if j == i else 0)
//for j, v_j in enumerate(v)]
//
//return (f(w) - f(v)) / h
//
//def estimate_gradient(f, v, h=0.00001):
//return [partial_difference_quotient(f, v, i, h)
//for i, _ in enumerate(v)]
//


//def step(v, direction, step_size):
//"""move step_size in the direction from v"""
func Step(v, direction [] float64, step_size float64 ) [] float64 {
	//return [v_i + step_size * direction_i
	//for v_i, direction_i in zip(v, direction)]
	//
	res := [] float64{}
	for i,_ := range v{
		res = append(res, v[i] + step_size*direction[i])
	}
	return res
}
//def sum_of_squares_gradient(v):
func Sum_of_squares_gradient( v []float64) []float64 {
	res := [] float64{}
	for _,j := range v{
		res = append(res, 2*j)
	}
	return res
	//return [2 * v_i for v_i in v]
	//
}



//def minimize_batch(target_fn, gradient_fn, theta_0, tolerance=0.000001):
//"""use gradient descent to find theta that minimizes target function"""
func Minimize_batch( target_fn func([] float64)float64, gradient_fn func([] float64) [] float64, theta_0[] float64, tolerance float64) []float64{

	//
	step_sizes := [] float64{ 100, 10, 1, 0.1, 0.01, 0.001, 0.0001, 0.00001}
	//step_sizes = [100, 10, 1, 0.1, 0.01, 0.001, 0.0001, 0.00001]
	//
	theta := make([] float64, len(theta_0))
	copy(theta,theta_0)
	//can't do the safe part
	value := target_fn(theta)
	//theta = theta_0                           # set theta to initial value
	//target_fn = safe(target_fn)               # safe version of target_fn
	//value = target_fn(theta)                  # value we're minimizing
	//
	for {
		gradient := gradient_fn(theta)
		next_thetas := make([][] float64, len(step_sizes))
		for i,step_size := range step_sizes{
			next_thetas[i] = Step(theta,gradient,- step_size)
		}
		//choose the one that minimizes the error function
		next_value := math.Inf(1)
		next_theta := next_thetas[0]
		for _,temp_next_theta := range next_thetas{
			temp_theta := target_fn(temp_next_theta)
			if   temp_theta < next_value{
				next_value = temp_theta
				next_theta = temp_next_theta
			}
		}
		//next_theta := floats.Min(target_fn(next_thetas))
		//next_value := target_fn((next_theta))
		// stop if we're "converging"
		//fmt.Println(value , next_value , tolerance, theta)
		if math.Abs(value - next_value) < tolerance {
			return theta
		}else {
			theta, value = next_theta, next_value
		}
	}
	//while True:
	//gradient = gradient_fn(theta)
	//next_thetas = [step(theta, gradient, -step_size)
	//for step_size in step_sizes]
	//
	//# choose the one that minimizes the error function
	//next_theta = min(next_thetas, key=target_fn)
	//next_value = target_fn(next_theta)
	//
	//# stop if we're "converging"
	//if abs(value - next_value) < tolerance:
	//return theta
	//else:
	//theta, value = next_theta, next_value
}
//def safe(f):
//"""define a new function that wraps f and return it"""
//def safe_f(*args, **kwargs):
//try:
//return f(*args, **kwargs)
//except:
//return float('inf')         # this means "infinity" in Python
//return safe_f
//
//
//#
//#
//# minimize / maximize batch
//#
//#
//
//def minimize_batch(target_fn, gradient_fn, theta_0, tolerance=0.000001):
//"""use gradient descent to find theta that minimizes target function"""
//
//step_sizes = [100, 10, 1, 0.1, 0.01, 0.001, 0.0001, 0.00001]
//
//theta = theta_0                           # set theta to initial value
//target_fn = safe(target_fn)               # safe version of target_fn
//value = target_fn(theta)                  # value we're minimizing
//
//while True:
//gradient = gradient_fn(theta)
//next_thetas = [step(theta, gradient, -step_size)
//for step_size in step_sizes]
//
//# choose the one that minimizes the error function
//next_theta = min(next_thetas, key=target_fn)
//next_value = target_fn(next_theta)
//
//# stop if we're "converging"
//if abs(value - next_value) < tolerance:
//return theta
//else:
//theta, value = next_theta, next_value
//
//def negate(f):
//"""return a function that for any input x returns -f(x)"""
//return lambda *args, **kwargs: -f(*args, **kwargs)
//
//def negate_all(f):
//"""the same when f returns a list of numbers"""
//return lambda *args, **kwargs: [-y for y in f(*args, **kwargs)]
//
//def maximize_batch(target_fn, gradient_fn, theta_0, tolerance=0.000001):
//return minimize_batch(negate(target_fn),
//negate_all(gradient_fn),
//theta_0,
//tolerance)
//
//#
//# minimize / maximize stochastic
//#
//
//def in_random_order(data):
//"""generator that returns the elements of data in random order"""
//indexes = [i for i, _ in enumerate(data)]  # create a list of indexes
//random.shuffle(indexes)                    # shuffle them
//for i in indexes:                          # return the data in that order
//yield data[i]
//
//def minimize_stochastic(target_fn, gradient_fn, x, y, theta_0, alpha_0=0.01):
//
//data = zip(x, y)
//theta = theta_0                             # initial guess
//alpha = alpha_0                             # initial step size
//min_theta, min_value = None, float("inf")   # the minimum so far
//iterations_with_no_improvement = 0
//
//# if we ever go 100 iterations with no improvement, stop
//while iterations_with_no_improvement < 100:
//value = sum( target_fn(x_i, y_i, theta) for x_i, y_i in data )
//
//if value < min_value:
//# if we've found a new minimum, remember it
//# and go back to the original step size
//min_theta, min_value = theta, value
//iterations_with_no_improvement = 0
//alpha = alpha_0
//else:
//# otherwise we're not improving, so try shrinking the step size
//iterations_with_no_improvement += 1
//alpha *= 0.9
//
//# and take a gradient step for each of the data points
//for x_i, y_i in in_random_order(data):
//gradient_i = gradient_fn(x_i, y_i, theta)
//theta = vector_subtract(theta, scalar_multiply(alpha, gradient_i))
//
//return min_theta
//
//def maximize_stochastic(target_fn, gradient_fn, x, y, theta_0, alpha_0=0.01):
//return minimize_stochastic(negate(target_fn),
//negate_all(gradient_fn),
//x, y, theta_0, alpha_0)
//
//if __name__ == "__main__":
//
//print "using the gradient"
//
//v = [random.randint(-10,10) for i in range(3)]
//
//tolerance = 0.0000001
//
//while True:
//#print v, sum_of_squares(v)
//gradient = sum_of_squares_gradient(v)   # compute the gradient at v
//next_v = step(v, gradient, -0.01)       # take a negative gradient step
//if distance(next_v, v) < tolerance:     # stop if we're converging
//break
//v = next_v                              # continue if we're not
//
//print "minimum v", v
//print "minimum value", sum_of_squares(v)
//print
//
//
//print "using minimize_batch"
//
//v = [random.randint(-10,10) for i in range(3)]
//
//v = minimize_batch(sum_of_squares, sum_of_squares_gradient, v)
//
//print "minimum v", v
//print "minimum value", sum_of_squares(v)


//func minimize_stochastic(target_fn, gradient_fn, x, y, theta_0, alpha_0):
//data = zip(x, y)    theta = theta_0                             # initial guess    alpha = alpha_0                             # initial step size    min_theta, min_value = None, float("inf")   # the minimum so far    iterations_with_no_improvement = 0

//func main() {
//	data := [] float64{2.0,3.0,5.0,7.0}
//	fmt.Println(in_random_order(data)[0])
//}