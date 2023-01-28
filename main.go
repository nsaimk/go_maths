package main
 
 import (
 	"fmt"
 	"io/ioutil"
 	"math"
 	"sort"
 	"strconv"
 	"strings"
 )
 
 // This makes the file readable as a graph with an x and y value
 func readFile(fname string) (nums []int, err error) {
 	b, err := ioutil.ReadFile(fname)
 	if err != nil {
 		return nil, err
 	}
 
 	lines := strings.Split(string(b), "\n")
 
 	nums = make([]int, 0, len(lines))
 
 	for _, l := range lines {
 
 		if len(l) == 0 {
 			continue
 		}
 
 		n, err := strconv.Atoi(l)
 		if err != nil {
 			return nil, err
 		}
 		nums = append(nums, n)
 	}
 
 	return nums, nil
 }
 
 func main() {
 	nums, err := readFile("data.txt")
 	if err != nil {
 		panic(err)
 	}
 	
 
 	sum := 0.0
 	Average := 0.0
 	Median := 0.0
 	
 	// this calculates the average
 	for i := 0; i < len(nums); i++ {
 		sum += float64(nums[i])
 		Average = sum / float64(len(nums))
 
 	}
 	// this calculates the median
 	sort.Ints(nums)
 	M := len(nums) / 2
 	if len(nums)%2 == 1 {
 		Median = float64(nums[M])
 	} else {
 		Median = (float64(nums[M-1]) + float64(nums[M])) / 2
 	}
 
 	
 
 	// total := 0
 	// for _, number := range nums {
 	// 	total += number - resultA
 	// }
 	// result = total * total
 	// variance := result / len(nums)
 	total := 0.0
 	for _, number := range nums {
 		total += math.Pow(float64(number)-float64(Average), 2)
 	}
 	variance := total / float64(len(nums))
 	sd := math.Sqrt(variance)
 	resultV := int(math.Round(variance))
 	resultS := int(math.Round(sd))
 	resultA := int(math.Round(Average))
 	resultM := int(math.Round(Median))
 
 	
 
 
 	fmt.Println("Average:", resultA)
 	fmt.Println("Median:", resultM)
 	fmt.Println("Variance:", resultV)
 	fmt.Println("Standard Deviation:", resultS)
 	
 } 



//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
