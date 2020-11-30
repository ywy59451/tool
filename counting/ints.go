package counting

func Ints(arr []int, rng *Range) {
	if len(arr) < 2 {
		return
	}
	var min, max int
	if rng == nil {
		min, max = getIntsRange(arr)
	} else {
		min, max = rng.Min, rng.Max
	}
	count := make([]int, max-min+1)
	for _, v := range arr {
		count[v-min]++
	}
	j := 0
	for i, c := range count {
		v := i + min
		for c > 0 {
			arr[j] = v
			j++
			c--
		}
	}
}

func getIntsRange(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for _, v := range arr {
		switch {
		case v > max:
			max = v
		case v < min:
			min = v
		}
	}
	return
}
