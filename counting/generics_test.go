package counting

import (
	"math/rand"
	"sort"
	"testing"
)

type elementSlice []element

var _ sort.Interface = elementSlice{}

func (p elementSlice) Len() int           { return len(p) }
func (p elementSlice) Less(i, j int) bool { return p[i].ComparedField < p[j].ComparedField }
func (p elementSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestGenericsSort(t *testing.T) {
	arr := make(elementSlice, 10000) // [-100,100)
	for {
		for i := range arr {
			arr[i].ComparedField = rand.Intn(200) - 100
		}
		if sort.IsSorted(arr) {
			continue
		}
		break
	}
	dst := make(elementSlice, 10000)
	genericsSort(dst, arr, nil)
	if !sort.IsSorted(dst) {
		t.Error("function genericsSort did not work correctly")
		return
	}
}

func BenchmarkGenericsSort(b *testing.B) {
	arr := make(elementSlice, 1e6)
	for i := range arr {
		arr[i].ComparedField = rand.Intn(2400)
	}
	dst := make(elementSlice, 1e6)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genericsSort(dst, arr, nil)
	}
}
