package main

import "fmt"

type HasName interface{
 getName()string
}

func sayHello(hasName HasName){
  fmt.Println("Hello,",hasName.getName())
}

type Person struct{
  Name string
}

func (person Person)getName() string{
  return person.Name
}

type Animal struct{
  Name string
}

func (animal Animal)getName() string{
  return animal.Name
}

func main(){
  var Solichin Person
  Solichin.Name = "Solichin"
  
  cat := Animal{
    Name : "Puniyem",
  }

  sayHello(cat)

  sayHello(Solichin)

}
