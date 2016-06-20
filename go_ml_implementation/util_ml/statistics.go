package util_ml
import (
	"sort"
	"math"
	"github.com/gonum/floats"

)

//from __future__ import division
//from collections import Counter
//from linear_algebra import sum_of_squares, dot
//import math
//
//num_friends = [100,49,41,40,25,21,21,19,19,18,18,16,15,15,15,15,14,14,13,13,13,13,12,12,11,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,8,8,8,8,8,8,8,8,8,8,8,8,8,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
//
//def make_friend_counts_histogram(plt):
//friend_counts = Counter(num_friends)
//xs = range(101)
//ys = [friend_counts[x] for x in xs]
//plt.bar(xs, ys)
//plt.axis([0,101,0,25])
//plt.title("Histogram of Friend Counts")
//plt.xlabel("# of friends")
//plt.ylabel("# of people")
//plt.show()
//
//num_points = len(num_friends)               # 204
//
//largest_value = max(num_friends)            # 100
//smallest_value = min(num_friends)           # 1
//
//sorted_values = sorted(num_friends)
//smallest_value = sorted_values[0]           # 1
//second_smallest_value = sorted_values[1]    # 1
//second_largest_value = sorted_values[-2]    # 49
//
//# this isn't right if you don't from __future__ import division
//return the mean
func Mean(x [] float64)float64 {
	return Sum(x) / float64(len(x))

	//def mean(x):
	//return sum(x) / len(x)
}//
//"""finds the 'middle-most' value of v"""
func Median (x [] float64) float64 {
//def median(v):
//n = len(v)
	n:= len(x)
	sorted_x := make([] float64, len(x))
	copy(sorted_x,x)
	sort.Float64s(sorted_x)
	mid_point := n/2
	//# if odd, return the middle value
	if n%2 ==1{
		return x[mid_point]
	}else {
		//# if even, return the average of the middle values
		lo := mid_point - 1
		hi := mid_point
		return (sorted_x[lo]+ sorted_x[hi])/2.0

	}
//sorted_v = sorted(v)
//midpoint = n // 2
//
//if n % 2 == 1:
//# if odd, return the middle value

//return sorted_v[midpoint]
//else:
//# if even, return the average of the middle values
//lo = midpoint - 1
//hi = midpoint
//return (sorted_v[lo] + sorted_v[hi]) / 2
//
}

//"""returns the pth-percentile value in x"""
func Quantile ( x [] float64, p float64) float64 {
	p_index := int(p * float64(len(x)) )
	sorted_x := make([] float64, len(x))
	copy(sorted_x,x)
	sort.Float64s(sorted_x)
	return sorted_x[p_index]

	//def quantile(x, p):
	//p_index = int(p * len(x))
	//return sorted(x)[p_index]
	//
}


//def mode(x):
//"""returns a list, might be more than one mode"""
func Mode(x [] float64) [] float64 {
	counts := make(map[float64]float64,len(x))
	counts_values := [] float64{}
	for _,j := range x{
		if _, ok := counts[j]; !ok {
			counts[j] = 0
		}else{
			counts[j] +=1
		}
	}

	for _,j := range counts{
		counts_values = append(counts_values, j)
	}

	max_count := floats.Max(counts_values)
	res := [] float64 {}
	for i,j := range counts{
		if j == max_count{
			res = append(res, i)
		}
	}
	sort.Float64s(res)
	return res
	//counts = Counter(x)
	//max_count = max(counts.values())
	//return [x_i for x_i, count in counts.iteritems()
	//if count == max_count]
	//
}
//# "range" already means something in Python, so we'll use a different name
//def data_range(x):
func Data_range(x [] float64) float64 {
	return floats.Max(x) - floats.Min(x)
	//return max(x) - min(x)
}

//return the sum of the slice
func Sum(x [] float64)float64{
	res :=  0.0
	for _,j := range x {
		res+=j
	}
	return res
}



//"""translate x by subtracting its mean (so the result has mean 0)"""
func De_mean( x [] float64) [] float64 {
	x_bar := Mean(x)
	res := [] float64{}
	for _,j := range x {
		res = append(res, j - x_bar)
	}
	//fmt.Println("chips",x_bar,res,"\n",x,"plus chips")
	return res
	//def de_mean(x):
	//"""translate x by subtracting its mean (so the result has mean 0)"""
	//x_bar = mean(x)
	//return [x_i - x_bar for x_i in x]
}//


//def variance(x):
//"""assumes x has at least two elements"""
func Variance( x []float64) float64 {
	//n = len(x)
	n := len(x)
	deviations := De_mean(x)
	return Sum_of_squares(deviations)/float64(n-1)
	//deviations = de_mean(x)
	//return sum_of_squares(deviations) / (n - 1)
	//
}

