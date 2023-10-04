package main

import "testing"

func TestDo(t *testing.T){
  if Do() == 3 {
    return
  }else{
    t.Fail()
  }
}
