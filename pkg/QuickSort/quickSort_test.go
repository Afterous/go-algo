package QuickSort

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(n))
	}
	return s
}

func TestQuickSort(t *testing.T){
	a := []int{12,8,32,19,92,89,90,45,67,55,72}
	out := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	Quicksort(a,0,len(a)-1)
	if reflect.DeepEqual(out,a) {
	} else {
		t.Errorf("No match")
	}
}
// Non-averaged version
func TestQuickSortTime(t *testing.T){
	for i := 0; i <= 100000; i=i+100 {
		s2 := generateSlice(i)
		//t2 := time.Now()
		Quicksort(s2,0,len(s2)-1)
		//elapsed := time.Since(t2)
		//fmt.Printf("%d,%d\n",len(s2),elapsed.Nanoseconds())
		//fmt.Printf("QuickSort: Slice length of: %d, took: %d ns\n",len(s2),elapsed.Nanoseconds())
	}
}

func BenchmarkQSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(2)
		Quicksort(s,0,len(s)-1)
	}
}

func BenchmarkQSort10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(10)
		Quicksort(s,0,len(s)-1)
	}
}

func BenchmarkQSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(100)
		Quicksort(s,0,len(s)-1)
	}
}

func BenchmarkQSort1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(1000)
		Quicksort(s,0,len(s)-1)

	}
}

func BenchmarkSQort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(10000)
		Quicksort(s,0,len(s)-1)
	}
}

func BenchmarkQSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(100000)
		Quicksort(s,0,len(s)-1)
	}
}

func BenchmarkQSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(1000000)
		Quicksort(s,0,len(s)-1)
	}
}