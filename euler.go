package main

func P1(n []int) int {
  s := 0
  for v := range n {
    if v % 3 == 0 || v % 5 == 0 {
      s += v
    }
  }
  return s
}
