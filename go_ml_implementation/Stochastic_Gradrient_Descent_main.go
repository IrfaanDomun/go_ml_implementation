package main

import (
	"fmt"
	"Machine-learning-Go/go_ml_implementation/util_ml"
	"Machine-learning-Go/go_ml_implementation/StochasticGradientDescent"
)
//
//if __name__ == "__main__":
//
func main(){

	fmt.Println( "using the gradient")
	//

	v := []float64{}
	for i:=0;i<3;i++{
		v = append(v, util_ml.Random(-10,10))
	}

	//
	tolerance := 0.0000001

	//
	//while True:
	//#print v, sum_of_squares(v)
	//gradient = sum_of_squares_gradient(v)   # compute the gradient at v
	//next_v = step(v, gradient, -0.01)       # take a negative gradient step
	//if distance(next_v, v) < tolerance:     # stop if we're converging
	//break
	//v = next_v                              # continue if we're not
	for {
		gradient := StochasticGradientDescent.Sum_of_squares_gradient(v)// compute the gradient at v
		next_v :=  StochasticGradientDescent.Step(v, gradient, -0.01)
		if util_ml.Distance(next_v,v) < tolerance{
			break
		}
		v = next_v
	}
	//
	fmt.Println("minimum v", v)
	fmt.Println( "minimum value", util_ml.Sum_of_squares(v))
	fmt.Println("\n")
	//
	//
	fmt.Println( "using minimize_batch")
	//
	v = []float64{}
	for i:=0;i<3;i++{
		v = append(v, util_ml.Random(-10,10))
	}
	//v = [random.randint(-10,10) for i in range(3)]
	//
	v = StochasticGradientDescent.Minimize_batch(util_ml.Sum_of_squares, StochasticGradientDescent.Sum_of_squares_gradient, v, tolerance)

	//v = minimize_batch(sum_of_squares, sum_of_squares_gradient, v)
	//
	fmt.Println( "minimum v", v)
	fmt.Println( "minimum value", util_ml.Sum_of_squares(v))

	//print "minimum v", v
	//print "minimum value", sum_of_squares(v)
	//
}