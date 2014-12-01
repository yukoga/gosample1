package main

import (
  "fmt"
  "github.com/yukoga/newmath"
)

func swap(x, y string) (string, string) {
  return y, x
}


func main() {
  fmt.Printf("Hello, Go world. Sqrt(2) = %v\n", newmath.Sqrt(2))
  fmt.Printf("Hello, Go world. power_n(2,0) = %v\n", newmath.Power_n(2,0))
  fmt.Printf("Hello, Go world. power_n(2,1) = %v\n", newmath.Power_n(2,1))
  fmt.Printf("Hello, Go world. power_n(2,2) = %v\n", newmath.Power_n(2,2))
  fmt.Printf("Hello, Go world. power_n(2,3) = %v\n", newmath.Power_n(2,3))
  a, b:=swap("Hello", "world")
  fmt.Printf("Hello, Go world. \n %s, %s. \n", a, b)
}
