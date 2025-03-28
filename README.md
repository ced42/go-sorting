# Go Sorting Performance Benchmarks

## Overview
This repository contains performance benchmarks comparing different sorting methods in Go, exploring the efficiency of various sorting approaches introduced across different versions of the language.

## Methods Compared
- `sort.Sort`
- `sort.Slice`
- `slices.SortFunc`
- `slices.Sort`

## Benchmark Results

### Performance (ns/op)

| Method           | 10 elem | 100 elem | 1000 elem | 10000 elem | 100000 elem |
|-----------------|---------|----------|-----------|------------|------------|
| sort.Sort       | 483.6   | 5,356    | 84,800    | 764,741    | 12,272,027 |
| sort.Slice      | 731.3   | 5,510    | 63,731    | 759,724    | 9,089,272  |
| slices.SortFunc | 213.3   | 4,147    | 53,060    | 614,163    | 7,153,118  |
| slices.Sort     | 159.9   | 2,920    | 39,238    | 449,679    | 5,679,881  |

### Memory Allocations (B/op)

| Method           | 10 elem | 100 elem | 1000 elem | 10000 elem | 100000 elem |
|-----------------|---------|----------|-----------|------------|------------|
| sort.Sort       | 24      | 24       | 24        | 24         | 24         |
| sort.Slice      | 56      | 56       | 56        | 56         | 56         |
| slices.SortFunc | 0       | 0        | 0         | 0          | 0          |
| slices.Sort     | 0       | 0        | 0         | 0          | 0          |

## Key Observations
- `slices.Sort` is consistently the fastest method
- `slices.SortFunc` provides a close second-best performance
- Newer slice methods (introduced in Go 1.21) significantly outperform older sorting approaches
- Zero memory allocations for `slices` package methods

## Origin

These benchmarks were initially developed through a collaborative conversation with Claude AI (Anthropic) to analyze and compare sorting performance in Go. The iterative process of developing and refining the benchmarks demonstrates the potential for AI-assisted performance analysis and code development.

## Requirements
- Go 1.21 or later

## Running Benchmarks

```bash
go test -bench=. -benchmem