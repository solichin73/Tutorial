package main

import "fmt"

func sumAll(numbers ...int)int{
  total := 0
  for _,value := range numbers{
    total += value
  }
  return total
}

func main(){
    total := sumAll(10,20,30,40,50,60,70,80)
    fmt.Println(total)

    slice := []int{10,20,30,40,50,60,70,80}
    total = sumAll(slice...)

    fmt.Println(total)

}
