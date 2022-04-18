package main

import "fmt"

func main(){
  var(
    angka1 = 12
    angka2 = 23
    result = angka2 + angka1
  )
  angka1++
  fmt.Println(result*angka2)
  fmt.Println(angka1)

  // ini contoh perbandingan

  fmt.Println(result > angka2)
  fmt.Println(result < angka2)

}
