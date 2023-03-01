package expenses

func Average(expenses... float32) float32 {
	return Sum(expenses...) / float32(len(expenses))
}

func Sum(expenses... float32) float32 {
	var sum float32

	for _, exp := range expenses {
		sum += exp
	}

	return sum
}

func Max(expenses... float32) float32 {
	var max float32

	for _, exp := range expenses {
		if exp > max {
			max = exp
		}
	}
	
	return max
}

func Min(expenses... float32) float32 {
	
	if len(expenses) == 0 {
		return 0
	}

	var min float32 = expenses[0]

	for _, exp := range expenses {
		if exp < min {
			min = exp
		}
	}
	
	return min
}