package main

import "fmt"

func main() {
  defer fmt.Println("Third line")
  fmt.Println("First line")
  fmt.Println("Second line")
}
