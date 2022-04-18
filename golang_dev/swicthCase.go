package main

import "fmt"

func main(){
  name := "julia nadiem"

  length := len(name)
  switch{
  case length > 10:
    fmt.Println("jenenge kedawan")
  case length > 5:
    fmt.Println("sek kedawan")
  default:
    fmt.Println("wes cukup")
  }
  

}
