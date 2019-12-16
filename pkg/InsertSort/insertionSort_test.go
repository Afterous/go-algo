package InsertSort

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

func TestInsertionSort(t *testing.T){
	a := []int{20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1,0}
	out := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	InsertionSort(a)
	if reflect.DeepEqual(out,a) {
	} else {
		t.Errorf("No match")
	}
}

/* Non-averaged version
func TestInsertSortTime(t *testing.T){
	for i := 0; i <= 50000; i=i+100 {
		s2 := generateSlice(i)
		t2 := time.Now()
		InsertionSort(s2)
		elapsed := time.Since(t2)
		fmt.Printf("%d,%d\n",len(s2),elapsed.Nanoseconds())
		//fmt.Printf("InsertionSort: Slice length of: %d, took: %d ns\n",len(s2),elapsed.Nanoseconds())
	}
}*/

func BenchmarkISort2(b *testing.B) {
	s := generateSlice(2)
	for n := 0; n < b.N; n++ {
		InsertionSort(s)
	}
}

func BenchmarkISort10(b *testing.B) {
	s := generateSlice(10)
	for n := 0; n < b.N; n++ {
		InsertionSort(s)
	}
}

func BenchmarkISort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(100)
		InsertionSort(s)
	}
}

func BenchmarkISort1000(b *testing.B) {
	s := generateSlice(1000)
	for n := 0; n < b.N; n++ {
		InsertionSort(s)
	}
}

func BenchmarkISort10000(b *testing.B) {
	s := generateSlice(10000)
	for n := 0; n < b.N; n++ {
		InsertionSort(s)
	}
}

func BenchmarkISort100000(b *testing.B) {
	s := generateSlice(100000)
	for n := 0; n < b.N; n++ {
		InsertionSort(s)
	}
}
/* Long time, dont bother.
func BenchmarkISort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateSlice(1000000)
		InsertionSort(s)
	}
*/
