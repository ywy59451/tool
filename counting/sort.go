package counting

type Interface interface {
	Len() int
	ComparedField(index int) int
}

type Range struct {
	Min, Max int
}

func Sort(data Interface, copyElement func(dstIndex, srcIndex int), rng *Range) {
	dataLen := data.Len()
	if dataLen < 0 {
		panic("data.Len() < 0")
	}
	if dataLen == 0 {
		return
	}
	if dataLen == 1 {
		copyElement(0, 0)
		return
	}
	var min, max int
	if rng == nil {
		min, max = getRange(data)
	} else {
		min, max = rng.Min, rng.Max
	}
	count := make([]int, max-min+1)
	for i := 0; i < dataLen; i++ {
		count[data.ComparedField(i)-min]++
	}
	total := 0
	for i, c := range count {
		count[i] = total
		total += c
	}
	for i := 0; i < dataLen; i++ {
		countKey := data.ComparedField(i) - min
		copyElement(count[countKey], i)
		count[countKey]++
	}
}

func getRange(data Interface) (min, max int) {
	min = data.ComparedField(0)
	max = min
	for i, l := 1, data.Len(); i < l; i++ {
		switch n := data.ComparedField(i); {
		case n > max:
			max = n
		case n < min:
			min = n
		}
	}
	return
}
