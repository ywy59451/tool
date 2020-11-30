package counting

// if you care about performance, you can copy this code into your project and modify element to meet your own needs

type element struct {
	ComparedField int
	Others        [1]byte
}

func genericsSort(dst, src []element, rng *Range) {
	if len(dst) < len(src) {
		panic("len(dst) < len(src)")
	}
	if len(src) < 0 {
		panic("len(src) < 0")
	}
	if len(src) == 0 {
		return
	}
	if len(src) == 1 {
		dst[0] = src[0]
		return
	}
	var min, max int
	if rng == nil {
		min, max = getGenericsRange(src)
	} else {
		min, max = rng.Min, rng.Max
	}
	count := make([]int, max-min+1)
	for _, v := range src {
		count[v.ComparedField-min]++
	}
	total := 0
	for i, c := range count {
		count[i] = total
		total += c
	}
	for _, v := range src {
		countKey := v.ComparedField - min
		dst[count[countKey]] = v
		count[countKey]++
	}
}

func getGenericsRange(arr []element) (min, max int) {
	min, max = arr[0].ComparedField, arr[0].ComparedField
	for _, v := range arr {
		switch n := v.ComparedField; {
		case n > max:
			max = n
		case n < min:
			min = n
		}
	}
	return
}
