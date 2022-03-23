package main

import "fmt"

func main() {
   var cmd string
   for true {
      fmt.Scan(&cmd)
      if cmd == "1"{
         fmt.Println("open menu")
      }  else if cmd == "2" {
         fmt.Println("close menu")
      }  else{
         fmt.Println("unkonwn command")
      }
   }
}