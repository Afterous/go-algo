package MergeSort

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

/* Non-averaged version
func TestMergeSortTime(t *testing.T){
	for i := 0; i <= 100000; i=i+100 {
		s2 := generateSlice(i)
		t2 := time.Now()
		MergeSort(s2,0,len(s2)-1)
		elapsed := time.Since(t2)
		fmt.Printf("%d,%d\n",len(s2),elapsed.Nanoseconds())
		//fmt.Printf("MergeSort: Slice length of: %d, took: %d ns\n",len(s2),elapsed.Nanoseconds())
	}
}*/

func TestMergeSort(t *testing.T){
	a := []int{20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1,0}
	out := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	MergeSort(a,0,len(a)-1)
	if reflect.DeepEqual(out,a) {
	} else {
		t.Errorf("No match")
	}
}

func BenchmarkMSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(2)
		MergeSort(s,0,len(s)-1)
	}
}

func BenchmarkMSort10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(10)
		MergeSort(s,0,len(s)-1)
	}
}

func BenchmarkMSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(100)
		MergeSort(s,0,len(s)-1)
	}
}

func BenchmarkMSort1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(1000)
		MergeSort(s,0,len(s)-1)
	}
}

func BenchmarkMSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(10000)
		MergeSort(s,0,len(s)-1)
	}
}

func BenchmarkMSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(100000)
		MergeSort(s,0,len(s)-1)

	}
}

func BenchmarkMSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(1000000)
		MergeSort(s,0,len(s)-1)

	}
}