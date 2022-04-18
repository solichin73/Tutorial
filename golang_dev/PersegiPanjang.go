package main

import "fmt"

func PersegPanjang(p int, l int)(luas int){
  var kelPersegiP int
  kelPersegiP = 2*(p+l)
  fmt.Println("Keliling Pp: ",kelPersegiP)

  luas = p * l
  return
}

func main(){
  fmt.Println("Luas Pp: ", PersegPanjang(20,30))
}

