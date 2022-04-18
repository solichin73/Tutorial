package main

import "fmt"

func main(){
  person := map[string]string {
    "names" : "muhamad solichin",
    "country" : "blitar",
  }

  fmt.Println(person)
  fmt.Println(person["names"])
  fmt.Println(person["country"])
}
