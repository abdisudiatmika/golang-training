package coba

import "fmt"

func prime(numbers ...int) int {
	var faktor int
	for i, numbers := range numbers {
		if numbers%i == 0 {
			faktor += faktor

		} else {
			faktor = faktor

		}
	}
	if faktor == 2 {
		fmt.Println(numbers)
	} else {
		fmt.Println(numbers)
	}

}

// //GetPrima sdadadadadadasd
func GetPrima(X int) int {

	for i := 1; i <= N; i++ {

	}

}
