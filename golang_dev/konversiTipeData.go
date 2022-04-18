package main

import "fmt"

func main(){
  var (
    nilai int32 = 100000
    nilai1 = int64(nilai)
    nilai2 = int8(nilai1)
  )

  fmt.Println(nilai)
  fmt.Println(nilai1)
  fmt.Println(nilai2)


}
