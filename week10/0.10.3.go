package Mymath

//Finds the avarage of a series of numbers
func Average(numbers...float64)float64 {
	var total float64
	for _, v := range numbers {
		total = total + v
	}
	return total / float64(len(numbers))
}