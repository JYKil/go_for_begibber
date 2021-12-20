package main

import "fmt"

func multiply(a,b int) int{
  return a*b
}

func main() {
  // var name string = "kilga"
  name := "kilga" // same syntax. It only works inside functions.  
  name = "JY"
  fmt.Println(name)
  fmt.Println(multiply(2,2))
}
