package array

func Sum(numbers []int) int {
	var sum int

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int 
	for _, numbers := range numbersToSum {
		sums  = append(sums, Sum(numbers)) 
	}

	return sums
}


func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int 

	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}


	}

	return sums
}
