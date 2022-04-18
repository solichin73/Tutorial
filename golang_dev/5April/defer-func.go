package main

import "fmt"

func endApp(){
  message := recover()
  if message != nil {
    fmt.Println("Error dg message : ", message)
  }
  fmt.Println("Aplikasi Selesai")
}


func runApp(error bool){
  defer endApp()
  if error{
    panic("ERROR")
  }
  fmt.Println("ini akan di eksekusi jika program oke")
}

func main(){
  runApp(true)
  fmt.Println("ini bagian akhir")
}
