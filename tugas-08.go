package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main(){
rand.Seed(time.Now().Unix())
runtime.GOMAXPROCS(2)

var message = make(chan int)

go kirimdata(message)
terimadata(message)
}

func kirimdata(ch chan<- int){
  for i:=0; true;i++{
    ch <- i
    time.Sleep(time.Duration(rand.Int()%10 +1) * time.Second)
  }
}

func terimadata(ch <- chan int){
  loop:
  for{
    select{
    case data := <-ch:
      fmt.Print("Apa Kabar Teman Teman ", data, "\n")
    case <-time.After(time.Second * 5):
      fmt.Println("Timeout, Apa Kabar Teman")
      break loop
    }
  }
}