//def standard_deviation(x):
func Standard_deviation(x [] float64) float64{
	return math.Sqrt(Variance(x))
	//return math.sqrt(variance(x))
}//
//def interquartile_range(x):
func Interquartile_range(x [] float64) float64 {
	return Quantile(x,0.75) - Quantile(x,0.25)
	//return quantile(x, 0.75) - quantile(x, 0.25)
	//
}
//####
//#
//# CORRELATION
//#
//#####
//
//daily_minutes = [1,68.77,51.25,52.08,38.36,44.54,57.13,51.4,41.42,31.22,34.76,54.01,38.79,47.59,49.1,27.66,41.03,36.73,48.65,28.12,46.62,35.57,32.98,35,26.07,23.77,39.73,40.57,31.65,31.21,36.32,20.45,21.93,26.02,27.34,23.49,46.94,30.5,33.8,24.23,21.4,27.94,32.24,40.57,25.07,19.42,22.39,18.42,46.96,23.72,26.41,26.97,36.76,40.32,35.02,29.47,30.2,31,38.11,38.18,36.31,21.03,30.86,36.07,28.66,29.08,37.28,15.28,24.17,22.31,30.17,25.53,19.85,35.37,44.6,17.23,13.47,26.33,35.02,32.09,24.81,19.33,28.77,24.26,31.98,25.73,24.86,16.28,34.51,15.23,39.72,40.8,26.06,35.76,34.76,16.13,44.04,18.03,19.65,32.62,35.59,39.43,14.18,35.24,40.13,41.82,35.45,36.07,43.67,24.61,20.9,21.9,18.79,27.61,27.21,26.61,29.77,20.59,27.53,13.82,33.2,25,33.1,36.65,18.63,14.87,22.2,36.81,25.53,24.62,26.25,18.21,28.08,19.42,29.79,32.8,35.99,28.32,27.79,35.88,29.06,36.28,14.1,36.63,37.49,26.9,18.58,38.48,24.48,18.95,33.55,14.24,29.04,32.51,25.63,22.22,19,32.73,15.16,13.9,27.2,32.01,29.27,33,13.74,20.42,27.32,18.23,35.35,28.48,9.08,24.62,20.12,35.26,19.92,31.02,16.49,12.16,30.7,31.22,34.65,13.13,27.51,33.2,31.57,14.1,33.42,17.44,10.12,24.42,9.82,23.39,30.93,15.03,21.67,31.09,33.29,22.61,26.89,23.48,8.38,27.81,32.35,23.84]
//
func Covariance(x,y [] float64) float64 {
	//def covariance(x, y):
	//n = len(x)
	n := len(x)
	//fmt.Println("x",x)
	//fmt.Println("y",y)
	//fmt.Println(floats.Dot(De_mean(x), De_mean(y)), (n - 1),De_mean(x), De_mean(y))

	return floats.Dot(De_mean(x),De_mean(y)) /float64(n-1)
	//return dot(de_mean(x), de_mean(y)) / (n - 1)
	//
}


//def correlation(x, y):
func Correlation(x,y [] float64) float64 {
	stdev_x := Standard_deviation(x)
	stdev_y := Standard_deviation(y)
	if stdev_x > 0.0 && stdev_y >0.0{
		return  Covariance(x,y)/ stdev_x / stdev_y
	}else {
		return 0.0 //if no variation, correlation is zero
	}
	//stdev_x = standard_deviation(x)
	//stdev_y = standard_deviation(y)
	//if stdev_x > 0 and stdev_y > 0:
	//return covariance(x, y) / stdev_x / stdev_y
	//else:
	//return 0 # if no variation, correlation is zero
	//
}


	//outlier = num_friends.index(100) # index of outlier
	//
	//num_friends_good = [x
	//for i, x in enumerate(num_friends)
	//if i != outlier]
	//
	//daily_minutes_good = [x
	//for i, x in enumerate(daily_minutes)
	//if i != outlier]
	//
	//
	//
	//if __name__ == "__main__":
	//
	//print "num_points", len(num_friends)
	//print "largest value", max(num_friends)
	//print "smallest value", min(num_friends)
	//print "second_smallest_value", sorted_values[1]
	//print "second_largest_value", sorted_values[-2]
	//print "mean(num_friends)", mean(num_friends)
	//print "median(num_friends)", median(num_friends)
	//print "quantile(num_friends, 0.10)", quantile(num_friends, 0.10)
	//print "quantile(num_friends, 0.25)", quantile(num_friends, 0.25)
	//print "quantile(num_friends, 0.75)", quantile(num_friends, 0.75)
	//print "quantile(num_friends, 0.90)", quantile(num_friends, 0.90)
	//print "mode(num_friends)", mode(num_friends)
	//print "data_range(num_friends)", data_range(num_friends)
	//print "variance(num_friends)", variance(num_friends)
	//print "standard_deviation(num_friends)", standard_deviation(num_friends)
	//print "interquartile_range(num_friends)", interquartile_range(num_friends)
	//
	//print "covariance(num_friends, daily_minutes)", covariance(num_friends, daily_minutes)
	//print "correlation(num_friends, daily_minutes)", correlation(num_friends, daily_minutes)
	//print "correlation(num_friends_good, daily_minutes_good)", correlation(num_friends_good, daily_minutes_good)
	//
