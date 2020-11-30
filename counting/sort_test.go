package counting

import (
	"math/rand"
	"sort"
	"testing"
)

type intSlice []int

func (p intSlice) Len() int                    { return len(p) }
func (p intSlice) ComparedField(index int) int { return p[index] }

func TestSort(t *testing.T) {
	arr := make([]int, 10000) // [-100,100)
	for {
		for i := range arr {
			arr[i] = rand.Intn(200) - 100
		}
		if sort.IntsAreSorted(arr) {
			continue
		}
		break
	}
	dst := make([]int, 10000)
	copyElement := func(dstIndex, srcIndex int) {
		dst[dstIndex] = arr[srcIndex]
	}
	Sort(intSlice(arr), copyElement, nil)
	if !sort.IntsAreSorted(dst) {
		t.Error("function Sort did not work correctly")
		return
	}
}

func BenchmarkSort(b *testing.B) {
	arr := make([]int, 1e6)
	for i := range arr {
		arr[i] = rand.Intn(2400)
	}
	dst := make([]int, 1e6)
	copyElement := func(dstIndex, srcIndex int) {
		dst[dstIndex] = arr[srcIndex]
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(intSlice(arr), copyElement, nil)
	}
}
