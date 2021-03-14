package main

import "fmt"

  // write a function called linearSearch that accepts an array of integers and a target. if the array contains the target, return it's index, otherwise return -1

func linearSearch(arr []int, target int) int{
  for k, v := range arr {
    if v == target {
      fmt.Println(k)
      return k
    }
  }
  fmt.Println(-1)
  return -1
}

func main() {
  linearSearch([]int{10,15,20,25,30}, 15)
}