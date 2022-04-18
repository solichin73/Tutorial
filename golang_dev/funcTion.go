package main

import "fmt"

func add(x int, y int)int{
    total := 0 
    total = x * y
    return total
}

func main(){
  sum := add(21,2)
  fmt.Println(sum)
}
