package sorting

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

// Int64Slice implements sort.Interface for []int64
type Int64Slice []int64

func (s Int64Slice) Len() int           { return len(s) }
func (s Int64Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Int64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Helper function to generate random int64 slice
func generateRandomInt64Slice(n int) []int64 {
	slice := make([]int64, n)
	for i := range slice {
		slice[i] = rand.Int63()
	}
	return slice
}

// Benchmark sort.Sort with different slice sizes
func BenchmarkSortSort_10(b *testing.B)     { benchmarkSortSort(10, b) }
func BenchmarkSortSort_100(b *testing.B)    { benchmarkSortSort(100, b) }
func BenchmarkSortSort_1000(b *testing.B)   { benchmarkSortSort(1000, b) }
func BenchmarkSortSort_10000(b *testing.B)  { benchmarkSortSort(10000, b) }
func BenchmarkSortSort_100000(b *testing.B) { benchmarkSortSort(100000, b) }

func benchmarkSortSort(size int, b *testing.B) {
	// Prepare data once, then reuse for each iteration
	data := make([]int64, size)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Copy random data for this iteration
		copy(data, generateRandomInt64Slice(size))
		b.StartTimer()
		
		sort.Sort(Int64Slice(data))
	}
}

// Benchmark sort.Slice with different slice sizes
func BenchmarkSortSlice_10(b *testing.B)     { benchmarkSortSlice(10, b) }
func BenchmarkSortSlice_100(b *testing.B)    { benchmarkSortSlice(100, b) }
func BenchmarkSortSlice_1000(b *testing.B)   { benchmarkSortSlice(1000, b) }
func BenchmarkSortSlice_10000(b *testing.B)  { benchmarkSortSlice(10000, b) }
func BenchmarkSortSlice_100000(b *testing.B) { benchmarkSortSlice(100000, b) }

func benchmarkSortSlice(size int, b *testing.B) {
	// Prepare data once, then reuse for each iteration
	data := make([]int64, size)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Copy random data for this iteration
		copy(data, generateRandomInt64Slice(size))
		b.StartTimer()
		
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}
}

// Benchmark slices.SortFunc with different slice sizes
func BenchmarkSlicesSortFunc_10(b *testing.B)     { benchmarkSlicesSortFunc(10, b) }
func BenchmarkSlicesSortFunc_100(b *testing.B)    { benchmarkSlicesSortFunc(100, b) }
func BenchmarkSlicesSortFunc_1000(b *testing.B)   { benchmarkSlicesSortFunc(1000, b) }
func BenchmarkSlicesSortFunc_10000(b *testing.B)  { benchmarkSlicesSortFunc(10000, b) }
func BenchmarkSlicesSortFunc_100000(b *testing.B) { benchmarkSlicesSortFunc(100000, b) }

func benchmarkSlicesSortFunc(size int, b *testing.B) {
	// Prepare data once, then reuse for each iteration
	data := make([]int64, size)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Copy random data for this iteration
		copy(data, generateRandomInt64Slice(size))
		b.StartTimer()
		
		slices.SortFunc(data, func(a, b int64) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		})
	}
}

// Benchmark slices.Sort with different slice sizes
func BenchmarkSlicesSort_10(b *testing.B)     { benchmarkSlicesSort(10, b) }
func BenchmarkSlicesSort_100(b *testing.B)    { benchmarkSlicesSort(100, b) }
func BenchmarkSlicesSort_1000(b *testing.B)   { benchmarkSlicesSort(1000, b) }
func BenchmarkSlicesSort_10000(b *testing.B)  { benchmarkSlicesSort(10000, b) }
func BenchmarkSlicesSort_100000(b *testing.B) { benchmarkSlicesSort(100000, b) }

func benchmarkSlicesSort(size int, b *testing.B) {
	// Prepare data once, then reuse for each iteration
	data := make([]int64, size)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Copy random data for this iteration
		copy(data, generateRandomInt64Slice(size))
		b.StartTimer()
		
		slices.Sort(data)
	}
}
