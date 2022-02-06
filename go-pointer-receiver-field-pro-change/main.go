package main

import (
	"fmt"
)

type sample struct {
  a string
  b int
  c int
}

func (sa *sample) sampleFunc(s string, i int, c int) {
  sa.a = s
  sa.b = i
  sa.c = c
}

func newSample() *sample {
  return &sample{}
}

func main() {
	sa := newSample()
  sa.sampleFunc("data", 30, 40)

  fmt.Printf("%v, %v, %v\n", sa.a, sa.b, sa.c)
  // data. 30, 40
}
