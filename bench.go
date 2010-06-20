// Copyright (c) 2010, Jonathan Wills (runningwild@gmail.com)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The benchmark from http://github.com/runningwild/go-btree
package main

import "rand"
import "time"
import "fmt"
import "flag"

var N int
var R int
func init() {
  flag.IntVar(&N, "size", 1000000, "Size of input arrays")
  flag.IntVar(&R, "runs", 10, "Number of times to repeat each test")
  flag.Parse()
}



func NewIntSet() SortedSet {
    intComp := func(i int, j int) int {
        return i - j;
    }
    
    return NewTree(intComp)
}

func Bench(d []int) []float {
  times := make([]float, 7)

  t := NewIntSet()

  start := time.Nanoseconds()
  for _,v := range d {
    t.Insert(v)
  }
  times[0] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  for _,v := range d {
    t.Insert(v)
  }
  times[1] = float(time.Nanoseconds() - start) / 1000000000.0
  
  start = time.Nanoseconds()
  for i := 0; i < len(d)/2; i++ {
    t.Remove(d[i])
  }
  times[2] = float(time.Nanoseconds() - start) / 1000000000.0
  
  start = time.Nanoseconds()
  for i := 0; i < len(d)/2; i++ {
    t.Remove(d[i])
  }
  times[3] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  for v := range d {
    t.Contains(v)
  }
  times[4] = float(time.Nanoseconds() - start) / 1000000000.0
  
  start = time.Nanoseconds()
  sum := 0
  t.Foreach(func (e int) { sum += e})
  times[5] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  a := make([]*Node, N);
  for i := 0; i < N; i++ {
      a[i] = NewNode(1);
  } 
  times[6] = float(time.Nanoseconds() - start) / 1000000000.0
  

  return times
}

func main() { 
  names := []string {
    "Unique Inserts",
    "Repeated Inserts",
    "Unique Deletes",
    "Repeated Deletes",
    "Queries",
    "Iterations",
    "NewNodes",    
  }

  d := rand.Perm(N)
  total := Bench(d)
  for i := 1; i < R; i++ {
    times := Bench(d)
    for j := range times {
      total[j] += times[j]
    }
  }
  for i := range total {
    total[i] /= float(R)
  }

  fmt.Printf("Using input size %d and averaged over %d runs.\n", N, R)
  fmt.Printf("%3.3f:\t%d\t%s\n", total[0], N, names[0])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[1], N, names[1])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[2], N/2, names[2])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[3], N/2, names[3])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[4], N, names[4])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[5], N, names[5])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[6], N, names[6])
}
