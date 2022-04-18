package main

import "fmt"

func main(){
  var(
    firstName = "Muhamad"
    lastName = "Solichin"
    result = firstName == lastName
    nilaiUjian = 75
    absensi = 50

  )

  fmt.Println("ini program perbandingan.")
  fmt.Println(result)
  fmt.Println("lulus atau tidak")
  fmt.Println(nilaiUjian >= 80 && absensi >= 65)
}
