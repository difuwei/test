package main

import "fmt"

func main(){ 
  fmt.Println("22");

  var arr = [...]int{1,2,4,12,44}
  fmt.Println(arr[3])

  arr2 :=[9]int{3,3,3,4,5,6}
  fmt.Println(arr2[3])

  var slice = arr2[1:2]
  fmt.Println(slice)

  var mm = map[int]string{1:"a",2:"b"}
  fmt.Println(mm[1])

  e,ok := mm[4]
  fmt.Println(e)
  fmt.Println(ok)

  ch := make(chan int,5)
  ch <- 3
  value := <- ch 
  fmt.Println(value)
}

