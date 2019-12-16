package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"

	"github.com/afterous/go-algo/pkg/BinarySearchTree"
	"testing"
)
var result bool // so compiler doesn't optimize compare() out.


// Helper
func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(n))
	}
	return s
}


// Compare if both are equal (same elements).
func TestCompareT(t *testing.T){
	bree := BinarySearchTree.Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := BinarySearchTree.NewNode(v)
		bree.BSTInsert(n)
	}
	tsB := time.Now()
	a := compare(bree,elements)
	tsA := time.Now()
	diff := tsA.Sub(tsB)
	fmt.Printf("Test: %s, Time it took in nanoseconds %d\n",t.Name(),diff.Nanoseconds())
	assert.True(t,a)
}

// Benchmark for above test
func BenchmarkCompareT10(b *testing.B){
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		bree := BinarySearchTree.Newbtree()
		for _,v := range elements {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}
		b.StartTimer()
		a := compare(bree,elements)
		result = a// so we don't optimize out compare
	}
}

// Compare if elements are not equal.
func TestCompareF10(t *testing.T){
	bree := BinarySearchTree.Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	mock := []int{1,2,3,4,5,6,7,8}
	for _,v := range elements {
		n := BinarySearchTree.NewNode(v)
		bree.BSTInsert(n)
	}
	tsB := time.Now()
	a := compare(bree,mock)
	tsA := time.Now()
	diff := tsA.Sub(tsB)
	fmt.Printf("Test: %s, Time it took in nanoseconds %d\n",t.Name(),diff.Nanoseconds())
	assert.False(t,a)
}

// Benchmark for above test
func BenchmarkCompareF10(b *testing.B){
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	mock := []int{1,2,3,4,5,6,7,8}
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		bree := BinarySearchTree.Newbtree()
		for _,v := range elements {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}
		b.StartTimer()
		a := compare(bree,mock)
		result = a// so we don't optimize out compare
	}
}

// Compare if length of arrays are different.
func TestCompareFLen(t *testing.T){
	bree := BinarySearchTree.Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	mock := []int{99}

	for _,v := range elements {
		n := BinarySearchTree.NewNode(v)
		bree.BSTInsert(n)
	}
	tsB := time.Now()
	a := compare(bree,mock)
	tsA := time.Now()
	diff := tsA.Sub(tsB)
	fmt.Printf("Test: %s, Time it took in nanoseconds %d\n",t.Name(),diff.Nanoseconds())
	assert.False(t,a)
}

// Benchmark for above test
func BenchmarkCompareLen10(b *testing.B){
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	mock := []int{99}
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		bree := BinarySearchTree.Newbtree()
		for _,v := range elements {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}
		b.StartTimer()
		a := compare(bree,mock)
		result = a// so we don't optimize out compare
	}
}

// Benchmarks with random slices.
func BenchmarkCompare10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		s := generateSlice(10)
		bree := BinarySearchTree.Newbtree()
		for _,v := range s {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}
		b.StartTimer()
		a := compare(bree,s)
		result = a// so we don't optimize out compare
	}
}

func BenchmarkCompare100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		s := generateSlice(100)
		bree := BinarySearchTree.Newbtree()
		for _,v := range s {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}

		b.StartTimer()
		a := compare(bree,s)
		result = a// so we don't optimize out compare
	}
}

func BenchmarkCompare1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		s := generateSlice(1000)
		bree := BinarySearchTree.Newbtree()
		for _,v := range s {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}

		b.StartTimer()
		a := compare(bree,s)
		result = a// so we don't optimize out compare
	}
}


func BenchmarkCompare10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer() // Don't record time for init.
		s := generateSlice(10000)
		bree := BinarySearchTree.Newbtree()
		for _,v := range s {
			n := BinarySearchTree.NewNode(v)
			bree.BSTInsert(n)
		}

		b.StartTimer()
		a := compare(bree,s)
		result = a// so we don't optimize out compare
	}
}