package main

import "fmt"

func main(){
  fmt.Println("Hello world!")

  res := Do()
  fmt.Println(res)
}

func Do()int{
  return 3 
}
