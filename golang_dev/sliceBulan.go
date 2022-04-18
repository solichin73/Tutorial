package main

import "fmt"

func main(){
  var months = [...]string{
    "januari",
    "februari",
    "maret",
    "april",
    "mei",
    "juni",
    "juli",
    "agustus",
    "september",
    "oktober",
    "nopember",
    "desember",
  }
  
  fmt.Println("ini array.")
  fmt.Println(len(months))
  fmt.Println(months[3:7])

  fmt.Println("ini slice.")

  newSlice := [5]int{1,2,3,4,5,}
  fmt.Println(newSlice)

}
