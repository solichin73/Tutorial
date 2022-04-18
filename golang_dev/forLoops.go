package main

import "fmt"

func main(){
  counter := 1

  for counter < 11 {
    fmt.Println("perulangan ke ", counter)
    counter+=2
  }
}
