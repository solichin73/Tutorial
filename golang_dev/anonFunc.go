package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist) {
  if blacklist(name){
    fmt.Println("You are Blocked", name)
  }else{
    fmt.Println("Welcome ,",name)
  }
}

func main(){
  blacklist := func(name string) bool{
    return name =="admin"
  }

  registerUser("admin",blacklist)
  registerUser("Solichin",blacklist)

  // bisa juga dengan ini

  registerUser("root", func(name string)bool{
    return name == "root"
  })
  registerUser("Solichin", func(name string)bool{
    return name == "root"
  })


}
