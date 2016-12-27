package main

import "testing"
import "fmt"

func TestP1(t *testing.T) {
	list := []int{}

  for i := 1; i <= 1000; i++ {
    list = append(list, i)
  }

	result := P1(list)

	if result == 0 {
    t.Error("Problem #1: Invalid sum of", result)
	} else {
    fmt.Printf("Problem #1: The sum of elements is %v\n", result)
  }
}
