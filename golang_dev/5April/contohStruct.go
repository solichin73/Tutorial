package main

import "fmt"

type Customer struct{
  Name, Address string
  Age int
}

//menambahkan func
func (customer Customer) sayHi(name string){
  fmt.Println("Hello,",name,"My name is",customer.Name)
}

func (a Customer) sayHuu(){
  fmt.Println("Huuu from",a.Name)
}

func main(){
  Solichin := Customer{
    Name : "Solichin",
    Address : "Indonesia",
    Age : 28,
  }
  fmt.Println(Solichin)
// atau dengan

 budi := Customer{"Budi","Cino",45}

 fmt.Println(budi)

 Solichin.sayHi("Yugi")

 Solichin.sayHuu()


}
