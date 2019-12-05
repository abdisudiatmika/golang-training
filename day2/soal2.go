package coba

import "fmt"

//GetMinMax sdasdas
func GetMinMax(numbers ...*int) (min int, max int) {
	//fmt.Println(numbers)
	min = *numbers[1]

	for _, value := range numbers {
		if *value > max {
			max = *value
		}
		if *value < min {
			min = *value
		}

	}

	fmt.Println("nilai minimum adalah ", min)
	fmt.Println("nilai maksimum adalah ", max)
	return

}
